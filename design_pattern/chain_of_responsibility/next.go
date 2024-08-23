package chain_of_responsibility

// 充当抽象类

type Next struct {
	NextHandler PatientHandler
}

func (n *Next) SetNext(handler PatientHandler) PatientHandler {
	n.NextHandler = handler
	return handler
}

func (n *Next) Execute(patient *Patient) error {
	if n.NextHandler == nil {
		return nil
	}

	if err := n.NextHandler.Do(patient); err != nil {
		return err
	}

	return n.NextHandler.Execute(patient)
}
