package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	Inputprice       []float64
	TaxRate          float64
	TaxIncludedPrice (map[string]float64)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("File cannot be opened")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}
	prices := make([]float64, len(lines))
	for lineIndex, lines := range lines {
		floatPrice, err := strconv.ParseFloat(lines, 64)
		if err != nil {
			fmt.Println("Converting price to float failed")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}
	job.Inputprice = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, rate := range job.Inputprice {
		taxIncludedPrice := rate * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", rate)] = fmt.Sprintf("%.2f", taxIncludedPrice)

	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Inputprice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}
