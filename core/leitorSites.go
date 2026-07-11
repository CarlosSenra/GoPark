package core

import (
	"io/ioutil"
	"strings"
)

func LerSitesDeArquivoTxt(Path string) []string {
	var sites []string

	arquivo, err := ioutil.ReadFile(Path)

	if err != nil {
		println("Ocorreu um erro ao abrir o arquivo:", err)
		return nil
	}

	conteudo := string(arquivo)

	linhas := strings.Split(conteudo, "\n")

	for _, linha := range linhas {
		linha = strings.TrimSpace(linha)
		if linha != "" {
			sites = append(sites, linha)
		}
	}
	return sites

}
