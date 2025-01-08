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

// Функция (только для дебага) подготовки массива задач. 7-ая и 8-ая задачи обрабатываются с ошибкой.
func taskPrepareFunc() []Task {
	var tasks []Task
	// var runTasksCount int32
	for k := 0; k < 10; k++ {
		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))

		tasks = append(tasks, func() error {
			time.Sleep(taskSleep)
			// atomic.AddInt32(&runTasksCount, 1)
			if k == 2 || k == 3 || k == 7 || k == 8 {
				err := fmt.Errorf("error from task %d", k)
				return err
			}
			return nil
		})
	}
	return tasks
}

// Функция меняет значение кода ошибки, если код отличен от nil.
func saveErrCode(mu *sync.Mutex, errCode error) error {
	mu.Lock()
	if errCode == nil {
		errCode = ErrErrorsLimitExceeded
	}
	mu.Unlock()
	return errCode
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	// Создадим канал обрабатываемых задач
	ch := make(chan Task)
	// Индекс отправленных в канал задач
	index := 0
	// Кол-во задач в массиве
	taskCnt := len(tasks)
	// Счетчик кол-ва ошибок обработки задач
	errCount := 0
	// Код ошибки функции Run
	var errCode error
	errCode = nil

	wg := sync.WaitGroup{}
	wg.Add(n)
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		// Функция обработки задач из канала
		go func(taskCh chan Task) error {
			defer wg.Done()
			for {
				select {
				case tt, ok := <-taskCh:
					if !ok {
						return nil
					}

					// fmt.Println("i = " + strconv.Itoa(i) + " Exec of go routine")
					// fmt.Println("Task tt was readen from channel of tasks = ", tt)
					// Если при обработке задачи возникла ошибка, то увеличим значение errCount
					var err = tt()
					if err != nil {
						mu.Lock()
						errCount++
						mu.Unlock()
						// fmt.Println("errCount = " + strconv.Itoa(errCount))
					}
					// return nil
				default:
					// Если предельное кол-во ошибок достигнуто, закрываем канал и выдаем ошибку
					mu.Lock()
					if errCount >= m {
						mu.Unlock()
						close(taskCh)
						errCode = saveErrCode(&mu, errCode)
						// fmt.Println("errCode = ", errCode)
						return ErrErrorsLimitExceeded
					}
					mu.Unlock()

					mu.Lock()
					if index >= taskCnt {
						close(taskCh)
						return nil
					}
					mu.Unlock()
					// Отправляем задачу в канал, если предельное кол-во ошибок еще не достигнуто
					// fmt.Println("index = " + strconv.Itoa(i) + " Push task to channel")
					taskCh <- tasks[index]
					mu.Lock()
					index++
					mu.Unlock()
				}
			}
		}(ch)
	}
	wg.Wait()

	return errCode
}
