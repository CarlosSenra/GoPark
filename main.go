package main

import (
	"fmt"
	"modulo/core"
)

func main() {
	sites := core.LerSitesDeArquivo("sites.txt")
	// sites := []string{"https://www.google.com.br", "https://www.youtube.com.br", "https://www.facebook.com.br"}
	for {
		core.MostrarMenu()
		var opcao int
		fmt.Scan(&opcao)
		core.SelecionarOpcaoMenu(opcao, sites)
	}
}
