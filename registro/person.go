package registro

import "strconv"

type Person struct {
	Id        int  `db:"id" `
	Firstname string `db:"firstname" `
	Lastname  string `db:"lastname" `
}

func (u *Person) TableName() string {
	return "people"
}

func (u *Person) PK() string {
	return "ID"
}

func (r *Person) Type() string {
	return "P00"
}

func (r *Person) ToArray() []string {

	return []string{
		"P00",
		strconv.Itoa(r.Id),
		r.Firstname,
		r.Lastname,
	}
}
