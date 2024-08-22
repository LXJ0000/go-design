package main

type Department interface {
	handle(*Patient)
	setNext(Department)
}
