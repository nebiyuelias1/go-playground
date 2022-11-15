// ex8
package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main ()   {
	strVar := "200"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}