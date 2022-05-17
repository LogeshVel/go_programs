package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Id   int    `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Age  int    `bson:"age"`
}

func main() {
	var a Person
	// ee, _ := json.Marshal(a)

	ee, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(ee))
	fmt.Printf("ee: %v\n", ee)

	st := reflect.TypeOf(a)
	fmt.Printf("st: %v\n", st.NumField())
	field0 := st.Field(0)
	fmt.Printf("field0: %v\n", field0)
	fmt.Printf("field0.Tag.Get(\"bson\"): %v\n", field0.Tag.Get("bson"))

}
