package main

import (
	"log"
	"time"
	"github.com/go-co-op/gocron"
)

func startHealthCheck() {
	sch := gocron.NewScheduler(time.Local)
	for _, host := range serverList {
		_, err := sch.Every(2).Second().Do(func(s *server) {
			IsHealthy := s.checkHealth()
			if IsHealthy {
				log.Printf("%s is Healthy\n", s.Name)
			} else {
				log.Printf("%s is not Healthy\n", s.Name)
			}
		}, host)
		if err != nil {
			log.Fatalln(err)
		}
	}
	sch.StartAsync()
}
