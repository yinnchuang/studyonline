package constant

var (
	// 资源种类
	Video         = 1001
	Courseware    = 1002
	Exercises     = 1003
	ProjectCase   = 1004
	Package       = 1005
	OtherResource = 1006
	// 数据集种类
	LLMApplication  = 2001
	ComputerVision  = 2002
	NaturalLanguage = 2003
	DataProcess     = 2004
)

func IfResourceCategory(code int) bool {
	return code == Video || code == Courseware || code == Exercises || code == ProjectCase || code == Package || code == OtherResource
}

func IfDatasetCategory(code int) bool {
	return code == LLMApplication || code == ComputerVision || code == NaturalLanguage || code == DataProcess
}
