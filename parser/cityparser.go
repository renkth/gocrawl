package parser

import (
	"crawlZhenai/engine"
	"fmt"
	"regexp"
)

const cityRe = `<a href="http://album.zhenai.com/u/([0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all{
		result.Items = append(result.Items, "User:"+string(c[2]))
		Url := fmt.Sprintf(`https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=%s&_=%s`,string(c[1]),string(c[1]))
		Url += `&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2F9d4dd220-c93e-48ca-bc7a-eec49bc0809c%2F0%2F0%2F297224893&MmEwMD=4aINzCGB78FZAHXla1ZBObm2RLjr_HEe6PSt2J94Qt_juCtNcq3ljdX_wdCUKQeGC4.RGNIA0rqsv2Oil7W95chtkFGiOmGt4VenxQDUk6qW1.X_PwEIpZABpbuIGmStRjeonZr679Yibj3A7yLmeqkmUVSbaSJ5lgG9kqwddCVn4h6rCm7xpUIZ5PIc6whzUxvoeTM8iW4U6nLxqsMlSKfLJGueNUKoh_ZGTyXxFDnOjdtt0tdz7BOh0_ZP34kTU1s5D7E1j_H885dbn_myHyiEINkGGQdC2zIAgseUywuOubu7aV1jFUqEbHDZeRaSB8CZbrmarDjdDVYMMqwsUj9q6jMvaPUWKXEw6JQOUyAuWrPvBoEO_hNUROAx9UTqDDpg`
		result.Requests = append(result.Requests ,engine.Request{
			Url:        Url, //string(c[1]),
			//ParserFunc: func(c []byte) engine.ParseResult {
			//	return ParseProfile(c, name)
			//},
			//ParserFunc: engine.NilParser,
			ParserFunc: ParseProfile,
		})
	}
	return result
}

