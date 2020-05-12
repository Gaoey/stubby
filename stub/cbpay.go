package stub

import (
	"net/http"

	"github.com/labstack/echo"
)

var CBPayStub = []Stubby{
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
			Body: "<xml>Test june</xml>",
		},
	},
	Stubby{
		ID:          GenerateID(),
		Name:        CBPAY,
		Description: "CBPAY Test 2",
		Request: Request{
			URL:    Routes[CBPAY],
			Method: http.MethodPost,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: ".*happy.*",
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: "<xml>Test happy </xml>",
		},
	},
	Stubby{
		ID:          GenerateID(),
		Name:        CBPAY,
		Description: "CBPAY Test 3",
		Request: Request{
			URL:    Routes[CBPAY],
			Method: http.MethodPost,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: ".*error.*",
		},
		Response: Response{
			Status: 408,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: "client time out",
		},
	},
}
