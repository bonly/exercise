package main
import("fmt"
        "reflect"
) 
func main() { 
	// iterate through the attributes of a Data Model instance
	for name, mtype := range attributes(&Dish{}) {// InstantiateData Modeland get attributes
		fmt.Printf("Name:  %s,  Type:  %s\n", name, mtype);
	}
}

// Data Model
type Dish struct{
	Id_Val int;
	Name   string;
	Origin string;
	Query  func();
};

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type{
	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type);

	typ := reflect.TypeOf(m);

	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem();
	}

	// Only structs are supported so return an empty result if the passed objectis nota struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind());
		return attrs;
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i);
		if !p.Anonymous {
			attrs[p.Name] = p.Type;
		}
	}

	return attrs;
}
