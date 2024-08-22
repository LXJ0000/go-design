package main

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) handle(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.handle(patient)
		return
	}
	fmt.Println("Reception registering patient")
	patient.registrationDone = true
	r.next.handle(patient)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}
