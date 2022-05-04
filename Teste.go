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

const monitoramento = 3
const delayMonitoramento = 2

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
			//SAIU COM SUCESSO!!
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1)
			//SAIDA COM ERROR!! DEU RUIM
		}
	}
}

func exibeIntroducao() {
	versao := 1.1
	fmt.Println("Olá, seja bem vindo!")
	fmt.Println("Versão do programa:", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("Selecione um Comando:")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando iniciado...")
	sites := leSitesArquivos()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, "em monitoramento:", site)
			testaSite(site)
		}
		time.Sleep(delayMonitoramento * time.Second)
		fmt.Println("")
	}
	fmt.Println("Fim do monitoramento...")
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code: ",
			resp.StatusCode)
	}
}

func leSitesArquivos() []string {
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
