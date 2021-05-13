package scylla

import (
	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func SelectQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Displaying Results:")
	q := session.Query("SELECT pid,pname,os,product_location FROM product_data")
	var pid, pname, os, product_location string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select avdb.product", zap.Error(err))
		}
	}()
	for it.Scan(&pid, &pname, &os, &product_location) {
		logger.Info("\t" + pid + " " + pname + ", " + os + ", " + product_location)
	}
}