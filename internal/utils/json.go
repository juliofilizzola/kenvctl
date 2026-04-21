package utils

import (
	"encoding/json"
	"os"
	"strings"
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

// ReadEnvFile lê um arquivo de envs no formato KEY=VALUE por linha e retorna um map[string]string
func ReadEnvFile(filePath string) (map[string]string, error) {
	result := make(map[string]string)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // ignora linhas inválidas
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		result[key] = value
	}
	return result, nil
}

// BuildJSONFromMap recebe um map[string]string e retorna um JSON string
func BuildJSONFromMap(data map[string]string) string {
	if data == nil {
		return "{}"
	}
	b, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(b)
}
