package db

import (
	"os"
	"path/filepath"
)

func VerifyExistingDBs() []string {
	dbPath := "./db"

	files, err := os.ReadDir(dbPath)
	if err != nil {
		panic(err)
	}

	var dbFiles []string
	dbFiles = append(dbFiles, "Criar nova conta")
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".db" {
			dbFiles = append(dbFiles, file.Name())
		}
	}

	return dbFiles
}
