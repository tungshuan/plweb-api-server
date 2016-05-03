package controller

import (
	"github.com/kataras/iris"
	"github.com/yuniii/plweb-api-server/model"
	"strconv"
)

func Index(c *iris.Context) {
	c.Write("plweb api server")
}

func GetCourse(c *iris.Context) {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		BadRequest(c, err)
		return
	}
	lessonId, err := strconv.Atoi(c.Param("lessonId"))
	if err != nil {
		BadRequest(c, err)
		return
	}
	courseXml, err := model.GetCourse(courseId, lessonId)
	if err != nil {
		BadRequest(c, err)
	}
	c.Write(courseXml)
}

func BadRequest(c *iris.Context, err error) {
	c.Write(err.Error())
}
