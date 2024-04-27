package company

import "fmt"

type repository struct {
	dbClient any
}

func (r repository) FindAllWithName(name string) {
	r.dbClient(fmt.Sprintf("SELECT * FROM companies WHERE name LIKE '%%s%'", name))
}
