package core

import (
	"fmt"
	"os"
)

func EscreverEmArquivo(Path string, Mensagem string) {
	arquivo, err := os.OpenFile(Path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de logs:", err)
		return
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(Mensagem + "\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo de logs:", err)
	}
}

func ExibirLogs(path string) {
	arquivo, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de logs:", err)
		return
	}
	defer arquivo.Close()

	dados, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de logs:", err)
		return
	}
	fmt.Println(string(dados))
}
