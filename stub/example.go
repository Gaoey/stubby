package stub

import (
	"net/http"

	"github.com/labstack/echo"
)

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
			Body: "{\"accountNo\":\"4111534872\",\"productGroup\":\"SAV\",\"productType\":\"2001\",\"productDisplayName\":\"Savings\",\"accountName\":\"ทีมทดสอบ สอง ทีมทดสอบ สอง\",\"accountStatus\":\"3\",\"ownerBranch\":\"411\",\"availableBalance\":7340744.36,\"ledgerBalanceAmount\":7340744.36,\"currencyCode\":\"THB\",\"accountNickname\":\"เปลี่ยนชื่อ\",\"isMainAccount\":true,\"permissions\":[\"ACCOUNT_ACTIVE\",\"PAY_LOAN_ENABLE\",\"PAYMENT_ENABLE\",\"REQUEST_STATEMENT\",\"WITHDRAWAL_ENABLE\",\"TOPUP_ENABLE\",\"DEPOSIT_ENABLE\",\"TRANSFER_IN_ENABLE\",\"TRANSFER_OUT_ENABLE\",\"VIEW_ACCOUNT\",\"VIEW_ACCOUNT_DETAIL\",\"VIEW_DIGITAL_PASSBOOK\",\"VIEW_SCHEDULE_TRANSACTION\",\"VIEW_TRANSACTION_HISTORY\"]}",
		},
	},
	Stubby{
		ID:          "example-2",
		Name:        EXAMPLE,
		Description: "DEPOSIT - Response Error",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: Request{
			URL:    Routes[EXAMPLE],
			Method: http.MethodGet,
			Body:   ".*",
		},
		Response: Response{
			Status: 500,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
		},
	},
}
