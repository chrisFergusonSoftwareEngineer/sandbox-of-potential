package json

import (
	"encoding/json"
	"fmt"
)

type LimitedFields struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

const JsonPayload = "{\"name\":\"Fergus\",\"id\":1,\"extraField\":\"Spark\"}"

func ParsePayloadWithExtraField() {
	var result LimitedFields

	err := json.Unmarshal([]byte(JsonPayload), &result)

	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("no errors")
	}
	fmt.Println()

	fmt.Printf("result: %v\n", result)
}
