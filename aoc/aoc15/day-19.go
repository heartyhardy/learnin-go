package aoc15

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const molecules string = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"

type replacementsList map[string][]string

func readReplacements(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func mapReplacements(reps []string) replacementsList {
	replacements := make(map[string][]string)
	for _, rep := range reps {
		fields := strings.FieldsFunc(rep, func(c rune) bool {
			return !unicode.IsLetter(c)
		})
		if _, ok := replacements[fields[0]]; !ok {
			replacements[fields[0]] = make([]string, 1)
			replacements[fields[0]][0] = fields[1]
		} else {
			replacements[fields[0]] = append(replacements[fields[0]], fields[1])
		}
	}
	return replacements
}

func findAllIndices(src, sub string) (indices []int) {
	offset := 0
	index := strings.Index(src, sub)
	for index > -1 {
		indices = append(indices, index+offset)
		offset += len(src[:index+len(sub)])
		src = src[index+len(sub):]
		index = strings.Index(src, sub)
	}
	return
}

func replaceAll(replacements replacementsList) int {
	uniques := make(map[string]string)
	for rk, rep := range replacements {
		indices := findAllIndices(molecules, rk)

		for _, index := range indices {
			prefix := molecules[:index]
			suffix := molecules[index+len(rk):]
			for _, comb := range rep {
				newMolecule := prefix + comb + suffix
				if _, ok := uniques[newMolecule]; !ok {
					uniques[newMolecule] = newMolecule
				}
			}
		}
	}
	return len(uniques)
}

//Run solution
func Run() {
	reps := readReplacements("./inputs/day-19.txt")
	replacements := mapReplacements(reps)
	combinations := replaceAll(replacements)
	fmt.Println(combinations)
}
