package main

type DebitCard struct {}

func (d *DebitCard) pay() string {
	return "Paying with debit card"
}