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

type X32Enum int32

func (e *X32Enum) Get() []any {
	return []any{*e}
}

func (e *X32Enum) Set(values []any, enumMap map[string]int) (int, error) {
	if len(values) < 1 {
		return 0, fmt.Errorf("Values must be of at least size 1")
	}
	switch v := values[0].(type) {
	case string:
		index, ok := enumMap[v]
		if !ok {
			return 0, fmt.Errorf("Enum value doesn't exist")
		}
		*e = X32Enum(index)
	case int32:
		*e = X32Enum(v)
	case float32:
		*e = X32Enum(v)
	default:
		return 0, fmt.Errorf("Provided value isn't of type string, int, or float")
	}
	return 1, nil
}

var X32EnumOn = e(r("OFF"), r("ON"))

type X32EnumOnType struct{ X32Enum }

func (e *X32EnumOnType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumOn)
}

var X32EnumColor = e(r("OFF"), r("RD"), r("GN"), r("YE"), r("BL"), r("MG"), r("CY"), r("WH"), r("OFFi"), r("RDi"), r("GNi"), r("YEi"), r("BLi"), r("MGi"), r("CYi"), r("WHi"))

type X32EnumColorType struct{ X32Enum }

func (e *X32EnumColorType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumColor)
}

var X32EnumSource = e(r("OFF"), r("In%s", n(0, 1, 32, 2)), r("Aux%s", n(0, 1, 6, 1)), r("USB L"), r("USB R"), r("Fx %s%s", n(1, 1, 4, 1), c(0, 'L', 'R')), r("Bus %s", n(0, 1, 16, 2)))

type X32EnumSourceType struct{ X32Enum }

func (e *X32EnumSourceType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumSource)
}

var X32EnumHpSlope = e(r("12"), r("18"), r("24"))

type X32EnumHpSlopeType struct{ X32Enum }

func (e *X32EnumHpSlopeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumHpSlope)
}

var X32EnumGateMode = e(r("EXP2"), r("EXP3"), r("EXP4"), r("GATE"), r("DUCK"))

type X32EnumGateModeType struct{ X32Enum }

func (e *X32EnumGateModeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumGateMode)
}

var X32EnumDynMode = e(r("COMP"), r("EXP"))

type X32EnumDynModeType struct{ X32Enum }

func (e *X32EnumDynModeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumDynMode)
}

var X32EnumDet = e(r("PEAK"), r("RMS"))

type X32EnumDetType struct{ X32Enum }

func (e *X32EnumDetType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumDet)
}

var X32EnumEnv = e(r("LIN"), r("LOG"))

type X32EnumEnvType struct{ X32Enum }

func (e *X32EnumEnvType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumEnv)
}

var X32EnumRatio = e(r("1.1"), r("1.3"), r("1.5"), r("2.0"), r("2.5"), r("3.0"), r("4.0"), r("5.0"), r("7.0"), r("10"), r("20"), r("100"))

type X32EnumRatioType struct{ X32Enum }

func (e *X32EnumRatioType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumRatio)
}

var X32EnumPos = e(r("PRE"), r("POST"))

type X32EnumPosType struct{ X32Enum }

func (e *X32EnumPosType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumPos)
}

var X32EnumFilterType = e(r("LC6"), r("LC12"), r("HC6"), r("HC12"), r("1.0"), r("2.0"), r("3.0"), r("5.0"), r("10.0"))

type X32EnumFilterTypeType struct{ X32Enum }

func (e *X32EnumFilterTypeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumFilterType)
}

var X32EnumSel = e(r("OFF"), r("FX%s%s", n(1, 1, 8, 1), c(0, 'L', 'R')), r("AUX%s", n(0, 1, 6, 1)))

type X32EnumSelType struct{ X32Enum }

func (e *X32EnumSelType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumSel)
}

var X32EnumEqType = e(r("Lcut"), r("LShv"), r("PEQ"), r("VEQ"), r("HShv"), r("HCut"), r("BU6"), r("BU12"), r("BS12"), r("LR12"), r("BU18"), r("BU24"), r("BS24"), r("LR24"))

type X32EnumEqTypeType struct{ X32Enum }

func (e *X32EnumEqTypeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumEqType)
}

var X32EnumMixType = e(r("IN/LC"), r("<-EQ"), r("EQ->"), r("PRE"), r("POST"), r("GRP"))

type X32EnumMixTypeType struct{ X32Enum }

func (e *X32EnumMixTypeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumMixType)
}

var X32EnumMonoMode = e(r("LR+M"), r("LCR"))

type X32EnumMonoModeType struct{ X32Enum }

func (e *X32EnumMonoModeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumMonoMode)
}

var X32EnumSoloSource = e(r("OFF"), r("LR"), r("LR+C"), r("LR PFL"), r("LR AFL"), r("AUX 56"), r("AUX 78"))

