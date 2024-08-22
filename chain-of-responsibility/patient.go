package main

type Patient struct {
	name              string
	registrationDone  bool // 前台
	doctorCheckUpDone bool // 医生
	medicineDone      bool // 药房
	paymentDone       bool // 收费
}
