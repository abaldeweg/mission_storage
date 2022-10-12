package mission

import (
	"encoding/json"
	"log"

	"github.com/abaldeweg/storage/export/html"
	"github.com/abaldeweg/storage/mission/create"
	"github.com/abaldeweg/storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "mission.json"

type Msg struct {
    Msg string `json:"msg"`
}

func init() {
    log.SetPrefix("mission: ")
    log.SetFlags(0)
}

func Show(c *gin.Context) {
    if !storage.Exists(filename) {
        c.AbortWithStatus(404)
        return
    }

    var d create.Logfile
    if err:= json.Unmarshal(storage.Read(filename), &d); err != nil {
        c.AbortWithStatus(404)
        return
    }

    c.JSON(200, d)
}

func Create(c *gin.Context) {
    create.Create()

    d := Msg{Msg: "SUCCESS"}

    c.JSON(200, d)
}

func Update(c *gin.Context) {
    var file html.Response
    if err := c.ShouldBind(&file); err != nil {
        c.AbortWithStatus(404)
        return
    }

    storage.Write(filename, file.Body)

    d := Msg{Msg: "SUCCESS"}

    c.JSON(200, d)
}

func HtmlExport(c *gin.Context) {
    d := html.Export()

    c.JSON(200, d)
}
