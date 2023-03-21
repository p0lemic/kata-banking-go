package account

import (
	"time"
)

type Operation struct {
	date    time.Time
	amount  string
	balance int
}
