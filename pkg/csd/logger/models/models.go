package models

import (
	"encoding/json"
	"github.com/rs/zerolog"
)

type HttpFields struct {
	TraceID      string `json:"traceID"`
	RemoteIP     string `json:"remoteIP"`
	Host         string `json:"host"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Protocol     string `json:"protocol"`
	RequestBody  string `json:"requestBody"`
	ResponseBody string `json:"responseBody"`
	StatusCode   int    `json:"statusCode"`
	UserAgent    string `json:"userAgent"`
	Latency      int64  `json:"latency"`
	Error        string `json:"error"`
	Stack        []byte `json:"stack"`
	APIKey       string `json:"APIKey"`
	UserID       int    `json:"userID"`
	ClientID     int    `json:"clientID"`
	APIID        int    `json:"APIID"`
}

func (lf *HttpFields) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("traceID", lf.TraceID).
		Str("remoteIP", lf.RemoteIP).
		Str("host", lf.Host).
		Str("method", lf.Method).
		Str("path", lf.Path).
		Str("protocol", lf.Protocol).
		Str("requestBody", lf.RequestBody).
		Str("responseBody", lf.ResponseBody).
		Int("statusCode", lf.StatusCode).
		Int64("latency", lf.Latency).
		Str("APIKey", lf.APIKey).
		Int("userID", lf.UserID).
		Int("clientID", lf.ClientID).
		Int("APIID", lf.APIID).
		Str("error", lf.Error)
}

type Database struct {
	Name  string `json:"name"`
	Query string `json:"query"`
	Took  string `json:"took"`
}

func (lf *Database) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("name", lf.Name).
		Str("query", lf.Query).
		Str("took", lf.Took)
}

type GRPC struct {
	MD           string `json:"md"`
	Service      string `json:"service"`
	Req          string `json:"req"`
	Resp         string `json:"resp"`
	Method       string `json:"method"`
	Duration     string `json:"duration"`
	ErrorMessage string `json:"errorMessage"`
}

func (lf *GRPC) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("MD", lf.MD).
		Str("Req", lf.Req).
		Str("Resp", lf.Resp).
		Str("Method", lf.Method).
		Str("Duration", lf.Duration).
		Str("ErrorMessage", lf.ErrorMessage)
}

type LogStructure struct {
	HTTP     zerolog.LogObjectMarshaler
	Database zerolog.LogObjectMarshaler
	GRPC     zerolog.LogObjectMarshaler
}

func (lf *LogStructure) MarshalZerologObject(e *zerolog.Event) {
	e.
		RawJSON("HTTP", UnsafeMarshalJSON(lf.HTTP)).
		RawJSON("Database", UnsafeMarshalJSON(lf.Database)).
		RawJSON("GRPC", UnsafeMarshalJSON(lf.GRPC))
}

func UnsafeMarshalJSON(value interface{}) []byte {
	v, _ := json.Marshal(value)
	if string(v) == "null" {
		return nil
	}
	return v
}
