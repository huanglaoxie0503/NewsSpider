package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct{}

type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (NilParser) Parser(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 实现接口定义打方法
func (f *FuncParser) Parser(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// 构造函数
func NewFuncParser(p ParserFunc, n string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   n,
	}
}
