package chain_of_responsibility

import "fmt"

type Cashier struct {
	Next
}

func (c *Cashier) Do(patient *Patient) error {
	if patient.PaymentDone {
		fmt.Println("patient has already paid")
		return nil
	}

	patient.PaymentDone = true
	fmt.Println("patient done")
	return nil
}
