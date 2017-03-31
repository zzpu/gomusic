package models
import(
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)
type Song struct {
	Id        int		        //id
	Name      string 		//歌曲名
	IdStr	  string		//网易资源标识
	Artists   string 		//作曲
	Album     string 		//专辑
	Extension string 		//扩展名
	Src       string		//资源地址
}
func NewSong() *Song {
	return &Song{}
}
func Insert(s *Song)  (err error){
	_,err=orm.NewOrm().Insert(s)
	return err
}
func Delete(id int) (err error){
	_, err = orm.NewOrm().QueryTable("song").Filter("id", id).Delete()
	return err

}
func Update(s *Song, fields ...string) (err error){
	_, err = orm.NewOrm().Update(s, fields...)
	return err
}
func Get(id int) *Song {
	s:=NewSong()
	err := orm.NewOrm().QueryTable("song").Filter("id", id).One(s)
	if err == nil{
		return  s
	}
	return nil
}
