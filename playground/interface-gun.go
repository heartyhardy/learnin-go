package main

import (
	"fmt"
)

func main() {
	gun_ak47 := &AK47{
		name:    "AK-47 Magnum",
		bullets: 15,
	}

	for gun_ak47.bullets > 0 {
		gun_ak47.Fire()
	}
}

type Firer interface {
	Fire() (int error)
}

type AK47 struct {
	name    string
	bullets int
}

func (gun *AK47) Fire() {
	if gun.bullets > 0 {
		gun.bullets--
		fmt.Printf("%v was fired! %v bullets left\n", gun.name, gun.bullets)
	} else {
		fmt.Printf("Cannot fire! Reload!\n")
	}
}
