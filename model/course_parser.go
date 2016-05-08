package model

import (
	"encoding/base64"
	"encoding/xml"
	"github.com/qiniu/iconv"
	"strings"
)

var (
	uselessFile = []string{".plwebenv", ".plwebtest", "diff", ".exe", "execdump", ".class", "typescript", "#save#"}
)

func ParseCourse(xmlContent string) (map[string]Quiz, error) {
	lesson := Lesson{}
	err := xml.Unmarshal([]byte(xmlContent), &lesson)
	if err != nil {
		return nil, err
	}

	var path, content string
	newXmlContent := Lesson{}

	cd, _ := iconv.Open("UTF-8", "big5")
	defer cd.Close()

	for i := 0; i < len(lesson.Files); i++ {
		path = lesson.Files[i].Path
		if !isUsefulFileType(path) {
			continue
		}

		content = DecodeStr(lesson.Files[i].Content)
		if strings.Contains(path, ".cond") {
			content = cd.ConvString(content)
		}
		newXmlContent.Files = append(newXmlContent.Files, File{path, content})
	}
	quizzes := groupQuizzes(newXmlContent)
	return quizzes, nil
}

func groupQuizzes(lesson Lesson) map[string]Quiz {
	files := lesson.Files
	quizzes := make(map[string]Quiz)

	for _, f := range files {
		dotIndex := strings.LastIndex(f.Path, ".")
		title := f.Path[:dotIndex]
		filetype := f.Path[dotIndex:]
		aQuiz := quizzes[title]
		aQuiz.Title = title

		if filetype == ".cond" {
			aQuiz.Stdout = f.Content
		} else if filetype == ".html" {
			aQuiz.Description = f.Content
		} else if filetype == ".part" {
			aQuiz.Part = f.Content
		} else {
			aQuiz.Ans = f.Content
		}
		quizzes[title] = aQuiz
	}

	return quizzes
}

func isUsefulFileType(path string) bool {
	if strings.Contains(path, ".html") || strings.Contains(path, ".cond") || strings.Contains(path, ".part") {
		return true
	}

	for i := range uselessFile {
		if strings.Contains(path, uselessFile[i]) {
			return false
		}
	}

	return true
}

func Decode(base64str string) []byte {
	result, _ := base64.StdEncoding.DecodeString(base64str)
	return result
}

func DecodeStr(base64str string) string {
	result := Decode(base64str)
	return string(result)
}
