package model

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"github.com/qiniu/iconv"
	"strings"
)

var (
	uselessFile = []string{".plwebenv", ".plwebtest", "diff", ".exe", "execdump", ".class", "typescript", "#save#"}
)

func ParseCourse(xmlContent string) ([]byte, error) {
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
	result, err := json.Marshal(newXmlContent)

	return result, nil
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
