package paysdk

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	var model= Model{}

	//method add
	model.Add("index1","value1")
	fmt.Printf("%#v",model)

	//method stringify
	str := model.ToString()
	if str !=`{"index1":"value1"}`{
		t.Errorf(`expected value is {"index1":"value1"},real data is %s`,str)
	}
	//method get
	data := model.Get("index1")
	if data != "value1" {
		t.Errorf(`expected value is "value1",real data is %s`,data)
	}
	//method delete
	model.Del("index1")
	if model.Get("index1") !="" {
		t.Errorf(`del func errror`)
	}

	var modeler Modeler= Model{}
	modeler.Add("index2","value2")
	data2 := modeler.Get("index2")
	if data2 != "value2" {
		t.Errorf(`expected value is "value1",real data is %s`,data2)
	}
}