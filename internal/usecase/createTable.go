package usecase

func (u *SqlUsecase) CreateTable() error {
	return u.Repository.CreateTable()
}
