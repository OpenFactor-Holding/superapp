package assemblers

import (
	"fmt"
	"github.com/OpenFactor-Holding/superapp/dtos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

const (
	httpRequest = "HTTP"
)

func AssembleAuditLog(c *gin.Context, audit dtos.Audit) dtos.AuditLog {

	requestHeaders := buildHeaders(c)

	return dtos.AuditLog{
		RecordID:              uuid.New(),
		RequestUri:            c.Request.URL.Path,
		RequestHttpMethod:     c.Request.Method,
		RequestHeaders:        requestHeaders,
		RequestBody:           audit.RequestBody,
		RequestIpAddress:      c.Request.RemoteAddr,
		RequestTimestamp:      time.Now(),
		RequestType:           httpRequest,
		RequestDeviceType:     c.Request.UserAgent(),
		RequestResponseBody:   audit.RequestResponseBody,
		RequestResponseStatus: audit.RequestResponseStatus,
		RequestUserID:         audit.RequestUserID,
		RequestServiceID:      audit.RequestServiceID,
		RequestEntityTypes:    audit.RequestEntityTypes,
		RequestActionType:     audit.RequestActionType,
	}
}

func AssembleErrorLog(c *gin.Context, error dtos.Error) dtos.ErrorLogs {

	requestHeaders := buildHeaders(c)

	return dtos.ErrorLogs{
		RecordID:              uuid.New(),
		RequestUri:            c.Request.URL.Path,
		RequestHttpMethod:     c.Request.Method,
		RequestHeaders:        requestHeaders,
		RequestBody:           error.RequestBody,
		RequestIpAddress:      c.Request.RemoteAddr,
		RequestTimestamp:      time.Now(),
		RequestType:           httpRequest,
		RequestDeviceType:     c.Request.UserAgent(),
		RequestResponseBody:   error.RequestResponseBody,
		RequestResponseStatus: error.RequestResponseStatus,
		RequestUserID:         error.RequestUserID,
		RequestServiceID:      error.RequestServiceID,
		ErrorCode:             error.ErrorCode,
		ErrorMessage:          error.ErrorMessage,
		ErrorFileName:         error.ErrorFileName,
		ErrorLineNumber:       error.ErrorLineNumber,
		ErrorMethodName:       error.ErrorMethodName,
		ErrorStackTrace:       error.ErrorStackTrace,
	}
}

func AssembleEventLog(c *gin.Context, event dtos.Event) dtos.EventLog {

	requestHeaders := buildHeaders(c)

	return dtos.EventLog{
		RecordID:              uuid.New(),
		RequestUri:            c.Request.URL.Path,
		RequestHttpMethod:     c.Request.Method,
		RequestHeaders:        requestHeaders,
		RequestBody:           event.RequestBody,
		RequestIpAddress:      c.Request.RemoteAddr,
		RequestTimestamp:      time.Now(),
		RequestType:           httpRequest,
		RequestDeviceType:     c.Request.UserAgent(),
		RequestResponseBody:   event.RequestResponseBody,
		RequestResponseStatus: event.RequestResponseStatus,
		RequestUserID:         event.RequestUserID,
		RequestServiceID:      event.RequestServiceID,
		TopicName:             event.TopicName,
		BrokerHost:            event.BrokerHost,
		BrokerPort:            event.BrokerPort,
		FileName:              event.FileName,
		MethodName:            event.MethodName,
		LineNumber:            event.LineNumber,
		LogLevel:              event.LogLevel,
		Message:               event.Message,
	}
}

func AssembleCommnLog(c *gin.Context, commn dtos.Communication) dtos.CommunicationLog {

	requestHeaders := buildHeaders(c)

	return dtos.CommunicationLog{
		RecordID:              uuid.New(),
		RequestUri:            c.Request.URL.Path,
		RequestHttpMethod:     c.Request.Method,
		RequestHeaders:        requestHeaders,
		RequestBody:           commn.RequestBody,
		RequestIpAddress:      c.Request.RemoteAddr,
		RequestTimestamp:      time.Now(),
		RequestType:           httpRequest,
		RequestDeviceType:     c.Request.UserAgent(),
		RequestResponseBody:   commn.RequestResponseBody,
		RequestResponseStatus: commn.RequestResponseStatus,
		RequestUserID:         commn.RequestUserID,
		RequestServiceID:      commn.RequestServiceID,
		ProviderID:            commn.ProviderID,
		ChannelID:             commn.ChannelID,
		ChannelName:           commn.ChannelName,
		ChannelType:           commn.ChannelType,
		GatewayName:           commn.GatewayName,
		GatewayIpAddress:      commn.GatewayIpAddress,
		GatewayPort:           commn.GatewayPort,
		GatewayEndpoint:       commn.GatewayEndpoint,
		AuthRequired:          commn.AuthRequired,
		AuthType:              commn.AuthType,
		AuthCredentials:       commn.AuthCredentials,
		DeliveryStatus:        commn.DeliveryStatus,
	}
}

func buildHeaders(c *gin.Context) map[string]interface{} {

	var requestHeaders = make(map[string]interface{})
	for name, headers := range c.Request.Header {
		for _, header := range headers {
			fmt.Printf("%v: %v\n", name, header)
			requestHeaders[name] = header
		}
	}
	return requestHeaders
}
