package internal

import (
	"time"

	"github.com/integrmais/bling"
)

func NormalizeDate(t time.Time) string {
	return t.Format(bling.DefaultTimeFormat)
}
