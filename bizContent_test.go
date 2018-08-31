package paysdk

import (
	"fmt"
	"testing"
)

func TestBizContent(t *testing.T) {
	var bizContent= BizContent{}

	//method add
	bizContent.Add("index1","value1")
	fmt.Printf("%#v",bizContent)

	//method stringify
	str := bizContent.ToString()
	if str !=`{"index1":"value1"}`{
		t.Errorf(`expected value is {"index1":"value1"},real data is %s`,str)
	}
	//method get
	data := bizContent.Get("index1")
	if data != "value1" {
		t.Errorf(`expected value is "value1",real data is %s`,data)
	}
	//method delete
	bizContent.Del("index1")
	if bizContent.Get("index1") !="" {
		t.Errorf(`del func errror`)
	}

	var bizContenter BizContenter= BizContent{}
	bizContenter.Add("index2","value2")
	data2 := bizContenter.Get("index2")
	if data2 != "value2" {
		t.Errorf(`expected value is "value1",real data is %s`,data2)
	}
}