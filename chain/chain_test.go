package chain

import "testing"

func TestChain(t *testing.T) {
	c1 := NewChain(&ProjectManager{})
	c2 := NewChain(&DeptManager{})
	c3 := NewChain(&GeneralManager{})

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	c1.HandleFeeRequest("junbao", 1200)

}
