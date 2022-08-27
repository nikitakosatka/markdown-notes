package repository

func Create(note *Note) error {
	return dbConnect.Insert(note)
}

func GetAll(notes *[]Note) error {
	return dbConnect.Model(notes).Select()
}

func Read(id string) (*Note, error) {
	note := &Note{ID: id}
	err := dbConnect.Select(note)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func Update(note *Note) error {
	_, err := dbConnect.Model(note).WherePK().Update()
	return err
}

func Remove(note *Note) error {
	return dbConnect.Delete(note)
}
