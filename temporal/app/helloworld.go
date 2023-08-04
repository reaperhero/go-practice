package app

import (
	"context"
	"errors"
	"fmt"
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"log"
	"math/rand"
	"time"
)

const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}

func MoneyTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {

	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        0, // unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	// Withdraw money.
	var withdrawOutput string

	withdrawErr := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &withdrawOutput)

	if withdrawErr != nil {
		return "", withdrawErr
	}

	// Deposit money.
	var depositOutput string

	depositErr := workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &depositOutput)

	if depositErr != nil {
		// The deposit failed; put money back in original account.

		var result string

		refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &result)

		if refundErr != nil {
			return "",
				fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
					input.TargetAccount, depositErr, input.SourceAccount, refundErr)
		}

		return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
			input.TargetAccount, input.SourceAccount, depositErr)
	}

	result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", withdrawOutput, depositOutput)
	return result, nil
}

// @@@SNIPSTART money-transfer-project-template-go-activity-withdraw
func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-deposit
func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	// Uncomment the next line and comment the one after that to simulate an unknown failure
	// confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return confirmation, err
}

// @@@SNIPEND

// @@@SNIPSTART money-transfer-project-template-go-activity-refund
func Refund(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Refunding $%v back into account %v.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
	bank := BankingService{"bank-api.example.com"}
	confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

type account struct {
	AccountNumber string
	Balance       int64
}

type bank struct {
	Accounts []account
}

func (b bank) findAccount(accountNumber string) (account, error) {

	for _, v := range b.Accounts {
		if v.AccountNumber == accountNumber {
			return v, nil
		}
	}

	return account{}, errors.New("account not found")
}

// InsufficientFundsError is raised when the account doesn't have enough money.
type InsufficientFundsError struct{}

func (m *InsufficientFundsError) Error() string {
	return "Insufficient Funds"
}

// InvalidAccountError is raised when the account number is invalid
type InvalidAccountError struct{}

func (m *InvalidAccountError) Error() string {
	return "Account number supplied is invalid"
}

// our mock bank
var mockBank = &bank{
	Accounts: []account{
		{AccountNumber: "85-150", Balance: 2000},
		{AccountNumber: "43-812", Balance: 0},
	},
}

// BankingService mocks interaction with a bank API. It supports withdrawals
// and deposits
type BankingService struct {
	// the hostname is to make it more realistic. This code does not
	// actually make any network calls.
	Hostname string
}

// Withdraw simulates a Withdrawwal from a bank.
// Acceptsthe account number (string), amount (int), and a reference ID (string)
// for idempotent transaction tracking.
// Returns a transaction id when successful
// Returns various errors based on amount and account number.
func (client BankingService) Withdraw(accountNumber string, amount int, referenceID string) (string, error) {
	acct, err := mockBank.findAccount(accountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	if amount > int(acct.Balance) {
		return "", &InsufficientFundsError{}
	}

	return generateTransactionID("W", 10), nil
}

// Deposit simulates a Withdrawwal from a bank.
// Accepts the account number (string), amount (int), and a reference ID (string)
// for idempotent transaction tracking.
// Returns a transaction id when successful
// Returns InvalidAccountError if the account is invalid
func (client BankingService) Deposit(accountNumber string, amount int, referenceID string) (string, error) {

	_, err := mockBank.findAccount(accountNumber)
	if err != nil {
		return "", &InvalidAccountError{}
	}

	return generateTransactionID("D", 10), nil
}

// DepositThatFails simulates an unknonw error.
func (client BankingService) DepositThatFails(accountNumber string, amount int, referenceID string) (string, error) {
	return "", errors.New("This deposit has failed.")
}

func generateTransactionID(prefix string, length int) string {
	randChars := make([]byte, length)
	for i := range randChars {
		allowedChars := "0123456789"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return prefix + string(randChars)
}
