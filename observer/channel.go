package main

type Channel interface {
	register(subscriber Subscriber)
	deregister(subscriber Subscriber)
	notifyAll()
}