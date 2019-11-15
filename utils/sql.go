package utils

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

func NullString(d string) sql.NullString {
	return sql.NullString{String: d, Valid: true}
}

func NullInt64(d int64) sql.NullInt64 {
	return sql.NullInt64{Int64: d, Valid: true}
}

func NullTime(d time.Time) pq.NullTime {
	return pq.NullTime{Time: d, Valid: true}
}

func NullBool(d bool) sql.NullBool {
	return sql.NullBool{Bool: d, Valid: true}
}
