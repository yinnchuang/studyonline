package init

import (
	"context"
	"log"
	"studyonline/service"
	"studyonline/util"

	"gopkg.in/ini.v1"

	"studyonline/dao/mysql"
	"studyonline/dao/redis"
	mylog "studyonline/log"
)

func adminInit() {
	cfg, err := ini.Load("./init/project.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	username := cfg.Section("admin").Key("username").String()
	password := cfg.Section("admin").Key("password").String()
	bcryptPassword, _ := util.GetPwd(password)
	err = service.ImportAdmin(context.Background(), username, string(bcryptPassword))
	if err != nil {
		log.Fatal("Fail to import admin: ", err)
	}
}

func Init() {
	mysql.Init()
	redis.Init()
	mylog.Init()
	adminInit()
}
