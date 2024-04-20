package before

import "fmt"

type Company struct {
	dbClient any
}

func (c Company) FindCompanyByName(name string) []string {
	// todo: validate name parameter
	// todo: SQL lib
	c.dbClient(fmt.Sprintf("SELECT name, description FROM companies WHERE NAME LIKE '%%s%'", name))
	// todo: display comnapny information

	return nil
}
