package core

import "os"

func MostrarMenu() {
	println("1 - Iniciar Monitoramento")
	println("2 - Exibir Logs")
	println("0 - Sair do Programa")
	println("Selecione um dos numeros acima.")
}

func SelecionarOpcaoMenu(opcao int, sites []string) {
	monitoramentos := 5
	delay := 5
	var opcaoSelecionada int = opcao

	switch opcaoSelecionada {
	case 1:
		InciarMonitoramento(sites, monitoramentos, delay)
	case 2:
		ExibirLogs()
	case 0:
		println("Saindo do programa...")
		os.Exit(0)
	default:
		println("Opção inválida. Tente novamente.")
		os.Exit(-1)
	}
}
