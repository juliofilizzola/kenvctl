package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrintInfo imprime uma mensagem informativa
func PrintInfo(msg string) {
	fmt.Println(msg)
}

// PrintError imprime uma mensagem de erro
func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
}

// PrintJSON imprime um struct em formato JSON
func PrintJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		PrintError(err)
		return
	}
	fmt.Println(string(jsonData))
}

// PrintChoices exibe múltiplas opções e retorna o índice escolhido pelo usuário
func PrintChoices(choices []string) int {
	if len(choices) == 0 {
		PrintError(fmt.Errorf("nenhuma opção disponível"))
		return -1
	}
	for i, choice := range choices {
		fmt.Printf("%d) %s\n", i+1, choice)
	}
	fmt.Print("Escolha uma opção: ")
	var selected int
	_, err := fmt.Scan(&selected)
	if err != nil || selected < 1 || selected > len(choices) {
		PrintError(fmt.Errorf("opção inválida"))
		return -1
	}
	return selected - 1
}

// AskValue solicita um valor ao usuário e retorna a entrada como string
func AskValue(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		PrintError(fmt.Errorf("erro ao ler valor: %v", err))
		return ""
	}
	return input
}
