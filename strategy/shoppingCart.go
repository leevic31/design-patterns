package main

type ShoppingCart struct {
	paymentType Payment
}

func (s *ShoppingCart) setPayment(p Payment) {
	s.paymentType = p
}

func (s *ShoppingCart) pay() string {
	return s.paymentType.pay()
}