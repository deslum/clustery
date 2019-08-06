package engines

type Engines interface {
	Scan() error
	Len() int
	GetCenterClusterById(clusterId int) (float64, float64)
}
