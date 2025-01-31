package usecase

func (u *SqlUsecase) IsPageVisited(url string) (bool, error) {
	cond, err := u.Repository.IsPageVisited(url)

	if err != nil {
		return false, err
	}

	return cond, nil
}
