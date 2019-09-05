package cmd

import (
	"fmt"
	"log"

	"github.com/l-lin/doggy-cli/add"
	"github.com/l-lin/doggy-cli/dog"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "ADD a dog",
		Run:   runAdd,
	}
	dogName, dogType string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&dogName, "name", "n", "", "name of the dog to create")
	addCmd.Flags().StringVarP(&dogType, "type", "t", "", "type of the dog to create")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("type")
}

func runAdd(cmd *cobra.Command, args []string) {
	s := add.Service{URL: providerURL + "/dogs"}
	dog, err := s.Add(dog.Dog{
		Name: dogName,
		Type: dogType,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dog)
}
