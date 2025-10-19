package utils

import (
	"net/http"
	"runtime"
)

func GetMemoryUsage() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return bToMb(m.Alloc)
}

func bToMb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
