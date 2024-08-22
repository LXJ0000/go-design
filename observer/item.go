package main

import (
	"fmt"
)

type Item struct {
	name      string
	observers map[Observer]chan struct{}
	inStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name:      name,
		observers: make(map[Observer]chan struct{}),
	}
}

func (i *Item) register(o Observer) chan struct{} {
	ch := make(chan struct{}, 10)
	i.observers[o] = ch
	return ch
}

func (i *Item) unregister(o Observer) {
	ch := i.observers[o]
	close(ch) // notify the observer to stop
	delete(i.observers, o)
}

func (i *Item) notifyAll() {
	for o := range i.observers {
		ch := i.observers[o]
		ch <- struct{}{}
		// o.update(i.name)
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
