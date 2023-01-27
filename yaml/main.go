package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/goccy/go-yaml" //使用这个库，实现 anchor 的的解释，golang 内置的不支持
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
	// func2()
	sample()
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

var funcMap = map[string]func() interface{}{
	"onesUUID4": func() interface{} {
		return "01234"
	},
}

var func1Map = map[string]func(interface{}) interface{}{
	"calTime": func(v interface{}) interface{} {
		vValue := v.(int64)
		return time.Now().Unix() + vValue
	},
}

type MyID string

// 自定义简单生成的能力
func (id *MyID) UnmarshalYAML(b []byte) error {
	v := funcMap[string(b)]()
	*id = MyID(v.(string))
	return nil
}

type Product struct {
	UUID string
}

// 用初始化的上下文，自定义数据的反序列化逻辑
func (product *Product) UnmarshalYAML(ctx context.Context, data []byte) error {
	v := ctx.Value("db")
	product.UUID = v.(string) + "uuid001"
	return nil
}

type SampleTime struct {
	TimeId    string `yaml:"time_id"`
	TimeValue int64  `yaml:"time_value"`
}

func (st *SampleTime) UnmarshalYAML(b []byte) error {
	//p := new(Project)
	//err := yaml.Unmarshal(b, &p)
	//if err != nil {
	//	panic(err)
	//}
	//v := func1Map[p.TimeId](p.TimeValue)
	//p.TimeValue = v.(int64)
	return nil
}

type Project struct {
	ID MyID `yaml:"id"`
	SampleTime
	Name string `yaml:"name"`
}

type Temp1 struct {
	Project Project `yaml:"project"`
	Tasks   []Task  `yaml:"tasks"`
}

type Task struct {
	ID      MyID     `yaml:"id"`
	Project *Project `yaml:"project"`
	Name    string   `yaml:"name"`
	Product *Product `yaml:"product"`
}

var _ context.Context = (*MyCtx)(nil)

type MyCtx struct {
	DB string
}

func (m MyCtx) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (m MyCtx) Done() <-chan struct{} {
	panic("implement me")
}

func (m MyCtx) Err() error {
	panic("implement me")
}

func (m MyCtx) Value(key interface{}) interface{} {
	return m.DB
}

func sample() {
	data, err := ioutil.ReadFile("/Users/tim/go/src/github.com/timliudream/go-test/yaml/test2.yaml")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	//var v interface{}
	var v Temp1
	//if err := yaml.Unmarshal(data, &v); err != nil {
	//	fmt.Printf("%v", err)
	//	return
	//}
	ctx := &MyCtx{DB: "mydb"}
	if err := yaml.UnmarshalContext(ctx, data, &v); err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v", v)
}
