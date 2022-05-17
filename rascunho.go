package main

import (
	"fmt"
	"map-read-write/registro"
	"map-read-write/repositorio"
	"map-read-write/writer"
	"strconv"
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

	birthDate, errTime := time.Parse("2006-01-01", "2001-08-01")
	fmt.Println(birthDate)
	fmt.Println(errTime)

	for i := 0; i < 100; i++ {

		//person 1
		pReg1 := personRepository.CreatePerson(registro.Person{
			Firstname:    "First Teste " + strconv.Itoa(i),
			Lastname:     "Lastname    " + strconv.Itoa(i),
			AnnualIncome: 2000000.5050404,
			BirthDate:    birthDate,
		})

		//person 2
		pReg2 := personRepository.CreatePerson(registro.Person{
			Firstname:    "Person 2" + strconv.Itoa(i),
			Lastname:     "Last person 2" + strconv.Itoa(i),
			AnnualIncome: 2000000.5050404,
			BirthDate:    birthDate,
		})

		//INSERT JSON - person 1
		personRepository.CreatePersonJSON(pReg1)

		strJsonPerson1 := personRepository.GetPersonJson(pReg1.Id)
		fmt.Println(strJsonPerson1)

		//INSERT JSON  person 2
		personRepository.CreatePersonJSON(pReg2)
		strJsonPerson2 := personRepository.GetPersonJson(pReg2.Id)
		fmt.Println(strJsonPerson2)

	}

	persons := personRepository.SelectAllPerson()
	writer.WriterPersons(persons)

}
