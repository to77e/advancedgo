package hw09

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"omitempty,address"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

type User struct {
	Id     int     `properties:"id"`
	Login  string  `properties:"login"`
	Name   string  `properties:"name,omitempty"`
	Email  string  `properties:"email,omitempty"`
	Age    int     `properties:"age,omitempty"`
	Active bool    `properties:"active"`
	Score  float32 `properties:"score"`
}

func Serialize(person interface{}) string {
	var result []string
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	for i := 0; i < t.NumField(); i++ {
		tFieldTag := t.Field(i).Tag
		value, present := tFieldTag.Lookup("properties")
		if !present || (strings.Contains(value, "omitempty") && v.Field(i).IsZero()) {
			continue
		}
		parts := strings.Split(value, ",")
		name := ""
		for _, part := range parts {
			if part != "omitempty" && part != "" {
				name = part
				break
			}
		}
		vField := v.Field(i)
		result = append(result, fmt.Sprintf("%s=%v", name, vField))
	}
	return strings.Join(result, "\n")
}
