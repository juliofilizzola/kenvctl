package aws

import (
	"fmt"
	"os/exec"
	"strings"
)

func CreateAwsSecret(name string, secretString string) error {
	cmd := exec.Command("aws", "secretsmanager", "create-secret", "--name", name, "--secret-string", secretString)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao criar secret na AWS: %v\n%s", err, string(out))
	}
	return nil
}

// UpdateAwsSecret atualiza o valor de um secret existente na AWS
func UpdateAwsSecret(secretId string, secretString string) error {
	cmd := exec.Command("aws", "secretsmanager", "update-secret", "--secret-id", secretId, "--secret-string", secretString)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao atualizar secret na AWS: %v\n%s", err, string(out))
	}
	return nil
}

// IsAwsCliInstalled verifica se a AWS CLI está instalada no sistema
func IsAwsCliInstalled() (string, error) {
	cmd := exec.Command("aws", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao executar aws --version: %w", err)
	}

	versionInfo := strings.TrimSpace(string(out))
	versionInfo = strings.TrimPrefix(versionInfo, "aws-cli/")
	version := strings.Fields(versionInfo)[0]

	return version, nil
}
