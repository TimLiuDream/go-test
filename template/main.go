package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/bangwork/ones-ai-api-common/utils/i18n"
)

// 仪表盘卡片
type Card struct {
	UUID          string `db:"uuid" json:"uuid"`
	TeamUUID      string `db:"team_uuid" json:"team_uuid"`
	DashboardUUID string `db:"dashboard_uuid" json:"dashboard_uuid"`
	Name          string `db:"name" json:"name"`
	Description   string `db:"description" json:"description"`
	Type          int    `db:"type" json:"type"`
	Config        string `db:"config" json:"config"`
	Status        int    `db:"status" json:"status"`
	Owner         string `db:"owner" json:"owner"`
	ObjectType    int    `db:"object_type" json:"object_type"`
	ObjectID      string `db:"object_id" json:"object_id"`
}

// show how to get map key
func main() {
	var path = "/Users/nuc/go/src/github.com/timliudream/go-test/template/tpl.json"
	t, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	t.Funcs(template.FuncMap{
		"tran": i18n.TG,
	})

	value := map[string]interface{}{"ComMap": map[string]string{"需求": "h"}}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, value)
	if err != nil {
		panic(err)
	}
	cardBytes := buf.Bytes()

	// 解析数据
	cards := make([]*Card, 0)
	err = json.Unmarshal(cardBytes, &cards)
	if err != nil {
		panic(err)
	}
	for _, card := range cards {
		fmt.Println(card.Config)
	}
}
