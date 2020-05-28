package route

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/Gaoey/stubby/scenario"
	"github.com/Gaoey/stubby/stub"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ResponseError struct {
	Message string `json:"message"`
}

func Controller(c echo.Context, stubName string, stub []stub.Stubby) error {
	logrus.Printf("%s - %s", c.Request().Method, c.Request().URL)

	var (
		httpMethodHasRequest = c.Request().Method == http.MethodPost || c.Request().Method == http.MethodPut || c.Request().Method == http.MethodPatch
		request              = c.Request().Body
		method               = c.Request().Method
		isScenario           = viper.GetBool("env.scenario")
	)

	// transform io.Reader from request to byte
	var bodyBytes []byte
	if request != nil {
		bodyBytes, _ = ioutil.ReadAll(request)
	}

	if isScenario {
		id := viper.GetString("scenario-id")
		sce := scenario.New(id, stub)
		s := sce.GetStubByName(stubName)
		var (
			status            = s.Response.Status
			response          = s.Response.Body
			headerContentType = s.Response.Header[echo.HeaderContentType]
		)
		return ResponseFormatter(c, headerContentType, status, response)
	}

	for _, s := range stub {

		if method != s.Request.Method {
			continue
		}

		if !s.Validate(c) {
			continue
		}

		var (
			status            = s.Response.Status
			response          = s.Response.Body
			headerContentType = s.Response.Header[echo.HeaderContentType]
		)

		if httpMethodHasRequest {
			// validate request
			matchBody, err := regexp.MatchString(s.Request.Body, string(bodyBytes))
			if err != nil {
				continue
			}

			if matchBody {
				return ResponseFormatter(c, headerContentType, status, response)
			}
		}

		if !httpMethodHasRequest {
			return ResponseFormatter(c, headerContentType, status, response)
		}
	}

	logrus.Println("no stubby to define")
	return c.JSON(http.StatusInternalServerError, ResponseError{Message: "holy shit!!! no stubby to define"})
}

func ResponseFormatter(c echo.Context, contentType string, status int, response string) error {
	switch contentType {
	case echo.MIMEApplicationXML:
		return c.XMLBlob(status, []byte(response))
	case echo.MIMEApplicationJSON:
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(response), &result); err != nil {
			log.Fatal(err)
		}
		return c.JSON(status, result)
	default:
		return c.String(status, response)
	}
}
