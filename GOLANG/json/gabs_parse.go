package json

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func ParseJsonWithGabs() {
	//algDefRaw := d.Get("algorithm_definition").(string)
	algDefRaw := "{\"name \":  \"archival-algorithm \", 		 \"numConcurrent \": 0, 		 \"numCpus \": 1, 		 \"numCpusArray \": null, 		 \"outputData \": [], 		 \"permissions \": { 		   \"delete \": [], 		   \"read \": [], 		   \"write \": [] 		}, 		 \"priority \": 7, 		 \"tag \": [], 		 \"taskGroup \": null, 		 \"timeout \": 0, 		 \"version \":  \"1.0.0 \", 		 \"containerPorts \": null, 		 \"skipDedupCheck \": false, \"dedupWindow \":  \" \"}"
	algDef, _ := gabs.ParseJSON([]byte(algDefRaw))

	algName := algDef.Path("name").Data().(string)
	algVersion := algDef.Path("version").Data().(string)

	fmt.Printf("Create algorithm: installing %s, version %s", algName, algVersion)
}
