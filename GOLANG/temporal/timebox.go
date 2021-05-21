package temporal

import (
	"fmt"
	"time"
)

func FindEndOfMonth() {
	year := 2019
	for month := 1; month <= 12; month++ {
		testDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

		testDate = testDate.Add(time.Hour * 24 * 28)

		for testDate.Month() == time.Month(month) {
			testDate = testDate.Add(time.Hour * 24)
		}

		testDate = testDate.Add(time.Hour * -24)

		fmt.Printf("The last day of [%s] is [%s]\n", time.Month(month).String(), testDate.Format("2006-01-02"))
	}
}
