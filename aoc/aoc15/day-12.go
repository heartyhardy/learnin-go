package aoc15

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const red = "red"

func readJSONData(filename string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return data
}

func sumAllNums(contents []byte) int {
	sum := 0
	re := regexp.MustCompile(`(-)*([0-9]+)`)
	nums := re.FindAll(contents, -1)
	for _, num := range nums {
		inum, _ := strconv.Atoi(string(num))
		sum += inum
	}
	return sum
}

func walk(m interface{}) int {
	sum := 0
	if _, ok := m.(string); ok {
	} else if num, ok := m.(float64); ok {
		sum += int(num)
	} else if arr, ok := m.([]interface{}); ok {
		for _, v := range arr {
			sum += walk(v)
		}
	} else if mm, ok := m.(map[string]interface{}); ok {
		for _, v := range mm {
			if str, ok := v.(string); ok && str == "red" {
				return 0
			}
			sum += walk(v)
		}
	} else {
		panic(fmt.Sprint("unhandled ", m, " ", reflect.TypeOf(m)))
	}
	return sum
}

func sumAllNumsNoReds(t interface{}) int {
	sum := 0
	if _, ok := t.(string); ok {
	} else if num, ok := t.(float64); ok {
		sum += int(num)
	} else if jsarr, ok := t.([]interface{}); ok {
		for _, jsv := range jsarr {
			sum += sumAllNumsNoReds(jsv)
		}
	} else if jsmap, ok := t.(map[string]interface{}); ok {
		for _, jsmv := range jsmap {
			if jstr, ok := jsmv.(string); ok && strings.ToLower(jstr) == red {
				return 0
			}
			sum += sumAllNumsNoReds(jsmv)
		}
	} else {
		panic("Error occured while parsing!")
	}
	return sum
}

//Run Solution
func Run() {
	contents := readJSONData("./inputs/day-12.txt")
	sum := sumAllNums(contents)
	var master interface{}
	json.Unmarshal(contents, &master)
	sumNoReds := sumAllNumsNoReds(master)

	fmt.Println("\n-- Day 12: JSAbacusFramework.io --")
	fmt.Printf("\n➕     Sum up all: %26v \n➕  🔴  Sum NO REDS: %24v\n", sum, sumNoReds)
	fmt.Println(`
	______                 
	|  _  \                
	| | | |___  _ __   ___ 
	| | | / _ \| '_ \ / _ \
	| |/ / (_) | | | |  __/
	|___/ \___/|_| |_|\___|`)
	fmt.Println("")

}
