package route

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestControllerCase(t *testing.T) {
	e := echo.New()
	for _, data := range MockHttpTest.Expected {
		req := httptest.NewRequest(data.Request.Method, data.Request.URL, strings.NewReader(data.Request.Body))
		for key, value := range data.Request.Header {
			req.Header.Set(key, value)
		}
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		method := Controller(c, "any", MockStub)
		if assert.NoError(t, method) {
			assert.Equal(t, data.ShouldBe.Status, rec.Code)
			assert.Equal(t, data.ShouldBe.Body+"\n", rec.Body.String())
		}
	}
}
