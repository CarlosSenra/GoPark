package core

import (
	"fmt"
	"net/http"
	"time"
)

func InciarMonitoramento(sites []string, timesRun int, sleepSeconds int) {

	fmt.Println("Monitoramento iniciado...")

	for i := 0; i < timesRun; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)
			testarSite(site)

		}
		fmt.Println("Monitoramento finalizado.\n Esperando", sleepSeconds, "segundos para o próximo monitoramento...")
		println(" ")
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

}

func testarSite(site string) (*http.Response, error) {

	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreou algum problema.", err)
		return nil, err
	}

	fmt.Println("Resposta do site:", response.Status)

	return response, nil
}
