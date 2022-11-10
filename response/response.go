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
	JSONParseErrorMessage = "failed to extract JSON from request body"
	NotFoundErrorCode     = "4040"
	NotFoundMessage = " not found"
	NotFoundMessageDescriptive       = "no record found for provided "
	FetchSuccess          = "success"
	FetchFailure          = "failed"
	ErrorLogs                = "ERROR_LOGS"
	AuditLogs                = "AUDIT_LOGS"
)

func JSONParseError(apiError dtos.APIError, ctx *gin.Context, reqBody []byte, serviceId string, userId string) dtos.APIResponse {
	apiResponse := dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusBadRequest),
		StatusMessage: JSONParseErrorMessage,
		Error: apiError,
	}
	var jsonMap map[string]interface{}
	json.Unmarshal(reqBody, &jsonMap)
	LogError(ctx, apiResponse, jsonMap, apiError, serviceId, userId)
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

func EntityFetchError(apiError dtos.APIError, ctx *gin.Context, serviceId string, userId string) dtos.APIResponse {
	apiResponse := dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusInternalServerError),
		StatusMessage: FetchFailure,
		Error: apiError,
	}
	LogError(ctx, apiResponse, make(map[string]interface{}), apiError, serviceId, userId)
	return apiResponse
}

func EntityPersistError(entityName string, apiError dtos.APIError, reqBody []byte, ctx *gin.Context, serviceId string, userId string) dtos.APIResponse {
	apiResponse := dtos.APIResponse{
		StatusCode:    strconv.Itoa(http.StatusInternalServerError),
		StatusMessage: PersistError + entityName,
		Error: apiError,
	}
	var jsonMap map[string]interface{}
	json.Unmarshal(reqBody, &jsonMap)
	LogError(ctx, apiResponse, jsonMap, apiError, serviceId, userId)
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
	errorBody dtos.APIError,
	serviceId string,
	userId string) {
	err := crash(errors.New(errorBody.ErrorMessage))

	stackTrace := err.(*errors.Error).ErrorStack()

	_, filename, line, _ := runtime.Caller(1)

	kafka.Publish(
		assemblers.AssembleErrorLog(
			ctx,
			dtos.Error{
				RequestBody:           requestBody,
				RequestResponseBody:   response,
				RequestResponseStatus: response.StatusCode,
				ErrorMessage:          errorBody.ErrorMessage,
				ErrorCode:             errorBody.ErrorCode,
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
