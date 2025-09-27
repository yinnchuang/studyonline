package constant

var (
	// 资源种类
	TeachingVideo    = 1001
	ElectricResource = 1002
	Courseware       = 1003
	Exercise         = 1004
	ProjectCases     = 1005
	// 数据集种类
	StructuredData     = 2001
	UnstructuredData   = 2002
	MedicalImagingData = 2003
	TimeSeriesData     = 2004
	OmicsData          = 2005
	MultimodalData     = 2006
)

func IfResourceCategory(code int) bool {
	return code == TeachingVideo || code == ElectricResource || code == Courseware || code == Exercise || code == ProjectCases
}

func IfDatasetCategory(code int) bool {
	return code == StructuredData || code == UnstructuredData || code == MedicalImagingData || code == TimeSeriesData || code == OmicsData || code == MultimodalData
}
