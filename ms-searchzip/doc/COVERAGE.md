# MS-SEARCHZIP COVERAGE

Abaixo temos o `report` da cobertura dos testes implementados no SearchZip (SERVIÇO A)

```plain-text
Running tests for ms-searchzip...
        github.com/vs0uz4/observability/ms-searchzip/cmd/searchzip              coverage: 0.0% of statements
        github.com/vs0uz4/observability/ms-searchzip/internal/infra/utils               coverage: 0.0% of statements
?       github.com/vs0uz4/observability/ms-searchzip/internal/infra/web/health  [no test files]
?       github.com/vs0uz4/observability/ms-searchzip/internal/service/contract  [no test files]
?       github.com/vs0uz4/observability/ms-searchzip/internal/usecase/contract  [no test files]
=== RUN   TestLoadConfigReadInConfigFails
--- PASS: TestLoadConfigReadInConfigFails (0.00s)
=== RUN   TestLoadConfigPanicOnUnmarshalError
    config_test.go:40: Panic successfully captured: Unmarshal error
--- PASS: TestLoadConfigPanicOnUnmarshalError (0.00s)
=== RUN   TestLoadConfigMissingRequiredConfigFails
--- PASS: TestLoadConfigMissingRequiredConfigFails (0.00s)
=== RUN   TestLoadConfig
--- PASS: TestLoadConfig (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-searchzip/config     0.342s  coverage: 100.0% of statements
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
=== RUN   TestNewUnexpectedWeatherServiceError
--- PASS: TestNewUnexpectedWeatherServiceError (0.00s)
=== RUN   TestNewFailedToCreateRequestError
--- PASS: TestNewFailedToCreateRequestError (0.00s)
=== RUN   TestNewFailedToMakeRequestError
--- PASS: TestNewFailedToMakeRequestError (0.00s)
=== RUN   TestNewFailedToDecodeResponseError
--- PASS: TestNewFailedToDecodeResponseError (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-searchzip/internal/domain    0.174s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-searchzip/internal/infra/web 0.363s  coverage: 100.0% of statements
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
2024/12/30 02:31:43 "INVALID /test HTTP/1.1" from [::1]:52646 - 405 0B in 7.833µs - Error: 
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
ok      github.com/vs0uz4/observability/ms-searchzip/internal/infra/web/webserver       2.019s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-searchzip/internal/infra/web/webserver/middleware    0.653s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-searchzip/internal/service   0.897s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-searchzip/internal/service/mock      0.936s  coverage: 100.0% of statements
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
=== RUN   TestGetWeatherLocationByCep/CEP_Service_Error
=== RUN   TestGetWeatherLocationByCep/WeahterZip_Service_Error
=== RUN   TestGetWeatherLocationByCep/Success
--- PASS: TestGetWeatherLocationByCep (0.00s)
    --- PASS: TestGetWeatherLocationByCep/Invalid_CEP (0.00s)
    --- PASS: TestGetWeatherLocationByCep/CEP_Not_Found (0.00s)
    --- PASS: TestGetWeatherLocationByCep/CEP_Service_Error (0.00s)
    --- PASS: TestGetWeatherLocationByCep/WeahterZip_Service_Error (0.00s)
    --- PASS: TestGetWeatherLocationByCep/Success (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/vs0uz4/observability/ms-searchzip/internal/usecase   1.014s  coverage: 100.0% of statements
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
ok      github.com/vs0uz4/observability/ms-searchzip/internal/usecase/mock      1.160s  coverage: 100.0% of statements
Calculating coverage for ms-searchzip...
ms-searchzip Coverage: 90.2%
---------------------------------------------------------------------------------------------------------------------
```