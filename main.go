package main

import (
	"context"
	"fmt"
	"github.com/procyon-projects/chrono"
	"log"
)

func main() {

	fmt.Println("Iniciando a execução do programa...")

	// taskScheduler := chrono.NewDefaultTaskScheduler()
	//_, err := taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
	//	log.Print("Fixed Delay Task")
	//	time.Sleep(3 * time.Second)
	//}, 5 * time.Second)
	//
	//if err == nil {
	//	log.Print("Task has been scheduled successfully.")
	//}

	taskScheduler := chrono.NewDefaultTaskScheduler()
	_, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
		log.Print("Scheduled Task With Cron")
	}, "0 16,17,18,20 22 * * * ")
	if err == nil {
		log.Print("Task has been scheduled successfully.")
	}


	for {
	}

}
