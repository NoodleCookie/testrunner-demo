package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	command := cobra.Command{
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("yes")
		},
	}
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
