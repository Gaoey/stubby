package main

import (
	"github.com/Gaoey/stubby/route"
	"github.com/Gaoey/stubby/stub"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	Routing(e, stub.CBPayStub)
	Routing(e, stub.CBSStub)

	a := e.Routes()
	for _, r := range a {
		logrus.Printf("------")
		logrus.Printf("%#v", r.Name)
		logrus.Printf("%#v", r.Path)
		logrus.Printf("%#v", r.Method)
	}
	e.Logger.Fatal(e.Start(":1323"))
}

func Routing(e *echo.Echo, stubs []stub.Stubby) error {
	for _, v := range stubs {
		e.Add(v.Request.Method, stub.Routes[v.Name], func(c echo.Context) error {
			return route.Controller(c, v)
		})
	}
	return nil
}
