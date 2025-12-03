package math

const PI float64 = 3.14

func Sum[T int | float64](a, b T) T {
	return a + b
}
