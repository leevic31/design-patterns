package main

type CreditCard struct {}

func (c *CreditCard) pay() string {
	return "Paying with credit card"
}