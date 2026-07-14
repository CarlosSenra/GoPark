package core

import (
	"fmt"
	"net/http"
	"time"
)

func InciarMonitoramento(sites []string, timesRun int, sleepSeconds int, PathLog string) {
	var mensagem string
	fmt.Println("Monitoramento iniciado...")

	for i := 0; i < timesRun; i++ {
		for idx, site := range sites {
			fmt.Println("Testando site", idx+1, ":", site)
			mensagem = time.Now().Format("02/01/2006 15:04:05") + " - Testando site " + site
			testarSite(site, PathLog)
			EscreverEmArquivo(PathLog, mensagem)
		}

		fmt.Println("Monitoramento finalizado.\n Esperando", sleepSeconds, "segundos para o próximo monitoramento...")
		println(" ")
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

}

func testarSite(site string, PathLog string) (*http.Response, error) {

	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreou algum problema.", err)
		mensagem := time.Now().Format("02/01/2006 15:04:05") + " - Ocorreu algum problema ao testar o site " + site + ": " + err.Error()
		EscreverEmArquivo(PathLog, mensagem)
		return nil, err
	}

	defer response.Body.Close()

	fmt.Println("Resposta do site:", response.Status)
	EscreverEmArquivo(PathLog, time.Now().Format("02/01/2006 15:04:05")+" - Resposta do site "+site+": "+response.Status)

	return response, nil
}
