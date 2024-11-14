/*
Copyright © 2024 NAME HERE <saiduzzamanzishan44@gmail.com>
*/
package main

import (
	"github.com/saiduzzaman44/cli-todo/cmd"
	"github.com/saiduzzaman44/cli-todo/db"
)

func main() {

	db.InitDB()
	defer db.CloseDB()

	cmd.Execute()
}
