package server

import (
	"db2/cms/db"
	"os"
)

func handleScripts() {
	argLength := len(os.Args[1:])
	if argLength == 1 {
		if os.Args[1] == "create" || os.Args[1] == "c" {
			db.RunScript("db/scripts/create_script.sql")
			os.Exit(0)
		}
		if os.Args[1] == "generate" || os.Args[1] == "g" {
			db.RunScript("db/scripts/generate_script.sql")
			os.Exit(0)
		}
	}

}
