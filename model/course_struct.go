package model

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
	Id         int        `xml:"id"`
	Title      string     `xml:"title"`
	Properties []Property `xml:"property"`
	Files      []File     `xml:"file"`
}
