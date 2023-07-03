package internal

import (
	"time"

	"github.com/aggitech/bling-sdk"
)

func NormalizeDate(t time.Time) string {
	return t.Format(bling.DefaultTimeFormat)
}
