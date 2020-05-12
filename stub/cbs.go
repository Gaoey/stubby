package stub

import (
	"net/http"

	"github.com/labstack/echo"
)

var CBSStub = []Stubby{
	Stubby{
		ID:          GenerateID(),
		Name:        CBS,
		Description: "CBS Test 1",
		Request: Request{
			URL:    Routes[CBPAY],
			Method: http.MethodPost,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: ".*june.*",
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: "<xml>CBS Test</xml>",
		},
	},
	Stubby{
		ID:          GenerateID(),
		Name:        CBS,
		Description: "CBS Test 2",
		Request: Request{
			URL:    Routes[CBPAY],
			Method: http.MethodGet,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: ".*nom.*",
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: "<xml>CBS Test2</xml>",
		},
	},
}
