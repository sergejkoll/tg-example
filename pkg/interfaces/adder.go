package interfaces

import "context"

// @tg jsonRPC-server log
// @tg http-prefix=api/v1
type Adder interface {
	SumTwoNumbers(ctx context.Context, first int, second int) (result int, err error)
}
