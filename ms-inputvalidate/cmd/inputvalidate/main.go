package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/vs0uz4/observability/ms-inputvalidate/config"
	_ "github.com/vs0uz4/observability/ms-inputvalidate/doc/swagger"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/service"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase"
	"github.com/vs0uz4/observability/pkg/tracing"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// @title Input Validation API
// @version 1.0
// @description This is a simple API for validating input data (zip code), forwarding it to the weather service and returning the weather for a given location.

// @contact.name Vitor Rodrigues
// @contact.url https://vsouza.rio.br/contact
// @contact.email me@vsouza.rio.br

// @license.name MIT
// @license.url https://opensource.org/license/MIT

// @host localhost:8000
// @BasePath /
func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	traceProvider, err := tracing.InitTracer("ms-inputvalidate", cfg.ZipKinUrl)
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer func() { _ = traceProvider.Shutdown(context.Background()) }()

	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	cpuService := service.NewCPUService()
	memoryService := service.NewMemoryService()
	uptimeService := service.NewUptimeService()
	weatherService := service.NewWeatherService(httpClient, cfg.WeatherZipAPIUrl)

	healthCheckUseCase := usecase.NewHealthCheckUseCase(cpuService, memoryService, uptimeService)
	weatherLocationByCepUseCase := usecase.NewWeatherLocationByCepUsecase(weatherService)

	handlerHealth := web.NewHealthHandler(healthCheckUseCase).GetHealth
	handlerInput := web.NewInputHandler(weatherLocationByCepUseCase).GetLocationWeatherByCep

	webserver := webserver.NewWebServer(cfg.WebServerPort)
	webserver.AddHandler("/health", handlerHealth, "GET")
	webserver.AddHandler("/", handlerInput, "POST")
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
