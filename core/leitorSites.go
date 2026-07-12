package core

import (
	"bufio"
	"os"
	"strings"
)

func LerSitesDeArquivo(Path string) []string {
	var sites []string
	arquivo, err := os.Open(Path)
	if err != nil {
		println("Ocorreu um erro ao abrir o arquivo:", err)
		return nil
	}
	defer arquivo.Close()

	leitor := bufio.NewScanner(arquivo)
	for leitor.Scan() {
		linha := strings.TrimSpace(leitor.Text())
		if linha != "" {
			sites = append(sites, linha)
		}
	}
	if err := leitor.Err(); err != nil {
		println("Erro ao ler arquivo:", err)
	}
	return sites
}
