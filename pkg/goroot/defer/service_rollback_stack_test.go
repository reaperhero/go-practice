package _defer

import (
	"fmt"
	"testing"
	"time"
)

const (
	StackStateSuccess = "success"
	StackStateFail    = "fail"
)

type StackService struct {
	finshChan   chan string
	rollBackFun *Stack
}

func (s *StackService) MonitorRollback() {
	go func() {
		finshState := <-s.finshChan
		switch finshState {
		case StateSuccess:
		case StateFail:
			for !s.rollBackFun.Empty() {
				callFun := s.rollBackFun.Pop().(func())
				callFun()
			}
		}
	}()
}

func newStackService() *StackService {
	return &StackService{
		finshChan:   make(chan string),
		rollBackFun: NewStack(),
	}
}

func (s *StackService) stackServiceA() error {
	defer s.stackRollbackServiceA()
	fmt.Println("start service a ....")
	return nil
}

func (s *StackService) stackRollbackServiceA() {
	s.rollBackFun.Push(func() { fmt.Println("rollback service a ....") })
}

func (s *StackService) stackServiceB() error {
	defer s.stackRollbackServiceB()
	fmt.Println("start service b ....")
	return nil
}

func (s *StackService) stackRollbackServiceB() {
	s.rollBackFun.Push(func() { fmt.Println("rollback service b ....") })
}

func (s *StackService) stackServiceC() error {
	defer s.stackRollbackServiceC()
	fmt.Println("start service c ....")
	return nil
}

func (s *StackService) stackRollbackServiceC() {
	s.rollBackFun.Push(func() { fmt.Println("rollback service c ....") })

}

func TestStackRollbak(t *testing.T) {
	s := newStackService()
	s.stackServiceA()
	s.stackServiceB()
	//s.stackServiceC()
	s.MonitorRollback()
	s.finshChan <- StackStateFail
	time.Sleep(time.Second * 1)
}
