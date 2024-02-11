package funcmap_test

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/funcmap"
	"github.com/stretchr/testify/require"
	"testing"
)

type IW struct {
	path        string
	separator   string
	casing      funcmap.CasingMode
	indexMode   funcmap.IndexHandling
	indexFormat funcmap.IndexFormat
	w           string
}

func TestFormatPath(t *testing.T) {

	iws := []IW{}

	for i, iw := range iws {
		actual := funcmap.FormatPath(iw.path, iw.separator, iw.casing, iw.indexMode, iw.indexFormat)
		require.Equal(t, iw.w, actual, fmt.Sprintf("[%d]", i))
	}
}
