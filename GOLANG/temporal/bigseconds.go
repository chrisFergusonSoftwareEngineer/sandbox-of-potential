package temporal

import (
	"fmt"
	"time"
)

func TryBigSeconds() {
	now := time.Now().UTC()

	for i := 0; i < 86400; i += 1234 {
		testTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, i, 0, time.UTC)

		fmt.Println(testTime.Format("2006-01-02 T 15:04:05"))
	}
}
