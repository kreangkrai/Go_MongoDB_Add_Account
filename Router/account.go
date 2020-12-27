package Router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kriangkrai/Mee/MongoDB/Controller"
	"github.com/kriangkrai/Mee/MongoDB/Models"
)

func GetAccount(c *gin.Context) {
	name := c.Params.ByName("name")
	data := Controller.ReadAccount(name)
	c.JSON(200, data)
}

func InsertAccount(c *gin.Context) {
	var input Models.AccountModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err)
	}
	_, err = Controller.InsertAccount(input)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(200, "Insert Success")
}
