package main

import (
	"fmt"
	"github.com/fatih/structs"
	"net/http"
)

func main() {
	Examples()
}

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

var (
	server = &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}
)

func Examples() {
	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	_ = structs.Map(server)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	_ = structs.Values(server)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	_ = structs.Names(server)

	// Convert the values of a struct to a []*Field
	// (see "Field methods" for more info about fields)
	_ = structs.Fields(server)

	// Return the struct name => "Server"
	_ = structs.Name(server)

	// Check if any field of a struct is initialized or not.
	_ = structs.HasZero(server)

	// Check if all fields of a struct is initialized or not.
	_ = structs.IsZero(server)

	// Check if server is a struct or a pointer to struct
	_ = structs.IsStruct(server)
}

func Fieldmethods() {
	s := structs.New(server)

	// Get the Field struct for the "Name" field
	name := s.Field("Name")

	// Get the underlying value,  value => "gopher"
	_ = name.Value().(string)

	// Set the field's value
	name.Set("another gopher")

	// Get the field's kind, kind =>  "string"
	name.Kind()

	// Check if the field is exported or not
	if name.IsExported() {
		fmt.Println("Name field is exported")
	}

	// Check if the value is a zero value, such as "" for string, 0 for int
	if !name.IsZero() {
		fmt.Println("Name is initialized")
	}

	// Check if the field is an anonymous (embedded) field
	if !name.IsEmbedded() {
		fmt.Println("Name is not an embedded field")
	}

	// Get the Field's tag value for tag name "json", tag value => "name,omitempty"
	_ = name.Tag("json")
}
