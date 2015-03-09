// Map with convenient methods for getting and setting values
package cmap

type C map[string]interface{}

// Get value as string.
// It will panic if value not exists or is wrong type
func (c C) Str(key string) string {
	return c[key].(string)
}

// Get value as string if exist, if not return default value
func (c C) StrOrDef(key string, def string) string {
	switch t := c[key].(type) {
	case string:
		return t
	default:
		return def
	}
}

// Get value as bool
// It will panic if value not exists or is wrong type
func (c C) Bool(key string) bool {
	return c[key].(bool)
}

// Get value as bool if exist, if not return default value
func (c C) BoolOrDef(key string, def bool) bool {
	switch t := c[key].(type) {
	case bool:
		return t
	default:
		return def
	}
}

// Get value as int
// It will panic if value not exists or is wrong type
func (c C) Int(key string) int {
	return c[key].(int)
}

// Get value as int if exist, if not return default value
func (c C) IntOrDef(key string, def int) int {
	switch t := c[key].(type) {
	case int:
		return t
	default:
		return def
	}
}

// Set value on map
func (c C) Set(key string, value interface{}) {
	c[key] = value
}
