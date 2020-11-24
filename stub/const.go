package stub

var (
	EXAMPLE                         = "EXAMPLE"
	TESCO_CUSTOMER_PLAFORM_REGISTER = "TESCO_CUSTOMER_PLAFORM_REGISTER"
)

var Routes = map[string]string{
	EXAMPLE:                         "/path/v1/:id",
	TESCO_CUSTOMER_PLAFORM_REGISTER: "/v1/customers/registration",
}

var MapStub = map[string][]Stubby{
	EXAMPLE:                         ExampleStub,
	TESCO_CUSTOMER_PLAFORM_REGISTER: RegistrationStub,
}
