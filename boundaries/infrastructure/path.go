package boundaries

type PathInfrastructure interface {
	Join(elem ...string) string
	Split(path string) (dir, file string)
}
