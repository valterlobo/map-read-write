package writer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"map-read-write/registro"
)

func WriterPersons(persons []registro.Person) {

	//var buffer *bytes.Buffer
	buffer := bytes.Buffer{}
	w := NewWriter(&buffer, ArrayToStringLine)
	for _ , p := range persons {
		newP := p
		w.AddRegister(&newP)
	}
	w.WriteRegister()
	err := ioutil.WriteFile("persons.txt", buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo gerado:" + "persons.txt")

}