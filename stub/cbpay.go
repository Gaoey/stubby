package stub

import (
	"net/http"

	"github.com/labstack/echo"
)

var CBPayStub = []Stubby{
	Stubby{
		Name:        CBPAY,
		Description: "Test 1",
		Request: Request{
			URL:    Route[CBPAY],
			Method: http.MethodPost,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: ".*",
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
