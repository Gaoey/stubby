#  Golang Stubby 

easily stub services, and scenarios service for integration testing.

feature
    - stub service
    - scenario

## Stub server usage 
1. create const at `./stub/const.go`

``` 
var (
	EXAMPLE = "EXAMPLE"
)

var Routes = map[string]string{
	EXAMPLE: "/path/v1/:id",
}

var MapStub = map[string][]Stubby{
	EXAMPLE: ExampleStub,
}
```

2. create stub file in `./stub` such as example.go
```
var ExampleStub = []Stubby{
	Stubby{
		ID:          "example-1",
		Name:        EXAMPLE,
		Description: "DEPOSIT - Response Success",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: Request{
			URL:    Routes[EXAMPLE],
			Method: http.MethodGet,
			Body:   ".*",
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
			Body: "{\"test\":{\"payload\": 2}}",
		},
	}
}
```


## Scenario server usage 
1. ???
