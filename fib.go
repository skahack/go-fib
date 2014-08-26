package fib

var memo map[int]int

func Fib(n int) <-chan int {
	ch := make(chan int)

	go func() {
		i := int(0)
		for {
			if n > 0 && i >= n+1 {
				break
			}
			ch <- fib(i)
			i++
		}
		defer close(ch)
	}()

	return ch
}

func fib(n int) int {
	v, ok := memo[n]
	if ok {
		return v
	}

	if n < 2 {
		memo[n] = n
	} else {
		memo[n] = fib(n-1) + fib(n-2)
	}

	return memo[n]
}

func init() {
	memo = make(map[int]int)
	memo[0] = 0
	memo[1] = 1
}
