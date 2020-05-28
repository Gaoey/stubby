package stub

var (
	ANYID   = "ANYID"
	DEPOSIT = "DEPOSIT"
)

var Routes = map[string]string{
	ANYID:   "/anyid/AnyIDESBActualAcct/service",
	DEPOSIT: "/deposit/v1/:id",
}

var MapStub = map[string][]Stubby{
	ANYID:   AnyIdStub,
	DEPOSIT: DepositStub,
}
