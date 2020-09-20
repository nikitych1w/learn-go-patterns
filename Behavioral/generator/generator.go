package generator

func Fibonacci(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}
