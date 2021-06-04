module github.com/olongfen/demo

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.7.1
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lib/pq v1.6.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olongfen/contrib v1.0.2
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.1
	github.com/subosito/gotenv v1.2.0
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	gorm.io/driver/clickhouse v0.1.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.10
)
