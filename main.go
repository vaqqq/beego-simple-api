package main

import (
	_ "GoApp/routers"
	"log"
	"sync"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5) // 1 req/s, max 5 burst
		visitors[ip] = limiter
	}
	return limiter
}

func main() {
	beego.BConfig.WebConfig.XSRFKey = beego.AppConfig.DefaultString("xsrfkey", "defaultKey")
	beego.BConfig.WebConfig.XSRFExpire = beego.AppConfig.DefaultInt("xsrfexpire", 5400)

	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		ip := ctx.Input.IP()
		limiter := getVisitor(ip)

		if !limiter.Allow() {
			log.Println("rate limit exceeded:", ip)
			ctx.ResponseWriter.WriteHeader(429)
			ctx.ResponseWriter.Write([]byte("Too Many Requests\n"))
		}
	})

	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.ResponseWriter.Header().Set("X-Frame-Options", "DENY")
		ctx.ResponseWriter.Header().Set("Content-Security-Policy", "default-src 'self'")
		ctx.ResponseWriter.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	})

	beego.Run()
}
