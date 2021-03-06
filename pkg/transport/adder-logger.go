// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/seniorGolang/dumper/viewer"
	"github.com/sergejkoll/tg-example/pkg/interfaces"
)

type loggerAdder struct {
	next interfaces.Adder
	log  zerolog.Logger
}

func loggerMiddlewareAdder(log zerolog.Logger) MiddlewareAdder {
	return func(next interfaces.Adder) interfaces.Adder {
		return &loggerAdder{
			log:  log,
			next: next,
		}
	}
}

func (m loggerAdder) SumTwoNumbers(ctx context.Context, first int, second int) (result int, err error) {
	log := m.log.With().Str("service", "Adder").Str("method", "sumTwoNumbers").Logger()
	if ctx.Value(headerRequestID) != nil {
		log = log.With().Interface("requestID", ctx.Value(headerRequestID)).Logger()
	}
	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"request": viewer.Sprintf("%+v", requestAdderSumTwoNumbers{
				First:  first,
				Second: second,
			}),
			"response": viewer.Sprintf("%+v", responseAdderSumTwoNumbers{Result: result}),
			"took":     time.Since(begin).String(),
		}
		if err != nil {
			log.Error().Err(err).Fields(fields).Msg("call sumTwoNumbers")
			return
		}
		log.Info().Fields(fields).Msg("call sumTwoNumbers")
	}(time.Now())
	return m.next.SumTwoNumbers(ctx, first, second)
}
