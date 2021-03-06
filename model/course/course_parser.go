package course

import (
	"encoding/base64"
	"encoding/xml"
	"github.com/qiniu/iconv"
	"strings"
)

var (
	uselessFile = []string{".plwebenv", ".plwebtest", "diff", ".exe", "execdump", ".class", "typescript", "#save#"}
)

func ParseCourse(xmlContent string) ([]Quiz, error) {
	lesson := Lesson{}
	err := xml.Unmarshal([]byte(xmlContent), &lesson)
	if err != nil {
		return nil, err
	}

	var path, content string
	newXMLContent := Lesson{}

	cd, _ := iconv.Open("UTF-8", "big5")
	defer cd.Close()

	for i := 0; i < len(lesson.Files); i++ {
		path = lesson.Files[i].Path
		if !isUsefulFileType(path) {
			continue
		}

		content = decodeStr(lesson.Files[i].Content)
		if strings.Contains(path, ".cond") {
			content = cd.ConvString(content)
		}
		newXMLContent.Files = append(newXMLContent.Files, File{path, content})
	}
	quizzes := groupQuizzes(newXMLContent)
	parseQid(lesson, quizzes)
	return mapToArray(quizzes), nil
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

func parseQid(lesson Lesson, quizzes map[string]Quiz) {
	var task Task
	seq := 1

	for i := 0; i < len(lesson.Tasks); i++ {
		task = lesson.Tasks[i]
		for j := 0; j < len(task.Properties); j++ {
			if task.Properties[j].Key != "ExName" {
				continue
			}
			for k := range quizzes {
				if k == task.Properties[j].Value {
					aQuiz := quizzes[k]
					aQuiz.Seq = seq
					seq++
					aQuiz.Qid = lesson.Tasks[i].ID
					quizzes[k] = aQuiz
				}
			}
		}
	}
}

func mapToArray(quizMap map[string]Quiz) []Quiz {
	result := make([]Quiz, len(quizMap))
	i := 0
	var seq int
	for _, val := range quizMap {
		seq = val.Seq - 1
		if seq < 0 {
			continue
		}
		result[seq] = val
		result[seq].ID = (i + 1)
		i++
	}
	trimmed := trimArray(result)
	return trimmed
}

func trimArray(quizzes []Quiz) []Quiz {
	var trimmed []Quiz
	for i := range quizzes {
		if quizzes[i].Seq > 0 {
			trimmed = append(trimmed, quizzes[i])
		}
	}
	return trimmed
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

func decode(base64str string) []byte {
	result, _ := base64.StdEncoding.DecodeString(base64str)
	return result
}

func decodeStr(base64str string) string {
	result := decode(base64str)
	return string(result)
}
