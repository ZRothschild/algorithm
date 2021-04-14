package main

import (
	"fmt"
)

type MapStr struct {
	Name string
}

func main() {

	// var (
	// 	mapTest = map[string]int64{"A": 1, "B.A": 2, "B.B": 3, "CC.D.E": 4, "CC.D.F": 5}
	// )
	// {
	// 	"A": 1,
	// 	"B": {
	// 		"A": 2,
	// 		"B": 3,
	// 	},
	// 	"CC": {
	// 		"D": {
	// 			"E": 4,
	// 			"F": 5
	// 		}
	// 	}
	// }

	// var maxStr []string
	// // var tmpStr string
	// slices := make(map[string][]string)
	// for k := range mapTest {
	// 	slices[k] = strings.Split(k,".")
	// 	//  sliceStr := strings.Split(k,".")
	//
	// 	for _ ,val := range slices[k] {
	// 		mapVal[val] = k
	//
	//
	// 	}
	// }

	str := []string{"A", "B", "C"}
	mapTmp := make(map[string]interface{})
	mapVal := make(map[string]interface{})
	for k, v := range str {
		// mapVal[v]= &k
		if k == 0 {
			mapTmp[v] = k
		} else {

			mapTmp = Test(mapTmp, v)
			// mapVal[v] = k
			// mapTmp[str[k-1]] = mapVal
			//
			// mapVal[v] = "GG"
			// // mapTmp = mapTmp[str[k-1]]
		}
	}

	// if len(mapKeys) == 0 {
	// 	mapVal[mapKeys[0]] = mapTest[k]
	// }else {
	//
	// }

	fmt.Printf("%#v", mapTmp)
}

func Test(mapTmp map[string]interface{}, v string, k int64) map[string]interface{} {
	mapTmp[str[k-1]]
}
