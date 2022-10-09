package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{Id: "1000", Name: "Eduardo", City: "Santiago", Zipcode: "100000", DateOfBirth: "1988-12-24", Status: "1"},
		{Id: "1001", Name: "Rob", City: "Santiago", Zipcode: "100000", DateOfBirth: "2000-12-24", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
