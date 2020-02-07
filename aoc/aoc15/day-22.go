package aoc15

import "fmt"

const (
	offensive = iota
	defensive
)

type player struct {
	hp      int
	mana    int
	armor   int
	effects []*effect
}

type enemy struct {
	hp      int
	armor   int
	dmg     int
	effects []*effect
}

type spell struct {
	name     string
	class    int
	dmg      int
	heal     int
	duration int
}

type effect struct {
	spell string
	dmg   int
	armor int
	mana  int
	heal  int
	turns int
}

// Is Player Alive
func (p *player) isAlive() bool {
	if p.hp <= 0 {
		return false
	}
	return true
}

// Player has Effects?
func (p *player) hasEffects() bool {
	if len(p.effects) > 0 {
		return true
	}
	return false
}

// Apply specified effect to Player
func (p *player) apply(efct *effect) {
	if efct.turns > 0 {
		p.hp += efct.heal
		p.mana += efct.mana
		p.armor += efct.armor
		efct.turns--
		if efct.turns-1 <= 0 {
			efct.turns = 0
			p.removeEffect(efct)
		}
	}
}

// If Player has effects apply then
func (p *player) applyEffects() {
	for _, efct := range p.effects {
		p.apply(efct)
	}
}

// Remove Effect
func (p *player) removeEffect(efct *effect) {
	for i := 0; i < len(p.effects); i++ {
		if p.effects[i].spell == efct.spell {
			copy(p.effects[i:], p.effects[i+1:])
			p.effects[len(p.effects)-1] = nil
			p.effects = p.effects[:len(p.effects)-1]
		}
	}
}

// Is Enemy Alive
func (e *enemy) isAlive() bool {
	if e.hp <= 0 {
		return false
	}
	return true
}

// Player has Effects?
func (e *enemy) hasEffects() bool {
	if len(e.effects) > 0 {
		return true
	}
	return false
}

// If Enemy has effects apply then
func (e *enemy) applyEffects() {
	for _, efct := range e.effects {
		e.apply(efct)
	}
}

// Apply specified effect to Player
func (e *enemy) apply(efct *effect) {
	if efct.turns > 0 {
		e.hp -= efct.dmg
		efct.turns--
		if efct.turns-1 <= 0 {
			efct.turns = 0
			e.removeEffect(efct)
		}
	}
}

// Remove Effect
func (e *enemy) removeEffect(efct *effect) {
	for i := 0; i < len(e.effects); i++ {
		if e.effects[i].spell == efct.spell {
			copy(e.effects[i:], e.effects[i+1:])
			e.effects[len(e.effects)-1] = nil
			e.effects = e.effects[:len(e.effects)-1]
		}
	}
}

// Attack player
func (e *enemy) attack(p *player) bool {
	p.hp -= e.dmg + p.armor
	if p.hp <= 0 {
		p.hp = 0
	}
	return p.isAlive()
}

//Make a clone of player
func (p *player) clone() *player {
	pclone := new(player)
	pclone.armor = p.armor
	pclone.hp = p.hp
	pclone.mana = p.mana
	pclone.effects = make([]*effect, len(p.effects))
	for i, efct := range p.effects {
		newEffect := new(effect)
		newEffect.armor = efct.armor
		newEffect.dmg = efct.dmg
		newEffect.heal = efct.heal
		newEffect.mana = efct.mana
		newEffect.spell = efct.spell
		newEffect.turns = efct.turns
		pclone.effects[i] = newEffect
	}
	return pclone
}

// Make a clone of enemy
func (e *enemy) clone() *enemy {
	eclone := new(enemy)
	eclone.armor = e.armor
	eclone.hp = e.hp
	eclone.dmg = e.dmg
	eclone.effects = make([]*effect, len(e.effects))
	for i, efct := range e.effects {
		newEffect := new(effect)
		newEffect.armor = efct.armor
		newEffect.dmg = efct.dmg
		newEffect.heal = efct.heal
		newEffect.mana = efct.mana
		newEffect.spell = efct.spell
		newEffect.turns = efct.turns
		eclone.effects[i] = newEffect
	}
	return eclone
}

//Check how many turns it takes to win/lose
// Takes current status into consideration
func (p *player) getTurnsBeforeDeath(e *enemy) (interface{}, int) {
	// consider enemy dmg per turn
	// consider player current effects
	// consider enemy current effects
	pl, en := p.clone(), e.clone()

	turn := 0
	for {

		//Increment Turn
		turn++

		//Player Turn
		fmt.Println(pl.hp, en.hp)
		//Has player got effects?
		if ok := pl.hasEffects(); ok {
			pl.applyEffects()
		}
		//Has enemy got effects?
		if ok := en.hasEffects(); ok {
			en.applyEffects()
		}
		if ok := en.isAlive(); !ok {
			return p, turn
		}

		//Increment turns
		turn++

		// Enemy Turn
		fmt.Println(pl.hp, en.hp)
		//Has player got effects?
		if ok := pl.hasEffects(); ok {
			pl.applyEffects()
		}
		//Has enemy got effects?
		if ok := en.hasEffects(); ok {
			en.applyEffects()
		}
		//is enemy still alive?
		if ok := en.isAlive(); ok {
			en.attack(pl)
		} else {
			//if not return winner and turns
			// return winner, turns
			return p, turn
		}
		//is player still alive?
		if ok := pl.isAlive(); ok {
			//Advance to next turn

		} else {
			//Return winner and turns
			return e, turn
		}
	}
}

func runSimulation() {
	p := new(player)
	p.hp = 50
	p.mana = 500
	p.armor = 0

	e := new(enemy)
	e.hp = 55
	e.dmg = 8
	e.armor = 0

	winner, turns := p.getTurnsBeforeDeath(e)
	if v, ok := winner.(*player); ok {
		fmt.Println(v, turns)
	}
	if v, ok := winner.(*enemy); ok {
		fmt.Println(v, turns)
	}
}

//Run Solution
func Run() {
	runSimulation()
}
