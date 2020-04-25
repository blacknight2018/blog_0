package utils

import "encoding/json"

func SetJSONKeyValue(d interface{}, key string, val interface{}) (interface{}, bool) {
	var m = make(map[string]interface{})
	js, err := json.Marshal(d)
	if nil != err {
		return nil, false
	}
	err = json.Unmarshal(js, &m)
	if nil != err {
		return nil, false
	}
	m[key] = val
	return m, true
}
