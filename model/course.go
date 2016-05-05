package model

import (
	"github.com/yuniii/plweb-api-server/db"
)

func GetCourse(courseId, lessonId int) (string, error) {
	stmt, err := db.DB.Prepare("SELECT text_xml FROM COURSE_FILE WHERE course_id= ? AND lesson_id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var xmlContent string
	err = stmt.QueryRow(courseId, lessonId).Scan(&xmlContent)
	if err != nil {
		return "", err
	}
	return xmlContent, nil
}
