package main

import (
	"github.com/dipesh23-apt/golang_api/cmd"
	"github.com/dipesh23-apt/golang_api/repo"
)

func init() {
	_, err := repo.InitialMigration()
	if err != nil {
		panic("Could not connect Database")
	}

}

func main() {
	cmd.Execute()
}
