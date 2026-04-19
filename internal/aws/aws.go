package aws

import (
	"fmt"
	"os/exec"
)

func CreateAwsSecret(name string, secretString string) error {
	cmd := exec.Command("aws", "secretsmanager", "create-secret", "--name", name, "--secret-string", secretString)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao criar secret na AWS: %v\n%s", err, string(out))
	}
	return nil
}
