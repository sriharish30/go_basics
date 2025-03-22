package file

import (
	"fmt"
	"os"
)

func Writefiles(Ebt, profit, ratio float64) {
	results := fmt.Sprintf("Ebt:%.1f\nprofit:%.2f\nratio:%.3f\n", Ebt, profit, ratio)
	os.WriteFile("Result.txt", []byte(results), 0644)
}
