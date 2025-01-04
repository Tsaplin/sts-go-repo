package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func main() {
	fmt.Println("hw05_parallel_execution - main start")

	tasksOfJob := taskPrepareFunc()
	Run(tasksOfJob, 5, 2)
	// time.Sleep(2 * time.Second)
	fmt.Println("Finish Count of active go routines in main = ", runtime.NumGoroutine())
	fmt.Println("hw05_parallel_execution - main finish")
}

// Создание канала обрабатываемых задач и его заполнение.
func generator(tasks []Task, channelCapacity int) chan Task {
	c := make(chan Task, channelCapacity)

	go func() {
		for _, task := range tasks {
			c <- task
		}
		close(c)
	}()

	return c
}

// Функция (только для дебага) подготовки массива задач. 7-ая и 8-ая задачи обрабатываются с ошибкой.
func taskPrepareFunc() []Task {
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
	ch := generator(tasks, len(tasks)+2)
	// Создадим канал для остановки данной функции из рутин.
	// Сделаем его буферизованным с кол-вом элементов больше кол-ва одновременно работающих горутин,
	// чтобы избежать блокировки горутин
	stopCh := make(chan bool, n+2)
	// Счетчик кол-ва ошибок обработки задач
	errCount := 0

	wg := sync.WaitGroup{}
	wg.Add(n)
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		// Функция обработки задач из канала
		go func(taskCh chan Task, stopWorkCh chan bool) error {
			defer wg.Done()
			for {
				select {
				case tt, ok := <-taskCh:
					// fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")

					// fmt.Println("Task tt was readen from channel of tasks = ", tt)

					if !ok {
						return nil
					}

					// Если при обработке задачи возникла ошибка, то увеличим значение errCount
					var err = tt()
					if err != nil {
						mu.Lock()
						errCount++
						// Если превышего предельно допустимое кол-во ошибок обработки задач, отправим "сигнал" о завершении работы рутин
						if errCount >= m {
							stopWorkCh <- true
							// close(stopWorkCh) тут закрывать канал не будем, т.к. в этот блок можем и не попасть
							// fmt.Println("Отправлен сигнал об остановке")
							// return ErrErrorsLimitExceeded
						}
						mu.Unlock()
						// fmt.Println("errCount = " + strconv.Itoa(errCount))
					}
				case val := <-stopWorkCh:
					// Получен сигнал об остановке
					// fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")
					// fmt.Println("Остановлен")
					if val {
						return ErrErrorsLimitExceeded
					}
					return nil
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
