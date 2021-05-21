package builtin

import (
	"fmt"
	"net/url"
)

func ParseTest() {
	u, err := url.Parse("http://192.168.0.2:1234")

	if nil != err {
		fmt.Printf(err.Error())
	}

	fmt.Printf(u.Hostname() + "   " + u.Port())
}
