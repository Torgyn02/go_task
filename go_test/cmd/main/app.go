package main

import "fmt"

func main() {
	// 	api := SdekAPI{
	// 		Token:      "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI",
	// 		TestMode:   true,
	// 		APIAddress: "https://api.cdek.ru/v2/oauth/token?parameters",
	// 	}
	// 	addrFrom := Address{
	// 		CountryCode: "RU",
	// 		Postcode:    "101000",
	// 		City:        "Москва",
	// 		Street:      "Cлавянский бульвар",
	// 		House:       "д.43",
	// }
	//    addrTo  :=

	addrFrom := " Россия, г.Москва, Cлавянский бульвар д.1"
	addrTo := "Россия, Воронежская обл., г.Воронеж, ул.Ленина д.43"
	size := Size{Weight: 5.0, Length: 10.0, Width: 15.0, Height: 20.0}
	auth := Auth{Token: "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI", IsTest: true, APIUrl: "https://api.edu.cdek.ru/v2/oauth/token?parameters"}

	// Call the function
	result, err := Calculate(addrFrom, addrTo, size, auth)

	// Check for errors
	if err != nil {
		fmt.Printf("Error calculating shipping: %s\n", err.Error())
		return
	}

	fmt.Printf("Shipping cost:\n")
	for _, fareCode := range result.FareCodes {
		fmt.Printf("  %s - %.2f rubles, delivery in %d to %d days\n", fareCode.TariffName, fareCode.DeliveryAmount, fareCode.PeriodMin, fareCode.PeriodMax)
	}
}
