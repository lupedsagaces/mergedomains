package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printBanner() {
	banner := `	              
 ._ _   _  ._ _   _         
 | | | (/_ | (_| (/_        
  _|  _  ._ _ _|_. o ._   _ 
 (_| (_) | | | (_| | | | _> 
                            
			By: lupedsagaces		
`

	fmt.Println(banner)
}

func main() {
	printBanner()
	// Solicitar ao usuário os nomes dos arquivos
	var file1Name, file2Name, outputFileName string
	fmt.Print("Digite o nome do primeiro arquivo (ex: bol_fulldomains200.txt): ")
	fmt.Scanln(&file1Name)
	fmt.Print("Digite o nome do segundo arquivo (ex: domains_fin.txt): ")
	fmt.Scanln(&file2Name)
	fmt.Print("Digite o nome do arquivo de saída (ex: merged_domains.txt): ")
	fmt.Scanln(&outputFileName)

	// Ler os domínios de ambos os arquivos
	bolDomains := make(map[string]struct{})
	finDomains := make(map[string]struct{})

	readDomains(file1Name, bolDomains)
	readDomains(file2Name, finDomains)

	// Encontrar domínios comuns e únicos
	var commonDomains []string
	var uniqueFinDomains []string

	for domain := range finDomains {
		if _, exists := bolDomains[domain]; exists {
			commonDomains = append(commonDomains, domain)
		} else {
			uniqueFinDomains = append(uniqueFinDomains, domain)
		}
	}

	// Contar a quantidade de domínios
	numCommon := len(commonDomains)
	numUniqueFin := len(uniqueFinDomains)

	// Mostrar os resultados
	fmt.Println("Domínios comuns:")
	for _, domain := range commonDomains {
		fmt.Println(domain)
	}
	fmt.Printf("\nQuantidade de domínios comuns: %d"+"\n", numCommon)

	fmt.Printf("\nDomínios únicos em %s:\n", file2Name)
	for _, domain := range uniqueFinDomains {
		fmt.Println(domain)
	}
	fmt.Printf("\nQuantidade de domínios únicos em %s: %d\n", file2Name, numUniqueFin)

	// Mesclar as duas listas e salvar em um novo arquivo
	mergedDomains := make(map[string]struct{})
	for domain := range bolDomains {
		mergedDomains[domain] = struct{}{}
	}
	for domain := range finDomains {
		mergedDomains[domain] = struct{}{}
	}

	// Salvar os domínios mesclados em um arquivo
	saveMergedDomains(outputFileName, mergedDomains)

	fmt.Printf("\nDomínios mesclados foram salvos em: %s\n", outputFileName)
}

// Função para ler os domínios de um arquivo e armazená-los em um mapa
func readDomains(filename string, domains map[string]struct{}) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo %s: %s\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			domains[line] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo %s: %s\n", filename, err)
	}
}

// Função para salvar os domínios mesclados em um arquivo
func saveMergedDomains(filename string, domains map[string]struct{}) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo %s: %s\n", filename, err)
		return
	}
	defer file.Close()

	for domain := range domains {
		file.WriteString(domain + "\n")
	}
}
