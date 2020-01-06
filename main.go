package main

import(
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	urls := []string{
		"https://bus.paas.santanderbr.pre.corp/bus-bloqueios-java",
		"https://cpt.paas.santanderbr.pre.corp/gest-bloqueios/gest-bloqueios/treasury-block-manager/v1/health",
		"https://cptocu.paas.santanderbr.pre.corp/gest-posicao",
		"https://cpt.homologacao.paas.santnaderbr.pre.corp/cat-produtos",
		"https://bus.paas.santanderbr.pre.corp/bus-audit",
	}

	for i := 0; i < len(urls); i++ {
		ping(urls[i])
	}
}

func ping(url string) {
	out, _ := exec.Command("ping", url).Output()
	if strings.Contains(string(out), "Destination Host Unreachable"); strings.Contains(string(out), "Ping request could not find host") {
		fmt.Println(url + " is down")
	} else {
		fmt.Println(url + " is up")
	}
}
