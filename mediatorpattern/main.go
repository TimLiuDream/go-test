package main

import "fmt"

func main() {
	syncMediator := new(Mediator)
	mysqlDatabase := new(MysqlDatabase)
	mysqlDatabase.mediator = syncMediator
	redisDatabase := new(RedisDatabase)
	redisDatabase.mediator = syncMediator
	esDatabase := new(ElasticSearchDatabase)
	esDatabase.mediator = syncMediator

	syncMediator.mysqlDatabase = mysqlDatabase
	syncMediator.redisDatabase = redisDatabase
	syncMediator.elasticSearchDatabase = esDatabase

	fmt.Println("\n---------mysql 添加数据 1，将同步到Redis和ES中-----------")
	mysqlDatabase.add("1")
	mysqlDatabase.query()
	redisDatabase.cache()
	esDatabase.count()
	fmt.Println("\n---------Redis添加数据 2，将不同步到其它数据库-----------")
	redisDatabase.add("2")
	mysqlDatabase.query()
	redisDatabase.cache()
	esDatabase.count()
	fmt.Println("\n---------ES 添加数据 3，只同步到 Mysql-----------")
	esDatabase.add("3")
	mysqlDatabase.query()
	redisDatabase.cache()
	esDatabase.count()
}
