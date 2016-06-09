package controller

import (
	"github.com/Yuniii/plweb-api-server/model/course"
	"github.com/Yuniii/plweb-api-server/model/report"
	"github.com/kataras/iris"
	"strconv"
)

func Index(c *iris.Context) {
	c.Write("plweb api server v0609")
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

	courseXML, err := course.GetCourse(courseID, lessonID)
	if !checkErr(err, c) {
		return
	}

	lesson, err := course.ParseCourse(courseXML)
	if !checkErr(err, c) {
		return
	}
	c.JSON(iris.StatusOK, lesson)
}

func SubmitCode(c *iris.Context) {
	classID, err := c.ParamInt("classID")
	if !checkErr(err, c) {
		return
	}

	courseID, err := c.ParamInt("courseID")
	if !checkErr(err, c) {
		return
	}

	lessonID, err := c.ParamInt("lessonID")
	if !checkErr(err, c) {
		return
	}

	qn, err := c.ParamInt("qn")
	if !checkErr(err, c) {
		return
	}

	code := c.PostFormValue("code")
	userID, err := strconv.Atoi(c.PostFormValue("uid"))
	t := c.PostFormValue("type")
	if !checkErr(err, c) {
		return
	}

	submission := report.UserSubmission{
		classID,
		courseID,
		lessonID,
		qn,
		userID,
		code,
		t,
	}
	err = report.SubmitCode(submission)

	if !checkErr(err, c) {
		c.Write(err.Error())
		return
	}

	c.Write("ok")
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
