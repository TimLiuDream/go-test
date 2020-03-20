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
