package main

import (
	"fmt"
	"reflect"
)

type Result struct {
	Parham    int64
	ParhamRAE float64
	ParhamRE  float64

	Elahe    int64
	ElaheRAE float64
	ElaheRE  float64
}

func (r Result) Compute(name string) Result {
	reflect.ValueOf(&r).Elem().FieldByName(name).SetInt(1378)
	reflect.ValueOf(&r).Elem().FieldByName(fmt.Sprintf("%sRAE", name)).SetFloat(78.07)
	reflect.ValueOf(&r).Elem().FieldByName(fmt.Sprintf("%sRE", name)).SetFloat(73.12)

	return r
}

func main() {
	r := Result{}.Compute("Parham").Compute("Elahe")

	fmt.Printf("%+v\n", r)
}
