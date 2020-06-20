package category

type adapterIn struct {
	portOut PortOut
}

func NewAdapterIn(portOut PortOut) PortIn {
	return &adapterIn{
		portOut,
	}
}

func (adapter *adapterIn) GetAll() ([]*Category, error) {
	return adapter.portOut.GetAllDB()
}
