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
	env, err := service.CreateEnv()
	if err != nil {
		return err
	}
	fmt.Println(env)
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
