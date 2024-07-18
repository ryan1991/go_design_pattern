package strategy

type PaymentBetter struct {
	context   *PaymentContext
	strategys []PaymentStrategy
}

func NewPaymentBetter(name, cardId string, money int, payType PAYTYPE, strategys []PaymentStrategy) *PaymentBetter {

	return &PaymentBetter{
		context: &PaymentContext{
			Name:    name,
			CardId:  cardId,
			Money:   money,
			PayType: payType,
		},

		strategys: strategys,
	}

}

func (p *PaymentBetter) Pay() {
	payType := p.context.PayType
	for _, policy := range p.strategys {
		if (policy).PayType() == payType {
			(policy).Pay(p.context)
		}

	}

}
