package access_control

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)


type requestLimit struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}


type RequestLimit interface {
	getRequest(ip string, limit int) *rate.Limiter
	RateLimitRequest(next http.Handler, limit int) http.Handler
}

func NewRequestLimit(limiter  *rate.Limiter) RequestLimit {
	return &requestLimit{
		limiter:  limiter,
	}
}

var (
	requests = make(map[string]*requestLimit)
 	mu sync.Mutex
)

func (request *requestLimit) getRequest(ip string, limit int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := requests[ip]
	if !exists {
		rt := rate.Every(24*time.Hour / 50)
		limiter := rate.NewLimiter(rt, limit)
		requests[ip] = &requestLimit{limiter, time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

func (request *requestLimit) RateLimitRequest(next http.Handler, limit int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		limiter := request.getRequest(ip, limit)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
