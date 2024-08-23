package chain_of_responsibility

import "fmt"

type Pharmacy struct {
	Next
}

func (p *Pharmacy) Do(patient *Patient) error {
	if patient.MedicineDone {
		fmt.Println("patient has done medicine")
		return nil
	}
	patient.MedicineDone = true
	fmt.Println("patient has buy medicine")
	return nil
}
