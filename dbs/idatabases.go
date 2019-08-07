package dbs

type IDatabases interface {
	GetData(query string) ([][]float64, error)
}
