package main

import (
	"engine"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	e := echo.New()
	e.GET("/", hello)
	e.GET("/search = <:request>", search)

	e.Start(":" + port)

}

func search(c echo.Context) error {
	var urlx string = c.Request().URL.String()
	var request []string = strings.Split(urlx, "/")
	var data = request[len(request)-1]
	parsed, err := url.QueryUnescape(data)
	if err != nil {
	}
	var input string = strings.Trim(parsed, "search = ")
	input = strings.Trim(input, "<>")

	var response []string = Engine.Perform_search(input)
	return c.JSON(http.StatusOK, response)

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hey World!")
}

//sd asd
