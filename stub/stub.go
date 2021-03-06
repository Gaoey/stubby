package stub

import "github.com/labstack/echo"

type Stubby struct {
	ID          string
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Validate    func(c echo.Context) bool `json:"middleware"`
	Request     Request                   `json:"request"`
	Response    Response                  `json:"response"`
}

type Request struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Query  map[string]string `json:"query"`
	Header map[string]string `json:"header"`
	Body   string            `json:"body"`
}

type Response struct {
	Status int               `json:"status"`
	Header map[string]string `json:"header"`
	Body   string            `json:"body"`
}
