package main

import (
	"fmt"
	"map-read-write/registro"
	"map-read-write/repositorio"
	"map-read-write/writer"
	"time"
)

/***
1  - Separar função para conexao = OK
2 -  Remover orm - não retorna a conexao padrao = OK
3 -  Criar banco de dados = OK
4 -  Prepare com arquivos SQL = OK
5 -  Refatorar para sctruct = OK
6 -  Limpeza codigo
*/
func main() {
	pathDatabase := repositorio.CreateDatabase()
	personRepository := repositorio.NewPersonRepository(pathDatabase)

	var pReg registro.Person
	pReg = personRepository.CreatePerson(registro.Person{
		Firstname: "First Teste",
		Lastname:  "Last Test",
	})

	pReg = personRepository.CreatePerson(registro.Person{
		Firstname: "First Teste 1",
		Lastname:  "Last Test 1",
	})

	birthDate, errTime := time.Parse("2006-01-02", "1975-06-01")
	fmt.Println(birthDate)
	fmt.Println(errTime)

	pReg = personRepository.CreatePerson(registro.Person{
		Firstname:    "First Teste 2",
		Lastname:     "Last Test 2",
		AnnualIncome: 2000000.5050404,
		BirthDate:    birthDate,
	})

	//INSERT JSON  JSON
	strJson := pReg.ToJson()
	fmt.Println(strJson)
	personRepository.CreatePersonJSON(pReg)

	persons := personRepository.SelectAllPerson()
	writer.WriterPersons(persons)
	p := personRepository.GetPerson(pReg.Id)
	fmt.Println(p)

	strJsonPerson := personRepository.GetPersonJson(pReg.Id)
	fmt.Println(strJsonPerson)

}
