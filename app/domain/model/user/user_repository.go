package user

type Repository interface {
	List() ([]Entity, error)
}
