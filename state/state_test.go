package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	ctx := Context{}
	ctx.state = &EmployeeStartState{}
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
}

func TestState2(t *testing.T) {
	ctx := Context{}
	ctx.state = &EmployeeStartState{}
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Reject()
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
	fmt.Println("current state: ", ctx.GetStateName())
	ctx.Approval()
}
