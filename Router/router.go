package Router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/Mee/MongoDB/Controller"
)

func SetupRouter() (*gin.Engine, string) {
	Controller.Connect()

	r := gin.Default()
	r.Use(Middleware())
	r.GET("/get/:device", Get)
	r.POST("/insert", Insert)
	r.PUT("/update", Update)
	r.DELETE("/delete/:device", Delete)

	//Account
	r.GET("/getname/:name", GetAccount)
	r.POST("insertaccount", InsertAccount)
	// port := "8080"
	// if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
	// 	port = os.Getenv("ASPNETCORE_PORT")
	// }
	//Port
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		//fmt.Println("No Port In Heroku" + port)
	}
	// return ":" + port
	return r, port
}
