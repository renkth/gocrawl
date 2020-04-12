package parser

import (
	"crawlZhenai/engine"
	"crawlZhenai/model"
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)


func ParseProfile(contents []byte) engine.ParseResult{
	//match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	profile := parseJson(contents)
	fmt.Println(profile)
	//if len(match) >= 2 {
	//	json := match[1][0]
	//	//fmt.Printf("json:%s\n",json)
	//	profile := parseJson(json)
	//	fmt.Println(profile)
	//	//profile.Name = name
	//	//result.Items = append(result.Items, profile)
	//	//fmt.Println(result)
	//}

	return result
}

func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析失败")
	}
	infos, err := res.Get("data").Get("basicInfo").Array()

	var profile model.Profile

	for k,v := range infos{
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}


	}
	return profile
}