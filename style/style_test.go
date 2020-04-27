package style

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

type SidebarDefaultSettings struct {
	Settings []*SidebarDefaultSetting `json:"settings"`
}

type SidebarDefaultSetting struct {
	Key               string `json:"key"`
	Value             string `json:"value"`
	ElementShow       bool   `json:"element_show"`
	IsShow            bool   `json:"is_show"`
	IsCanUpdateText   bool   `json:"is_can_update_text"`
	IsCanUpdateIsShow bool   `json:"is_can_update_is_show"`
	Type              string `json:"type"`
}

func ReadStyleJson() *SidebarDefaultSettings {
	data, err := ioutil.ReadFile("/Users/tim/go/src/github.com/timliudream/golangtraining/style/style.json")
	if err != nil {
		log.Fatalln(err)
	}
	sidebarDefaultSettings := new(SidebarDefaultSettings)
	err = json.Unmarshal(data, &sidebarDefaultSettings)
	if err != nil {
		log.Fatalln(err)
	}
	return sidebarDefaultSettings
}

func GetMenusOld() {
	sidebarDefaultSettings := ReadStyleJson()
	for _, setting := range sidebarDefaultSettings.Settings {
		if setting.Type == "menu" {
			// fmt.Printf("key: %s, value: %s\n", setting.Key, setting.Value)
		}
	}
}

var sidebarDefaultSettings *SidebarDefaultSettings

func GetMenusNew() {
	if sidebarDefaultSettings == nil {
		sidebarDefaultSettings = ReadStyleJson()
	}
	for _, setting := range sidebarDefaultSettings.Settings {
		if setting.Type == "menu" {
			// fmt.Printf("key: %s, value: %s\n", setting.Key, setting.Value)
		}
	}
}

func BenchmarkStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMenusNew()
	}
}
