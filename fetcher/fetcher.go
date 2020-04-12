package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetch(url string)([]byte, error) {
	//resp, err := http.Get(url)
	//if err != nil{
	//	return nil, err
	//}
	request, _ := http.NewRequest(http.MethodGet,url,nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//request.Header.Add("cookie","sid=bed99bb6-f9a7-4e6a-bc7f-cf7c440e8d3e; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1586100231; FSSBBIl1UgzbN7N443S=EqBb.YW0J1J7FA2kvKahgB_Lhf15HnObvTWQoh_RluYHw6UhkN0DdFYs5fiYDutm; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1586143952; FSSBBIl1UgzbN7N443T=4SIyN6GdLIF7ViXV0sZdvvmTwNjp5iEzuOSFfW9.3h__66tyDr3VX5XjR5Co1wean_.x9LIqaqqkO9O8qzW0_khFo70.1UC5fkIfQ0f7qp7Hi_ynbBbE9_E9rDbeoZ_.OoHlBxeFlwWFf04SctJMerzuq7QeGhkAze1ekmlSiQ8m4BNxfP5iZb5fDx1y6iLn0nR9WtDV1Sp8qZ1FRWc6hnnplpQqDEu9yWW4Z.NRXGdtcv0tMgpyNcCteFdTKKZnoXrgpGabT.gCd471j2f.PQcxnwEGj2u6ZvihsJW2CNMFEc0LsQA_cZ5PIt.YKaZOdn4vNhx2MIgwGBT9bShQ9le52XU75chwt0v..aYlBhltdXruINyYAEICOJcsRRu83lOQ")
	resp,_:=http.DefaultClient.Do(request)
	time.Sleep(1*time.Second)

	defer resp.Body.Close()



	if resp.StatusCode != http.StatusOK {
		//b, _ := ioutil.ReadAll(resp.Body)
		//ioutil.WriteFile("a.html",b,0666)
		//fmt.Printf("%s",b)
		return nil, fmt.Errorf("error:status code %s", resp.StatusCode)
	}

	// 指定页面编码为utf-8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err!= nil {
		log.Printf("Ftcher error:%v", err)
		return unicode.UTF8
	}

	e, _,_ := charset.DetermineEncoding(bytes, "")
	return e
}
