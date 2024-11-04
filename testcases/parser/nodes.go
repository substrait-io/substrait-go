package parser

type TestFileHeader struct {
	Version     string
	IncludedURI string
}

type TestCase struct {
	BaseURI   string
	GroupDesc string
	FuncName  string
	Args      []*CaseLiteral
	Result    *CaseLiteral
}

type TestFile struct {
	Header    TestFileHeader
	TestCases []*TestCase
}
