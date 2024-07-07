package models

type TextMessage string

type TextResponse string

type RequestType struct {
	Resume string `json:"resume"`
	Tags []string `json:"tags"`
}