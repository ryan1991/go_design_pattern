package strategy

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PAYTYPE int

const (
	CASH PAYTYPE = iota
	BANK
)

type PaymentContext struct {
	Name, CardId string
	Money        int
	PayType      PAYTYPE
}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)

	PayType() PAYTYPE
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},

		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type CashPay struct{}

func (*CashPay) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

func (*CashPay) PayType() PAYTYPE {
	return CASH
}

type BankPay struct{}

func (*BankPay) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank\n", ctx.Money, ctx.Name)
}

func (*BankPay) PayType() PAYTYPE {
	return BANK
}
