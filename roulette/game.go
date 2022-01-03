package roulette

import (
	"math/rand"
	"time"
)

type betTypeInfo struct {
	Numbers        []int
	OddsMultiplier int
}

var straightMultiplier = 35

var betTypesMap = map[string]betTypeInfo{
	"0":  {[]int{0}, straightMultiplier},
	"1":  {[]int{1}, straightMultiplier},
	"2":  {[]int{2}, straightMultiplier},
	"3":  {[]int{3}, straightMultiplier},
	"4":  {[]int{4}, straightMultiplier},
	"5":  {[]int{5}, straightMultiplier},
	"6":  {[]int{6}, straightMultiplier},
	"7":  {[]int{7}, straightMultiplier},
	"8":  {[]int{8}, straightMultiplier},
	"9":  {[]int{9}, straightMultiplier},
	"10": {[]int{10}, straightMultiplier},
	"11": {[]int{11}, straightMultiplier},
	"12": {[]int{12}, straightMultiplier},
	"13": {[]int{13}, straightMultiplier},
	"14": {[]int{14}, straightMultiplier},
	"15": {[]int{15}, straightMultiplier},
	"16": {[]int{16}, straightMultiplier},
	"17": {[]int{17}, straightMultiplier},
	"18": {[]int{18}, straightMultiplier},
	"19": {[]int{19}, straightMultiplier},
	"20": {[]int{20}, straightMultiplier},
	"21": {[]int{21}, straightMultiplier},
	"22": {[]int{22}, straightMultiplier},
	"23": {[]int{23}, straightMultiplier},
	"24": {[]int{24}, straightMultiplier},
	"25": {[]int{25}, straightMultiplier},
	"26": {[]int{26}, straightMultiplier},
	"27": {[]int{27}, straightMultiplier},
	"28": {[]int{28}, straightMultiplier},
	"29": {[]int{29}, straightMultiplier},
	"30": {[]int{30}, straightMultiplier},
	"31": {[]int{31}, straightMultiplier},
	"32": {[]int{32}, straightMultiplier},
	"33": {[]int{33}, straightMultiplier},
	"34": {[]int{34}, straightMultiplier},
	"35": {[]int{35}, straightMultiplier},
	"36": {[]int{36}, straightMultiplier},
}

type SpinWheelFunc func() int

func SpinWheel() int {
	seed := rand.NewSource(time.Now().UnixNano())
	_rand := rand.New(seed)
	return _rand.Intn(37)
}

func calculateWinnings(_bet Bet, winningNumber int) int {
	_betTypeInfo := betTypesMap[_bet.Type]
	betNumbers := _betTypeInfo.Numbers

	if contains(betNumbers, winningNumber) {
		return (_bet.Size * _betTypeInfo.OddsMultiplier) + _bet.Size
	}
	return 0
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
