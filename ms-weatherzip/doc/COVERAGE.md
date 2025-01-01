# MS-WEATHERZIP COVERAGE

Abaixo temos os `report` de execução e cobertura dos testes implementados no WeatherZip (SERVIÇO B)

Execução da Suíte de Testes

```plain-text
Running tests for WeatherZip...
        github.com/vs0uz4/observability/ms-weatherzip/cmd/weatherzip            coverage: 0.0% of statements
        github.com/vs0uz4/observability/ms-weatherzip/doc/swagger               coverage: 0.0% of statements
?       github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/health [no test files]
=== RUN   TestLoadConfigReadInConfigFails
--- PASS: TestLoadConfigReadInConfigFails (0.00s)
=== RUN   TestLoadConfigPanicOnUnmarshalError
    config_test.go:43: Panic successfully captured: Unmarshal error
--- PASS: TestLoadConfigPanicOnUnmarshalError (0.00s)
=== RUN   TestLoadConfigMissingRequiredConfigFails
--- PASS: TestLoadConfigMissingRequiredConfigFails (0.00s)
=== RUN   TestLoadConfig
--- PASS: TestLoadConfig (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/config    0.199s  coverage: 100.0% of statements
?       github.com/vs0uz4/observability/ms-weatherzip/internal/service/contract [no test files]
?       github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/contract [no test files]
=== RUN   TestCepResponsePopulateFromMap
=== RUN   TestCepResponsePopulateFromMap/Valid_Data
=== RUN   TestCepResponsePopulateFromMap/CEP_Not_Found
=== RUN   TestCepResponsePopulateFromMap/Invalid_CEP_Data
=== RUN   TestCepResponsePopulateFromMap/Invalid_Street_Data
=== RUN   TestCepResponsePopulateFromMap/Invalid_Neighborhood_Data
=== RUN   TestCepResponsePopulateFromMap/Invalid_Location_Data
=== RUN   TestCepResponsePopulateFromMap/Invalid_UF_Data
--- PASS: TestCepResponsePopulateFromMap (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/Valid_Data (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/CEP_Not_Found (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/Invalid_CEP_Data (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/Invalid_Street_Data (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/Invalid_Neighborhood_Data (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/Invalid_Location_Data (0.00s)
    --- PASS: TestCepResponsePopulateFromMap/Invalid_UF_Data (0.00s)
=== RUN   TestNewUnexpectedStatusCodeError
--- PASS: TestNewUnexpectedStatusCodeError (0.00s)
=== RUN   TestNewFailedToCreateRequestError
--- PASS: TestNewFailedToCreateRequestError (0.00s)
=== RUN   TestNewFailedToMakeRequestError
--- PASS: TestNewFailedToMakeRequestError (0.00s)
=== RUN   TestNewFailedToDecodeResponseError
--- PASS: TestNewFailedToDecodeResponseError (0.00s)
=== RUN   TestNewFailedToMapResponseError
--- PASS: TestNewFailedToMapResponseError (0.00s)
=== RUN   TestWeatherResponsePopulateFromMap
=== RUN   TestWeatherResponsePopulateFromMap/Valid_Data
=== RUN   TestWeatherResponsePopulateFromMap/Invalid_Location_Structure
=== RUN   TestWeatherResponsePopulateFromMap/Invalid_Current_Structure
--- PASS: TestWeatherResponsePopulateFromMap (0.00s)
    --- PASS: TestWeatherResponsePopulateFromMap/Valid_Data (0.00s)
    --- PASS: TestWeatherResponsePopulateFromMap/Invalid_Location_Structure (0.00s)
    --- PASS: TestWeatherResponsePopulateFromMap/Invalid_Current_Structure (0.00s)
PASS
coverage: 97.6% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/domain   0.348s  coverage: 97.6% of statements
=== RUN   TestHealthHandlerGetHealth
--- PASS: TestHealthHandlerGetHealth (0.00s)
=== RUN   TestHealthHandlerGetHealthErrorFromUseCase
--- PASS: TestHealthHandlerGetHealthErrorFromUseCase (0.00s)
=== RUN   TestHealthHandlerGetHealthErrorEncodingResponse
--- PASS: TestHealthHandlerGetHealthErrorEncodingResponse (0.00s)
=== RUN   TestWeatherHandler
=== RUN   TestWeatherHandler/CEP_Inválido
=== RUN   TestWeatherHandler/CEP_Não_Encontrado
=== RUN   TestWeatherHandler/Erro_no_Serviço_de_Clima
=== RUN   TestWeatherHandler/Sucesso
=== RUN   TestWeatherHandler/Erro_no_JSON_Encode
--- PASS: TestWeatherHandler (0.00s)
    --- PASS: TestWeatherHandler/CEP_Inválido (0.00s)
    --- PASS: TestWeatherHandler/CEP_Não_Encontrado (0.00s)
    --- PASS: TestWeatherHandler/Erro_no_Serviço_de_Clima (0.00s)
    --- PASS: TestWeatherHandler/Sucesso (0.00s)
    --- PASS: TestWeatherHandler/Erro_no_JSON_Encode (0.00s)
=== RUN   TestNewWeatherHandlerInitialization
--- PASS: TestNewWeatherHandlerInitialization (0.00s)
PASS
coverage: 97.8% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web        0.351s  coverage: 97.8% of statements
=== RUN   TestAddHandler
--- PASS: TestAddHandler (0.00s)
=== RUN   TestWebServerLifecycle
=== RUN   TestWebServerLifecycle/Uptime_must_not_be_nil
=== RUN   TestWebServerLifecycle/Valid_HEAD_Handler
=== RUN   TestWebServerLifecycle/Valid_GET_Handler
=== RUN   TestWebServerLifecycle/Valid_POST_Handler
=== RUN   TestWebServerLifecycle/Valid_PUT_Handler
=== RUN   TestWebServerLifecycle/Valid_PATCH_Handler
=== RUN   TestWebServerLifecycle/Valid_DELETE_Handler
--- PASS: TestWebServerLifecycle (0.51s)
    --- PASS: TestWebServerLifecycle/Uptime_must_not_be_nil (0.00s)
    --- PASS: TestWebServerLifecycle/Valid_HEAD_Handler (0.00s)
    --- PASS: TestWebServerLifecycle/Valid_GET_Handler (0.00s)
    --- PASS: TestWebServerLifecycle/Valid_POST_Handler (0.00s)
    --- PASS: TestWebServerLifecycle/Valid_PUT_Handler (0.00s)
    --- PASS: TestWebServerLifecycle/Valid_PATCH_Handler (0.00s)
    --- PASS: TestWebServerLifecycle/Valid_DELETE_Handler (0.00s)
=== RUN   TestAddHandlerWithDuplicateMethods
--- PASS: TestAddHandlerWithDuplicateMethods (0.00s)
=== RUN   TestInvalidMethods
=== RUN   TestInvalidMethods/Invalid_Method
2025/01/01 17:16:15 "INVALID /test HTTP/1.1" from [::1]:55736 - 405 0B in 4.917µs - Error: 
--- PASS: TestInvalidMethods (0.50s)
    --- PASS: TestInvalidMethods/Invalid_Method (0.00s)
=== RUN   TestWebServerStop
=== RUN   TestWebServerStop/Stop_After_Start
=== RUN   TestWebServerStop/Stop_Without_Start
--- PASS: TestWebServerStop (0.50s)
    --- PASS: TestWebServerStop/Stop_After_Start (0.50s)
    --- PASS: TestWebServerStop/Stop_Without_Start (0.00s)
=== RUN   TestWebServerErrorScenarios
=== RUN   TestWebServerErrorScenarios/Start_With_Invalid_Port
=== RUN   TestWebServerErrorScenarios/Run_Without_Start
=== RUN   TestWebServerErrorScenarios/Run_Start_Without_Error
--- PASS: TestWebServerErrorScenarios (0.00s)
    --- PASS: TestWebServerErrorScenarios/Start_With_Invalid_Port (0.00s)
    --- PASS: TestWebServerErrorScenarios/Run_Without_Start (0.00s)
    --- PASS: TestWebServerErrorScenarios/Run_Start_Without_Error (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver      1.978s  coverage: 100.0% of statements
=== RUN   TestResponseRecorderWriteHeader
=== RUN   TestResponseRecorderWriteHeader/Non-200_status_with_no_bytes_written
=== RUN   TestResponseRecorderWriteHeader/200_status_with_no_bytes_written
=== RUN   TestResponseRecorderWriteHeader/Non-200_status_with_bytes_written
=== RUN   TestResponseRecorderWriteHeader/200_status_with_bytes_written
--- PASS: TestResponseRecorderWriteHeader (0.00s)
    --- PASS: TestResponseRecorderWriteHeader/Non-200_status_with_no_bytes_written (0.00s)
    --- PASS: TestResponseRecorderWriteHeader/200_status_with_no_bytes_written (0.00s)
    --- PASS: TestResponseRecorderWriteHeader/Non-200_status_with_bytes_written (0.00s)
    --- PASS: TestResponseRecorderWriteHeader/200_status_with_bytes_written (0.00s)
=== RUN   TestResponseRecorderWrite
--- PASS: TestResponseRecorderWrite (0.00s)
=== RUN   TestResponseRecorderWriteError
--- PASS: TestResponseRecorderWriteError (0.00s)
=== RUN   TestErrorLogger
=== RUN   TestErrorLogger/Logs_Error_for_404
=== RUN   TestErrorLogger/Logs_Error_for_500
=== RUN   TestErrorLogger/Does_Not_Log_for_200
--- PASS: TestErrorLogger (0.00s)
    --- PASS: TestErrorLogger/Logs_Error_for_404 (0.00s)
    --- PASS: TestErrorLogger/Logs_Error_for_500 (0.00s)
    --- PASS: TestErrorLogger/Does_Not_Log_for_200 (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware   0.615s  coverage: 100.0% of statements
=== RUN   TestCepServiceCreateRequest
=== RUN   TestCepServiceCreateRequest/Request_Creation_Error
--- PASS: TestCepServiceCreateRequest (0.00s)
    --- PASS: TestCepServiceCreateRequest/Request_Creation_Error (0.00s)
=== RUN   TestCepServiceExecuteRequest
=== RUN   TestCepServiceExecuteRequest/Request_Execution_Error
--- PASS: TestCepServiceExecuteRequest (0.00s)
    --- PASS: TestCepServiceExecuteRequest/Request_Execution_Error (0.00s)
=== RUN   TestCepServiceUnexpectedStatusCode
=== RUN   TestCepServiceUnexpectedStatusCode/Unexpected_Status_Code
--- PASS: TestCepServiceUnexpectedStatusCode (0.00s)
    --- PASS: TestCepServiceUnexpectedStatusCode/Unexpected_Status_Code (0.00s)
=== RUN   TestCepServiceDecodeResponse
=== RUN   TestCepServiceDecodeResponse/Failed_to_Decode_Response
--- PASS: TestCepServiceDecodeResponse (0.00s)
    --- PASS: TestCepServiceDecodeResponse/Failed_to_Decode_Response (0.00s)
=== RUN   TestCepServiceDecodeResponseError
--- PASS: TestCepServiceDecodeResponseError (0.00s)
=== RUN   TestCepServicePopulateMapError
--- PASS: TestCepServicePopulateMapError (0.00s)
=== RUN   TestCepServiceGetCepData
=== RUN   TestCepServiceGetCepData/Valid_CEP
=== RUN   TestCepServiceGetCepData/CEP_Not_Found
=== RUN   TestCepServiceGetCepData/Invalid_CEP_Format
--- PASS: TestCepServiceGetCepData (0.00s)
    --- PASS: TestCepServiceGetCepData/Valid_CEP (0.00s)
    --- PASS: TestCepServiceGetCepData/CEP_Not_Found (0.00s)
    --- PASS: TestCepServiceGetCepData/Invalid_CEP_Format (0.00s)
=== RUN   TestGetCPUStats
--- PASS: TestGetCPUStats (0.00s)
=== RUN   TestGetCPUStatsError
--- PASS: TestGetCPUStatsError (0.00s)
=== RUN   TestGetMemoryStats
--- PASS: TestGetMemoryStats (0.00s)
=== RUN   TestGetMemoryStatsError
--- PASS: TestGetMemoryStatsError (0.00s)
=== RUN   TestNewUptimeService
--- PASS: TestNewUptimeService (0.00s)
=== RUN   TestGetUptime
--- PASS: TestGetUptime (0.10s)
=== RUN   TestWeatherServiceCreateRequest
=== RUN   TestWeatherServiceCreateRequest/Request_Creation_Error
--- PASS: TestWeatherServiceCreateRequest (0.00s)
    --- PASS: TestWeatherServiceCreateRequest/Request_Creation_Error (0.00s)
=== RUN   TestWeatherServiceExecuteRequest
=== RUN   TestWeatherServiceExecuteRequest/Request_Execution_Error
--- PASS: TestWeatherServiceExecuteRequest (0.00s)
    --- PASS: TestWeatherServiceExecuteRequest/Request_Execution_Error (0.00s)
=== RUN   TestWeatherServiceStatusCodeHandling
=== RUN   TestWeatherServiceStatusCodeHandling/Unexpected_Status_Code
--- PASS: TestWeatherServiceStatusCodeHandling (0.00s)
    --- PASS: TestWeatherServiceStatusCodeHandling/Unexpected_Status_Code (0.00s)
=== RUN   TestWeatherServiceDecodeResponse
=== RUN   TestWeatherServiceDecodeResponse/Failed_to_Decode_Response
--- PASS: TestWeatherServiceDecodeResponse (0.00s)
    --- PASS: TestWeatherServiceDecodeResponse/Failed_to_Decode_Response (0.00s)
=== RUN   TestWeatherServiceDecodeResponseError
--- PASS: TestWeatherServiceDecodeResponseError (0.00s)
=== RUN   TestWeatherServiceBadRequestHandling
=== RUN   TestWeatherServiceBadRequestHandling/Unexpected_Bad_Request
--- PASS: TestWeatherServiceBadRequestHandling (0.00s)
    --- PASS: TestWeatherServiceBadRequestHandling/Unexpected_Bad_Request (0.00s)
=== RUN   TestWeatherServiceUnexpectedStatusCode
=== RUN   TestWeatherServiceUnexpectedStatusCode/Unexpected_Status_Code
--- PASS: TestWeatherServiceUnexpectedStatusCode (0.00s)
    --- PASS: TestWeatherServiceUnexpectedStatusCode/Unexpected_Status_Code (0.00s)
=== RUN   TestWeatherServiceGetWeatherData
=== RUN   TestWeatherServiceGetWeatherData/Valid_Location
=== RUN   TestWeatherServiceGetWeatherData/Location_Not_Found
--- PASS: TestWeatherServiceGetWeatherData (0.00s)
    --- PASS: TestWeatherServiceGetWeatherData/Valid_Location (0.00s)
    --- PASS: TestWeatherServiceGetWeatherData/Location_Not_Found (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/service  0.883s  coverage: 100.0% of statements
=== RUN   TestMockHTTPClientDo
--- PASS: TestMockHTTPClientDo (0.00s)
=== RUN   TestMockCepService
=== RUN   TestMockCepService/Zipcode_is_Valid
=== RUN   TestMockCepService/Zipcode_Not_Found
--- PASS: TestMockCepService (0.00s)
    --- PASS: TestMockCepService/Zipcode_is_Valid (0.00s)
    --- PASS: TestMockCepService/Zipcode_Not_Found (0.00s)
=== RUN   TestMockCPUService
--- PASS: TestMockCPUService (0.00s)
=== RUN   TestMockMemoryService
--- PASS: TestMockMemoryService (0.00s)
=== RUN   TestMockWeatherService
=== RUN   TestMockWeatherService/Valid_Location
=== RUN   TestMockWeatherService/Invalid_Location
--- PASS: TestMockWeatherService (0.00s)
    --- PASS: TestMockWeatherService/Valid_Location (0.00s)
    --- PASS: TestMockWeatherService/Invalid_Location (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock     0.921s  coverage: 100.0% of statements
=== RUN   TestGetHealth
--- PASS: TestGetHealth (0.00s)
=== RUN   TestGetHealthCpuServiceError
--- PASS: TestGetHealthCpuServiceError (0.00s)
=== RUN   TestGetHealthMemoryServiceError
--- PASS: TestGetHealthMemoryServiceError (0.00s)
=== RUN   TestGetHealthMessageWhenNotPass
--- PASS: TestGetHealthMessageWhenNotPass (0.00s)
=== RUN   TestIsNumeric
--- PASS: TestIsNumeric (0.00s)
=== RUN   TestNewWeatherByCepUsecase
--- PASS: TestNewWeatherByCepUsecase (0.00s)
=== RUN   TestGetWeatherByCep
=== RUN   TestGetWeatherByCep/Invalid_CEP
=== RUN   TestGetWeatherByCep/CEP_Not_Found
=== RUN   TestGetWeatherByCep/Weather_Service_Error
=== RUN   TestGetWeatherByCep/Success
--- PASS: TestGetWeatherByCep (0.00s)
    --- PASS: TestGetWeatherByCep/Invalid_CEP (0.00s)
    --- PASS: TestGetWeatherByCep/CEP_Not_Found (0.00s)
    --- PASS: TestGetWeatherByCep/Weather_Service_Error (0.00s)
    --- PASS: TestGetWeatherByCep/Success (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/usecase  1.009s  coverage: 100.0% of statements
=== RUN   TestMockHealthCheckUseCase
--- PASS: TestMockHealthCheckUseCase (0.00s)
=== RUN   TestMockWeatherByCepUsecase
=== RUN   TestMockWeatherByCepUsecase/Success
=== RUN   TestMockWeatherByCepUsecase/Failure
--- PASS: TestMockWeatherByCepUsecase (0.00s)
    --- PASS: TestMockWeatherByCepUsecase/Success (0.00s)
    --- PASS: TestMockWeatherByCepUsecase/Failure (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/mock     1.153s  coverage: 100.0% of statements
```

