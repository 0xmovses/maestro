package shared

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Features map[string]bool `default:"useCorsMiddleware:false,validateTimestamps:true,validateAllowlist:true,validateBlocklist:true,validateSigs:true"`
}

type Database struct {
	Conn    *pgxpool.Pool
	Context context.Context
	Name    string
	Env     *string
}

type PitchRequest struct {
	Pitches []int `json:"pitches"`
}

type VoicingsRequest struct {
	Pitches []int `json:"pitches"`
	Amount  int   `json:"amount"`
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
