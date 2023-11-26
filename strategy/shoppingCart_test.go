package main

import "testing"

func TestSetPaymentForCreditCard(t *testing.T) {
	cart := &ShoppingCart{}
	creditCard := &CreditCard{}
	cart.setPayment(creditCard)
	if cart.paymentType != creditCard {
		t.Fatalf("Payment type is incorrect")
	}
}

func TestSetPaymentForDebitCard(t *testing.T) {
	cart := &ShoppingCart{}
	debitCard := &DebitCard{}
	cart.setPayment(debitCard)
	if cart.paymentType != debitCard {
		t.Fatalf("Payment type is incorrect")
	}
}