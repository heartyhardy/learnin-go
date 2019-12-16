package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// //Resource struct
// type Resource struct {
// 	id  string
// 	val int
// }

// // func parse(file string) map[string][]string {
// // 	data, _ := ioutil.ReadFile(file)
// // 	sfunc := strings.Split(strings.TrimSpace(string(data)), "\n")
// // 	prodmap := map[string][]string{}
// // 	prodqty := map[string]int{}
// // 	for _, funct := range sfunc {

// // 		sparams := strings.Split(funct, " => ")
// // 		resource := strings.Split(sparams[len(sparams)-1], " ")
// // 		resid := resource[1]
// // 		resVal, _ := strconv.Atoi(resource[0])
// // 		prodqty[resid] = resVal
// // 		params := []string{}

// // 		for _, param := range sparams[:len(sparams)-1] {

// // 			params = append(params, strings.Split(param, ",")...)
// // 			subparams := []string{}

// // 			for _, subparam := range params {

// // 				subparams = append(subparams, strings.Split(subparam, "  ")...)
// // 				subparamresources := []Resource{}
// // 				for _, subresource := range subparams {
// // 					subparamresources:= append()
// // 				}
// // 			}
// // 		}
// // 		prodmap[resid] = params
// // 	}
// // 	return prodmap
// // }

// func parse(file string) map[string][]Resource {
// 	data, _ := ioutil.ReadFile(file)
// 	remcommas := strings.ReplaceAll(string(data), ",", "")
// 	remarrow := strings.ReplaceAll(remcommas, " => ", "  ")
// 	sfunc := strings.Split(remarrow, "\n")

// 	resources := map[string][]Resource{}

// 	for _, sparam := range sfunc {
// 		sparam := strings.ReplaceAll(sparam, "  ", " ")
// 		sparams := strings.Split(sparam, " ")
// 		params := []Resource{}
// 		for i := 0; i < len(sparams)-1; i = i + 2 {
// 			val, _ := strconv.Atoi(sparams[i])
// 			resource := &Resource{id: sparams[i+1], val: val}
// 			params = append(params, *resource)
// 		}
// 		resources[sparams[len(sparams)-1]] = params
// 	}
// 	return resources
// }

// func swarm(wg *sync.WaitGroup, prod string, prodmap map[string][]Resource) {
// 	defer wg.Done()
// 	fmt.Println(prod)
// 	if params, ok := prodmap[prod]; ok && len(params) > 2 {
// 		//prodval := params[len(params)-1:][0].val
// 		params = params[:len(params)-1]

// 		for _, param := range params {
// 			wg.Add(1)
// 			go swarm(wg, param.id, prodmap)
// 		}
// 	} else {
// 		//fmt.Println(prod, prodmap[prod][len(prodmap[prod])-1].val)
// 		//prodmap[prod][len(prodmap[prod])-1].val
// 	}

// }

// func main() {
// 	var wg sync.WaitGroup
// 	resources := parse("day-14-input.txt")
// 	//fmt.Println(resources)
// 	wg.Add(1)
// 	go swarm(&wg, "FUEL", resources)
// 	wg.Wait()
// }
