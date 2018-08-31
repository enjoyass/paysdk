package paysdk
import (
	"encoding/json"
)

type BizContent map[string]string

type BizContenter interface {
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
func (bc BizContent) Get(key string)string{
	if bc == nil {
		return ""
	}
	value,ok :=bc[key]
	if !ok {
		return ""
	}
	return value
}
// Add adds the value to key. It appends to any existing
// values associated with key.
func (bc BizContent) Add(key,value string){
	bc[key]=value
}
// Set sets the key to value. It replaces any existing
// values.
func (bc BizContent)Set(key,value string){
	bc[key]=value
}
// Del deletes the values associated with key
func (bc BizContent)Del(key string){
	delete(bc,key)
}
// stringify
func (bc BizContent)ToString()string{
	mjson,_ :=json.Marshal(bc)
	mString :=string(mjson)
	return mString
}