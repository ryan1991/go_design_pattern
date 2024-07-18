package abstract_factory

import "testing"

func TestRdbFactory(t *testing.T) {

	var factory = &RDBDAOFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()

}
