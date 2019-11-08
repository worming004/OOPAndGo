package main

import (
	"fmt"
	"io"
	"strings"
)

type Todo struct{ Id int }

func (td Todo) String() string {
	fmt.Sprintf("%v", td.Id)
}

type Parser interface {
	ParseBytes(str io.Reader) Todo
}

type XmlParser struct{}

func (p XmlParser) ParseStream(str io.Reader) Todo {
	result Todo
	xml.Unmarshal()
	return Todo{Id: 1}
}

type JsonParser struct{}

func (p JsonParser) ParseStream(str io.Reader) Todo {
	//
	return Todo{Id: 1}
}

func main() {
	xmlReader := strings.NewReader("<Todo><Id>1</Id></Todo>")
	jsonReader := strings.NewReader("\"Todo\": {\"Id\": 2}")

	xmlParser := XmlParser{}
	jsonParser := JsonParser{}

	todo := xmlParser.ParseStream(xmlReader)
	fmt.Println(todo)
	todo = jsonParser.ParseStream(jsonReader)
	fmt.Println(todo)
}
