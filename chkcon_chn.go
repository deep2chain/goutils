package goutils

import (
	"fmt"
)

type Worker struct {
	data    chan string
	quit    chan bool
	stopped bool
	success int
	failure int
}

func (w *Worker) Success() int {
	return w.success
}

func (w *Worker) IsExecuted() bool {
	return w.success != 0 && w.failure != 0
}

func (w *Worker) Stop() {
	w.quit <- true
	w.stopped = true
}

func (w *Worker) Run() {
	w.eventloop()
}

func (w *Worker) Receive(data string) {
	fmt.Printf("recieve(%v) \n", data)
	w.data <- data

}
func (w *Worker) eventloop() {
	for {
		if w.stopped {
			return
		}
		select {
		case task := <-w.data:
			_, err := GetWebSite(task)
			if err != nil {
				fmt.Printf("task(%v) failed(%v)\n", task, err)
				w.failure += 1
			} else {
				fmt.Println(task, "is okay")
				w.success += 1
			}
			if w.stopped {
				return
			}
		case <-w.quit:
			return
		}
	}
}

func Run(urls []string) int {
	cnt := len(urls)
	workers := make([]Worker, cnt)
	for i := range workers {
		go func(i int) {
			workers[i].Run()
		}(i)
	}
	go func() {
		sched(urls, workers)
	}()
	// for {
	// 	if len(workers) < 1 {
	// 		break
	// 	}
	// 	for i := range workers {
	// 		if workers[i].IsExecuted() {
	// 			workers[i].Stop()
	// 		}
	// 	}
	// }
	stopall(workers)
	return sumup(workers)

}

func sched(tasks []string, workers []Worker) {
	fmt.Printf("on schedul\n")
	if len(tasks) != len(workers) {
		fmt.Printf("len(tasks)%d == len(workers:)%d expected\n", len(tasks), len(workers))
		return
	}
	for i := range workers {
		fmt.Printf("task(%v) \n", i)
		workers[i].Receive(tasks[i])
	}
}

func stopall(workers []Worker) {
	for i := range workers {
		workers[i].Stop()
	}
}

func sumup(workers []Worker) int {
	var success int
	for i := range workers {
		success += workers[i].Success()
	}
	return success
}
