package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll()([]Customer,error)  {
	return s.customers, nil	
}

func NewCustomeRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "MRaj", City: "Hyderabad", Zipcode: "500075", DateofBirth: "2000-01-07", Status: "1"},
		{Id: "1002", Name: "MRaj", City: "Hyderabad", Zipcode: "500075", DateofBirth: "2000-01-07", Status: "1"},
		{Id: "1003", Name: "MRaj", City: "Hyderabad", Zipcode: "500075", DateofBirth: "2000-01-07", Status: "1"},
		{Id: "1004", Name: "MRaj", City: "Hyderabad", Zipcode: "500075", DateofBirth: "2000-01-07", Status: "1"},
		{Id: "1005", Name: "MRaj", City: "Hyderabad", Zipcode: "500075", DateofBirth: "2000-01-07", Status: "1"},

	}

	return CustomerRepositoryStub{customers: customers}
}