package mictorio

// A partr da ocupação dos mictórios retorna os lugares disponíveis sem um mijão do lado
// Recebe um parâmetro um array de booleanos aonde a true informa que o mictório naquela
// posição está ocupado.
func lugaresParaMijar(ocupados []bool) []bool {

	// Definindo quantidade de mictórios existentes
	// Observe o uso de acentos nas variáveis, Go aceita UTF-8 no nomes das variáveis
	quantidadeMictórios := len(ocupados)

	// Pré-inicializa um slice to mesmo tamanho que o ocupados. Por ser um slice
	// de booleanos todos os valores serã pré-inicializados como false.
	disponibilidade := make([]bool, quantidadeMictórios, quantidadeMictórios)

	// Itera sobre todos os ocupados. A palavra chave range retorna os índice e o
	// valor do slice. Como apenas precisamos do índice para executar a lógica,
	// utilizaremos "_" para ignorar o valor.
	for índice, _ := range ocupados {
		var próximo = índice + 1
		var anterior = índice - 1

		if !ocupados[índice] {
			disponibilidade[índice] = true

			if próximo < quantidadeMictórios && ocupados[próximo] {
				disponibilidade[índice] = false
			}

			if anterior >= 0 && ocupados[anterior] {
				disponibilidade[índice] = false
			}
		}
	}

	return disponibilidade
}
