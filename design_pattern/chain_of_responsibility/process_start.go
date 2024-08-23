package chain_of_responsibility

type Start struct {
	Next
}

func (s *Start) Do(patient *Patient) error {
	return nil
}
