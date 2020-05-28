package stub

import (
	"net/http"

	"github.com/labstack/echo"
)

var AnyIdStub = []Stubby{
	{
		ID:          "anyid1",
		Name:        ANYID,
		Description: "Any ID Verify - Response Success",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: Request{
			URL:    Routes[ANYID],
			Method: http.MethodPost,
			Body:   ".*lookupActualAcctReq.*",
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
			Body: "{\"headerResp\":{\"reqID\":\"1589350552\",\"reqDtm\":\"2020-05-13 13:15:52.431\",\"txnRefID\":\"2020051312411446\",\"service\":\"LookupActualAcct\",\"statusCd\":\"00000\"},\"lookupActualAcctResp\":{\"toDisplayName\":\"บัญชีทดสอบ\",\"toAcctName\":\"TEST111300089\",\"toAcctID\":\"1113000898\",\"toBankCd\":\"014\",\"switchFeeAmt\":0.00,\"toBankFeeAmt\":3.00,\"customerFeeAmt\":0.00,\"retrievalRefNo\":\"013413486964\",\"availBalAmt\":7340744.36,\"fromAcctID\":\"4111534872\",\"fromAcctType\":\"2001\",\"amt\":1.00,\"ledgerAmt\":7340744.36}}",
		},
	},
	{
		ID:          "anyid2",
		Name:        ANYID,
		Description: "Any ID Confirm - Response Success",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: Request{
			URL:    Routes[ANYID],
			Method: http.MethodPost,
			Body:   ".*transferActualAcctReq.*",
		},
		Response: Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
			Body: "{\"headerResp\":{\"reqID\":\"1589350556\",\"reqDtm\":\"2020-05-13 13:15:56.814\",\"txnRefID\":\"2020051312411447\",\"service\":\"TransferActualAcct\",\"statusCd\":\"00000\"},\"transferActualAcctResp\":{\"financialTxnRef\":\"22020051301695680\",\"availBalAmt\":7340743.36,\"ledgerAmt\":7340743.36}}",
		},
	},
	{
		ID:          "anyid3",
		Name:        ANYID,
		Description: "Any ID Confirm - Response Failure",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: Request{
			URL:    Routes[ANYID],
			Method: http.MethodPost,
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
