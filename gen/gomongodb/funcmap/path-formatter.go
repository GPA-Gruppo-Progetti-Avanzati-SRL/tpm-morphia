package funcmap

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/attributes"
	"strings"
)

const (
	IndexIjk     = "indexIjk"
	IndexSprintf = "indexSprintf"

	CasingNone    = "none"
	CamelCase     = "camelCase"
	LowerCase     = "lowerCase"
	UpperCase     = "upperCase"
	ClassifyCase  = "classify"
	DasherizeCase = "dasherize"
)

type PathFormatter struct {
	path                  attributes.PathInfo
	casing                string
	indexFormat           string
	suppressIndex         bool
	trimSuffixPlaceHolder bool
	separator             string
}

func NewPathFormatter(p attributes.PathInfo) *PathFormatter {
	pf := PathFormatter{
		path: p, indexFormat: IndexIjk, casing: CasingNone, separator: ".",
	}

	return &pf
}

func (pf *PathFormatter) WithCasing(c string) *PathFormatter {
	pf.casing = c
	return pf
}

func (pf *PathFormatter) WithSeparator(s string) *PathFormatter {
	pf.separator = s
	return pf
}

func (pf *PathFormatter) WithSprintfFormat() *PathFormatter {
	pf.indexFormat = IndexSprintf
	return pf
}

func (pf *PathFormatter) WithSuppressCollectionsPlaceHolders() *PathFormatter {
	pf.suppressIndex = true
	return pf
}

func (pf *PathFormatter) WithTrimSuffixPlaceHolder() *PathFormatter {
	pf.trimSuffixPlaceHolder = true
	return pf
}

func (pf *PathFormatter) getIndexVariableName(offsetChar int, variablePosition int) string {
	if variablePosition < 0 {
		return ""
	}
	return pf.adaptCasing(string(rune(offsetChar + variablePosition)))
}

func (pf *PathFormatter) adaptCasing(s string) string {

	switch pf.casing {
	case CamelCase:
		s = util.Classify(s) // TODO move to Camelize.
	case UpperCase:
		s = strings.ToUpper(s)
	case ClassifyCase:
		s = util.Classify(s)
	case DasherizeCase:
		s = util.Dasherize(s)
	case LowerCase:
		s = strings.ToLower(s)
	}

	return s
}

func (pf *PathFormatter) FormatPath() string {
	return pf.format(pf.path.Path)
}

func (pf *PathFormatter) FormatBsonPath() string {
	return pf.format(pf.path.BsonPath)
}

func (pf *PathFormatter) format(aPath string) string {

	if aPath == "" {
		return ""
	}

	if pf.trimSuffixPlaceHolder {
		aPath = strings.TrimSuffix(aPath, ".[]")
	}

	pathSegments := strings.Split(aPath, ".")

	var stb strings.Builder
	// totalNumberOfBrackets := strings.Count(aName, "[") + strings.Count(aName, "%s")
	numberOfParts := 0
	numberOfArrayBrackets := 0
	numberOfMapBrackets := 0

	for _, s := range pathSegments {

		if s == "" {
			continue
		}

		switch s {
		case "[]":
			if !pf.suppressIndex {
				if numberOfParts > 0 && pf.separator != "" {
					stb.WriteString(pf.separator)
				}

				if pf.indexFormat == IndexIjk {
					stb.WriteString(pf.getIndexVariableName('i', numberOfArrayBrackets))
				} else {
					stb.WriteString("%d")
				}

				numberOfArrayBrackets++
				numberOfParts++
			}

		case "%s":
			if numberOfParts > 0 && pf.separator != "" {
				stb.WriteString(pf.separator)
			}
			if pf.indexFormat == IndexIjk {
				stb.WriteString(pf.getIndexVariableName('s', numberOfMapBrackets))
			} else {
				stb.WriteString("%s")
			}
			numberOfMapBrackets++
			numberOfParts++

		default:
			if numberOfParts > 0 && pf.separator != "" {
				stb.WriteString(pf.separator)
			}

			stb.WriteString(pf.adaptCasing(s))
			numberOfParts++
		}

	}

	return stb.String()
}
