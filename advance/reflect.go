package advance

import (
	"fmt"
	"reflect"
)

// Iterate struct fields:
func dump(v interface{}){
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}

// Struct tags - metadata read by libraries
type User struct {
	Name string `json:"name" db:"user_name" validate:"required"`
	Email string `json:"email"`
}

func reflection()  {
	f, _ := reflect.TypeOf(User{}).FieldByName("Name")
	f.Tag.Get("json")
	f.Tag.Get("db")
	f.Tag.Get("validate")

	// Setting fields (must use a pointer + elem):
	p := &User{}
	rv := reflect.ValueOf(p).Elem()
	rv.FieldByName("Name").SetString("Ada")
}