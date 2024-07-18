package state

import "fmt"

type EmployeeStartState struct{}

func (*EmployeeStartState) Approval(ctx *Context) {
	fmt.Println("员工发起审批")
	ctx.SetState(&LeaderApprovalState{})

}

func (*EmployeeStartState) Reject(ctx *Context) {
	fmt.Println("员工审批被拒绝")
}

func (*EmployeeStartState) GetName() string {
	return "EmployeeStartState"
}
