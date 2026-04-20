package utils

import (
	"strings"
)

// TableItem representa uma linha da tabela com campos tipados
type TableItem struct {
	Namespace string
	Name      string
	Type      string
	Data      string
	Age       string
}

// ParseTableLine divide uma linha em um TableItem
func ParseTableLine(line string) TableItem {
	fields := strings.Fields(line)
	item := TableItem{}
	if len(fields) > 0 {
		item.Namespace = fields[0]
	}
	if len(fields) > 1 {
		item.Name = fields[1]
	}
	if len(fields) > 2 {
		item.Type = fields[2]
	}
	if len(fields) > 3 {
		item.Data = fields[3]
	}
	if len(fields) > 4 {
		item.Age = fields[4]
	}
	return item
}

// ParseTableLines separa uma string em linhas não vazias
func ParseTableLines(table string) []string {
	lines := strings.Split(table, "\n")
	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
