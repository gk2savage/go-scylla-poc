package main

import (
	"go-scylla-poc/internal/log"
	"go-scylla-poc/internal/scylla"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func main() {
	logger := log.CreateLogger("info")

	cluster := scylla.CreateCluster(gocql.Quorum, "avdb", "scylla-node1", "scylla-node2", "scylla-node3")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		logger.Fatal("unable to connect to scylla", zap.Error(err))
	}
	defer session.Close()

	scylla.SelectQuery(session, logger)
	insertQuery(session, logger)
	scylla.SelectQuery(session, logger)
	deleteQuery(session, logger)
	scylla.SelectQuery(session, logger)
}

func insertQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Inserting qsecurity")
	if err := session.Query("INSERT INTO product_data (pid,pname,os,product_location) VALUES ('AR121P','qsecurity','Windows', 'http://www.av.com/qsec')").Exec(); err != nil {
		logger.Error("insert avdb.product_data", zap.Error(err))
	}
}

func deleteQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Deleting qsecurity")
	if err := session.Query("DELETE FROM product_data WHERE pid = 'AR121P' and pname = 'qsecurity'").Exec(); err != nil {
		logger.Error("delete avdb.product_data", zap.Error(err))
	}
}
