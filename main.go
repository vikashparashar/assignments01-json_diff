package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("_____________________________________________________________")
	fmt.Println("Difference Between New Old Json And New Json")
	fmt.Println("List Of Keys Missing In Both Json")
	oldkey := JsonToSlice("old.json")
	newkey := JsonToSlice("new.json")
	diffStr := difference(oldkey, newkey)

	for _, diffVal := range diffStr {
		fmt.Println("Missing Key : ", diffVal)
	}
	fmt.Println("_____________________________________________________________")
	fmt.Println("Difference Between New TestJson1 And TestJson2")
	fmt.Println("List Of Keys Missing In Both Json")
	testkey1 := JsonToSlice("testjson1.json")
	testkey2 := JsonToSlice("testjson2.json")
	diffStr2 := difference(testkey1, testkey2)

	for _, diffVal2 := range diffStr2 {
		fmt.Println("Missing Key : ", diffVal2)
	}
	fmt.Println("_____________________________________________________________")

}

func difference(s []string, s2 []string) []string {
	diffStr := []string{}
	m := map[string]int{}

	for _, v := range s {
		m[v] = 1
	}
	for _, k := range s2 {
		m[k] = m[k] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return diffStr
}

func JsonToSlice(filename string) []string {
	var oldmap = map[string]string{}
	// var newmap = map[string]string{}
	var strkey []string
	// var newkey []string
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println("failed to read data from file")
	}
	// new, err := os.ReadFile("new.json")
	// if err != nil {
	// 	log.Println("failed to read data from file")
	// }

	if err = json.Unmarshal(data, &oldmap); err != nil {
		log.Println("failed to unmarshal data ")
	}
	// if err = json.Unmarshal(new, &newmap); err != nil {
	// 	log.Println("failed to unmarshal data ")
	// }

	// fmt.Println(oldmap["a"])
	// for i, _ := range newmap {
	// 	newkey = append(newkey, i)

	// }
	for j, _ := range oldmap {
		strkey = append(strkey, j)
	}

	return strkey

}
