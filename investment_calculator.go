package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationrate = 6.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	year := 10.0

	fmt.Print("investmentAmount :")
	fmt.Scan(&investmentAmount)

	fmt.Print("expected return rate :")
	fmt.Scan(&expectedReturnRate)

	futurevalue := investmentAmount * math.Pow(1+expectedReturnRate/100, year)
	realfuturevalue := futurevalue / math.Pow(1+inflationrate/100, year)
	fmt.Println(futurevalue)
	fmt.Println(realfuturevalue)

}
