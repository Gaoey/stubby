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
		stub                 = stub.MapStub[v.Name]
		request              = c.Request().Body
		method               = c.Request().Method
	)

	// transform io.Reader from request to byte
	var bodyBytes []byte
	if request != nil {
		bodyBytes, _ = ioutil.ReadAll(request)
	}

	for _, s := range stub {
		// validate method
		if method != s.Request.Method {
			continue
		}

		logrus.Printf("doing %s ", s.Description)

		var (
			status   = s.Response.Status
			response = s.Response.Body
		)

		if httpMethodHasRequest {
			// validate request
			matchBody, err := regexp.MatchString(s.Request.Body, string(bodyBytes))
			if err != nil {
				continue
			}

			if matchBody {
				return c.String(status, response)
			}
		}

		if !httpMethodHasRequest {
			return c.String(status, response)
		}
	}

	logrus.Printf("no stubby to define")
	return c.JSON(http.StatusInternalServerError, ResponseError{Message: "no stubby for this"})
}
