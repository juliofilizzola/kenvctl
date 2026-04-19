package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("add called")
	return nil
}

func List(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("list called")
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
