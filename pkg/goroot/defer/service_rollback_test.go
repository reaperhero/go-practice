package _defer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	StateSuccess = "success"
	StateFail    = "fail"
)

type Service struct {
	finshState string
	cond       *sync.Cond
}

func newService() *Service {
	return &Service{
		finshState: "",
		cond:       sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Service) serviceA() error {
	defer s.rollbackServiceA()
	fmt.Println("start service a ....")
	return nil
}

func (s *Service) rollbackServiceA() {
	func() {
		s.cond.L.Lock()
		s.cond.Wait()
		switch s.finshState {
		case StateSuccess:
			fmt.Println("result no need roll a ....")
		case StateFail:
			fmt.Println("rollback service a ....")
		}
		s.cond.L.Unlock()
	}()
}

func (s *Service) serviceB() error {
	defer s.rollbackServiceB()
	fmt.Println("start service b ....")
	return nil
}

func (s *Service) rollbackServiceB() {
	go func() {
		s.cond.L.Lock()
		s.cond.Wait()
		switch s.finshState {
		case StateSuccess:
			fmt.Println("result no need roll b ....")
		case StateFail:
			fmt.Println("rollback service b ....")
		}
		s.cond.L.Unlock()
	}()
}

func (s *Service) serviceC() error {
	defer s.rollbackServiceC()
	fmt.Println("start service c ....")
	return nil
}

func (s *Service) rollbackServiceC() {
	go func() {
		s.cond.L.Lock()
		s.cond.Wait()
		switch s.finshState {
		case StateSuccess:
			fmt.Println("result no need roll c ....")
		case StateFail:
			fmt.Println("rollback service c ....")
		}
		s.cond.L.Unlock()
	}()
}

func TestRollbakc(t *testing.T) {
	s := newService()
	s.serviceA()
	s.serviceB()
	s.serviceC()

	time.Sleep(time.Second*1)
	s.finshState = StateSuccess

	s.cond.Broadcast()
	time.Sleep(time.Second*1)
}
