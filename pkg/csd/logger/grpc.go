package logger

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var (
	Marshaller          = &jsonpb.Marshaler{}
	TimestampLog        = true
	ServiceField        = "service"
	ServiceLog          = true
	MethodField         = "method"
	MethodLog           = true
	DurationField       = "dur"
	DurationLog         = true
	IPField             = "ip"
	IPLog               = true
	MetadataField       = "md"
	MetadataLog         = true
	UserAgentField      = "ua"
	UserAgentLog        = true
	ReqField            = "req"
	ReqLog              = true
	RespField           = "resp"
	RespLog             = true
	MaxSize             = 2048000
	CodeField           = "code"
	MsgField            = "msg"
	DetailsField        = "details"
	UnaryMessageDefault = "unary"
)

type GRPCLog struct {
	Ctx    context.Context
	Method string
	Start  time.Time
	Req    interface{}
	Resp   interface{}
	Err    error
}

func (g *GRPCLog) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString(ServiceField, path.Dir(g.Method)[1:])
	enc.AddString(MethodField, path.Base(g.Method))
	enc.AddString(DurationField, fmt.Sprintf("%v", time.Since(g.Start)))
	enc.AddString(ReqField, LogBody(g.Req))
	enc.AddString(RespField, LogBody(g.Resp))
	enc.AddString(CodeField, LogStatusError(g.Err))

	if IPLog {
		if p, ok := peer.FromContext(g.Ctx); ok {
			enc.AddString(IPField, p.Addr.String())
		}
	}
	if MetadataLog {
		enc.AddString(MetadataField, LogMetadata(g.Ctx))
	}
	if UserAgentLog {
		if md, ok := metadata.FromIncomingContext(g.Ctx); ok {
			enc.AddString(UserAgentField, strings.Join(md.Get("user-agent"), ""))
		}
	}
	return nil
}

func NewUnaryServerInterceptorWithLogger(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		now := time.Now()
		resp, err := handler(ctx, req)

		grpcLog := &GRPCLog{
			Ctx:    ctx,
			Method: info.FullMethod,
			Start:  now,
			Req:    req,
			Resp:   resp,
			Err:    err,
		}

		if err != nil {
			logger.Error("grpc error", zap.Object("grpc", grpcLog))
		} else {
			logger.Info("grpc request", zap.Object("grpc", grpcLog))
		}
		return resp, err
	}
}

func LogBody(req interface{}) string {
	if b := GetRawJSON(req); b != nil {
		return b.String()
	}
	return ""
}

func GetRawJSON(i interface{}) *bytes.Buffer {
	if pb, ok := i.(proto.Message); ok {
		b := &bytes.Buffer{}
		if err := Marshaller.Marshal(b, pb); err == nil && b.Len() < MaxSize {
			return b
		}
	}
	return nil
}

func LogMetadata(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		var result strings.Builder
		for k, v := range md {
			result.WriteString(fmt.Sprintf("%s: %s; ", k, strings.Join(v, ",")))
		}
		return result.String()
	}
	return ""
}

func LogStatusError(err error) string {
	if err == nil {
		return ""
	}
	statusErr := status.Convert(err)
	return fmt.Sprintf(
		"code: %d; message: %s; details: %s",
		statusErr.Code(),
		statusErr.Message(),
		statusErr.Details(),
	)
}
