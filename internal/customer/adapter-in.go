package customer

type adapterIn struct {
	portOut PortOut
}

func NewAdapterIn(portOut PortOut) PortIn {
	return &adapterIn{
		portOut,
	}
}

func (adapter *adapterIn) GetAll() ([]*Customer, error) {
	return adapter.portOut.GetAllDB()
}

func (adapter *adapterIn) Create(customer Customer) error {
	return adapter.portOut.CreateDB(customer)
}
