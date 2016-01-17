package main

import (
	"fmt"
	"github.com/hulucat/utils"
	"net/http"
	"runtime"
	"time"
	"tripwar/handlers"
)

func main() {
	utils.Infof("TripWar start to run.")
	go regTasks()
	s := &http.Server{
		Addr:         ":9092",
		Handler:      handlers.HttpRouter{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}

func regTasks() {
	timer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer.C:
			continue
			fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
		}
	}

}
