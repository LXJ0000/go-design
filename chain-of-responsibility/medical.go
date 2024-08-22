package main

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) handle(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.handle(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.handle(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
