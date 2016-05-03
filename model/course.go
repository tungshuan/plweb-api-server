package model

import (
	"github.com/yuniii/plweb-api-server/db"
)

type Course struct {
	COURSE_ID,
	LESSON_ID int
	TEXT_XML string
}

func GetCourse(courseId, lessonId int) string {
	stmt, err := db.DB.Prepare("SELECT text_xml FROM COURSE_FILE WHERE course_id= ? AND lesson_id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var xmlContent string
	err = stmt.QueryRow(courseId, lessonId).Scan(&xmlContent)
	if err != nil {
		panic(err)
	}
	return xmlContent
}
