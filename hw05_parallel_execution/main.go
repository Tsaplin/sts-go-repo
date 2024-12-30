package main

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func main() {
	fmt.Println("hw05_parallel_execution - main start")

	var tasksOfJob []Task
	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)

	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)

	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)

	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)
	tasksOfJob = append(tasksOfJob, nil)

	Run(tasksOfJob, 5, 2)
	fmt.Println("hw05_parallel_execution - main finish")
}

// Создание канала обрабатываемых задач и его заполнение
func generator(tasks []Task) <-chan Task {
	c := make(chan Task)

	go func() {
		for _, task := range tasks {
			c <- task
		}
		close(c)
	}()

	return c
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	fmt.Println("Start Count of active go routines = ", runtime.NumGoroutine())
	// Создадим канал обрабатываемых задач и заполним его
	ch := generator(tasks)
	fmt.Println("Канал=", ch)

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		// Функция обработки задач из канала
		go func(c <-chan Task) {
			for t := range c {
				fmt.Println("i = "+strconv.Itoa(i)+" Treatment of Task. Task=", t)
				fmt.Println("i = "+strconv.Itoa(i)+" Count of active go routines = ", runtime.NumGoroutine())
			}
			defer wg.Done()
		}(ch)
	}
	wg.Wait()

	//int countActiveRoutines = runtime.NumGoroutine();
	fmt.Println("Finish Count of active go routines = ", runtime.NumGoroutine())
	return nil
}
