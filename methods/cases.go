package methods

import "fmt"

func Cases(n int) string {
	remainder := n % 10

	var res string
	switch remainder {
	case 1:
		res = fmt.Sprintf("%d компьютер", n)
	case 2, 3, 4:
		res = fmt.Sprintf("%d компьютера", n)
	case 5, 6, 7, 8, 9:
		res = fmt.Sprintf("%d компьютеров", n)
	}

	return res
}
