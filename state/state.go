package state

type State interface {
	Approval(ctx *Context)
	Reject(ctx *Context)

	//当前状态
	GetName() string
}
