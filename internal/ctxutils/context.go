package ctxutils

import (
	"context"
	"time"
)

type contextKey string

const (
	localTimeKey contextKey = "localTime"
)

func NewCTX() context.Context {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ctx = addLocalTimeToContext(ctx)

	return ctx

}

func addLocalTimeToContext(ctx context.Context) context.Context {

	localTime := time.Now().Local().Format("2006-01-02 15:04:05")
	return context.WithValue(ctx, localTimeKey, localTime)
}
