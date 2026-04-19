package kube

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetAllSecrets executa 'kubectl get secrets -A -o json' e retorna a saída como string
func GetAllSecrets() (string, error) {
	cmd := exec.Command("kubectl", "get", "secrets", "-A", "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao buscar secrets: %v\n%s", err, string(out))
	}
	return string(out), nil
}

// GetSecret executa 'kubectl get secret <name> -n <namespace> -o json' e retorna a saída como string
func GetSecret(namespace, name string) (string, error) {
	cmd := exec.Command("kubectl", "get", "secret", name, "-n", namespace, "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao buscar secret: %v\n%s", err, string(out))
	}
	return string(out), nil
}

// CreateSecretFromLiterals executa 'kubectl create secret generic <name> --from-literal=...'
func CreateSecretFromLiterals(namespace, name string, data map[string]string) error {
	args := []string{"create", "secret", "generic", name}
	for k, v := range data {
		args = append(args, "--from-literal="+k+"="+v)
	}
	if namespace != "" {
		args = append(args, "-n", namespace)
	}
	cmd := exec.Command("kubectl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao criar secret: %v\n%s", err, string(out))
	}
	return nil
}

// UpdateSecretFromLiterals faz update ou cria secret usando dry-run + apply
func UpdateSecretFromLiterals(namespace, name string, data map[string]string) error {
	args := []string{"create", "secret", "generic", name, "--dry-run=client", "-o", "yaml"}
	for k, v := range data {
		args = append(args, "--from-literal="+k+"="+v)
	}
	if namespace != "" {
		args = append(args, "-n", namespace)
	}
	cmd := exec.Command("kubectl", args...)
	yamlOut, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao gerar yaml da secret: %v\n%s", err, string(yamlOut))
	}
	applyCmd := exec.Command("kubectl", "apply", "-f", "-")
	applyCmd.Stdin = strings.NewReader(string(yamlOut))
	out, err := applyCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao aplicar secret: %v\n%s", err, string(out))
	}
	return nil
}

// GetAllSecretsByLabel executa 'kubectl get secrets -A -l <label> -o json' e retorna a saída como string
func GetAllSecretsByLabel(label string) (string, error) {
	cmd := exec.Command("kubectl", "get", "secrets", "-A", "-l", label, "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao buscar secrets por label: %v\n%s", err, string(out))
	}
	return string(out), nil
}

// GetAllSecretsTable executa 'kubectl get secrets --all-namespaces' e retorna a saída tabular como string
func GetAllSecretsTable() (string, error) {
	cmd := exec.Command("kubectl", "get", "secrets", "--all-namespaces")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao buscar secrets (tabela): %v\n%s", err, string(out))
	}
	return string(out), nil
}
