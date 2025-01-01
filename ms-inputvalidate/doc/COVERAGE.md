# MS-INPUTVALIDATE COVERAGE

Abaixo temos o `report` da cobertura dos testes implementados no InputValidate (SERVIÇO A)

```plain-text
Running tests for InputValidate...
        github.com/vs0uz4/observability/ms-inputvalidate/cmd/inputvalidate              coverage: 0.0% of statements
        github.com/vs0uz4/observability/ms-inputvalidate/doc/swagger            coverage: 0.0% of statements
        github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/utils           coverage: 0.0% of statements
?       github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/health      [no test files]
?       github.com/vs0uz4/observability/ms-inputvalidate/internal/service/contract      [no test files]
?       github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/contract      [no test files]
=== RUN   TestLoadConfigReadInConfigFails
--- PASS: TestLoadConfigReadInConfigFails (0.00s)
=== RUN   TestLoadConfigPanicOnUnmarshalError
    config_test.go:41: Panic successfully captured: Unmarshal error
--- PASS: TestLoadConfigPanicOnUnmarshalError (0.00s)
=== RUN   TestLoadConfigMissingRequiredConfigFails
--- PASS: TestLoadConfigMissingRequiredConfigFails (0.00s)
=== RUN   TestLoadConfig
--- PASS: TestLoadConfig (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/config 0.468s  coverage: 100.0% of statements
=== RUN   TestNewUnexpectedStatusCodeError
--- PASS: TestNewUnexpectedStatusCodeError (0.00s)
=== RUN   TestNewUnexpectedWeatherServiceError
--- PASS: TestNewUnexpectedWeatherServiceError (0.00s)
=== RUN   TestNewFailedToCreateRequestError
--- PASS: TestNewFailedToCreateRequestError (0.00s)
=== RUN   TestNewFailedToMakeRequestError
--- PASS: TestNewFailedToMakeRequestError (0.00s)
=== RUN   TestNewFailedToDecodeResponseError
--- PASS: TestNewFailedToDecodeResponseError (0.00s)
=== RUN   TestNewInvalidZipCodeDetailsError
--- PASS: TestNewInvalidZipCodeDetailsError (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/domain        0.628s  coverage: 100.0% of statements
=== RUN   TestHealthHandlerGetHealth
--- PASS: TestHealthHandlerGetHealth (0.00s)
=== RUN   TestHealthHandlerGetHealthErrorFromUseCase
--- PASS: TestHealthHandlerGetHealthErrorFromUseCase (0.00s)
=== RUN   TestHealthHandlerGetHealthErrorEncodingResponse
--- PASS: TestHealthHandlerGetHealthErrorEncodingResponse (0.00s)
=== RUN   TestGetLocationWeatherByCep
=== RUN   TestGetLocationWeatherByCep/Erro_no_JSON_Encode
=== RUN   TestGetLocationWeatherByCep/CEP_Inválido
=== RUN   TestGetLocationWeatherByCep/CEP_Não_Encontrado
=== RUN   TestGetLocationWeatherByCep/Erro_no_Serviço_de_CEP
=== RUN   TestGetLocationWeatherByCep/Erro_no_Serviço_de_Clima
=== RUN   TestGetLocationWeatherByCep/Sucesso
--- PASS: TestGetLocationWeatherByCep (0.00s)
    --- PASS: TestGetLocationWeatherByCep/Erro_no_JSON_Encode (0.00s)
    --- PASS: TestGetLocationWeatherByCep/CEP_Inválido (0.00s)
    --- PASS: TestGetLocationWeatherByCep/CEP_Não_Encontrado (0.00s)
    --- PASS: TestGetLocationWeatherByCep/Erro_no_Serviço_de_CEP (0.00s)
    --- PASS: TestGetLocationWeatherByCep/Erro_no_Serviço_de_Clima (0.00s)
    --- PASS: TestGetLocationWeatherByCep/Sucesso (0.00s)
=== RUN   TestNewInputHandlerInitialization
--- PASS: TestNewInputHandlerInitialization (0.00s)
=== RUN   TestGetLocationWeatherByCepInvalidJSON
--- PASS: TestGetLocationWeatherByCepInvalidJSON (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web     0.641s  coverage: 100.0% of statements
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
2025/01/01 17:16:12 "INVALID /test HTTP/1.1" from [::1]:55719 - 405 0B in 1.5µs - Error: 
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
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver   2.271s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware        0.909s  coverage: 100.0% of statements
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
=== RUN   TestWeatherServiceDecodeResponse
=== RUN   TestWeatherServiceDecodeResponse/Failed_to_Decode_Response
--- PASS: TestWeatherServiceDecodeResponse (0.00s)
    --- PASS: TestWeatherServiceDecodeResponse/Failed_to_Decode_Response (0.00s)
=== RUN   TestWeatherServiceDecodeResponseError
--- PASS: TestWeatherServiceDecodeResponseError (0.00s)
=== RUN   TestWeatherServiceReadResponseBodyError
--- PASS: TestWeatherServiceReadResponseBodyError (0.00s)
=== RUN   TestWeatherServiceUnexpectedErrors
=== RUN   TestWeatherServiceUnexpectedErrors/Unexpected_Internal_Server_Error
=== RUN   TestWeatherServiceUnexpectedErrors/Unexpected_Unprocessable_Entity_Error
=== RUN   TestWeatherServiceUnexpectedErrors/Unexpected_Not_Found_Error
=== RUN   TestWeatherServiceUnexpectedErrors/Unexpected_Weather_Service_Error
--- PASS: TestWeatherServiceUnexpectedErrors (0.00s)
    --- PASS: TestWeatherServiceUnexpectedErrors/Unexpected_Internal_Server_Error (0.00s)
    --- PASS: TestWeatherServiceUnexpectedErrors/Unexpected_Unprocessable_Entity_Error (0.00s)
    --- PASS: TestWeatherServiceUnexpectedErrors/Unexpected_Not_Found_Error (0.00s)
    --- PASS: TestWeatherServiceUnexpectedErrors/Unexpected_Weather_Service_Error (0.00s)
=== RUN   TestWeatherServiceGetWeatherData
=== RUN   TestWeatherServiceGetWeatherData/12345678
=== RUN   TestWeatherServiceGetWeatherData/00000000
=== RUN   TestWeatherServiceGetWeatherData/123
--- PASS: TestWeatherServiceGetWeatherData (0.00s)
    --- PASS: TestWeatherServiceGetWeatherData/12345678 (0.00s)
    --- PASS: TestWeatherServiceGetWeatherData/00000000 (0.00s)
    --- PASS: TestWeatherServiceGetWeatherData/123 (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/service       1.072s  coverage: 100.0% of statements
=== RUN   TestMockHTTPClientDo
--- PASS: TestMockHTTPClientDo (0.00s)
=== RUN   TestMockCPUService
--- PASS: TestMockCPUService (0.00s)
=== RUN   TestMockMemoryService
--- PASS: TestMockMemoryService (0.00s)
=== RUN   TestMockWeatherService
=== RUN   TestMockWeatherService/Success
=== RUN   TestMockWeatherService/Invalid_ZipCode_Failure
=== RUN   TestMockWeatherService/ZipCode_Not_Found_Failure
=== RUN   TestMockWeatherService/Unexpected_Error_Failure
--- PASS: TestMockWeatherService (0.00s)
    --- PASS: TestMockWeatherService/Success (0.00s)
    --- PASS: TestMockWeatherService/Invalid_ZipCode_Failure (0.00s)
    --- PASS: TestMockWeatherService/ZipCode_Not_Found_Failure (0.00s)
    --- PASS: TestMockWeatherService/Unexpected_Error_Failure (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/service/mock  1.115s  coverage: 100.0% of statements
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
=== RUN   TestNewWeatherLocationByCepUsecase
--- PASS: TestNewWeatherLocationByCepUsecase (0.00s)
=== RUN   TestGetWeatherLocationByCep
=== RUN   TestGetWeatherLocationByCep/Invalid_CEP
=== RUN   TestGetWeatherLocationByCep/CEP_Not_Found
=== RUN   TestGetWeatherLocationByCep/WeahterZip_Service_Error
=== RUN   TestGetWeatherLocationByCep/Success
--- PASS: TestGetWeatherLocationByCep (0.00s)
    --- PASS: TestGetWeatherLocationByCep/Invalid_CEP (0.00s)
    --- PASS: TestGetWeatherLocationByCep/CEP_Not_Found (0.00s)
    --- PASS: TestGetWeatherLocationByCep/WeahterZip_Service_Error (0.00s)
    --- PASS: TestGetWeatherLocationByCep/Success (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase       1.191s  coverage: 100.0% of statements
=== RUN   TestMockHealthCheckUseCase
--- PASS: TestMockHealthCheckUseCase (0.00s)
=== RUN   TestMockWeatherLocationByCepUsecase
=== RUN   TestMockWeatherLocationByCepUsecase/Success
=== RUN   TestMockWeatherLocationByCepUsecase/Invalid_ZipCode_Failure
=== RUN   TestMockWeatherLocationByCepUsecase/ZipCode_Not_Found_Failure
=== RUN   TestMockWeatherLocationByCepUsecase/Unexpected_Error_Failure
--- PASS: TestMockWeatherLocationByCepUsecase (0.00s)
    --- PASS: TestMockWeatherLocationByCepUsecase/Success (0.00s)
    --- PASS: TestMockWeatherLocationByCepUsecase/Invalid_ZipCode_Failure (0.00s)
    --- PASS: TestMockWeatherLocationByCepUsecase/ZipCode_Not_Found_Failure (0.00s)
    --- PASS: TestMockWeatherLocationByCepUsecase/Unexpected_Error_Failure (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/mock  1.305s  coverage: 100.0% of statements
```

