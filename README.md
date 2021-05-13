# go-scylla-poc

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://docs.scylladb.com/using-scylla/drivers/cql-drivers/scylla-go-driver/)

### Proof of Concept
Implementation of ScyllaDB with golang

<img src="https://www.scylladb.com/wp-content/uploads/BlinkJump_Animation_07_Shadow_193b6c-1.gif" width="152" height="191"/>

### Implementation and Pre-requisites

Scylladb setup via Docker
https://hub.docker.com/r/scylladb/scylla/

Zap Logger
https://github.com/uber-go/zap
https://sunitc.dev/2019/05/27/adding-uber-go-zap-logger-to-golang-project/#1.0

GoCQL
https://docs.scylladb.com/using-scylla/drivers/cql-drivers/scylla-go-driver/
https://github.com/scylladb/gocql

```
go get github.com/gocql/gocql
go get go.uber.org/zap
```

Created Keyspace and Tables in Scylla cluster and interacted using Golang to do basic query and CRUD operations.

### Source of References
https://www.scylladb.com/2018/02/12/gocql-gocqlx-scylla/

https://university.scylladb.com/courses/the-mutant-monitoring-system-training-course/lessons/golang-and-scylla-part-1/

https://docs.scylladb.com/getting-started/ddl/
