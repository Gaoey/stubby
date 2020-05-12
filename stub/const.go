package stub

var (
	CBPAY = "CBPAY"
	CBS   = "CBS"
)

var Routes = map[string]string{
	CBPAY: "/billpresentment-ws/BillPresentment",
	CBS:   "/cbs",
}

var MapStub = map[string][]Stubby{
	CBPAY: CBPayStub,
	CBS:   CBSStub,
}
