package utils

import (
	"bufio"
	"os"
)

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func Pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func GetFactorsExceptSelf(n int) []int {
	factors := []int{1}

	for i := 2; i*i <= n; i++ {
		if n%i==0 {
			factors = append(factors,i)
			if n/i != i {
				factors = append(factors,n/i)
			}
		}
	}
	return factors
}
