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

	//var tasksOfJob []Task
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
	//time.Sleep(2 * time.Second)
	fmt.Println("Finish Count of active go routines in main = ", runtime.NumGoroutine())
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

// Создание мапы с номерами "ошибочных" значений индекса (ключа)
// func fillErrIndexMap() map[int]int {
// 	errIndexMap := make(map[int]int)
// 	errIndexMap[3] = 1
// 	errIndexMap[4] = 1
// 	errIndexMap[5] = 1

// 	return errIndexMap
// }

// Функция (только для дебага) обработки задач. 7-ая и 8-ая задачи обрабатываются с ошибкой
func taskTreatmentFunc() []Task {
	var tasks []Task
	for k := 0; k < 10; k++ {
		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))

		tasks = append(tasks, func() error {
			time.Sleep(taskSleep)
			//atomic.AddInt32(&runTasksCount, 1)
			if k == 7 || k == 8 {
				err := fmt.Errorf("error from task %d", k)
				return err
			} else {
				return nil
			}
		})
	}
	return tasks
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	fmt.Println("Start Count of active go routines = ", runtime.NumGoroutine())
	// Создадим канал обрабатываемых задач и заполним его
	ch := generator(tasks)
	fmt.Println("Канал=", ch)

	// Создадим канал для остановки данной функции из рутин
	stopCh := make(chan bool)

	//errIndexMap := fillErrIndexMap()
	errCount := 0

	wg := sync.WaitGroup{}
	wg.Add(n)
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		// Функция обработки задач из канала
		go func(c <-chan Task, stopChRead <-chan bool, stopChWrite chan<- bool) error {
			//fmt.Println("Iteration i = " + strconv.Itoa(i) + " n=" + strconv.Itoa(n))
			wg.Done()
			for {
				select {
				default:
					// Выполнение работы в горутине
					//defer wg.Done()
					fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")
					// Если индекс относится к одному из списка "ошибочных", то выходим из функции
					isErrorIndex := true // errIndexMap[i]
					if isErrorIndex {
						mu.Lock()
						errCount++
						mu.Unlock()

						fmt.Println("errCount = " + strconv.Itoa(errCount))

						return nil
					}

					if errCount >= m {
						stopChWrite <- true
						fmt.Println("Отправлен сигнал об остановке")
						return ErrErrorsLimitExceeded
					}

					// Обработка задач канала
					for t := range c {
						fmt.Println("i = "+strconv.Itoa(i)+" Treatment of Task. Task=", t)
						//fmt.Println("i = "+strconv.Itoa(i)+" Count of active go routines = ", runtime.NumGoroutine())
					}
				case <-stopChRead:
					// Получен сигнал об остановке
					//defer wg.Done()
					//fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")
					//fmt.Println("Остановлен")
					return ErrErrorsLimitExceeded
				}
			}
		}(ch, stopCh, stopCh)

	}
	wg.Wait()

	//time.Sleep(2 * time.Second)

	fmt.Println("Finish Count of active go routines in function Run = ", runtime.NumGoroutine())
	return nil
}