type X32EnumSoloSourceType struct{ X32Enum }

func (e *X32EnumSoloSourceType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumSoloSource)
}

var X32EnumSoloMode = e(r("PFL"), r("AFL"))

type X32EnumSoloModeType struct{ X32Enum }

func (e *X32EnumSoloModeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumSoloMode)
}

var X32EnumTalkSource = e(r("INT"), r("EXT"))

type X32EnumTalkSourceType struct{ X32Enum }

func (e *X32EnumTalkSourceType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumTalkSource)
}

var X32EnumFsel = e(r("F1"), r("F2"))

type X32EnumFselType struct{ X32Enum }

func (e *X32EnumFselType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumFsel)
}

var X32EnumOscType = e(r("SINE"), r("PINK"), r("WHITE"))

type X32EnumOscTypeType struct{ X32Enum }

func (e *X32EnumOscTypeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumOscType)
}

var X32EnumDest = e(r("MixBus%s", n(0, 1, 16, 2)), r("L"), r("R"), r("L+R"), r("Matrix%s", n(0, 1, 6, 1)))

type X32EnumDestType struct{ X32Enum }

func (e *X32EnumDestType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumDest)
}

var X32EnumInRouting = e(r("AN1-8"), r("AN9-16"), r("AN17-24"), r("AN25-32"), r("A1-8"), r("A9-16"), r("A17-24"), r("A25-32"), r("A33-40"), r("A41-48"), r("B1-8"), r("B9-16"), r("B17-24"), r("B25-32"), r("B33-40"), r("B41-48"), r("CARD1-8"), r("CARD9-16"), r("CARD17-24"), r("CARD25-32"))

type X32EnumInRoutingType struct{ X32Enum }

func (e *X32EnumInRoutingType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumInRouting)
}

var X32EnumInAux = e(r("AUX1-4"), r("AN1-2"), r("AN1-4"), r("AN1-6"), r("A1-2"), r("A1-4"), r("A1-6"), r("B1-2"), r("B1-4"), r("B1-6"), r("CARD1-2"), r("CARD1-4"), r("CARD1-6"))

type X32EnumInAuxType struct{ X32Enum }

func (e *X32EnumInAuxType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumInAux)
}

var X32EnumRouting = e(r("AN1-8"), r("AN9-16"), r("AN17-24"), r("AN25-32"), r("A1-8"), r("A9-16"), r("A17-24"), r("A25-32"), r("A33-40"), r("A41-48"), r("B1-8"), r("B9-16"), r("B17-24"), r("B25-32"), r("B33-40"), r("B41-48"), r("CARD1-8"), r("CARD9-16"), r("CARD17-24"), r("CARD25-32"), r("OUT1-8"), r("OUT9-16"), r("P161-8"), r("P16 9-16"), r("AUX1-6/Mon"), r("AuxIN1-6/TB"))

type X32EnumRoutingType struct{ X32Enum }

func (e *X32EnumRoutingType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumRouting)
}

var X32EnumOutRoutingLow = e(r("AN1-4"), r("AN9-12"), r("AN17-20"), r("AN25-28"), r("A1-4"), r("A9-12"), r("A17-20"), r("A25-28"), r("A33-36"), r("A41-44"), r("B1-4"), r("B9-12"), r("B17-20"), r("B25-28"), r("B33-46"), r("B41-44"), r("CARD1-4"), r("CARD9-12"), r("CARD17-20"), r("CARD25-28"), r("OUT1-4"), r("OUT9-12"), r("P161-4"), r("P169-12"), r("AUX/CR"), r("AUX/TB"))

type X32EnumOutRoutingLowType struct{ X32Enum }

func (e *X32EnumOutRoutingLowType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumOutRoutingLow)
}

var X32EnumOutRoutingHigh = e(r("AN5-8"), r("AN13-16"), r("AN21-24"), r("AN29-32"), r("A5-8"), r("A13-16"), r("A21-24"), r("A29-32"), r("A37-40"), r("A45-48"), r("B5-8"), r("B13-16"), r("B21-24"), r("B29-32"), r("B37-40"), r("B45-48"), r("CARD5-8"), r("CARD13-16"), r("CARD21-24"), r("CARD29-32"), r("OUT5-8"), r("OUT13-16"), r("P165-8"), r("P1613-16"), r("AUX/CR"), r("AUX/TB"))

type X32EnumOutRoutingHighType struct{ X32Enum }

func (e *X32EnumOutRoutingHighType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumOutRoutingHigh)
}

var X32EnumClockrate = e(r("48K"), r("44K1"))

type X32EnumClockrateType struct{ X32Enum }

func (e *X32EnumClockrateType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumClockrate)
}

var X32EnumClocksource = e(r("INT"), r("AES50A"), r("AES50B"), r("Exp Card"))

type X32EnumClocksourceType struct{ X32Enum }

