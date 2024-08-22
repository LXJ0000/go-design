package main

type Subject interface {
	register(Observer) chan struct{}
	unregister(Observer)
	notifyAll()
}
