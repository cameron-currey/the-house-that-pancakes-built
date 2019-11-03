package restaurant

import (
	"../pancake"
)

type Waiter interface {
	FlipDatStack(stack pancake.Stack) int
}
