# cmap
[![Build Status](https://travis-ci.org/ivpusic/cmap.svg?branch=master)](https://travis-ci.org/ivpusic/cmap)
[![GoDoc](https://godoc.org/github.com/ivpusic/cmap?status.svg)](https://godoc.org/github.com/ivpusic/cmap)

Map with convenient methods for getting and setting values
#Why?
Because checking if key in map exists, setting it to default value, and converting it to right type can be....yeah, annoying. 
## Example
### Classical way
```Go
// pass config
conf := make(map[string]interface{})
conf["key1"] = "some value"
conf["key2"] = 5678
ParseConfig(conf)

// parse config
func ParseConfig(conf map[string]interface{}) {
	// default value
	setting1 := "some default"
	if val, ok := conf["key1"]; ok {
		setting1 = val.(string)
	}

	// default value
	setting2 := 1234
	if val, ok := conf["key2"]; ok {
		setting2 = val.(int)
	}
}
```

### ``cmap`` way
```Go
// pass config
ParseConfig(cmap.C{
  "key1": "bla",
  "key2": 456,
})

// parse config
func ParseConfig(conf cmap.C) {
  	// get string value on key `key1` or return default value (`some default` in this case)
	setting1 := conf.StrOrDef("key1", "some default")
	// the same, just getting int
	setting2 := conf.IntOrDef("key2", 1234)
}
```

Find out more on [godoc](https://godoc.org/github.com/ivpusic/cmap)

# License
*MIT*
