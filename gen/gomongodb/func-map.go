package gomongodb

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var PackageNamePathForm = regexp.MustCompile("^(./)?[a-z\\-A-Z_0-9]+(/[a-z\\-A-Z_0-9]+)*$")
var PackageNameDotForm = regexp.MustCompile("^[a-z\\-A-Z_0-9]+(\\.[a-z\\-A-Z_0-9]+)*$")

func IsPathForm(s string) bool {
	return PackageNamePathForm.Match([]byte(s))
}

func IsDottedForm(s string) bool {
	return PackageNameDotForm.Match([]byte(s))
}

func SliceSegmentedName(s string) []string {

	if IsPathForm(s) {
		return strings.Split(s, "/")
	}

	if IsDottedForm(s) {
		return strings.Split(s, ".")
	}

	return nil
}

func ToCapitalCase(s string) string {
	return strings.Title(s)
}

/*
 * Utility Functions
 */
func NameWellFormed(n string) bool {
	return FieldNameWellFormed(n)
}

func FieldNameWellFormed(n string) bool {

	if n == "" {
		return false
	}
	return true
}

type FormatMode string

const (
	none          FormatMode = "none"
	camelCase                = "camelCase"
	lowerCase                = "lowerCase"
	upperCase                = "upperCase"
	classifyCase             = "classify"
	dasherizeCase            = "dasherize"
	indexIjk                 = "indexIjk"
	indexSprintf             = "indexSprintf"

	suppress    = "suppress"
	index       = "index"
	indexWoLast = "indexWoLast"
)

func FormatIdentifier(aName string, aSeparator string, aCasingMode FormatMode, indexHandling FormatMode, indexFormat FormatMode) string {

	if aName == "" {
		return ""
	}

	if indexHandling == indexWoLast {
		aName = strings.TrimSuffix(aName, ".[]")
		indexHandling = index
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

			if (actualFormattingMode != suppress || s == "%s") && numberOfParts > 0 && aSeparator != "" {
				stb.WriteString(aSeparator)
			}

			switch actualFormattingMode {
			case index:
				if indexFormat == indexIjk {
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

			case suppress:
				if s != "[]" {
					if indexFormat == indexIjk {
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

func getIndexVariableName(offsetChar int, variablePosition int, casingMode FormatMode) string {
	if variablePosition < 0 {
		return ""
	}

	return adaptCasing(string(rune(offsetChar+variablePosition)), casingMode)
}

func AppendToNamespace(s1 string, s2 string, sep string) string {

	if s1 != "" {
		return strings.Join([]string{s1, s2}, sep)
	}

	return s2
}

func FirstToLower(s string) string {
	if len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r != utf8.RuneError || size > 1 {
			lo := unicode.ToLower(r)
			if lo != r {
				s = string(lo) + s[size:]
			}
		}
	}
	return s
}

func adaptCasing(s string, casing FormatMode) string {

	switch casing {
	case camelCase:
		s = ToCapitalCase(s)

	case upperCase:
		s = strings.ToUpper(s)
	case classifyCase:
		s = util.Classify(s)
	case dasherizeCase:
		s = util.Dasherize(s)
	case lowerCase:
		s = strings.ToLower(s)
	}

	return s
}

func criteriaMethodVarParams(p string, withType bool, commaHandling string) string {

	if !(strings.Contains(p, "[]") || strings.Contains(p, "%s")) {
		if strings.Contains(commaHandling, "addonempty") {
			return ", "
		}

		return ""
	}

	var sb strings.Builder

	arr := strings.Split(p, ".")

	numIjk := 0
	numStw := 0
	for _, s := range arr {
		switch s {
		case "[]":
			if (numIjk + numStw) > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString("ndx")
			sb.WriteRune(rune('I' + numIjk))
			if withType {
				sb.WriteString(" int")
			}
			numIjk++
		case "%s":
			if (numIjk + numStw) > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString("key")
			sb.WriteRune(rune('S' + numStw))
			if withType {
				sb.WriteString(" string")
			}
			numStw++
		default:
			// No-op
		}
	}

	if strings.Contains(commaHandling, "before") {
		return ", " + sb.String()
	} else if strings.Contains(commaHandling, "after") {
		return sb.String() + ", "
	}

	return sb.String()
}

/*
func updateMethodVarParams(p string, withType bool, commaHandling string) string {

	if strings.Contains(p, "[]") || strings.Contains(p, "%s") {

		var sb strings.Builder

		arr := strings.Split(p, ".")

		numIjk := 0
		numStw := 0
		for _, s := range arr {
			if s == "[]" {
				if (numIjk + numStw) > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString("ndx")
				sb.WriteRune(rune('I' + numIjk))
				if withType {
					sb.WriteString(" int")
				}
				numIjk++
			} else if s == "%s" {
				if (numIjk + numStw) > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString("key")
				sb.WriteRune(rune('S' + numStw))
				if withType {
					sb.WriteString(" string")
				}
				numStw++
			}
		}

		if strings.Contains(commaHandling, "before") {
			return ", " + sb.String()
		} else if strings.Contains(commaHandling, "after") {
			return sb.String() + ", "
		}

		return sb.String()
	}

	if strings.Contains(commaHandling, "addonempty") {
		return ", "
	}

	return ""
}
*/
