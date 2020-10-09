package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Test struct {
	All *All `yaml:"all"`
}

type All struct {
	Hosts    *Hosts    `yaml:"hosts"`
	Children *Children `yaml:"children"`
}

type Hosts struct {
	Node1 *Node1 `yaml:"node1"`
}

type Node1 struct {
	AnsibleHost string `yaml:"ansible_host"`
	Ip          string `yaml:"ip"`
	AccessIp    string `yaml:"access_ip"`
}

type Children struct {
	KubeMaster *KubeMaster `yaml:"kube-master"`
}

type KubeMaster struct {
	KubeMasterHosts *KubeMasterHosts `yaml:"hosts"`
}

type KubeMasterHosts struct {
	Node1 string `yaml:"node1"`
}

func main() {
	func2()
}

func func1() {
	bytes, err := ioutil.ReadFile("/Users/tim/go/src/github.com/timliudream/golangtraining/yaml/test.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	test := new(Test)
	err = yaml.Unmarshal(bytes, &test)
	if err != nil {
		log.Fatalln(err)
	}

	bytes, err = yaml.Marshal(&test)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytes))
}

func func2() {
	// 因为原来的yaml的password是没有单引号的，会导致解析错误，所以加上单引号
	bytes, err := ioutil.ReadFile("/Users/tim/go/src/github.com/timliudream/go-test/yaml/test1.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatalln(err)
	}
	m1 := m["inbound"].(map[interface{}]interface{})
	fmt.Printf("%+v\n", m1["password"])
	fmt.Println(m)
}
