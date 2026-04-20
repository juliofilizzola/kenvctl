package utils

import (
	"encoding/json"
)

// BuildJSONKeyValue retorna uma string JSON no formato '{"chave": "valor"}'
func BuildJSONKeyValue(key, value string) string {
	m := map[string]string{key: value}
	b, err := json.Marshal(m)
	if err != nil {
		return "{}"
	}
	return string(b)
}
