package model

import "strings"

type Universe string

const (
	Marvel   Universe = "MARVEL"
	DC       Universe = "DC"
	DCMarvel Universe = "DC|MARVEL"
)

func CheckUniverse(universe Universe) bool {
	universe = Universe(strings.ToUpper(string(universe)))
	switch universe {
	case Marvel, DC, DCMarvel:
		return true
	}
	return false
}
