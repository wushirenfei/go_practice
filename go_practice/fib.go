// :date 2021/7/12

package go_practice

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}
