package chain_of_responsibility

type PatientHandler interface {
	Execute(*Patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*Patient) error
}
