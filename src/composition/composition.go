package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// )

// type Todo struct{ Id int }

// func (td Todo) String() string {
// 	return fmt.Sprintf("%v", td.Id)
// }

// type Parser interface {
// 	ParseBytes(content []byte) (Todo, error)
// }

// type XmlParser struct{}

// func (p XmlParser) ParseStream(content []byte) (Todo, error) {
// 	result := Todo{}
// 	err := xml.Unmarshal(content, &result)
// 	return result, err
// }

// type JsonParser struct{}

// func (p JsonParser) ParseStream(content []byte) (Todo, error) {
// 	result := Todo{}
// 	err := json.Unmarshal(content, &result)
// 	return result, err
// }

// func main() {
// 	xmlReader := []byte("<Todo><Id>1</Id></Todo>")
// 	jsonReader := []byte(`{"Id": 2}`)

// 	xmlParser := XmlParser{}
// 	jsonParser := JsonParser{}

// 	todo, err := xmlParser.ParseStream(xmlReader)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(todo)
// 	todo, err = jsonParser.ParseStream(jsonReader)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(todo)
// }
