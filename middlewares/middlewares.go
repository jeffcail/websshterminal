package middlewares

import (
	"runtime"
	"strings"

	"github.com/google/uuid"

	"github.com/c/websshterminal.io/ubzer"
	"go.uber.org/zap"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	UIDKey = "Service"
)

// RequestLog
func RequestLog() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4<<10)
					length := runtime.Stack(stack, false)
					ubzer.MLog.Error("程序崩溃", zap.String("崩溃日志", string(stack[:length])))
				}
			}()
			if !strings.HasPrefix(context.Path(), "/api/") {
				ubzer.MLog.Info("请求开始", zap.Any(context.Request().RequestURI, "网页请求"))
				return handlerFunc(context)
			}
			uid := uuid.New().String()
			context.Set(UIDKey, uid)
			err := handlerFunc(context)
			return err
		}
	}
}

var DefaultBodyDumpConfig = middleware.BodyDumpConfig{
	Skipper: BodyDumpDefaultSkipper,
	Handler: func(context echo.Context, bytes []byte, bytes2 []byte) {
		if !strings.HasPrefix(context.Path(), "/api/") {
			return
		}
		uid := context.Get(UIDKey).(string)
		ubzer.MLog.Info("请求结束", zap.String("请求UID", uid), zap.String(context.Request().RequestURI, string(bytes2)))
	},
}

func BodyDumpDefaultSkipper(c echo.Context) bool {
	if !strings.HasPrefix(c.Path(), "/api/") {
		return true
	}
	return false
}
