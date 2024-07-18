package factory_method

// Go中不存在继承 所以使用匿名组合来实现
// 在值接收器和指针接收器之间做出选择并不总是那么简单
// 值接收器操作副本， 指针接收器操作原始数据
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {

	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b

}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

type MinusOperatorFactory struct {
}

func (MinusOperatorFactory) Create() Operator {

	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
