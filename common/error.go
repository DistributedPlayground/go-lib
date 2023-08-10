package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogError(c echo.Context, err error, handlerMsg string) {
	lg := c.Get("logger").(*zerolog.Logger)
	lg.Error().Stack().Err(err).Msg(handlerMsg)
}

func LogDPError(c echo.Context, err error, handlerMsg string) {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	tracer, ok := errors.Cause(err).(stackTracer)
	if !ok {
		log.Warn().Str("error", err.Error()).Msg("error does not implement stack trace")
		return
	}

	cause := errors.Cause(err)
	st := tracer.StackTrace()

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Warn().Msg("service name is missing from env")
	}

	if IsLocalEnv() {
		st2 := fmt.Sprintf("\nSTACK TRACE:\n%+v: [%+v ]\n\n", cause.Error(), st)
		st2 = strings.ReplaceAll(st2, "/"+serviceName+"/", "")
		fmt.Print(st2)
		return
	}

	LogError(c, err, handlerMsg)
}

func DPError(err error, optionalMsg ...string) error {
	if err == nil {
		return nil
	}

	concat := ""

	for _, msgs := range optionalMsg {
		concat += msgs + " "
	}

	if errors.Cause(err) == nil || errors.Cause(err).Error() == err.Error() {
		return errors.Wrap(errors.New(err.Error()), concat)
	}

	return errors.Wrap(err, concat)
}
