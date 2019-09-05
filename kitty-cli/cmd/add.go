package cmd

import (
	"fmt"
	"log"

	"github.com/l-lin/kitty-cli/add"
	"github.com/l-lin/kitty-cli/cat"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "ADD a cat",
		Run:   runAdd,
	}
	catName, catType string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&catName, "name", "n", "", "name of the pet to create")
	addCmd.Flags().StringVarP(&catType, "type", "t", "", "type of the pet to create")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("type")
}

func runAdd(cmd *cobra.Command, args []string) {
	s := add.Service{URL: providerURL + "/cats"}
	cat, err := s.Add(cat.Cat{
		Name: catName,
		Type: catType,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cat)
}
