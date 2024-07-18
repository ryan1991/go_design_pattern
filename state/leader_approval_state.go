package state

import "fmt"

type LeaderApprovalState struct{}

func (*LeaderApprovalState) Approval(ctx *Context) {
	fmt.Println("领导审批成功")
	ctx.SetState(&FinacialApprovalState{})

}

func (*LeaderApprovalState) Reject(ctx *Context) {
	fmt.Println("领导审批拒绝")
	ctx.SetState(&EmployeeStartState{})
}

func (*LeaderApprovalState) GetName() string {
	return "LeaderApprovalState"
}
