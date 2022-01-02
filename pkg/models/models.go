package models

type Cli struct {
	Name  string
	Kind  string
	State string
}
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline string `json:"inline"`
}
type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	Fields      []Field `json:"fields"`
}
