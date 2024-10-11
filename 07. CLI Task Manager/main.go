package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/shounaklohokare/gophercises/task/cmd"
	"github.com/shounaklohokare/gophercises/task/db"
)

func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")

	must(db.Init(dbPath))

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
