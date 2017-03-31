package models
import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
const DB_USR="root"
const DB_PASSWD="123456"
const DB_HOST="127.0.0.1"
const DB_PORT="3306"
const DB_NAME="gomusic"
func Init()  {

	dsn := DB_USR + ":" + DB_PASSWD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8"

	//注册数据库
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		beego.Error("Failed to connect database,code:",err)
		return
	}

	//注册模型
	orm.RegisterModel(new(Song))

	//是否开启调试模式
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = false
	}

	//如果不存在表会建立表
	orm.RunSyncdb("default", false, true)
}
