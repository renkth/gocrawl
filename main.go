package main

import (
	"crawlZhenai/engine"
	"crawlZhenai/parser"
	"crawlZhenai/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       	url,
	//	ParserFunc: parser.ParseCityList,
	//})
}
