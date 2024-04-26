package methods

import (
	"fmt"
	"strconv"
)

func MultiplicationTable(n int) {
	table := make([][]string, n)
	for i := 0; i < n; i++ {
		table[i] = make([]string, n)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			table[i-1][j-1] = strconv.Itoa(i * j)
		}
	}

	for i := 1; i <= 1; i++ {
		fmt.Print(" ")

		for j := 0; j < len(table[0][n-1])-len(table[0][i-1]); j++ {
			fmt.Print(" ")
		}

		for j := 0; j < n; j++ {
			spaceAmount := len(table[len(table)-1][j]) - len(table[i-1][j])

			for k := 0; k < spaceAmount; k++ {
				fmt.Print(" ")
			}

			fmt.Printf(" %s", table[i-1][j])
		}
	}

	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%d", i)
		for j := 0; j < len(table[0][n-1])-len(table[0][i-1]); j++ {
			fmt.Print(" ")
		}

		for j := 0; j < n; j++ {
			spaceAmount := len(table[len(table)-1][j]) - len(table[i-1][j])

			for k := 0; k < spaceAmount; k++ {
				fmt.Print(" ")
			}

			fmt.Printf(" %s", table[i-1][j])
		}

		fmt.Println()
	}
}
