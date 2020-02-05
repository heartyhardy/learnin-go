package aoc15

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	_ = iota
	weapon
	armor
	misc
)

const (
	_ = iota
	dmg
	def
)

type item struct {
	key      string
	val      int
	class    int
	cost     int
	category int
}

type gameObject struct {
	key   string
	hp    int
	dmg   int
	armor int
}

type shop map[int]*item

func readItemsList(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	replaced := strings.ReplaceAll(string(data), "Weapons:    Cost  Damage  Armor", "")
	replaced = strings.ReplaceAll(replaced, "Armor:      Cost  Damage  Armor", "")
	replaced = strings.ReplaceAll(replaced, "Rings:      Cost  Damage  Armor", "")
	return strings.Split(replaced, "\n")
}

func makeItemsList(itemsData []string) shop {
	items := make(shop)
	for i, itm := range itemsData {
		props := strings.Fields(itm)
		if len(props) <= 0 {
			continue
		}
		newItem := new(item)
		if len(props) == 5 {
			v1, _ := strconv.Atoi(props[3])
			v2, _ := strconv.Atoi(props[4])
			if v1 > v2 {
				newItem.val = v1
				newItem.class = dmg
			} else {
				newItem.val = v2
				newItem.class = def
			}
			newItem.key = props[0] + props[1]
			newItem.cost, _ = strconv.Atoi(props[2])
			newItem.category = misc
			items[i] = newItem
			continue
		}
		newItem.key = props[0]
		v1, _ := strconv.Atoi(props[2])
		v2, _ := strconv.Atoi(props[3])
		if v1 > v2 {
			newItem.val = v1
			newItem.class = dmg
			newItem.category = weapon
		} else {
			newItem.val = v2
			newItem.class = def
			newItem.category = armor
		}
		newItem.key = props[0]
		newItem.cost, _ = strconv.Atoi(props[1])
		items[i] = newItem
	}
	return items
}

func sortByCost(items shop) []*item {
	itemsList := make([]*item, len(items))
	i := 0
	for _, v := range items {
		itemsList[i] = v
		i++
	}
	sort.Slice(itemsList, func(i, j int) bool {
		return itemsList[i].cost > itemsList[j].cost
	})
	return itemsList
}

func filterByCategory(items []*item, category int) []*item {
	filtered := make([]*item, len(items))
	i := 0
	for _, v := range items {
		if v.category == category {
			filtered[i] = v
			i++
		}
	}
	return filtered[:i]
}

func filterByClass(items []*item, class int) []*item {
	filtered := make([]*item, len(items))
	i := 0
	for _, v := range items {
		if v.class == class {
			filtered[i] = v
			i++
		}
	}
	return filtered[:i]
}

func equipGear(player *gameObject, items []*item) {
	for _, itm := range items {
		switch itm.category {
		case weapon:
			player.dmg += itm.val
		case armor:
			player.armor += itm.val
		case misc:
			if itm.class == dmg {
				player.dmg += itm.val
			} else if itm.class == def {
				player.armor += itm.val
			}
		}
	}
}

func getGearCost(items []*item) (cost int) {
	for _, itm := range items {
		cost += itm.cost
	}
	return
}

func reset(player *gameObject, enemy *gameObject) {
	player.hp = 100
	player.dmg = 0
	player.armor = 0

	enemy.hp = 104
}

func findOptimalGear(items []*item) {
	weapons := filterByCategory(items, weapon)
	armor := filterByCategory(items, armor)
	misc := filterByCategory(items, misc)
	miscDmg := filterByClass(misc, dmg)
	miscDef := filterByClass(misc, def)

	player := &gameObject{key: "player", hp: 100, dmg: 0, armor: 0}
	enemy := &gameObject{key: "enemy", hp: 104, dmg: 8, armor: 1}

	mincost, maxcost := math.MaxInt32, math.MinInt32
	for w := 0; w < len(weapons); w++ {
		for a := 0; a < len(armor); a++ {
			for rl := 0; rl < len(miscDmg); rl++ {
				for rr := 0; rr < len(miscDef); rr++ {
					reset(player, enemy)
					buy := make([]*item, 4)
					buy[0] = weapons[w]
					buy[1] = armor[a]
					buy[2] = miscDmg[rl]
					buy[3] = miscDef[rr]
					cost := getGearCost(buy)
					equipGear(player, buy)
					winner, _ := deathMatch(player, enemy)

					if winner.key == "player" {
						if mincost > cost {
							mincost = cost
						}
					} else if winner.key == "enemy" {
						if maxcost < cost {
							maxcost = cost
						}
					}
				}
			}
		}
	}

	fmt.Println("\nOptimal Gear Cost: ", mincost, "\nMax Gear Cost: ", maxcost)
}

func deathMatch(player, enemy *gameObject) (winner *gameObject, turns int) {
	turns = 1
	for {
		enemy.hp -= (player.dmg - enemy.armor)
		if enemy.hp <= 0 {
			return player, turns
		}
		player.hp -= (enemy.dmg - player.armor)
		if player.hp <= 0 {
			return enemy, turns
		}
	}
}

//Run Solution
func Run() {
	itemData := readItemsList("./inputs/day-21.txt")
	items := makeItemsList(itemData)
	sorted := sortByCost(items)
	findOptimalGear(sorted)
}
