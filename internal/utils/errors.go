package utils

import "fmt"

// NewError cria um erro padronizado a partir de uma mensagem
func NewError(msg string) error {
	return fmt.Errorf(msg)
}
