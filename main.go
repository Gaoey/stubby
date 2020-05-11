package main

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/Gaoey/stubby/stub"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	Routing(e, stub.CBPayStub)

	e.Logger.Fatal(e.Start(":1323"))
}

func Routing(e *echo.Echo, stubs []stub.Stubby) error {
	for _, v := range stubs {
		switch v.Request.Method {
		case http.MethodPost:
			e.POST(stub.Route["CBPAY"], func(c echo.Context) error {
				return Controller(c, v)
			})
		}
	}

	return nil
}

func Controller(c echo.Context, v stub.Stubby) error {
	// make body as byte to convert as string
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	matchBody, err := regexp.MatchString(v.Request.Body, string(bodyBytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if matchBody {
		return c.String(v.Response.Status, v.Response.Body)
	}

	logrus.Printf("no stubby to define")
	return c.JSON(http.StatusInternalServerError, ResponseError{Message: "no stubby for this"})
}
