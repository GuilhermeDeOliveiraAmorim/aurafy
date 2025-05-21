package repositoriesinterfaces

type ImageRepositoryInterface interface {
	SaveImage(pathImage string) (string, error)
}
