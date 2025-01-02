package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func main() {
	fmt.Println("hw05_parallel_execution - main start")

	// var tasksOfJob []Task
	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)

	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)

	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)

	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)
	// tasksOfJob = append(tasksOfJob, nil)

	tasksOfJob := taskTreatmentFunc()
	Run(tasksOfJob, 5, 1)
	// time.Sleep(2 * time.Second)
	fmt.Println("Finish Count of active go routines in main = ", runtime.NumGoroutine())
	fmt.Println("hw05_parallel_execution - main finish")
}

// Создание канала обрабатываемых задач и его заполнение.
func generator(tasks []Task) chan Task {
	c := make(chan Task)

	go func() {
		for _, task := range tasks {
			c <- task
		}
		close(c)
	}()

	return c
}

// Функция (только для дебага) обработки задач. 7-ая и 8-ая задачи обрабатываются с ошибкой.
func taskTreatmentFunc() []Task {
	var tasks []Task
	// var runTasksCount int32
	for k := 0; k < 10; k++ {
		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))

		tasks = append(tasks, func() error {
			time.Sleep(taskSleep)
			// atomic.AddInt32(&runTasksCount, 1)
			if k == 7 || k == 8 {
				err := fmt.Errorf("error from task %d", k)
				return err
			}
			return nil
		})
	}
	return tasks
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	fmt.Println("Start Count of active go routines = ", runtime.NumGoroutine())

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	// Создадим канал обрабатываемых задач и заполним его
	ch := generator(tasks)
	fmt.Println("Канал=", ch)

	// Создадим канал для остановки данной функции из рутин
	stopCh := make(chan bool)
	errCount := 0

	wg := sync.WaitGroup{}
	wg.Add(n)
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		// Функция обработки задач из канала
		go func(taskCh chan Task, stopWorkCh chan bool) error {
			// fmt.Println("Iteration i = " + strconv.Itoa(i) + " n=" + strconv.Itoa(n))
			wg.Done()
			for {
				select {
				default:
					// Выполнение работы в горутине

					// defer wg.Done()
					fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")
					tt, ok := <-taskCh
					fmt.Println("Task tt was readen from channel of tasks = ", tt)

					// Если при обработке задачи возникла ошибка, то увеличим значение errCount
					if !ok {
						mu.Lock()
						errCount++
						mu.Unlock()
						fmt.Println("errCount = " + strconv.Itoa(errCount))
					}

					if errCount >= m {
						stopWorkCh <- true
						// close(stopWorkCh)
						fmt.Println("Отправлен сигнал об остановке")
						return ErrErrorsLimitExceeded
					}

					return nil
				case <-stopWorkCh:
					// Получен сигнал об остановке
					// defer wg.Done()
					// fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")
					fmt.Println("Остановлен")
					return ErrErrorsLimitExceeded
				}
			}
		}(ch, stopCh)
	}
	wg.Wait()
	close(stopCh)

	time.Sleep(2 * time.Second)

	fmt.Println("Finish Count of active go routines in function Run = ", runtime.NumGoroutine())
	return nil
}
