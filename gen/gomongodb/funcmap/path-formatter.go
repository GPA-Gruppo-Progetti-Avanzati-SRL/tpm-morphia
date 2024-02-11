package funcmap

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"strings"
)

type CasingMode string
type IndexHandling string
type IndexFormat string

const (
	CasingModeNone      CasingMode = "none"
	CasingCamelCase     CasingMode = "camelCase"
	CasingLowerCase     CasingMode = "lowerCase"
	CasingUpperCase     CasingMode = "upperCase"
	CasingClassifyCase  CasingMode = "classify"
	CasingDasherizeCase CasingMode = "dasherize"

	IndexFormatIndexIjk     IndexFormat = "indexIjk"
	IndexFormatIndexSprintf IndexFormat = "indexSprintf"

	IndexHandlingSuppress    IndexHandling = "suppress"
	IndexHandlingIndex       IndexHandling = "index"
	IndexHandlingIndexWoLast IndexHandling = "indexWoLast"
)

func FormatPath(aName string, aSeparator string, aCasingMode CasingMode, indexHandling IndexHandling, indexFormat IndexFormat) string {

	if aName == "" {
		return ""
	}

	if indexHandling == IndexHandlingIndexWoLast {
		aName = strings.TrimSuffix(aName, ".[]")
		indexHandling = IndexHandlingIndex
	}

	nameComponents := strings.Split(aName, ".")

	var stb strings.Builder
	// totalNumberOfBrackets := strings.Count(aName, "[") + strings.Count(aName, "%s")
	numberOfParts := 0
	numberOfArrayBrackets := 0
	numberOfMapBrackets := 0

	for _, s := range nameComponents {

		if s == "" {
			continue
		}

		actualFormattingMode := indexHandling
		if s == "[]" || s == "%s" {

			/*
				if (numberOfArrayBrackets + numberOfMapBrackets) == (totalNumberOfBrackets-1) &&
					(anIndexingHandlingMode == indexIjkWoLast || anIndexingHandlingMode == indexSprintfWoLast) {
					actualFormattingMode = suppress
				}
			*/

			if (actualFormattingMode != IndexHandlingSuppress || s == "%s") && numberOfParts > 0 && aSeparator != "" {
				stb.WriteString(aSeparator)
			}

			switch actualFormattingMode {
			case IndexHandlingIndex:
				if indexFormat == IndexFormatIndexIjk {
					if s == "[]" {
						stb.WriteString(getIndexVariableName('i', numberOfArrayBrackets, aCasingMode))
					} else {
						stb.WriteString(getIndexVariableName('s', numberOfMapBrackets, aCasingMode))
					}
				} else {
					if s == "[]" {
						stb.WriteString("%d")
					} else {
						stb.WriteString("%s")
					}
				}

				numberOfParts++

			case IndexHandlingSuppress:
				if s != "[]" {
					if indexFormat == IndexFormatIndexIjk {
						stb.WriteString(getIndexVariableName('s', numberOfMapBrackets, aCasingMode))
					} else {
						stb.WriteString("%s")
					}
					numberOfParts++
				}

			default:
				stb.WriteString(s)
				numberOfParts++
			}

			if s == "[]" {
				numberOfArrayBrackets++
			} else {
				numberOfMapBrackets++
			}

		} else {
			if numberOfParts > 0 && aSeparator != "" {
				stb.WriteString(aSeparator)
			}

			stb.WriteString(adaptCasing(s, aCasingMode))
			numberOfParts++
		}
	}

	s := stb.String()
	// fmt.Printf("%-30s %s %-10s %-10s %-10s = %s\n", aName, aSeparator, aCasingMode, indexHandling, indexFormat, stb.String())
	return s
}

func getIndexVariableName(offsetChar int, variablePosition int, casingMode CasingMode) string {
	if variablePosition < 0 {
		return ""
	}

	return adaptCasing(string(rune(offsetChar+variablePosition)), casingMode)
}

func adaptCasing(s string, casing CasingMode) string {

	switch casing {
	case CasingCamelCase:
		s = util.Classify(s)

	case CasingUpperCase:
		s = strings.ToUpper(s)
	case CasingClassifyCase:
		s = util.Classify(s)
	case CasingDasherizeCase:
		s = util.Dasherize(s)
	case CasingLowerCase:
		s = strings.ToLower(s)
	}

	return s
}
