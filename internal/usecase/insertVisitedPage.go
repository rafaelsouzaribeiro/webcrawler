package usecase

func (u *SqlUsecase) InsertVisitedPage(url string) error {
	return u.Repository.InsertVisitedPage(url)
}
