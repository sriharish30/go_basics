package prices

import (
	"example/practice/conversion"
	"example/practice/filehandling"
	"fmt"
)

type TaxIncludedPriceJob struct {
	Inputprice       []float64
	TaxRate          float64
	TaxIncludedPrice (map[string]string)
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filehandling.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringtoFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
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

	//after using the json file this [fmt.Println(result)] line can be deleted and they change like this

	job.TaxIncludedPrice = result
	filehandling.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Inputprice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}
