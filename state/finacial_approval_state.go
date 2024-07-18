package state

import "fmt"

type FinacialApprovalState struct{}

func (*FinacialApprovalState) Approval(ctx *Context) {
	fmt.Println("财务审批成功")
	ctx.SetState(&EndApprovalState{})

}

func (*FinacialApprovalState) Reject(ctx *Context) {
	fmt.Println("财务审批拒绝")
	ctx.SetState(&LeaderApprovalState{})
}

func (*FinacialApprovalState) GetName() string {
	return "FinacialApprovalState"
}
