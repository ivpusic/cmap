# cmap
Map with convenient methods for getting and setting values

#Why?
I hate doing this:
```Go
// pass config
conf := make(map[string]interface{})
conf["key1"] = "some value"
conf["key2"] = 5678
ParseConfig(conf)

// parse config
func ParseConfig(conf map[string]interface{}) {
	setting1 := "some default"
	if val, ok := conf["key1"]; ok {
		setting1 = val.(string)
	}

	setting2 := 1234
	if val, ok := conf["key2"]; ok {
		setting2 = val.(int)
	}
}
```

I would like to do something like this:
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
With ``cmap`` I can do it.

# License
*MIT*
