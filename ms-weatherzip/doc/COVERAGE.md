# MS-WEATHERZIP COVERAGE

Abaixo temos o `report` da cobertura dos testes implementados no WeatherZip (SERVIÇO B)

```plain-text
Running tests for ms-weatherzip...
        github.com/vs0uz4/observability/ms-weatherzip/cmd/weatherzip            coverage: 0.0% of statements
?       github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/health [no test files]
?       github.com/vs0uz4/observability/ms-weatherzip/internal/service/contract [no test files]
?       github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/contract [no test files]
=== RUN   TestLoadConfigReadInConfigFails
--- PASS: TestLoadConfigReadInConfigFails (0.00s)
=== RUN   TestLoadConfigPanicOnUnmarshalError
    config_test.go:42: Panic successfully captured: Unmarshal error
--- PASS: TestLoadConfigPanicOnUnmarshalError (0.00s)
=== RUN   TestLoadConfigMissingRequiredConfigFails
--- PASS: TestLoadConfigMissingRequiredConfigFails (0.00s)
=== RUN   TestLoadConfig
--- PASS: TestLoadConfig (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/config    0.661s  coverage: 100.0% of statements
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
=== RUN   TestWeatherResponsePopulateFromMap
=== RUN   TestWeatherResponsePopulateFromMap/Valid_Data
=== RUN   TestWeatherResponsePopulateFromMap/Invalid_Location_Structure
=== RUN   TestWeatherResponsePopulateFromMap/Invalid_Current_Structure
--- PASS: TestWeatherResponsePopulateFromMap (0.00s)
    --- PASS: TestWeatherResponsePopulateFromMap/Valid_Data (0.00s)
    --- PASS: TestWeatherResponsePopulateFromMap/Invalid_Location_Structure (0.00s)
    --- PASS: TestWeatherResponsePopulateFromMap/Invalid_Current_Structure (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/domain   0.469s  coverage: 100.0% of statements
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
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web        0.692s  coverage: 100.0% of statements
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
2024/12/30 03:26:34 "INVALID /test HTTP/1.1" from [::1]:54041 - 405 0B in 8.5µs - Error: 
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
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver      2.342s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware   0.994s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/service  1.279s  coverage: 100.0% of statements
=== RUN   TestMockHTTPClientDo
--- PASS: TestMockHTTPClientDo (0.00s)
=== RUN   TestMockCepService
--- PASS: TestMockCepService (0.00s)
=== RUN   TestMockCPUService
--- PASS: TestMockCPUService (0.00s)
=== RUN   TestMockMemoryService
--- PASS: TestMockMemoryService (0.00s)
=== RUN   TestMockWeatherService
--- PASS: TestMockWeatherService (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/service/mock     1.337s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/usecase  1.487s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/mock     1.311s  coverage: 100.0% of statements
Calculating coverage for ms-weatherzip...
ms-weatherzip Coverage: 90.4%
```
