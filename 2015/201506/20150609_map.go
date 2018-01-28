package main

import (
	"fmt"
)

type DATA_TYPE uint32

const (
	BUILDING DATA_TYPE = iota
	BUILDINGTREE
	CARD
	CARDSET
	ITEM
	LANCH
	OPTION
	PROJECTILE
	RES
	SCENE
	SKILL
	STATE
	TASK
	TOOL
	TRIGGER
)

type ResData struct {
	Id                string
	Name              string
	ModelId           string
	ModelScale        float64
	EffectId          string
	EffectHandlePoint string
	EffectPlayTimes   int64
	EffectPlayTime    float64
	EffectScale       float64
	AudioId           string
	AudioPlayTimes    int64
	AudioPlayTime     float64
	NoCache           bool
}

type Other struct {
	Id   string
	Name string
}

func main() {
	all := map[string]string{
		"a": "astring",
		"b": "bstring",
	}
	fmt.Printf("%#v\n", all)

	info := map[DATA_TYPE]interface{}{
		BUILDINGTREE: &ResData{
			Id: "10",
		},
		TOOL: &Other{
			Id:   "otherId",
			Name: "OtherName",
		},
	}

	fmt.Printf("%#v\n", info)
}
