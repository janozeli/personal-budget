package tui

import (
	"fmt"
	"main/db"

	"github.com/rivo/tview"
)

func ShowDataBaseOptions(app *tview.Application) {
	existingDBs := db.VerifyExistingDBs()
	if len(existingDBs) > 0 {
		list := tview.NewList()
		for _, db := range existingDBs {
			list.AddItem(db, "", 0, func() {
				selectedDB := db
				fmt.Println("Banco de dados selecionado:", selectedDB)
				app.Stop()
			})
		}
		list.AddItem("Criar novo banco de dados", "", 0, func() {
			fmt.Println("Criando novo banco de dados...")
			app.Stop()
		})

		app.SetRoot(list, true).Run()
	} else {
		form := tview.NewForm()
		form.AddInputField("Nome do banco de dados:", "", 20, nil, nil)
		form.AddButton("Criar", func() {
			dbName := form.GetFormItem(0).(*tview.InputField).GetText()
			fmt.Println("Criando banco de dados:", dbName)
			app.Stop()
		})
		app.SetRoot(form, true).Run()
	}
}