Cobertura dos Testes

```plain-text
github.com/vs0uz4/observability/ms-weatherzip/cmd/weatherzip/main.go:32:                                        main                            0.0%
github.com/vs0uz4/observability/ms-weatherzip/cmd/weatherzip/main.go:79:                                        enjoy                           0.0%
github.com/vs0uz4/observability/ms-weatherzip/config/config.go:18:                                              LoadConfig                      100.0%
github.com/vs0uz4/observability/ms-weatherzip/doc/swagger/docs.go:221:                                          init                            0.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/cep.go:13:                                        PopulateFromMap                 100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/errors.go:28:                                     NewUnexpectedStatusCodeError    100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/errors.go:32:                                     NewInvalidZipCodeDetailsError   0.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/errors.go:36:                                     NewFailedToCreateRequestError   100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/errors.go:40:                                     NewFailedToMakeRequestError     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/errors.go:44:                                     NewFailedToDecodeResponseError  100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/errors.go:48:                                     NewFailedToMapResponseError     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/domain/weather.go:32:                                    PopulateFromMap                 100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/health_handler.go:14:                          NewHealthHandler                100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/health_handler.go:26:                          GetHealth                       100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/weather_handler.go:30:                         NewWeatherHandler               100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/weather_handler.go:45:                         GetWeatherByCep                 97.1%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware/error_logger.go:16:       WriteHeader                     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware/error_logger.go:24:       Write                           100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware/error_logger.go:30:       ReadError                       100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware/error_logger.go:34:       WriteError                      100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware/error_logger.go:38:       ErrorLogger                     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/server.go:24:                        NewWebServer                    100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/server.go:45:                        setupDependencies               100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/server.go:49:                        AddHandler                      100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/server.go:57:                        Start                           100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/server.go:80:                        Run                             100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/server.go:90:                        Stop                            100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/cep.go:23:                                       NewCepService                   100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/cep.go:30:                                       GetLocation                     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/cpu_stats.go:19:                                 NewCPUService                   100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/cpu_stats.go:25:                                 roundToOneDecimal               100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/cpu_stats.go:29:                                 GetCPUStats                     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/memory_stats.go:17:                              NewMemoryService                100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/memory_stats.go:23:                              roundToOneDecimal               100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/memory_stats.go:27:                              GetMemoryStats                  100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock/http_client.go:11:                          Do                              100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock/mock_cep.go:13:                             GetLocation                     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock/mock_cpu_stats.go:7:                        GetCPUStats                     100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock/mock_memory_stats.go:7:                     GetMemoryStats                  100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock/mock_weather.go:13:                         GetWeather                      100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/uptime.go:13:                                    NewUptimeService                100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/uptime.go:19:                                    GetUptime                       100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/weather.go:35:                                   NewWeatherService               100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/service/weather.go:44:                                   GetWeather                      100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/health_check.go:20:                              NewHealthCheckUseCase           100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/health_check.go:24:                              GetHealth                       100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/mock/mock_health_check.go:9:                     GetHealth                       100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/mock/mock_weather_by_cep.go:13:                  GetWeatherByCep                 100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/weather_by_cep.go:17:                            NewWeatherByCepUsecase          100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/weather_by_cep.go:24:                            GetWeatherByCep                 100.0%
github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/weather_by_cep.go:50:                            isNumeric                       100.0%
total:                                                                                                          (statements)                    89.9%
```