package response

import (
	"encoding/json"
	"github.com/OpenFactor-Holding/superapp/assemblers"
	"github.com/OpenFactor-Holding/superapp/dtos"
	"github.com/OpenFactor-Holding/superapp/kafka"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	PersistError          = "an error occurred while creating "
	PersistSuccess        = "successfully created a new "
	PersistErrorCode      = "5000"
	JSONParseErrorMessage = "failed to extract JSON from request body"
	JSONParseErrorCode    = "4000"
	NotFoundErrorCode     = "4040"
	NotFoundMessage = " not found"
	NotFoundMessageDescriptive       = "no record found for provided "
	FetchSuccess          = "success"
	FetchFailure          = "failed"
	ErrorLogs                = "ERROR_LOGS"
	AuditLogs                = "AUDIT_LOGS"
	InternalServerError = "internal server error occurred"
	JsonParseError = "could not parse the request body"
)

func JSONParseError(err interface{}, ctx *gin.Context, reqBody []byte, serviceId string, userId string) dtos.APIResponse {
	var apiError = err.(dtos.APIError)
	apiResponse := dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusBadRequest),
		StatusMessage: JSONParseErrorMessage,
		Error: dtos.APIError{
			ErrorCode:    JSONParseErrorCode,
			ErrorMessage: JsonParseError,
			ErrorDetails: apiError.ErrorDetails,
		},
	}
	var jsonMap map[string]interface{}
	json.Unmarshal(reqBody, &jsonMap)
	LogError(ctx, apiResponse, jsonMap, apiError.ErrorMessage, serviceId, userId)
	return apiResponse
}

func EntityPersistSuccess(entityName string, data interface{}) dtos.APIResponse {
	return dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusCreated),
		StatusMessage: PersistSuccess + entityName,
		Data:          data,
	}
}

func EntityFetchSuccess(data interface{}) dtos.APIResponse {
	return dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusOK),
		StatusMessage: FetchSuccess,
		Data:          data,
	}
}

func EntityFetchError(err error, ctx *gin.Context, serviceId string, userId string) dtos.APIResponse {
	apiResponse := dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusInternalServerError),
		StatusMessage: FetchFailure,
		Error: dtos.APIError{
			ErrorCode:    PersistErrorCode,
			ErrorMessage: InternalServerError,
		},
	}
	LogError(ctx, apiResponse, make(map[string]interface{}), err, serviceId, userId)
	return apiResponse
}

func EntityPersistError(entityName string, error error, reqBody []byte, ctx *gin.Context, serviceId string, userId string) dtos.APIResponse {
	apiResponse := dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusInternalServerError),
		StatusMessage: PersistError + entityName,
		Error: dtos.APIError{
			ErrorCode:    PersistErrorCode,
			ErrorMessage: InternalServerError,
		},
	}
	var jsonMap map[string]interface{}
	json.Unmarshal(reqBody, &jsonMap)
	LogError(ctx, apiResponse, jsonMap, error, serviceId, userId)
	return apiResponse
}

func EntityNotFoundError(entityName string) dtos.APIResponse {
	return dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusNotFound),
		StatusMessage: entityName + NotFoundMessage,
		Error: dtos.APIError{
			ErrorCode:    NotFoundErrorCode,
			ErrorMessage: NotFoundMessageDescriptive + entityName +"_id",
		},
	}
}

func LogError(
	ctx *gin.Context,
	response dtos.APIResponse,
	requestBody map[string]interface{},
	errorBody error,
	serviceId string,
	userId string) {
	err := crash(errorBody)

	stackTrace := err.(*errors.Error).ErrorStack()

	_, filename, line, _ := runtime.Caller(1)
	var apiError = response.Error.(dtos.APIError)

	kafka.Publish(
		assemblers.AssembleErrorLog(
			ctx,
			dtos.Error{
				RequestBody:           requestBody,
				RequestResponseBody:   response,
				RequestResponseStatus: response.StatusCode,
				ErrorMessage:          apiError.ErrorMessage,
				ErrorCode:             apiError.ErrorCode,
				ErrorStackTrace:       stackTrace,
				ErrorLineNumber:       strconv.Itoa(line),
				ErrorFileName:         filename,
				RequestServiceID:      serviceId,
				RequestUserID:         userId,
			}),
		os.Getenv(ErrorLogs))
}
func crash(err error) error {
	return errors.Errorf(err.Error())
}

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func LogAudit(
	ctx *gin.Context,
	apiResponse dtos.APIResponse,
	requestBody map[string]interface{},
	userId string,
	serviceId string,
	actionType string,
	entityTypes string,
) {
	kafka.Publish(
		assemblers.AssembleAuditLog(
			ctx,
			dtos.Audit{
				RequestEntityTypes:    strings.Split(entityTypes, ","),
				RequestActionType:     actionType,
				RequestBody:           requestBody,
				RequestResponseBody:   apiResponse,
				RequestResponseStatus: apiResponse.StatusCode,
				RequestServiceID:      serviceId,
				RequestUserID:         userId,
			}),
		os.Getenv(AuditLogs))
}
