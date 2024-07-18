package chain

import "fmt"

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(rc *RequestChain) {
	r.successor = rc
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		r.Manager.HandleFeeRequest(name, money)
	}

	if r.successor != nil {

		//* 这里不能用
		// r.successor.Manager.HandleFeeRequest(name, money)
		r.successor.HandleFeeRequest(name, money)
	}

	return false
}

func (r *RequestChain) HaveRight(money int) bool {
	return true
}

type ProjectManager struct{}

func NewChain(manager Manager) *RequestChain {
	return &RequestChain{
		Manager: manager,
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money <= 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("Project manager permit %s %d fee request\n", name, money)
	return true
}

type DeptManager struct{}

func (*DeptManager) HaveRight(money int) bool {
	return money > 500 && money <= 1000
}

func (*DeptManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("dept manager permit %s %d fee request\n", name, money)
	return true
}

type GeneralManager struct{}

func (*GeneralManager) HaveRight(money int) bool {
	return money > 1000
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	fmt.Printf("general manager permit %s %d fee request\n", name, money)
	return true
}
