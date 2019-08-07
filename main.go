package main

import (
	"clustery/engines"
	"github.com/golang/glog"

	"clustery/consts"
	"clustery/dbs"
)

func main() {
	glog.Flush()

	database, err := dbs.Connect(
		"clickhouse",
		"",
	)

	if err != nil {
		glog.Fatalf("dbs.Connect() error %v", err)
	}

	data, err := dbs.NewSQL(database).GetData(consts.QueryTemplate)
	if err != nil {
		glog.Fatalf("database.GetData() error %v", err)
	}

	dbscan := engines.NewDBSCAN(data, consts.Eps, consts.EpsNeighborhood)
	if err = dbscan.Scan(); err != nil {
		glog.Fatalf("dbscan.Scan() error %v", err)
	}

	kmeans := engines.NewKMeans(data, consts.Parts)
	if err = kmeans.Scan(); err != nil {
		glog.Fatalf("kmeans.Scan() error %v", err)
	}

	for _, engine := range []engines.Engines{dbscan, kmeans} {
		if err = engine.Scan(); err != nil {
			glog.Fatalf("engine.Scan() error %v", err)
		}

		for clusterID := 0; clusterID < kmeans.Len(); clusterID++ {
			var x, y = engine.GetCenterClusterById(clusterID)
			glog.Infof("ClusterID %v X: %v Y: %v", clusterID, x, y)
		}

	}
}
