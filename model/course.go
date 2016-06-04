package model

import (
	"github.com/Yuniii/plweb-api-server/db"
)

func GetCourse(courseID, lessonID int) (string, error) {
	stmt, err := db.DB.Prepare("SELECT text_xml FROM COURSE_FILE WHERE course_id= ? AND lesson_id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var xmlContent string
	err = stmt.QueryRow(courseID, lessonID).Scan(&xmlContent)
	if err != nil {
		return "", err
	}
	return xmlContent, nil
}
