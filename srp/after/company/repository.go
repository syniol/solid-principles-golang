package company

import "fmt"

type Repository struct {
	dbClient any
}

func (r Repository) FindAllWithName(name string) {
	r.dbClient(fmt.Sprintf("SELECT * FROM companies WHERE name LIKE '%%s%'", name))
}
