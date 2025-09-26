package constant

var (
	MaxResourceSize   int64 = 2 << 30   // 资源2G限制
	MaxCoverSize      int64 = 10 << 20  // 封面10M限制
	MaxDatasetSize    int64 = 2 << 30   // 数据集1G限制
	MaxHomeworkSize   int64 = 200 << 20 // 作业限制200M
	MaxSubmissionSize int64 = 200 << 20 // 提交作业限制200M
)
