package workerpool

import "sync"

type Result[T any] struct {
	Value T
	Err   error
}

type In[T any] chan<- func() (T, error)
type Out[T any] <-chan Result[T]

// Start creates a worker pool with n number of workers.
// Workers will process the functions sent to the input channel.
// The result of the function call will be sent to the output channel.
func Start[T any](numberOfWorkers int) (In[T], Out[T]) {
	in := make(chan func() (T, error))
	out := make(chan Result[T])

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		// start a worker in a goroutine
		go func() {
			for fn := range in {
				v, err := fn()
				out <- Result[T]{Value: v, Err: err}
			}
			wg.Done()
		}()
	}

	// close the output channel when the processing has finished
	go func() {
		wg.Wait()
		close(out)
	}()

	return in, out
}
