package main

import "fmt"

const Million = 1000000

func main() {
	var mz_teorico, mz_observado, erro_ppm float32

	for {
		fmt.Printf("Digite o m/z teórico: ")
		fmt.Scan(&mz_teorico)

		fmt.Printf("Digite o m/z observado: ")
		fmt.Scan(&mz_observado)

		erro_ppm = (mz_observado - mz_teorico) / mz_teorico * Million

		fmt.Printf("Erro (ppm) = %v\n", erro_ppm)

		fmt.Println("Pressione Ctrl + C para sair.")
	}
}
