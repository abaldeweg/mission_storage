package mission

import (
	"encoding/json"
	"log"

	"github.com/abaldeweg/mission_storage/export/html"
	"github.com/abaldeweg/mission_storage/mission/create"
	"github.com/abaldeweg/mission_storage/storage"
	"github.com/gin-gonic/gin"
)

var filename = "mission.json"

type Msg struct {
	Msg string `json:"msg"`
}

type Response struct {
	Type string `json:"type"`
	Body string `json:"body" binding:"required"`
}

func Show(c *gin.Context) {
	if !storage.Exists(filename) {
		c.AbortWithStatus(404)
		return
	}

	var d create.Logfile
	if err := json.Unmarshal(storage.Read(filename), &d); err != nil {
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
	var file Response
	if err := c.ShouldBind(&file); err != nil {
		c.AbortWithStatus(404)
		return
	}

	storage.Write(filename, file.Body)

	d := Msg{Msg: "SUCCESS"}

	c.JSON(200, d)
}

func HtmlExport(c *gin.Context) {
	b, err := html.Export(storage.Read(filename))
	if err != nil {
		log.Fatal(err)
	}

	d := Response{Type: "html", Body: b}

	c.JSON(200, d)
}
