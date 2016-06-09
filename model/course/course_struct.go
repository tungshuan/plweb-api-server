package course

import (
	"encoding/xml"
)

type Property struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type File struct {
	Path    string `xml:"path"`
	Content string `xml:"content"`
}

type Lesson struct {
	XMLName    xml.Name   `xml:"project"`
	ID         int        `xml:"id"`
	Title      string     `xml:"title"`
	Properties []Property `xml:"property"`
	Files      []File     `xml:"file"`
	Tasks      []Task     `xml:"task"`
}

type Quiz struct {
	ID          int    `json:"id"`
	Qid         int    `json:"qid"`
	Seq         int    `json:"seq"`
	Title       string `json:"title"`
	Stdout      string `json:"stdout"`
	Description string `json:"description"`
	Ans         string `json:"ans"`
	Part        string `json:"part"`
}

type Task struct {
	ID         int        `xml:"id"`
	Properties []Property `xml:"property"`
}
