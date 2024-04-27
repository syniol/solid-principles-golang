package company

import (
	"database/sql"
)

type repository struct {
	dbClient *sql.DB
}

func (r repository) FindAllWithName(name string) ([]string, error) {
	res, err := r.dbClient.Query("SELECT name FROM companies WHERE name LIKE '?%'", name)
	if err != nil {
		return nil, err
	}

	var companyNames []string
	err = res.Scan(&companyNames)
	if err != nil {
		return nil, err
	}

	return companyNames, nil
}
