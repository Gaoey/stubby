package stub

var (
	EXAMPLE = "EXAMPLE"
)

var Routes = map[string]string{
	EXAMPLE: "/path/v1/:id",
}

var MapStub = map[string][]Stubby{
	EXAMPLE: ExampleStub,
}
