package application

func (i interactor) Tags() ([]string, error) {
	return i.tagsRW.GetAll()
}
