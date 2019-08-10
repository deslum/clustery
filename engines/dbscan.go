package engines

import (
	"fmt"
	gpoint "github.com/smira/go-point-clustering"
)

// DBSCAN engine
type DBSCAN struct {
	points          gpoint.PointList
	eps             float64
	epsNeighborhood int
	clusters        []gpoint.Cluster
}

func NewDBSCAN(data [][]float64, eps float64, epsNeighborhood int) *DBSCAN {
	var points = make(gpoint.PointList, len(data))

	for i := 0; i < len(data); i++ {
		points[i][0] = data[i][0]
		points[i][1] = data[i][1]
	}

	return &DBSCAN{
		points:          points,
		eps:             eps,
		epsNeighborhood: epsNeighborhood,
		clusters:        make([]gpoint.Cluster, 0),
	}
}

func (o *DBSCAN) Scan() error {
	clusters, _ := gpoint.DBScan(o.points, o.eps, o.epsNeighborhood)
	if len(clusters) < 1 {
		return fmt.Errorf("clusters length less 1")
	}

	o.clusters = clusters

	return nil
}

func (o *DBSCAN) GetCenterClusterById(clusterId int) (float64, float64) {
	if clusterId < o.Len() {
		center, _, _ := o.clusters[clusterId].CentroidAndBounds(o.points)
		return center[0], center[1]
	}

	return 0, 0
}

func (o *DBSCAN) GetPointsByClusterId(clusterId int) []int {
	if clusterId < o.Len() {
		return o.clusters[clusterId].Points
	}

	return make([]int, 0)
}

func (o *DBSCAN) Len() int {
	return len(o.clusters)
}
