package cmd

import (
	"fmt"
	"log"

	"github.com/l-lin/doggy-cli/list"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "GET a dog",
		Run:   runGet,
	}
	id int64
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().Int64VarP(&id, "id", "i", 0, "id of the dog to fetch")
	addCmd.MarkFlagRequired("id")
}

func runGet(cmd *cobra.Command, args []string) {
	s := list.Service{URL: providerURL + "/dogs"}
	dog, err := s.Get(id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dog)
}
