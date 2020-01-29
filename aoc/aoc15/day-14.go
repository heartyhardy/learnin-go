package aoc15

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type raindeers []*rainDeer

type rainDeer struct {
	name      string
	isRunning bool
	runtime   int
	speed     int
	sleep     int
	distance  int
	points    int
}

func readRacerStats(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	trimmed := strings.Trim(massReplace(data), "")
	content := strings.Split(trimmed, "\n")
	return content
}

func massReplace(data []byte) string {
	replaced := strings.ReplaceAll(string(data), "can fly", "")
	replaced = strings.ReplaceAll(replaced, "but then must rest", "")
	replaced = strings.ReplaceAll(replaced, "seconds", "")
	replaced = strings.ReplaceAll(replaced, "for", "")
	replaced = strings.ReplaceAll(replaced, "km/s", "")
	replaced = strings.ReplaceAll(replaced, ".", "")
	replaced = strings.ReplaceAll(replaced, ",", "")
	return replaced
}

func collectRacerStats(content []string) raindeers {
	racers := make([]*rainDeer, len(content))
	for i, line := range content {
		fields := strings.Fields(line)
		racer := new(rainDeer)
		racer.name = fields[0]
		racer.speed, _ = strconv.Atoi(fields[1])
		racer.runtime, _ = strconv.Atoi(fields[2])
		racer.sleep, _ = strconv.Atoi(fields[3])
		racers[i] = racer
	}
	return racers
}

func letItRace(deers raindeers, dt int) rainDeer {
	for _, raindeer := range deers {
		tsegments := dt / (raindeer.runtime + raindeer.sleep)
		tremainder := dt % (raindeer.runtime + raindeer.sleep)
		di := raindeer.runtime * raindeer.speed * tsegments
		if dr := (tremainder / raindeer.runtime); dr > 0 {
			raindeer.distance = di + (1 * raindeer.speed * raindeer.runtime)
			continue
		}
		raindeer.distance = di + (tremainder%raindeer.runtime)*raindeer.speed
	}
	sort.Slice(deers, func(i, j int) bool {
		return deers[i].distance > deers[j].distance
	})
	return *deers[0]
}

func letItRacePerSec(deers raindeers, dt int) rainDeer {
	for i := 1; i <= dt; i++ {
		for _, raindeer := range deers {
			tsegments := i / (raindeer.runtime + raindeer.sleep)
			tremainder := i % (raindeer.runtime + raindeer.sleep)
			di := raindeer.runtime * raindeer.speed * tsegments
			if dr := (tremainder / raindeer.runtime); dr > 0 {
				raindeer.distance = di + (1 * raindeer.speed * raindeer.runtime)
				continue
			}
			raindeer.distance = di + (tremainder%raindeer.runtime)*raindeer.speed
		}
		sort.Slice(deers, func(i, j int) bool {
			return deers[i].distance > deers[j].distance
		})
		for k := 0; k < len(deers); k++ {
			if deers[k].distance != deers[0].distance {
				break
			}
			deers[k].points++
		}
	}
	sort.Slice(deers, func(i, j int) bool {
		return deers[i].distance+deers[i].points > deers[j].distance+deers[j].points
	})
	return *deers[0]
}

//Run Solution
func Run() {
	content := readRacerStats("./inputs/day-14.txt")
	raindeerRacers := collectRacerStats(content)
	winner := letItRace(raindeerRacers, 2503)
	newWinner := letItRacePerSec(raindeerRacers, 2503)

	fmt.Printf("\n\t -- Day 14: Reindeer Olympics --\n")
	fmt.Printf("\n\t🦌  %-10v ran \u001b[33m%20v Km\u001b[0m \n", winner.name, winner.distance)
	fmt.Printf("\n\t🦌  %-10v ran %20v Km and earned \u001b[33m%v points!\u001b[0m \n", newWinner.name, newWinner.distance+newWinner.points, newWinner.points)
	fmt.Println(`
	______
	|  _  \
	| | | |___  _ __   ___
	| | | / _ \| '_ \ / _ \
	| |/ / (_) | | | |  __/
	|___/ \___/|_| |_|\___|`)
	fmt.Println("")
}
