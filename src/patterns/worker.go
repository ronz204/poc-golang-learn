package patterns

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Trabajador %d procesando tarea %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func DemoWorkerPool() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for id := 1; id <= 3; id++ {
		go worker(id, jobs, results)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Resultado recibido: %d\n", result)
	}
}
