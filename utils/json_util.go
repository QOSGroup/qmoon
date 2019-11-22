package utils

import (
	"strings"
)


func FindAddressInJson(jsonStr string) ([]string) {
	//addrTags := []string{"address", "sender", "receiver", "addr", "operator", "validator", "proposer", "depositor", "voter"}
	addrPrifixes := []string{"\"qosval", "\"qosacc", "\"qoscons"}
	jsonStr = strings.Trim(jsonStr, "\n\t ")
	result := make([]string, 0)
	for _, prif := range addrPrifixes {
		ll := strings.Split(jsonStr,prif)
		if len(ll) <= 1 {
			continue
		}
		for i := 1; i<len(ll); i ++ {
			end := strings.Index(ll[i], "\"")
			if end<=0 {
				continue
			}
			result = append(result, prif[1:] + ll[i][:end])
		}
	}
	return result
}
