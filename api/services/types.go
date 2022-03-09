package services

type (
	instanceRepository interface {
		GetAllNames(clusterName string) ([]string, error)
	}

	nodeRepository interface {
		GetAllNames(clusterName string) ([]string, error)
	}
)
