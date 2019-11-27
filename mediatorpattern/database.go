package main

import "fmt"

const (
	Mysql         = "mysql"
	Redis         = "redis"
	ElasticSearch = "elasticsearch"
)

// Colleague interface
type Database interface {
	addData(data string)
	add(data string)
}

type SimpleDatabase struct {
	mediator *Mediator
}

type MysqlDatabase struct {
	SimpleDatabase
	dataList []string
}

func (mysql *MysqlDatabase) addData(data string) {
	fmt.Println("mysql add data: " + data)
	if mysql.dataList == nil {
		mysql.dataList = make([]string, 0)
	}
	mysql.dataList = append(mysql.dataList, data)
}

func (mysql *MysqlDatabase) add(data string) {
	mysql.addData(data)
	mysql.mediator.sync(Mysql, data)
}

func (mysql *MysqlDatabase) query() {
	fmt.Println("mysql query data: ", mysql.dataList)
}

type RedisDatabase struct {
	SimpleDatabase
	dataList []string
}

func (redis *RedisDatabase) addData(data string) {
	fmt.Println("redis add data: " + data)
	if redis.dataList == nil {
		redis.dataList = make([]string, 0)
	}
	redis.dataList = append(redis.dataList, data)
}

func (redis *RedisDatabase) add(data string) {
	redis.addData(data)
	redis.mediator.sync(Redis, data)
}

func (redis *RedisDatabase) cache() {
	fmt.Println("redis cache data: ", redis.dataList)
}

type ElasticSearchDatabase struct {
	SimpleDatabase
	dataList []string
}

func (elasticsearch *ElasticSearchDatabase) addData(data string) {
	fmt.Println("elasticsearch add data: " + data)
	if elasticsearch.dataList == nil {
		elasticsearch.dataList = make([]string, 0)
	}
	elasticsearch.dataList = append(elasticsearch.dataList, data)
}

func (elasticsearch *ElasticSearchDatabase) add(data string) {
	elasticsearch.addData(data)
	elasticsearch.mediator.sync(ElasticSearch, data)
}

func (elasticsearch *ElasticSearchDatabase) count() {
	fmt.Printf("elasticsearch count: %d ,query data: %v \n", len(elasticsearch.dataList), elasticsearch.dataList)
}
