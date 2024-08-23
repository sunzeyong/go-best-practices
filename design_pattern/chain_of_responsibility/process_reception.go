package chain_of_responsibility

import "fmt"

type Reception struct {
	Next
}

func (r *Reception) Do(patient *Patient) error {
	if patient.RegistrationDone {
		fmt.Println("patient has reception")
		return nil
	}

	patient.RegistrationDone = true
	fmt.Println("patient reception")
	return nil
}
