package state

import (
	"fmt"
	"iter"
)

type replacePart interface {
	getIndex() int
	getIter() iter.Seq[string]
}

type numReplacePart struct {
	index  int
	start  int
	end    int
	length int
}

func (p numReplacePart) getIndex() int {
	return p.index
}

func (p numReplacePart) getIter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := p.start; i <= p.end; i++ {
			next := fmt.Sprintf("%0*d", p.length, i)
			if !yield(next) {
				return
			}
		}
	}
}

func n(index, start, end, length int) numReplacePart {
	return numReplacePart{
		index:  index,
		start:  start,
		end:    end,
		length: length,
	}
}

type charReplacePart struct {
	index int
	chars []rune
}

func (p charReplacePart) getIndex() int {
	return p.index
}

func (p charReplacePart) getIter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, char := range p.chars {
			if !yield(string(char)) {
				return
			}
		}
	}
}

func c(index int, chars ...rune) charReplacePart {
	return charReplacePart{
		index: index,
		chars: chars,
	}
}

type replacer struct {
	fmtString    string
	replaceParts []replacePart
}

func r(fmtString string, replaceParts ...replacePart) replacer {
	return replacer{
		fmtString:    fmtString,
		replaceParts: replaceParts,
	}
}

type iterator struct {
	replacePart replacePart
	next        func() (string, bool)
	stop        func()
}

func e(values ...replacer) map[string]int {
	options := make([]string, 0)

	for _, repl := range values {
		numParts := len(repl.replaceParts)
		parts := make([]any, numParts)
		iterators := make([]iterator, numParts)
		order := make([]int, numParts)

		for index, part := range repl.replaceParts {
			iter, done := iter.Pull(part.getIter())
			iterators[index] = iterator{
				replacePart: part,
				next:        iter,
				stop:        done,
			}
			next, _ := iter()
			parts[index] = next
			order[part.getIndex()] = index
		}

	loop:
		for {
			options = append(options, fmt.Sprintf(repl.fmtString, parts...))

			ok := false
			for index := 0; !ok; index++ {
				if index >= numParts {
					break loop
				}

				iterIndex := order[index]
				iterator := iterators[iterIndex]
				var part string
				part, ok = iterator.next()
				if !ok {
					iterator.stop()
					next, done := iter.Pull(iterator.replacePart.getIter())
					iterators[iterIndex].next = next
					iterators[iterIndex].stop = done
					part, _ = next()
				}
				parts[iterIndex] = part
			}
		}

		for _, iterator := range iterators {
			iterator.stop()
		}
	}

	result := make(map[string]int)
	for index, option := range options {
		result[option] = index
	}
	return result
}

var X32Enums = map[string]map[string]int{
	"On":          e(r("OFF"), r("ON")),
	"Color":       e(r("OFF"), r("RD"), r("GN"), r("YE"), r("BL"), r("MG"), r("CY"), r("WH"), r("OFFi"), r("RDi"), r("GNi"), r("YEi"), r("BLi"), r("MGi"), r("CYi"), r("WHi")),
	"Source":      e(r("OFF"), r("In%s", n(0, 1, 32, 2)), r("Aux%s", n(0, 1, 6, 1)), r("USB L"), r("USB R"), r("Fx %s%s", n(1, 1, 4, 1), c(0, 'L', 'R')), r("Bus %s", n(0, 1, 16, 2))),
	"HpSlope":     e(r("12"), r("18"), r("24")),
	"GateMode":    e(r("EXP2"), r("EXP3"), r("EXP4"), r("GATE"), r("DUCK")),
	"DynMode":     e(r("COMP"), r("EXP")),
	"Det":         e(r("PEAK"), r("RMS")),
	"Env":         e(r("LIN"), r("LOG")),
	"Ratio":       e(r("1.1"), r("1.3"), r("1.5"), r("2.0"), r("2.5"), r("3.0"), r("4.0"), r("5.0"), r("7.0"), r("10"), r("20"), r("100")),
	"Pos":         e(r("PRE"), r("POST")),
	"FilterType":  e(r("LC6"), r("LC12"), r("HC6"), r("HC12"), r("1.0"), r("2.0"), r("3.0"), r("5.0"), r("10.0")),
	"Sel":         e(r("OFF"), r("FX%s%s", n(1, 1, 8, 1), c(0, 'L', 'R')), r("AUX%s", n(0, 1, 6, 1))),
	"EqType":      e(r("LCut"), r("LShv"), r("PEQ"), r("VEQ"), r("HShv"), r("HCut"), r("BU6"), r("BU12"), r("BS12"), r("LR12"), r("BU18"), r("BU24"), r("BS24"), r("LR24")),
	"MixType":     e(r("IN/LC"), r("<-EQ"), r("EQ->"), r("PRE"), r("POST"), r("GRP")),
	"MonoMode":    e(r("LR+M"), r("LCR")),
	"SoloSource":  e(r("OFF"), r("LR"), r("LR+C"), r("LR PFL"), r("LR AFL"), r("AUX 56"), r("AUX 78")),
	"SoloMode":    e(r("PFL"), r("AFL")),
	"TalkSource":  e(r("INT"), r("EXT")),
	"Fsel":        e(r("F1"), r("F2")),
	"OscType":     e(r("SINE"), r("PINK"), r("WHITE")),
	"Dest":        e(r("MixBus%s", n(0, 1, 16, 2)), r("L"), r("R"), r("L+R"), r("Matrix%s", n(0, 1, 6, 1))),
	"InRouting":   e(r("AN1-8"), r("AN9-16"), r("AN17-24"), r("AN25-32"), r("A1-8"), r("A9-16"), r("A17-24"), r("A25-32"), r("A33-40"), r("A41-48"), r("B1-8"), r("B9-16"), r("B17-24"), r("B25-32"), r("B33-40"), r("B41-48"), r("CARD1-8"), r("CARD9-16"), r("CARD17-24"), r("CARD25-32")),
	"InAux":       e(r("AUX1-4"), r("AN1-2"), r("AN1-4"), r("AN1-6"), r("A1-2"), r("A1-4"), r("A1-6"), r("B1-2"), r("B1-4"), r("B1-6"), r("CARD1-2"), r("CARD1-4"), r("CARD1-6")),
	"Routing":     e(r("AN1-8"), r("AN9-16"), r("AN17-24"), r("AN25-32"), r("A1-8"), r("A9-16"), r("A17-24"), r("A25-32"), r("A33-40"), r("A41-48"), r("B1-8"), r("B9-16"), r("B17-24"), r("B25-32"), r("B33-40"), r("B41-48"), r("CARD1-8"), r("CARD9-16"), r("CARD17-24"), r("CARD25-32"), r("OUT1-8"), r("OUT9-16"), r("P161-8"), r("P16 9-16"), r("AUX1-6/Mon"), r("AuxIN1-6/TB")),
	"OutRouting":  e(),
	"Clockrate":   e(r("48K"), r("44K1")),
	"Clocksource": e(r("INT"), r("AES50A"), r("AES50B"), r("Exp Card")),
}
