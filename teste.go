package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 6

func main() {
	introducaoAplicacao()

	for {
		escolherMenu()
		comando := escolherComando()

		switch comando {
		case "1":
			iniciarMonitoramento()
		case "2":
			fmt.Println("Exibindo Logs...")
		case "0":
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não existente")
			os.Exit(125)
		}
	}

}
func introducaoAplicacao() {
	fmt.Println("Hello World")
	nome := "Brenda"
	versao := 1.0
	idade := 23
	fmt.Println("Olá,", nome, "Sua idade é:", idade)
	fmt.Println("Este programa está na versão", versao)

}

func escolherComando() string {
	var escolherComando string
	fmt.Scan(&escolherComando)
	fmt.Println("O comando escolhido foi:", escolherComando)
	return escolherComando
}

func escolherMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSiteArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Minute)

	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSiteArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}
