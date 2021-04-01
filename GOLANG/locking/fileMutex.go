package locking

import (
	"fmt"
	"sync"
)

type OnceLock struct {
	sync.Once
	alpha int
}

func TryOnceLocking() {
	done := make(chan bool)

	for alpha := 1; alpha < 6; alpha++ {
		alphaOnceLock := new(OnceLock)
		alphaOnceLock.alpha = alpha

		for i := 0; i < 1000; i++ {
			go alphaOnceLock.Do(func() { doTheThing(alphaOnceLock.alpha, i); done <- true })
		}
	}

	for i := 1; i < 6; i++ {
		<-done
	}
}

func doTheThing(alpha int, instance int) {
	fmt.Printf("##### [%s] executed by [%d]\n", string('A'-1+alpha), instance)
}
