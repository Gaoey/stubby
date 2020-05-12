package route

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

func Controller(c echo.Context, v stub.Stubby) error {
	logrus.Printf("%s - %s", c.Request().Method, c.Request().URL)

	var (
		httpMethodHasRequest = c.Request().Method == http.MethodPost || c.Request().Method == http.MethodPut || c.Request().Method == http.MethodPatch
	)

	stub := stub.MapStub[v.Name]
	for _, s := range stub {
		// validate method
		if c.Request().Method != s.Request.Method {
			continue
		}

		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		}

		logrus.Printf("doing %s ", s.Description)

		var (
			status   = s.Response.Status
			response = s.Response.Body
		)
		if httpMethodHasRequest {
			// validate request
			matchBody, err := regexp.MatchString(v.Request.Body, string(bodyBytes))
			if err != nil {
				continue
			}

			if matchBody {
				return c.String(status, response)
			}
		}

		return c.String(status, response)
	}

	logrus.Printf("no stubby to define")
	return c.JSON(http.StatusInternalServerError, ResponseError{Message: "no stubby for this"})
}
