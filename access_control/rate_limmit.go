package access_control


import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)


type request struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}


var requests = make(map[string]*request)
var mu sync.Mutex

func getRequest(ip string, limit int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := requests[ip]
	if !exists {
		rt := rate.Every(24*time.Hour / 50)
		limiter := rate.NewLimiter(rt, limit)
		requests[ip] = &request{limiter, time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

func throttle(next http.Handler, limit int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		limiter := getRequest(ip, limit)
		fmt.Println(limiter.Allow())
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
