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






//after some changes the code will look like this for prics.go


package prices

import (
	"example/practice/conversion"
	"example/practice/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"Tax_Rate"`
	Inputprice       []float64           `json:"Input_price"`
	TaxIncludedPrice map[string]string   `json:"Tax_Included_Price"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringtoFloat(lines)

	if err != nil {
		return err
	}

	job.Inputprice = prices
	return nil

}
//this function have included the go routines

func (job *TaxIncludedPriceJob) Process(donechan chan bool){
	err := job.LoadData()
	if err != nil {
		//return err
	}

	result := make(map[string]string)
	for _, rate := range job.Inputprice {
		taxIncludedPrice := rate * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", rate)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrice = result
	job.IOManager.WriteResult(job)
	donechan<-true //go routine will send a signal to the channel whhen the job is done 

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:  iom,
		Inputprice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}
