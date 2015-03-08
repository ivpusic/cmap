package cmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestCMap() C {
	return C{
		"keystr":  "value",
		"keyint":  123,
		"keybool": true,
	}
}

func TestGetStr(t *testing.T) {
	m := getTestCMap()

	assert.Equal(t, "value", m.Str("keystr"))
	assert.Equal(t, "value", m.StrOrDef("keystr", "default"))
	assert.Equal(t, "default", m.StrOrDef("keystrunknown", "default"))
}

func TestGetInt(t *testing.T) {
	m := getTestCMap()

	assert.Equal(t, 123, m.Int("keyint"))
	assert.Equal(t, 123, m.IntOrDef("keyint", 567))
	assert.Equal(t, 567, m.IntOrDef("keystrunknown", 567))
}

func TestGetBool(t *testing.T) {
	m := getTestCMap()

	assert.Equal(t, true, m.Bool("keybool"))
	assert.Equal(t, true, m.BoolOrDef("keybool", true))
	assert.Equal(t, false, m.BoolOrDef("keystrunknown", false))
}

func TestSet(t *testing.T) {
	m := getTestCMap()
	m.Set("key", "value")
	assert.Equal(t, "value", m.Str("key"))
}
