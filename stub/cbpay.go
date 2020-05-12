package stub

import (
	"net/http"

	"github.com/labstack/echo"
)

var CBPayStub = []Stubby{
	Stubby{
		ID:          GenerateID(),
		Name:        CBPAY,
		Description: "CBPAY Test 2",
		Request: Request{
			URL:    Routes[CBPAY],
			Method: http.MethodGet,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: "<xml>Test 2</xml>",
		},
	},
	Stubby{
		ID:          GenerateID(),
		Name:        CBPAY,
		Description: "CBPAY Test 1",
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
			Body: "<xml>Test</xml>",
		},
	},
}