Cobertura dos testes

```plain-text
github.com/vs0uz4/observability/ms-inputvalidate/cmd/inputvalidate/main.go:33:                                  main                                    0.0%
github.com/vs0uz4/observability/ms-inputvalidate/cmd/inputvalidate/main.go:79:                                  enjoy                                   0.0%
github.com/vs0uz4/observability/ms-inputvalidate/config/config.go:16:                                           LoadConfig                              100.0%
github.com/vs0uz4/observability/ms-inputvalidate/doc/swagger/docs.go:232:                                       init                                    0.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/domain/errors.go:34:                                  NewUnexpectedStatusCodeError            100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/domain/errors.go:38:                                  NewUnexpectedWeatherServiceError        100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/domain/errors.go:42:                                  NewFailedToCreateRequestError           100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/domain/errors.go:46:                                  NewFailedToMakeRequestError             100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/domain/errors.go:50:                                  NewFailedToDecodeResponseError          100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/domain/errors.go:54:                                  NewInvalidZipCodeDetailsError           100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/utils/validations_utils.go:3:                   IsNumeric                               100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/health_handler.go:14:                       NewHealthHandler                        100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/health_handler.go:26:                       GetHealth                               100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/input_handler.go:24:                        NewInputHandler                         100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/input_handler.go:40:                        GetLocationWeatherByCep                 100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware/error_logger.go:16:    WriteHeader                             100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware/error_logger.go:24:    Write                                   100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware/error_logger.go:30:    ReadError                               100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware/error_logger.go:34:    WriteError                              100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware/error_logger.go:38:    ErrorLogger                             100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/server.go:24:                     NewWebServer                            100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/server.go:45:                     setupDependencies                       100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/server.go:49:                     AddHandler                              100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/server.go:57:                     Start                                   100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/server.go:80:                     Run                                     100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/server.go:90:                     Stop                                    100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/cpu_stats.go:19:                              NewCPUService                           100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/cpu_stats.go:25:                              roundToOneDecimal                       100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/cpu_stats.go:29:                              GetCPUStats                             100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/memory_stats.go:17:                           NewMemoryService                        100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/memory_stats.go:23:                           roundToOneDecimal                       100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/memory_stats.go:27:                           GetMemoryStats                          100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/mock/http_client.go:11:                       Do                                      100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/mock/mock_cpu_stats.go:7:                     GetCPUStats                             100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/mock/mock_memory_stats.go:7:                  GetMemoryStats                          100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/mock/mock_weather.go:13:                      GetWeather                              100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/uptime.go:13:                                 NewUptimeService                        100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/uptime.go:19:                                 GetUptime                               100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/weather.go:34:                                NewWeatherService                       100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/weather.go:41:                                GetWeather                              100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/service/weather.go:90:                                parseErrorResponse                      100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/health_check.go:20:                           NewHealthCheckUseCase                   100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/health_check.go:24:                           GetHealth                               100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/mock/mock_health_check.go:9:                  GetHealth                               100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/mock/mock_weather-location_by_cep.go:13:      GetWeatherLocationByCep                 100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/weather-location_by_cep.go:17:                NewWeatherLocationByCepUsecase          100.0%
github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/weather-location_by_cep.go:23:                GetWeatherLocationByCep                 100.0%
total:                                                                                                          (statements)                            87.7%
```
