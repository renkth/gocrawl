package main

import (
	"crawlZhenai/engine"
	"crawlZhenai/parser"
	"log"
)

func check(err error){
	if err != nil {
		log.Println(err)
	}
	return
}

func main(){
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:       	url,
		ParserFunc: parser.ParseCityList,
	})
}