package command

import (
	"fmt"

	"github.com/juliofilizzola/kenvctl/internal/service"
	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("add called")

	return nil
}

func List(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("list called")
	_, err = service.CreateEnv()
	if err != nil {
		return err
	}
	return nil
}

func Remove(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("remove called")
	return nil
}

func Set(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("set called")
	return nil
}

func AddNewEnv(cmd *cobra.Command, args []string) (err error) {
	_, err = service.AddNewEnv()
	return err
}

func IsValidAWSCLI(cmd *cobra.Command, args []string) (err error) {
	_, err = service.ValidAwsCLI()
	return err
}

func IsValidKubeCLI(cmd *cobra.Command, args []string) (err error) {
	_, err = service.ValidKubeCLI()
	return err
}
