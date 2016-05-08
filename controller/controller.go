package controller

import (
	"github.com/Yuniii/plweb-api-server/model"
	"github.com/kataras/iris"
)

func Index(c *iris.Context) {
	c.Write("plweb api server")
}

func GetCourse(c *iris.Context) {
	courseId, err := c.ParamInt("courseId")
	if !checkErr(err, c) {
		return
	}

	lessonId, err := c.ParamInt("lessonId")
	if !checkErr(err, c) {
		return
	}

	courseXml, err := model.GetCourse(courseId, lessonId)
	if !checkErr(err, c) {
		return
	}

	lesson, err := model.ParseCourse(courseXml)
	if !checkErr(err, c) {
		return
	}
	c.Write(string(lesson))
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
