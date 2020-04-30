package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name     string
	Input    int
	Output   int
	WorkFunc func(int) int
}

type Tasks []Task

func square(in int) int {
	return in * in
}

func startTasks(tasks Tasks) <-chan Task {
	tasksCh := make(chan Task)
	go func() {
		defer close(tasksCh)
		for _, task := range tasks {
			tasksCh <- task
		}
	}()

	return tasksCh
}

func startWorker(task <-chan Task, id int) <-chan Task {
	workerCh := make(chan Task)

	go func() {
		defer close(workerCh)
		for t := range task {
			duration := 5
			// simulate slow operations
			fmt.Printf("[Worker %d] Consuming slow task... %s, it will take %d seconds\n", id, t.Name, duration)
			time.Sleep(time.Second * time.Duration(duration))
			t.Output = t.WorkFunc(t.Input)
			// send the result to worker channel
			workerCh <- t
		}
	}()

	return workerCh
}

func collect(workersCh []<-chan Task) <-chan Task {
	var wg sync.WaitGroup
	resultCh := make(chan Task, 100)
	wg.Add(len(workersCh))
	for _, workerCh := range workersCh {
		go func(workersCh <-chan Task) {
			defer wg.Done()
			for workerCh := range workersCh {
				// block until receieve the result from consumer
				resultCh <- workerCh
			}
		}(workerCh)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("time elapsed %v\n", time.Since(start))
	}
}

func main() {
	defer elapsed()()

	// generate tasks
	taskCount := 5
	tasks := Tasks{}
	for i := 1; i <= taskCount; i++ {
		tasks = append(tasks, Task{Name: fmt.Sprintf("Task %d", i), Input: i, WorkFunc: square})
	}
	tasksCh := startTasks(tasks)

	// fan out tasks
	workerCount := len(tasks)
	taskResultCh := make([]<-chan Task, 0, workerCount)
	for i := 0; i < workerCount; i++ {
		taskResultCh = append(taskResultCh, startWorker(tasksCh, i))
	}

	// fan in tasks
	resultCh := collect(taskResultCh)
	for result := range resultCh {
		fmt.Printf("-> %s is completed, input = %d, output = %d\n", result.Name, result.Input, result.Output)
	}
}
