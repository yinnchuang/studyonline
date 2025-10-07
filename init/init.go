package init

import (
	"studyonline/dao/mysql"
	"studyonline/dao/redis"
	"studyonline/log"
)

func Init() {
	mysql.Init()
	redis.Init()
	log.Init()
}
