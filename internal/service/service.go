package service

import (
	"github.com/juliofilizzola/kenvctl/internal/aws"
	"github.com/juliofilizzola/kenvctl/internal/kube"
	"github.com/juliofilizzola/kenvctl/internal/output"
	"github.com/juliofilizzola/kenvctl/internal/utils"
)

func CreateEnv() (bool, error) {
	output.PrintInfo("Buscando secretes")
	table, err := kube.GetAllSecretsTable()
	if err != nil {
		return false, err
	}
	lines := utils.ParseTableLines(table)
	index := output.PrintChoices(lines)

	if index == -1 {
		return false, nil
	}

	items := utils.ParseTableLine(lines[index])

	output.PrintInfo("Secret selecionado: " + items.Name)

	output.PrintInfo("informe a env para o secret selecionado")
	env := output.AskValue("env")

	if env == "" {
		output.PrintInfo("env não foi informada")
		env = output.AskValue("env")

		if env == "" {
			output.PrintError(utils.NewError("Env não informada"))
			return false, nil
		}
	}

	output.PrintInfo("env informada")
	output.PrintInfo("informe o valor da env para o secret selecionado")
	value := output.AskValue("valor")

	if value == "" {
		output.PrintInfo("valor não foi informada")
		value = output.AskValue("valor")
		if value != "" {
			output.PrintError(utils.NewError("Valor não informada"))
			return false, nil
		}
	}

	secret := utils.BuildJSONKeyValue(env, value)

	err = aws.CreateAwsSecret(items.Name, secret)
	if err != nil {
		output.PrintError(err)
		return false, err
	}

	err = kube.CreateSecretFromLiterals(items.Namespace, items.Name, map[string]string{env: value})
	if err != nil {
		output.PrintError(err)
		return false, err
	}

	return true, nil
}

func AddNewEnv() (bool, error) {
	output.PrintInfo("Adicionando nova env")

	tables, err := kube.GetAllSecretsTable()
	if err != nil {
		return false, err
	}
	lines := utils.ParseTableLines(tables)
	index := output.PrintChoices(lines)

	items := utils.ParseTableLine(lines[index])

	output.PrintInfo("Secret selecionado: " + items.Name)

	output.PrintInfo("informe a nova env para o secret selecionado")

	env := output.AskValue("env")
	if env == "" {
		output.PrintInfo("env não foi informada")
		env = output.AskValue("env")
		if env == "" {
			output.PrintError(utils.NewError("Env não informada"))
			return false, nil
		}
	}

	output.PrintInfo("env informada")

	output.PrintInfo("informe o valor da nova env para o secret selecionado")
	value := output.AskValue("valor")

	if value == "" {
		output.PrintInfo("valor não foi informada")
		value = output.AskValue("valor")

		if value == "" {
			output.PrintError(utils.NewError("Valor não informada"))
			return false, nil
		}
	}

	secret := utils.BuildJSONKeyValue(env, value)

	err = aws.UpdateAwsSecret(items.Name, secret)

	if err != nil {
		output.PrintError(err)
		return false, err
	}

	err = kube.UpdateSecretFromLiterals(items.Namespace, items.Name, map[string]string{env: value})
	if err != nil {
		output.PrintError(err)
		return false, err
	}

	return true, nil
}

// CreateEnvFromFile lê um arquivo de envs e cria a secret no AWS e no Kubernetes
func CreateEnvFromFile(filePath string) (bool, error) {
	output.PrintInfo("Buscando secrets")
	table, err := kube.GetAllSecretsTable()
	if err != nil {
		return false, err
	}
	lines := utils.ParseTableLines(table)
	index := output.PrintChoices(lines)
	if index == -1 {
		return false, nil
	}
	items := utils.ParseTableLine(lines[index])
	output.PrintInfo("Secret selecionado: " + items.Name)

	output.PrintInfo("Lendo envs do arquivo: " + filePath)
	data, err := utils.ReadEnvFile(filePath)
	if err != nil {
		output.PrintError(err)
		return false, err
	}
	secret := utils.BuildJSONFromMap(data)

	err = aws.CreateAwsSecret(items.Name, secret)
	if err != nil {
		output.PrintError(err)
		return false, err
	}
	err = kube.CreateSecretFromLiterals(items.Namespace, items.Name, data)
	if err != nil {
		output.PrintError(err)
		return false, err
	}
	return true, nil
}

// UpdateEnvFromFile lê um arquivo de envs e atualiza a secret no AWS e no Kubernetes
func UpdateEnvFromFile(filePath string) (bool, error) {
	output.PrintInfo("Buscando secrets")
	table, err := kube.GetAllSecretsTable()
	if err != nil {
		return false, err
	}
	lines := utils.ParseTableLines(table)
	index := output.PrintChoices(lines)
	if index == -1 {
		return false, nil
	}
	items := utils.ParseTableLine(lines[index])
	output.PrintInfo("Secret selecionado: " + items.Name)

	output.PrintInfo("Lendo envs do arquivo: " + filePath)
	data, err := utils.ReadEnvFile(filePath)
	if err != nil {
		output.PrintError(err)
		return false, err
	}
	secret := utils.BuildJSONFromMap(data)

	err = aws.UpdateAwsSecret(items.Name, secret)
	if err != nil {
		output.PrintError(err)
		return false, err
	}
	err = kube.UpdateSecretFromLiterals(items.Namespace, items.Name, data)
	if err != nil {
		output.PrintError(err)
		return false, err
	}
	return true, nil
}
