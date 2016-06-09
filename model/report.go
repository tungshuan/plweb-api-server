package model

import (
	"bytes"
	"github.com/Yuniii/plweb-api-server/db"
)

func SubmitCode(subm UserSubmission) error {
	existsSql, updateSql, insertSql := buildSubmitSQL()
	exists, err := reportExists(existsSql, subm)
	if err != nil {
		return err
	}

	if exists {
		err = updateReport(updateSql, subm)
	} else {
		err = insertReport(insertSql, subm)
	}

	if err != nil {
		return err
	}

	return nil
}

func reportExists(sql string, subm UserSubmission) (bool, error) {
	var result int
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(subm.ClassID, subm.CourseID, subm.LessonID, subm.Qn, subm.UserID).Scan(&result)
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func insertReport(sql string, subm UserSubmission) error {
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subm.ClassID, subm.CourseID, subm.LessonID, subm.Qn, subm.UserID, subm.Code, subm.Type)
	if err != nil {
		return err
	}

	return nil
}

func updateReport(sql string, subm UserSubmission) error {
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subm.Code, subm.Type, subm.ClassID, subm.CourseID, subm.LessonID, subm.Qn, subm.UserID)
	if err != nil {
		return err
	}

	return nil
}

func buildSubmitSQL() (string, string, string) {
	var b1, b2, b3 bytes.Buffer
	where := " WHERE CLASS_ID=? AND COURSE_ID=? AND LESSON_ID=? AND QUESTION_ID=? AND USER_ID=?"

	b1.WriteString("SELECT EXISTS(SELECT * FROM ST_REPORTS")
	b1.WriteString(where + ")")

	b2.WriteString("UPDATE ST_REPORTS SET code=?, type=?")
	b2.WriteString(where)

	b3.WriteString("INSERT INTO ST_REPORTS(class_id, course_id, lesson_id, question_id, user_id, code, type) ")
	b3.WriteString("VALUES(?, ?, ?, ?, ?, ?, ?)")
	return b1.String(), b2.String(), b3.String()
}
