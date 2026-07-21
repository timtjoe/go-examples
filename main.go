package main
import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Ada", Age: 36}
    // use reflect.TypeOf and reflect.ValueOf to iterate
    rv := reflect.ValueOf(p)
    rt := reflect.TypeOf(p)

    for i := 0; i < rt.NumField(); i++ {
        field := rt.Field(i)
        value := rv.Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, value)
    }
}
