package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) handle(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.handle(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.handle(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}
