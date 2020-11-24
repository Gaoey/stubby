package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Gaoey/stubby/route"
	"github.com/Gaoey/stubby/stub"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initViper()
	e := echo.New()

	for _, eachStub := range stub.MapStub {
		Routing(e, eachStub)
	}

	a := e.Routes()
	fmt.Println("Provided path")
	for _, r := range a {
		logrus.Printf("%s - %s", r.Method, r.Path)
	}

	port := viper.GetString("app.port")
	e.Logger.Fatal(e.Start(":" + port))
}

func Routing(e *echo.Echo, stubs []stub.Stubby) error {
	for _, v := range stubs {
		e.Add(v.Request.Method, stub.Routes[v.Name], func(c echo.Context) error {
			stub := stub.MapStub[v.Name]
			return route.Controller(c, v.Name, stub)
		})
	}
	return nil
}

func initViper() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
