package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed gormSerializer.tpl
var httpTemplate string

type EnumGen struct {
	EnumName string
}

func NewEnumGen() *EnumGen {
	return &EnumGen{}
}

func (s *EnumGen) execute(name string) string {
	s.EnumName = name
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
