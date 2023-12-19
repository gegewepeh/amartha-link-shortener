package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/kr/pretty"
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

// Log message with context
func Log(ctx context.Context, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)

	if ctx != nil {
		reqID := GetRequestID(ctx)

		if reqID != "" {
			msg = fmt.Sprintf("[%s] %s", reqID, msg)
		}
	}

	log.Println(msg)
}

func PrettyLog(ctx context.Context, a interface{}) {
	fmt.Printf("%# v \n", pretty.Formatter(a))
}

// SetRequestID for a http call context
func SetRequestID(ctx context.Context, reqID string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contextKeyRequestID, reqID)
}

// GetRequestID of a http call context
func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(contextKeyRequestID).(string); ok {
		return reqID
	}

	return ""
}
