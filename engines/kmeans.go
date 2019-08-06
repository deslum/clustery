package engines

import (
	"fmt"
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

type Kmeans struct {
	kmeans   kmeans.Kmeans
	points   clusters.Observations
	clusters clusters.Clusters
	parts    int
}

func NewKMeans(data [][]float64, parts int) *Kmeans {

	var points clusters.Observations

	for i := 0; i < len(data); i++ {
		points = append(points, clusters.Coordinates{
			data[i][0],
			data[i][1],
		})
	}

	return &Kmeans{
		kmeans: kmeans.New(),
		parts:  parts,
		points: points,
	}
}

func (o *Kmeans) Scan() error {
	clusters, err := o.kmeans.Partition(o.points, o.parts)
	if err != nil {
		return err
	}

	if len(clusters) <= 1 {
		return fmt.Errorf("clusters length less 1")
	}

	o.clusters = clusters

	return nil
}

func (o *Kmeans) GetCenterClusterById(clusterId int) (float64, float64) {
	if clusterId < o.Len() {
		return o.clusters[clusterId].Center[0], o.clusters[clusterId].Center[1]
	}

	return 0, 0
}

func (o *Kmeans) Len() int {
	return len(o.clusters)
}
