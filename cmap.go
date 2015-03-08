// Map with convenient methods for getting and setting values
package cmap

type C map[string]interface{}

// Get value as string
func (c C) Str(key string) string {
	return c[key].(string)
}

// Get value as string if exist, if not return default value
func (c C) StrOrDef(key string, def string) string {
	if val, ok := c[key]; ok {
		return val.(string)
	}

	return def
}

// Get value as bool
func (c C) Bool(key string) bool {
	return c[key].(bool)
}

// Get value as bool if exist, if not return default value
func (c C) BoolOrDef(key string, def bool) bool {
	if val, ok := c[key]; ok {
		return val.(bool)
	}

	return def
}

// Get value as int
func (c C) Int(key string) int {
	return c[key].(int)
}

// Get value as int if exist, if not return default value
func (c C) IntOrDef(key string, def int) int {
	if val, ok := c[key]; ok {
		return val.(int)
	}

	return def
}

// Set value on map
func (c C) Set(key string, value interface{}) {
	c[key] = value
}
