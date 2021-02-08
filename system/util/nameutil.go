package util

import (
	"regexp"
	"strings"
)

var PackageNamePathForm = regexp.MustCompile("^(./)?[a-zA-Z_0-9]+(/[a-zA-Z_0-9]+)*$")
var PackageNameDotForm = regexp.MustCompile("^[a-zA-Z_0-9]+(\\.[a-zA-Z_0-9]+)*$")

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
	none               FormatMode = "none"
	camelCase                     = "camelCase"
	lowerCase                     = "lowerCase"
	upperCase                     = "upperCase"
	suppress                      = "suppress"
	indexIjk                      = "indexIjk"
	indexSprintf                  = "indexSprintf"
	indexIjkWoLast                = "indexIjkWoLast"
	indexSprintfWoLast            = "indexSprintfWoLast"
)

func FormatIdentifier(aName string, aSeparator string, aComponentFormatMode FormatMode, anIndexingFormatMode FormatMode) string {

	if aName == "" {
		return ""
	}

	nameComponents := strings.Split(aName, ".")

	var stb strings.Builder
	totalNumberOfBrackets := strings.Count(aName, "[")
	numberOfParts := 0
	numberOfBrackets := 0

	for _, s := range nameComponents {

		if s == "" {
			continue
		}

		actualFormattingMode := anIndexingFormatMode
		if s == "[]" {
			if numberOfBrackets == (totalNumberOfBrackets-1) &&
				(anIndexingFormatMode == indexIjkWoLast || anIndexingFormatMode == indexSprintfWoLast) {
				actualFormattingMode = suppress
			}

			if actualFormattingMode != suppress && numberOfParts > 0 && aSeparator != "" {
				stb.WriteString(aSeparator)
			}

			switch actualFormattingMode {
			case indexIjkWoLast:
				fallthrough
			case indexIjk:
				stb.WriteString(getIndexVariableName(numberOfBrackets, totalNumberOfBrackets))
				numberOfParts++

			case indexSprintfWoLast:
				fallthrough
			case indexSprintf:
				stb.WriteString("%d")
				numberOfParts++

			case suppress:

			default:
				stb.WriteString(s)
				numberOfParts++
			}

			numberOfBrackets++

		} else {
			if numberOfParts > 0 && aSeparator != "" {
				stb.WriteString(aSeparator)
			}

			switch aComponentFormatMode {
			case camelCase:
				stb.WriteString(ToCapitalCase(s))

			case upperCase:
				stb.WriteString(strings.ToUpper(s))

			case lowerCase:
				stb.WriteString(strings.ToLower(s))

			default:
				stb.WriteString(s)
			}

			numberOfParts++
		}
	}

	return stb.String()
}

func getIndexVariableName(variablePosition int, totalNumberOfVariables int) string {
	if variablePosition < 0 {
		return ""
	}

	return string('i' + variablePosition)
}

func AppendToNamespace(s1 string, s2 string, sep string) string {

	if s1 != "" {
		return strings.Join([]string{s1, s2}, sep)
	}

	return s2
}
