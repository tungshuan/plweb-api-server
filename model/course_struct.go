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

type Quiz struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Stdout      string `json:"stdout"`
	Description string `json:"description"`
	Ans         string `json:"ans"`
	Part        string `json:"part"`
}
