package state

import "fmt"

type EndApprovalState struct{}

func (*EndApprovalState) Approval(ctx *Context) {
	fmt.Println("流程审批结束")
}

func (*EndApprovalState) Reject(ctx *Context) {
	fmt.Println("流程审批结束")
}

func (*EndApprovalState) GetName() string {
	return "EndApprovalState"
}
