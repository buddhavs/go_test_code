package util

import (
	"fmt"
	"reflect"
)

func PrintType(in interface{}) {
	fmt.Println(reflect.TypeOf(in))
}
