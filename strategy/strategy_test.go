package strategy

import "testing"

func TestPayment(t *testing.T) {
	var payment = NewPayment("junbao", "", 100, &CashPay{})
	payment.Pay()

	payment = NewPayment("junbao", "", 200, &BankPay{})
	payment.Pay()
}

func TestPaymentBetter(t *testing.T) {
	// strategys := make([]PaymentStrategy, 2)
	// strategys[0] = &CashPay{}
	// strategys[1] = &BankPay{}

	payment := NewPaymentBetter("junbao", "", 100, BANK, []PaymentStrategy{&CashPay{}, &BankPay{}})

	payment.Pay()
}
