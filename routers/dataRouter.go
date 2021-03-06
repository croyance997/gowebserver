package routers

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	// "program/com.ypc/helloGin/model"
	"webserverRaspberry/processexe"
)

// RaspData 记录从树莓派传过来的数据
type RaspData struct {
	id    string
	time  string
	tem   string
	hum   string
	shake string
}

// DataCache 暂时缓存树莓派的数据
var DataCache RaspData

func init() {
	log.Println(">>>> controller user init <<<<")
	// log.Println(">>>> get database connection start <<<<")
	// db = database.GetDataBase()
}

// RasData 接收从树莓派传过来的数据
func RasData(context *gin.Context) {
	println(">>>> get data from raspberry <<<<")

	(DataCache).id = context.Query("id")
	(DataCache).tem = context.Request.URL.Query().Get("tem")
	(DataCache).hum = context.Request.URL.Query().Get("hum")
	(DataCache).shake = context.Request.URL.Query().Get("shake")
	(DataCache).time = time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("raspi get data from %s, data is %s\n", (DataCache).id, (DataCache).tem)

	context.JSON(200, gin.H{
		"result": "ok",
	})

}

// WebData 给前端返回数据
func WebData(context *gin.Context) {
	println(">>>> Ruturn data to web <<<<")

	// (DataCache).id = context.Query("id")
	// (DataCache).tem = context.Request.URL.Query().Get("tem")
	// (DataCache).hum = context.Request.URL.Query().Get("hum")
	// (DataCache).shake = context.Request.URL.Query().Get("shake")
	// (DataCache).time = time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("web get data from %s, data is %s\n", (DataCache).id, (DataCache).tem)

	context.JSON(200, gin.H{
		"id":    (DataCache).id,
		"tem":   (DataCache).tem,
		"hum":   (DataCache).hum,
		"shake": (DataCache).shake,
		"time":  (DataCache).time,
	})

}

func ProcessData(context *gin.Context) {
	println(">>>> Get process data <<<<")

	// (DataCache).id = context.Query("id")
	// (DataCache).tem = context.Request.URL.Query().Get("tem")
	// (DataCache).hum = context.Request.URL.Query().Get("hum")
	// (DataCache).shake = context.Request.URL.Query().Get("shake")
	// (DataCache).time = time.Now().Format("2006-01-02 15:04:05")
	processexe.Executeprocess( map[string]string { "o1":"wwww"  }  )
	

	context.JSON(200, gin.H{
		"result": "ok",
	})

}


// func checkError(e error) {
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// }
