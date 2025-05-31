package main

import (
	"example/practice/prices"
)

func main() {

	Taxrate := []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range Taxrate {
		JobPrice := prices.NewTaxIncludedPriceJob(rate)
		JobPrice.Process()
	}

}
//after using the cmdmanager package and iomanager package the program will look like this 


package main

import (
	"example/practice/filehandling"
	"example/practice/prices"
	"fmt"
)

func main() {

	Taxrate := []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range Taxrate {
		fh := filehandling.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		//cmdm := cmdmanager.New()
		JobPrice := prices.NewTaxIncludedPriceJob(fh, rate)
		err := JobPrice.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}

}




