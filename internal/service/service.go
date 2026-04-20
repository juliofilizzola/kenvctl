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
	// todo: criar env no kubernates e na aws

	return false, nil
}
