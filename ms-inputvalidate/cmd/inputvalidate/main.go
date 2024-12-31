package main

import (
	"fmt"
	"net/http"

	"github.com/vs0uz4/observability/ms-inputvalidate/config"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/service"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{}

	cpuService := service.NewCPUService()
	memoryService := service.NewMemoryService()
	uptimeService := service.NewUptimeService()
	weatherService := service.NewWeatherService(httpClient, cfg.WeatherZipAPIUrl)

	healthCheckUseCase := usecase.NewHealthCheckUseCase(cpuService, memoryService, uptimeService)
	weatherLocationByCepUseCase := usecase.NewWeatherLocationByCepUsecase(weatherService)

	handlerRoot := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("Enjoy the silence!")); err != nil {
			http.Error(w, "Unable to write response", http.StatusInternalServerError)
		}
	}

	handlerHealth := web.NewHealthHandler(healthCheckUseCase).GetHealth
	handlerInput := web.NewInputHandler(weatherLocationByCepUseCase).GetLocationWeatherByCep

	webserver := webserver.NewWebServer(cfg.WebServerPort)
	webserver.AddHandler("/health", handlerHealth, "GET")
	webserver.AddHandler("/", handlerInput, "POST")
	webserver.AddHandler("/", handlerRoot, "GET")

	fmt.Println("Starting web server on port", cfg.WebServerPort)
	webserver.Start()
	webserver.Run()
}
