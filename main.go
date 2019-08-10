package main

import (
	"clustery/config"
	"clustery/engines"
	"flag"
	"github.com/golang/glog"

	"clustery/consts"
	"clustery/dbs"
)

func main() {
	flag.Parse()
	glog.Flush()

	cfg, err := config.NewJsonConfig().Fill()
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info(cfg)

	//database, err := dbs.Connect(
	//	"clickhouse",
	//	"",
	//)
	//
	//if err != nil {
	//	glog.Fatalf("dbs.Connect() error %v", err)
	//}

	data, err := dbs.NewMock().GetData("")
	if err != nil {
		glog.Fatalf("database.GetData() error %v", err)
	}

	var alg engines.Engines
	switch cfg.Algorithm.Name {
	case "dbscan":
		alg = engines.NewDBSCAN(data, cfg.Algorithm.Epsilon, cfg.Algorithm.MinPoints)
	case "kmeans":
		alg = engines.NewKMeans(data, consts.Parts)
	default:
		glog.Fatal("Algorithm %v not exist", cfg.Algorithm.Name)
	}

	if err = alg.Scan(); err != nil {
		glog.Fatalf("kmeans.Scan() error %v", err)
	}

	for clusterID := 0; clusterID < alg.Len(); clusterID++ {
		var x, y = alg.GetCenterClusterById(clusterID)
		glog.Infof("ClusterID %v X: %v Y: %v", clusterID, x, y)
	}

}
