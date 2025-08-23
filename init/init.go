package init

import (
	"studyonline/dao/mysql"
	"studyonline/dao/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