func (e *X32EnumClocksourceType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumClocksource)
}

var X32EnumFxTypeSourced = e(r("HALL"), r("AMBI"), r("RPLT"), r("ROOM"), r("CHAM"), r("PLAT"), r("VREV"), r("VRM"), r("GATE"), r("RVRS"), r("DLY"), r("3TAP"), r("4TAP"), r("CRS"), r("FLNG"), r("PHAS"), r("DIMC"), r("FILT"), r("ROTA"), r("PAN"), r("SUB"), r("D/RV"), r("CR/R"), r("FL/R"), r("D/CR"), r("D/FL"), r("MODD"), r("GEQ2"), r("GEQ"), r("TEQ2"), r("TEQ"), r("DES2"), r("DES"), r("P1A"), r("P1A2"), r("PQ5"), r("PQ5S"), r("WAVD"), r("LIM"), r("CMB"), r("CMB2"), r("FAC"), r("FAC1M"), r("FAC2"), r("LEC"), r("LEC2"), r("ULC"), r("ULC2"), r("ENH2"), r("ENH"), r("EXC2"), r("EXC"), r("IMG"), r("EDI"), r("SON"), r("AMP2"), r("AMP"), r("DRV2"), r("DRV"), r("PIT2"), r("PIT"))

type X32EnumFxTypeSourcedType struct{ X32Enum }

func (e *X32EnumFxTypeSourcedType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumFxTypeSourced)
}

var X32EnumFxSource = e(r("INS"), r("MIX1", n(0, 1, 16, 1)), r("M/C"))

type X32EnumFxSourceType struct{ X32Enum }

func (e *X32EnumFxSourceType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumFxSource)
}

var X32EnumFxType = e(r("GEQ2"), r("GEQ"), r("TEQ2"), r("TEQ"), r("DES2"), r("DES"), r("P1A"), r("P1A2"), r("PQ5"), r("PQ5S"), r("WAVD"), r("LIM"), r("FAC"), r("FAC1M"), r("FAC2"), r("LEC"), r("LEC2"), r("ULC"), r("ULC2"), r("ENH2"), r("ENH"), r("EXC2"), r("EXC"), r("IMG"), r("EDI"), r("SON"), r("AMP2"), r("AMP"), r("DRV2"), r("DRV"), r("PHAS"), r("FILT"), r("PAN"), r("SUB"))

type X32EnumFxTypeType struct{ X32Enum }

func (e *X32EnumFxTypeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumFxType)
}

var X32EnumOutputSource = e(r("OFF"), r("Main L"), r("Main R"), r("M/C"), r("MixBus %s", n(0, 1, 16, 2)), r("Matrix %s", n(0, 1, 6, 1)), r("DirectOut Ch %s", n(0, 1, 32, 2)), r("DirectOut Aux %s", n(0, 1, 8, 1)), r("DirectOut FX %s%s", n(1, 1, 4, 1), c(0, 'L', 'R')), r("Monitor L"), r("Monitor R"), r("Talkback"))

type X32EnumOutputSourceType struct{ X32Enum }

func (e *X32EnumOutputSourceType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumOutputSource)
}

var X32EnumOutputPos = e(r("IN/LC"), r("IN/LC+M"), r("<-EQ"), r("<-EQ+M"), r("EQ->"), r("EQ->+M"), r("PRE"), r("PRE+M"), r("POST"))

type X32EnumOutputPosType struct{ X32Enum }

func (e *X32EnumOutputPosType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumOutputPos)
}

var X32EnumIQGroup = e(r("OFF"), r("A"), r("B"))

type X32EnumIQGroupType struct{ X32Enum }

func (e *X32EnumIQGroupType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumIQGroup)
}

var X32EnumIQSpeaker = e(r("none"), r("iQ8"), r("iQ10"), r("iQ12"), r("iQ15"), r("iQ15B"), r("iQ18B"))

type X32EnumIQSpeakerType struct{ X32Enum }

func (e *X32EnumIQSpeakerType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumIQSpeaker)
}

var X32EnumIQEq = e(r("Linear"), r("Live"), r("Speech"), r("Playback"), r("User"))

type X32EnumIQEqType struct{ X32Enum }

func (e *X32EnumIQEqType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumIQEq)
}

var X32EnumShowControl = e(r("CUES"), r("SCENES"), r("SNIPPETS"))

type X32EnumShowControlType struct{ X32Enum }

func (e *X32EnumShowControlType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumShowControl)
}

var X32EnumClockMode = e(r("24h"), r("12h"))

type X32EnumClockModeType struct{ X32Enum }

func (e *X32EnumClockModeType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumClockMode)
}

var X32EnumInv = e(r("NORM"), r("INV"))

type X32EnumInvType struct{ X32Enum }

func (e *X32EnumInvType) Set(values ...any) (int, error) {
	return e.X32Enum.Set(values, X32EnumInv)
}
