package schema

import "strings"

// Tag Methods
type Tag struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Value   string `json:"value,omitempty" yaml:"value,omitempty"`
	Options string `json:"options,omitempty" yaml:"options,omitempty"`
}

func (t *Tag) String() string {

	if t.Value == "" && len(t.Options) == 0 {
		return ""
	}

	var stb strings.Builder
	stb.WriteString(t.Name)

	stb.WriteString(":\"")
	if t.Value != "" {
		stb.WriteString(t.Value)
	}

	if t.Options != "" {
		stb.WriteRune(',')
		stb.WriteString(t.Options)
	}
	stb.WriteString("\"")

	return stb.String()
}

func NewTag(tn string, s string) Tag {

	t := Tag{Name: tn}
	if s != "" {
		ndx := strings.Index(s, ",")
		if ndx < 0 {
			t.Value = s
		} else {
			t.Value = s[:ndx]
			t.Options = s[ndx+1:]
		}
	}

	return t
}

func FindTag(tags []Tag, tn string) Tag {
	for _, t := range tags {
		if strings.ToLower(t.Name) == strings.ToLower(tn) {
			return t
		}
	}

	return Tag{}
}
