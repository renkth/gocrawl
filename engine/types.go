package engine

// 解析后返回结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url string    // 解析的url
	ParserFunc func([]byte) ParseResult //处理url所要用的函数
}

func NilParser([] byte) ParseResult{
	return ParseResult{}
}