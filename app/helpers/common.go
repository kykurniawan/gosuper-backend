package helpers

import (
	"database/sql"
	"time"
)

func NilOrTIme(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}

	return nil
}
