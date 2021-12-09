package day8

import (
	"sort"
	"strconv"
	"strings"
)

type Display struct {
	digitsSum int
	sigToNum  map[string]int
	numToSig  map[int]string
}

func NewDisplay() *Display {
	return &Display{0, make(map[string]int, 0), make(map[int]string, 0)}
}

func (d *Display) Parse(signals []string, segments []string) {
	d.init(signals)

	displayedNumber := ""
	for _, seg := range segments {
		displayedNumber += strconv.Itoa(d.getNum(seg))
	}

	number, _ := strconv.Atoi(displayedNumber)
	d.digitsSum += number
}

func (d Display) init(signals []string) {
	d.initBasicSignals(signals)

	for _, sig := range signals {
		switch len(sig) {
		case 5:
			if len(minus(sig, d.numToSig[1])) == 3 {
				d.mapSigToNum(sig, 3)
			} else if len(plus(sig, d.numToSig[4])) == 7 {
				d.mapSigToNum(sig, 2)
			} else {
				d.mapSigToNum(sig, 5)
			}
		case 6:
			if len(plus(sig, d.numToSig[7])) == 7 {
				d.mapSigToNum(sig, 6)
			} else if len(minus(sig, d.numToSig[4])) == 2 {
				d.mapSigToNum(sig, 9)
			} else {
				d.mapSigToNum(sig, 0)
			}
		}
	}
}

func (d Display) initBasicSignals(signals []string) {
	for _, sig := range signals {
		switch len(sig) {
		case 2:
			d.mapSigToNum(sig, 1)
		case 3:
			d.mapSigToNum(sig, 7)
		case 4:
			d.mapSigToNum(sig, 4)
		case 7:
			d.mapSigToNum(sig, 8)
		}
	}
}

func minus(a, b string) []string {
	set := make(map[string]struct{}, 0)

	for _, v := range strings.Split(b, "") {
		set[v] = struct{}{}
	}

	keys := make([]string, 0)

	for _, v := range strings.Split(a, "") {
		if _, ok := set[v]; !ok {
			keys = append(keys, v)
		}
	}

	return keys
}

func plus(a, b string) []string {
	set := make(map[string]struct{}, 0)

	for _, v := range strings.Split(a, "") {
		set[v] = struct{}{}
	}

	for _, v := range strings.Split(b, "") {
		set[v] = struct{}{}
	}

	keys := make([]string, 0)

	for k := range set {
		keys = append(keys, k)
	}

	return keys
}

func (d Display) mapSigToNum(sig string, num int) {
	s := sortString(sig)
	d.sigToNum[s] = num
	d.numToSig[num] = s
}

func sortString(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	sorted := strings.Join(split, "")
	return sorted
}

func (d *Display) addNumber(number int) {
	d.digitsSum += number
}

func (d Display) SumOfNumbers() int {
	return d.digitsSum
}

func (d Display) getNum(seg string) int {
	return d.sigToNum[sortString(seg)]
}
