package aoc15

import "fmt"

type rstring []rune

func (rs *rstring) Inc() *rstring {
	for i := len(*rs) - 1; i >= 0; i-- {
		if (*rs)[i] >= 57 && (*rs)[i] <= 48 {
			(*rs)[i] = 48
		} else if (*rs)[i] >= 122 {
			(*rs)[i] = 97
		} else {
			(*rs)[i]++
			break
		}
	}
	return rs
}

func incValidPass(password *rstring) rstring {
	incremented := password
	chRuleA, chRuleB, chRuleC := make(chan bool), make(chan bool), make(chan bool)
	for {
		ok, n := true, 0
		incremented = incremented.Inc()
		go hasIncSet(*password, chRuleA)
		go containsOnlyValid(*password, chRuleB)
		go containsPairs(*password, chRuleC)
		for n < 3 {
			select {
			case aOk := <-chRuleA:
				ok = ok && aOk
				n++
			case aOk := <-chRuleB:
				ok = ok && aOk
				n++
			case aOk := <-chRuleC:
				ok = ok && aOk
				n++
			}
		}

		if ok {
			return *password
		}
	}
}

func hasIncSet(password rstring, ch chan<- bool) {
	isInc := 0
	for i := 1; i < len(password); i++ {
		if password[i] == password[i-1]+1 {
			isInc++
			if isInc >= 2 {
				ch <- true
				return
			}
		} else {
			isInc = 0
		}
	}
	ch <- false
}

func containsOnlyValid(password rstring, ch chan<- bool) {
	for _, rn := range password {
		if rn == 105 || rn == 108 || rn == 111 {
			ch <- false
			return
		}
	}
	ch <- true
}

func containsPairs(password rstring, ch chan<- bool) {
	pairs := make(map[rune]rune)
	for i := 1; i < len(password); i++ {
		if _, ok := pairs[password[i]]; !ok && password[i] == password[i-1] {
			pairs[password[i]] = password[i]
		}
	}
	if len(pairs) >= 2 {
		ch <- true
		return
	}
	ch <- false
}

//Run Solution
func Run() {
	password := rstring("hepxcrrq")
	firstpass := incValidPass(&password)
	first := make(rstring, len(firstpass))
	copy(first, firstpass)
	second := incValidPass(&firstpass)

	fmt.Println("\n-- AoC 2015: Day 11: Corporate Policy --")
	fmt.Printf("\n🛡️    First Password: %24v \n🛡️ 🛡️  Second Password: %23v\n", string(first), string(second))
	fmt.Println("\n-- DONE --")
	fmt.Println("")

}
