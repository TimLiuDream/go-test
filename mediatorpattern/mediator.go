package main

// Mediator interface
type MediatorInterface interface {
	sync(databaseName string, data string)
}

type Mediator struct {
	mysqlDatabase         *MysqlDatabase
	redisDatabase         *RedisDatabase
	elasticSearchDatabase *ElasticSearchDatabase
}

func (m *Mediator) sync(databaseName string, data string) {
	switch databaseName {
	case Mysql:
		m.redisDatabase.addData(data)
		m.elasticSearchDatabase.addData(data)
	case Redis:
	case ElasticSearch:
		m.mysqlDatabase.addData(data)
	}
}
