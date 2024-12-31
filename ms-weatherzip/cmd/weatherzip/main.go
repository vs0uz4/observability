package main

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/vs0uz4/observability/ms-weatherzip/config"
	_ "github.com/vs0uz4/observability/ms-weatherzip/doc/swagger"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/service"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/usecase"
)

// @title Weather Zip API
// @version 1.0
// @description This is a simple API that, on receiving a (zip code), performs a query for the weather in that location.

// @contact.name Vitor Rodrigues
// @contact.url https://vsouza.rio.br/contact
// @contact.email me@vsouza.rio.br

// @license.name MIT
// @license.url https://opensource.org/license/MIT

// @host localhost:8001
// @BasePath /
func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{}

	cpuService := service.NewCPUService()
	memoryService := service.NewMemoryService()
	uptimeService := service.NewUptimeService()
	cepService := service.NewCepService(httpClient, cfg.CepAPIUrl)
	weatherService := service.NewWeatherService(httpClient, cfg.WeatherAPIUrl, cfg.WeatherAPIKey, cfg.WeatherAPILanguage)

	healthCheckUseCase := usecase.NewHealthCheckUseCase(cpuService, memoryService, uptimeService)
	wheaterByCepUseCase := usecase.NewWeatherByCepUsecase(cepService, weatherService)

	handlerHealth := web.NewHealthHandler(healthCheckUseCase).GetHealth
	handlerWeather := web.NewWeatherHandler(wheaterByCepUseCase).GetWeatherByCep

	webserver := webserver.NewWebServer(cfg.WebServerPort)
	webserver.AddHandler("/weather/{cep}", handlerWeather, "GET")
	webserver.AddHandler("/health", handlerHealth, "GET")
	webserver.AddHandler("/", enjoy, "GET")
	webserver.AddHandler("/swagger/*", httpSwagger.WrapHandler, "GET")

	fmt.Println("Starting web server on port", cfg.WebServerPort)
	webserver.Start()
	webserver.Run()
}

// HandlerRoot é o handler que exibe uma saudação
// @Summary Exibe uma saudação
// @Description Exibe a saudação "Enjoy the silence!"
// @Tags Root
// @Produce  plain
// @Success 200 {string} string "Enjoy the silence!"
// @Failure 500 {string} string "Unable to write response"
// @Router / [get]
func enjoy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Enjoy the silence!")); err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}
