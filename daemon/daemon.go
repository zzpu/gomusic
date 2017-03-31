package daemon
import(
	"net/http"
	"fmt"
	"io/ioutil"
)
const MUSIC_URL="http://music.163.com/api/song/detail/?id=99999999&ids=%5B99999999%5D&csrf_token="

func Init(){

	response,err:= http.Get(MUSIC_URL)
	if err != nil{
		fmt.Println("Err:%s",err)
	}
	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
