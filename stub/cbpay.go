package stub

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var CBPayStub = []Stubby{
	Stubby{
		Description: "Test 1",
		Request: Request{
			URL:    "/billpresentment-ws/BillPresentment",
			Method: http.MethodPost,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationXML,
			},
			Body: "*",
		},
	},
}
