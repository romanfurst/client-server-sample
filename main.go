package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

const PORT = "8080"

func main() {

	echoServer := echo.New()
	echoServer.HideBanner = true

	echoServer.Logger.Infof("Start listening on port %s", PORT)

	echoServer.GET("/server-client", callAndReturnResponse)
	echoServer.GET("/server", returnResponse)

	err := echoServer.Start(":" + PORT)
	if err != nil {
		echoServer.Logger.Fatal(err)
	}

}

func callAndReturnResponse(c echo.Context) error {

	copyHeaders := make(map[string]string)
	for key, values := range c.Request().Header {
		for _, value := range values {
			copyHeaders[key] = value
		}
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(copyHeaders).
		SetBody("{\"message\":\"this is request payload\"}").
		Get(os.Getenv("client_target"))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(resp.StatusCode(), string(resp.Body()))

}

func returnResponse(c echo.Context) error {
	return c.String(http.StatusOK, "{\"message\":\"this is response payload\"}")
}
