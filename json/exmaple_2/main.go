package main

import (
	"fmt"
	"reflect"
)

type TagPerson struct {
	Id   int    `json:"id"`
	Name string `json:"full_name"`
	Age  int    `json:"age"`
}

func main() {

	var b TagPerson
	printJSONTags(b)

}

func printJSONTags(tp TagPerson) {
	st := reflect.TypeOf(tp)
	numFields := st.NumField()
	fmt.Println("JSON Tags  of", st)
	fmt.Printf("Field Name -> JSON Tag\n")
	for i := 0; i < numFields; i++ {
		field := st.Field(i)
		fmt.Printf("%-10v -> %v\n", field.Name, field.Tag.Get("json"))
	}
}
