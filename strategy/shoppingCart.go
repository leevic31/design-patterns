package main

type ShoppingCart struct {
	paymentType Payment
}

func (s *ShoppingCart) setPayment(p Payment) {
	s.paymentType = p
}