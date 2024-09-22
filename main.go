package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name string `json:"name"`
	Age  int    `json:"xyz"`
	City string `json:"-"`//omit this field
	Mobile int `json:",string"`//int convert into string
	Salary float64 `json:",omitempty"`//omitempty field
}

func main() {
	man := Man{Name: "Raju", Age: 23, City: "pune",Mobile: 9625188594}
	fmt.Println(man)
	jdata, _ := json.Marshal(man)
	fmt.Println(string(jdata))

}
