package logging

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	sdk_log "github.com/monetha/reputation-go-sdk/log"
)

const (
	timeFormat = "2006-01-02T15:04:05-0700"
	errorKey   = "LOG_ERROR"
)

// Fun is a function for logging with context properties
var Fun = sdk_log.Fun(logFun)

func logFun(msg string, ctx ...interface{}) {
	if len(ctx) > 0 {
		props := make(map[string]interface{})

		if len(ctx)%2 != 0 {
			ctx = append(ctx, nil, errorKey, "Normalized odd number of arguments by adding nil")
		}

		for i := 0; i < len(ctx); i += 2 {
			k, ok := ctx[i].(string)
			if !ok {
				props[errorKey] = fmt.Sprintf("%+v is not a string key", ctx[i])
			}
			props[k] = formatJSONValue(ctx[i+1])
		}

		bs, _ := json.Marshal(props)
		msg = msg + "\t" + string(bs)
	}

	log.Println(msg)
}

func formatJSONValue(value interface{}) interface{} {
	value = formatShared(value)
	switch value.(type) {
	case int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32, uint64, string:
		return value
	default:
		return fmt.Sprintf("%+v", value)
	}
}

func formatShared(value interface{}) (result interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if v := reflect.ValueOf(value); v.Kind() == reflect.Ptr && v.IsNil() {
				result = "nil"
			} else {
				panic(err)
			}
		}
	}()

	switch v := value.(type) {
	case time.Time:
		return v.Format(timeFormat)

	case error:
		return v.Error()

	case fmt.Stringer:
		return v.String()

	default:
		return v
	}
}
