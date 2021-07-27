package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	chTasks := make(chan Task)
	chErrSignal := make(chan struct{}, len(tasks))
	done := make(chan struct{})

	careAboutErrors := m > 0

	var errCounter int
	go func(care bool) {
		defer close(chTasks)
		for _, task := range tasks {
			select {
			case <-chErrSignal:
				if care {
					errCounter++
				}
			default:
			}

			if care && errCounter >= m {
				close(done)
				return
			}
			chTasks <- task
		}
	}(careAboutErrors)

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(done, chTasks, chErrSignal, &wg)
	}
	wg.Wait()

	close(chErrSignal)
	if careAboutErrors && errCounter >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(done <-chan struct{}, tasks <-chan Task, errSignals chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
		}

		select {
		case <-done:
			return
		case t, ok := <-tasks:
			if !ok {
				return
			}
			err := t()
			if err != nil {
				errSignals <- struct{}{}
				break
			}
		}
	}
}
