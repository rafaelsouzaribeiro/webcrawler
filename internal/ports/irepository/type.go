package irepository

type IRepository interface {
	CreateTable() error
	IsPageVisited(url string) (bool, error)
	InsertVisitedPage(url string) error
}
