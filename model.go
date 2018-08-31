package paysdk
import (
	"encoding/json"
)

type Model map[string]string

type Modeler interface {
	Get(key string)string
	Add(key,value string)
	Set(key,value string)
	Del(key string)
	ToString() string

}
// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (m Model) Get(key string)string{
	if m == nil {
		return ""
	}
	value,ok :=m[key]
	if !ok {
		return ""
	}
	return value
}
// Add adds the value to key. It appends to any existing
// values associated with key.
func (m Model) Add(key,value string){
	m[key]=value
}
// Set sets the key to value. It replaces any existing
// values.
func (m Model)Set(key,value string){
	m[key]=value
}
// Del deletes the values associated with key
func (m Model)Del(key string){
	delete(m,key)
}
// stringify
func (m Model)ToString()string{
	mjson,_ :=json.Marshal(m)
	mString :=string(mjson)
	return mString
}