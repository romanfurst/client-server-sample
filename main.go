package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

const PORT = "8080"

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func main() {

	echoServer := echo.New()
	echoServer.HideBanner = true

	logger.Info().Msgf("Start listening on port %s", PORT)

	echoServer.GET("/server-client", callAndReturnResponse)
	echoServer.GET("/server", returnResponse)

	go randomErrLog()
	go randomlog()

	err := echoServer.Start(":" + PORT)
	if err != nil {
		echoServer.Logger.Fatal(err)
	}

}

func callAndReturnResponse(c echo.Context) error {

	logger.Info().Msgf("request on path /server-client received with headers: %v", c.Request().Header)

	copyHeaders := make(map[string]string)
	for key, values := range c.Request().Header {
		for _, value := range values {
			copyHeaders[key] = value
		}
	}

	client := resty.New()
	logger.Info().Msgf("calling %s service...", os.Getenv("client_target"))
	resp, err := client.R().
		SetHeaders(copyHeaders).
		SetBody(fmt.Sprintf("{\"test\":\"this is request payload\",\"hostame\":\"%s\",\"foo\":\"bar\"}", os.Getenv("HOSTNAME"))).
		Get(os.Getenv("client_target"))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	logger.Info().Msgf("returning response from client call")

	return c.String(resp.StatusCode(), string(resp.Body()))

}

func returnResponse(c echo.Context) error {

	logger.Info().Msgf("request on path /server received with headers: %v", c.Request().Header)
	logger.Info().Msg("returning static response...")

	return c.String(http.StatusOK, fmt.Sprintf("{\"test\":\"this is response payload\",\"hostame\":\"%s\",\"foo\":\"bar\"}", os.Getenv("HOSTNAME")))
}

func randomlog() {

	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		logger.Info().Msgf("Random log from %s service", os.Getenv("HOSTNAME"))
	}
}

func randomErrLog() {
	errStreamLog := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		errStreamLog.Error().Msgf("Just random error log from %s service", os.Getenv("HOSTNAME"))
	}
}
