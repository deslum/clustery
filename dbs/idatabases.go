package dbs

type IDatabases interface {
	GetData() ([][]float64, error)
}
