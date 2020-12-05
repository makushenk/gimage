package boundaries

type ImageInfrastructure interface {
	GenerateThumbnail(source, target string, x, y int, width, height int) error
}
