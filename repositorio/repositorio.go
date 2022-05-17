package repositorio

import (
	"database/sql"
	"map-read-write/registro"

	_ "github.com/mattn/go-sqlite3"
)

type PersonRepository struct {
	database *sql.DB
}

func NewPersonRepository(dbFilepath string) *PersonRepository {

	database := GetConnection(dbFilepath)
	repoPerson := PersonRepository{database: database}
	return &repoPerson
}

func (pRepo PersonRepository) CreatePerson(pReg registro.Person) registro.Person {

	statement, errorPrepare := pRepo.database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	result, _ := statement.Exec(pReg.Firstname, pReg.Lastname)

	if errorPrepare != nil {
		panic(errorPrepare)
	}
	idP, errorResult := result.LastInsertId()

	if errorResult != nil {
		panic(errorResult)
	}
	pReg.Id = int(idP)
	return pReg

}

func (pRepo PersonRepository) CreatePersonJSON(pReg registro.Person) string {

	statement, errorPrepare := pRepo.database.Prepare("INSERT INTO people_json (ID,JSON_VALUE) VALUES (?,?)")
	if errorPrepare != nil {
		panic(errorPrepare)
	}
	strJson := pReg.ToJson()
	_, errorExec := statement.Exec(pReg.Id, strJson)

	if errorExec != nil {
		panic(errorPrepare)
	}
	return strJson
}

func (pRepo PersonRepository) GetPerson(id int) registro.Person {

	pReg := registro.Person{}
	err := pRepo.database.QueryRow("select  id, firstname, lastname from people where id = ?", id).Scan(&pReg.Id, &pReg.Firstname, &pReg.Lastname)
	if err != nil {
		panic(err)
	}
	return pReg
}

func (pRepo PersonRepository) GetPersonJson(id int) string {

	strJsonPerson := ""
	err := pRepo.database.QueryRow("select JSON_VALUE from people_json  where ID= ?", id).Scan(&strJsonPerson)
	if err != nil {
		panic(err)
	}
	return strJsonPerson
}

func (pRepo PersonRepository) SelectAllPerson() []registro.Person {

	var arrayPerson = []registro.Person{}
	row, err := pRepo.database.Query("select  id, firstname, lastname  from people")
	if err != nil {
		panic(err)
	}
	defer row.Close()
	for row.Next() {
		person := registro.Person{}
		row.Scan(&person.Id, &person.Firstname, &person.Lastname)
		arrayPerson = append(arrayPerson, person)
	}
	return arrayPerson
}

func (pRepo PersonRepository) Close() {

	pRepo.database.Close()
}
