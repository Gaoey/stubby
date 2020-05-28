package route

import (
	"net/http"

	"github.com/Gaoey/stubby/stub"
	"github.com/labstack/echo"
)

type MockData struct {
	Expected []Expected
}

type Expected struct {
	Request  stub.Request
	Receive  string
	ShouldBe stub.Response
}

var MockHttpTest = MockData{
	Expected: []Expected{
		{
			Request:  MockStub[0].Request,
			Receive:  "{\"message\":\"xxxThisIsShouldBeTest1xxx\"}",
			ShouldBe: MockStub[0].Response,
		},
		{
			Request:  MockStub[1].Request,
			Receive:  "{\"message\":\"test2\", \"payload\":{\"data\":\"ThisIsShouldBeTest2\"}}",
			ShouldBe: MockStub[1].Response,
		},
		{
			Request:  MockStub[2].Request,
			Receive:  "{\"message\":\"test2\", \"payload\":{\"data\":\"ThisIsShouldBeTest2\"}}",
			ShouldBe: MockStub[2].Response,
		},
	},
}

var MockStub = []stub.Stubby{
	{
		ID:          "id-test-1",
		Name:        "Test1",
		Description: "this is for test 1",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: stub.Request{
			URL:    "/http-test",
			Method: http.MethodPost,
			Body:   ".*ThisIsShouldBeTest1.*",
		},
		Response: stub.Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
			Body: "{\"response\":\"test1\"}",
		},
	},
	{
		ID:          "id-test-2",
		Name:        "Test2",
		Description: "this is for test 2",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: stub.Request{
			URL:    "/http-test",
			Method: http.MethodPost,
			Body:   ".*ThisIsShouldBeTest2.*",
		},
		Response: stub.Response{
			Status: 200,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
			Body: "{\"response\":\"test2\"}",
		},
	},
	{
		ID:          "id-test-3",
		Name:        "Test3",
		Description: "test for error case",
		Validate: func(c echo.Context) bool {
			return true
		},
		Request: stub.Request{
			URL:    "/http-test",
			Method: http.MethodPost,
			Body:   ".*",
		},
		Response: stub.Response{
			Status: 500,
			Header: map[string]string{
				echo.HeaderContentType: echo.MIMEApplicationJSON,
			},
			Body: "{}",
		},
	},
}
