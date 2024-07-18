package state

type Context struct {
	// State  也可以ctx.State ?
	state State
}

func (ctx *Context) SetState(state State) {
	ctx.state = state
}

func (ctx *Context) Approval() {
	ctx.state.Approval(ctx)
}

func (ctx *Context) Reject() {
	ctx.state.Reject(ctx)
}

func (ctx *Context) GetStateName() string {
	return ctx.state.GetName()
}
