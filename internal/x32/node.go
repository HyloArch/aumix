package x32

import (
	"strconv"
	"strings"
)

func DecodeNode(nodeValue string) (string, []any) {
	index := strings.IndexRune(nodeValue, ' ')
	address := nodeValue[:index]

	values := make([]any, 0, 10)

	nodeValue = nodeValue[index+1:]
	inString := true
	for split := range strings.SplitSeq(nodeValue, "\"") {
		inString = !inString
		if inString {
			values = append(values, split)
			continue
		}

		for split2 := range strings.SplitSeq(split, " ") {
			trimmed := strings.TrimSpace(split2)
			if len(trimmed) == 0 {
				continue
			}

			var value any
			if trimmed[0] == '%' {
				if val32, err := strconv.ParseInt(trimmed[1:], 2, 32); err == nil {
					value = int32(val32)
				} else {
					value = trimmed
				}
			} else if val32, err := strconv.ParseFloat(trimmed, 32); err == nil {
				value = float32(val32)
			} else {
				value = trimmed
			}

			values = append(values, value)
		}

	}

	return address, values
}
