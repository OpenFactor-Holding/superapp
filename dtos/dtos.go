package dtos

import (
	"github.com/google/uuid"
	"time"
)

type BaseDto struct {
	RequestTimestamp  time.Time `json:"request_timestamp"`
	RequestUri        string    `json:"request_uri"`
	RequestHttpMethod string    `json:"request_http_method"`
	RequestHeaders    string    `json:"request_headers"`
	RequestIpAddress  string    `json:"request_ip_address"`
	RequestUserID     string    `json:"request_user_id"`
	RequestServiceID  string    `json:"request_service_id"`
}

type AuditLog struct {
	RecordID              uuid.UUID `json:"record_id"`
	RequestTimestamp      time.Time `json:"request_timestamp"`
	RequestUri            string    `json:"request_uri"`
	RequestHeaders        string    `json:"request_headers"`
	RequestHttpMethod     string    `json:"request_http_method"`
	RequestIpAddress      string    `json:"request_ip_address"`
	RequestUserID         string    `json:"request_user_id"`
	RequestServiceID      string    `json:"request_service_id"`
	RequestResponseStatus string    `json:"request_response_status"`
	RequestDeviceType     string    `json:"request_device_type"`
	RequestEntityType     string    `json:"request_entity_type"`
	RequestActionType     string    `json:"request_action_type"`
	RequestBody           string    `json:"request_body"`
	RequestResponseBody   string    `json:"request_response_body"`
	RequestType           string    `json:"request_type"`
}

type ErrorLogs struct {
	RecordID          uuid.UUID `json:"record_id"`
	RequestTimestamp  time.Time `json:"request_timestamp"`
	RequestUri        string    `json:"request_uri"`
	RequestHeaders    string    `json:"request_headers"`
	RequestHttpMethod string    `json:"request_http_method"`
	RequestIpAddress  string    `json:"request_ip_address"`
	RequestUserID     string    `json:"request_user_id"`
	RequestServiceID  string    `json:"request_service_id"`
	ErrorFileName     string    `json:"error_file_name"`
	ErrorLineNumber   string    `json:"error_line_number"`
	ErrorMethodName   string    `json:"error_method_name"`
	ErrorTimestamp    time.Time `json:"error_timestamp"`
	ErrorStackTrace   string    `json:"error_stack_trace"`
	ErrorCode         string    `json:"error_code"`
	ErrorMessage      string    `json:"error_message"`
}

type CommunicationLog struct {
	RecordID          uuid.UUID `json:"record_id"`
	RequestTimestamp  time.Time `json:"request_timestamp"`
	RequestUri        string    `json:"request_uri"`
	RequestHeaders    string    `json:"request_headers"`
	RequestHttpMethod string    `json:"request_http_method"`
	RequestIpAddress  string    `json:"request_ip_address"`
	RequestUserID     string    `json:"request_user_id"`
	RequestServiceID  string    `json:"request_service_id"`
	ProviderID        string    `json:"provider_id"`
	ChannelID         string    `json:"channel_id"`
	ChannelName       string    `json:"channel_name"`
	ChannelType       string    `json:"channel_type"`
	GatewayName       string    `json:"gateway_name"`
	GatewayIpAddress  string    `json:"gateway_ip_address"`
	GatewayPort       string    `json:"gateway_port"`
	GatewayEndpoint   string    `json:"gateway_endpoint"`
	AuthRequired      string    `json:"auth_required"`
	AuthType          string    `json:"auth_type"`
	AuthCredentials   string    `json:"auth_credentials"`
	DeliveryStatus    string    `json:"delivery_status"`
}

type EventLog struct {
	RecordID          uuid.UUID `json:"record_id"`
	RequestTimestamp  time.Time `json:"request_timestamp"`
	RequestUri        string    `json:"request_uri"`
	RequestHeaders    string    `json:"request_headers"`
	RequestHttpMethod string    `json:"request_http_method"`
	RequestIpAddress  string    `json:"request_ip_address"`
	RequestUserID     string    `json:"request_user_id"`
	RequestServiceID  string    `json:"request_service_id"`
	TopicName         string    `json:"topic_name"`
	BrokerHost        string    `json:"broker_host"`
	BrokerPort        string    `json:"broker_port"`
	FileName          string    `json:"file_name"`
	MethodName        string    `json:"method_name"`
	LineNumber        string    `json:"line_number"`
	LogLevel          string    `json:"log_level"`
	Message           string    `json:"message"`
}
type APIResponse struct {
	StatusCode    string      `json:"status_code,omitempty"`
	StatusMessage string      `json:"status_message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	Error         interface{} `json:"error,omitempty"`
}
type APIError struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorDetails string `json:"error_details,omitempty"`
}