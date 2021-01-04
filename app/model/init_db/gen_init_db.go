package model

import (
	"fmt"

	"github.com/olongfen/contrib/log"
	_ "github.com/olongfen/demo/app/model/admin"
	"github.com/olongfen/demo/app/model/common"
	_ "github.com/olongfen/demo/app/model/demo"
	_ "github.com/olongfen/demo/app/model/region"
	_ "github.com/olongfen/demo/app/model/user"
	"github.com/olongfen/demo/app/setting"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	var (
		err            error
		dataSourceName string
		dialector      gorm.Dialector
	)
	model_common.ModelLog = log.NewLogFile(log.ParamLog{Path: setting.Global.FilePath.LogDir + "/" + "models", Stdout: setting.DevEnv, P: setting.Global.FilePath.LogPatent})
	switch setting.Global.DB.Driver {
	case "postgres":
		dataSourceName = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", setting.Global.DB.Username,
			setting.Global.DB.Password, setting.Global.DB.Host, setting.Global.DB.Port, setting.Global.DB.DatabaseName)
		//dataSourceName = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", setting.Global.DB.Driver, setting.Global.DB.Username,
		//	setting.Global.DB.Password, setting.Global.DB.Host, setting.Global.DB.Port, setting.Global.DB.DatabaseName)
		dialector = postgres.Open(dataSourceName)
	case "mysql":
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%scharset=utf8mb4&parseTime=True&loc=Local", setting.Global.DB.Username,
			setting.Global.DB.Password, setting.Global.DB.Host, setting.Global.DB.Port, setting.Global.DB.DatabaseName)
		dialector = mysql.Open(dataSourceName)
	case "sqlite":
		dialector = sqlite.Open(setting.Global.DB.Source)
	case "sqlserver":
		dataSourceName = fmt.Sprintf("%s://%s:%s@%s:%s?database=%s", setting.Global.DB.Driver, setting.Global.DB.Username,
			setting.Global.DB.Password, setting.Global.DB.Host, setting.Global.DB.Port, setting.Global.DB.DatabaseName)
		dialector = sqlserver.Open(dataSourceName)
	case "clickhouse":
		dataSourceName = fmt.Sprintf("tcp://%s:%sdatabase=%s&username=%s&password=%s&read_timeout=10&write_timeout=30", setting.Global.DB.Host, setting.Global.DB.Port,
			setting.Global.DB.DatabaseName, setting.Global.DB.Username, setting.Global.DB.Password)
		dialector = clickhouse.Open(dataSourceName)
	default:
		log.Fatalln("dose not support this sql driver >>> ", setting.Global.DB.Driver)
	}

	if model_common.DB, err = gorm.Open(dialector, &gorm.Config{Logger: logger.New(model_common.ModelLog, logger.Config{
		Colorful: true})}); err != nil {
		logrus.Fatal(err)
	}
	if setting.DevEnv {
		model_common.DB = model_common.DB.Debug()
		err = model_common.DB.AutoMigrate(model_common.Tables...)
		if err != nil {
			panic(err)
		}
	}

	log.Infoln("database init success !")
}
