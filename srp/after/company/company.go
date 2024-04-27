package company

func FindCompanyWithName(
	name string,
) ([]string, error) {
	return findCompanyWithName(
		name,
		ValidateCompanyName,
		nil,
	)
}

func findCompanyWithName(
	name string,
	validateCompanyName func(string) error,
	repository *repository,
) ([]string, error) {
	err := validateCompanyName(name)
	if err != nil {
		return nil, err
	}

	repository.FindAllWithName(name)

	return nil, nil
}
