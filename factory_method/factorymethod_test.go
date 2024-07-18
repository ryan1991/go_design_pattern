package factory_method

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {

	//var ()支持一次声明多个变量和初始化
	var (
		factory OperatorFactory
	)
	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("result is err")
	}

	factory = MinusOperatorFactory{}
	if compute(factory, 2, 1) != 1 {
		t.Fatal("result is err")
	}

}
