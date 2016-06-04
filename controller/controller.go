package controller

import (
	"github.com/Yuniii/plweb-api-server/model"
	"github.com/kataras/iris"
)

func Index(c *iris.Context) {
	c.Write("plweb api server")
}

func GetCourse(c *iris.Context) {
	courseID, err := c.ParamInt("courseID")
	if !checkErr(err, c) {
		return
	}

	lessonID, err := c.ParamInt("lessonID")
	if !checkErr(err, c) {
		return
	}

	courseXML, err := model.GetCourse(courseID, lessonID)
	if !checkErr(err, c) {
		return
	}

	lesson, err := model.ParseCourse(courseXML)
	if !checkErr(err, c) {
		return
	}
	c.JSON(iris.StatusOK, lesson)
}

func checkErr(err error, c *iris.Context) bool {
	if err != nil {
		logErr(err)
		c.Panic()
		return false
	}
	return true
}

func logErr(err error) {
	logger := iris.Logger()
	logger.Printf("!!ERROR!! %s\n", err)
}
