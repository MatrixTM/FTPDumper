package Core

import (
	"context"
	"sync"
	"time"
)

// Task represents a unit of work to be executed by the worker pool.
type Task func()

// WorkerPool is a struct that manages a pool of worker goroutines.
type WorkerPool struct {
	maxWorkers int
	taskQueue  chan Task
	wg         sync.WaitGroup
}

// New creates a new WorkerPool with the specified maximum number of workers.
func New(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		taskQueue:  make(chan Task),
	}
}

// Start initializes the worker pool and starts the worker goroutines.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for task := range wp.taskQueue {
				task()
			}
		}()
	}
}

// Stop gracefully shuts down the worker pool.
func (wp *WorkerPool) Stop() {
	close(wp.taskQueue)
	wp.wg.Wait()
}

// Submit adds a new task to the worker pool's task queue.
func (wp *WorkerPool) Submit(task Task) {
	wp.taskQueue <- task
}

// IsFull returns true if the worker pool's task queue is full.
func (wp *WorkerPool) IsFull() bool {
	return len(wp.taskQueue) >= cap(wp.taskQueue)
}

// SubmitWithTimeout adds a new task to the worker pool's task queue with a timeout.
// If the task is not completed within the specified timeout duration, it is canceled.
func (wp *WorkerPool) SubmitWithTimeout(task Task, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})

	wp.taskQueue <- func() {
		defer close(done)
		task()
	}

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
