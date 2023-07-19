package sheduler

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
)

func HealthCheckShedule() {
	cr := cron.New()
	fmt.Println("Health Check")
	cr.AddFunc("*/5 * * * * *", func() {
		_, err := http.Get("http://localhost:8080/service3/healthcheck")
		if err != nil {
		}
	})
	cr.Start()
	select {}
}
