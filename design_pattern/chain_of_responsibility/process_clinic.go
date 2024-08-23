package chain_of_responsibility

import "fmt"

type Clinic struct {
	Next
}

func (c *Clinic) Do(patient *Patient) error {
	if patient.DoctorCheckUpDone {
		fmt.Println("doctor has check up")
		return nil
	}

	patient.DoctorCheckUpDone = true
	fmt.Println("doctor check done")
	return nil
}
