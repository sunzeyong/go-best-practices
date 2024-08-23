package chain_of_responsibility

// Patient 传递中的请求体
type Patient struct {
	Name string
	RegistrationDone bool
	DoctorCheckUpDone bool
	MedicineDone bool
	PaymentDone bool
}
