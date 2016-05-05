package controller

import (
	//"fmt"
	"github.com/kataras/iris"
	"github.com/yuniii/plweb-api-server/model"
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
	//c.Write(fmt.Sprintf("%+v", lesson))
	/*err = c.XML(lesson)
	if !checkErr(err, c) {
		return
	}*/
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
