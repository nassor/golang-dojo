package mictorio

import (
	"reflect"
	"testing"
)

// Estrutura contendo a situação de um mictório;
type Mictório struct {
	ocupados        []bool
	disponibilidade []bool
}

// Testa a função lugaresParaMijar([]bool) em diversas situações diferentes
func TestLugaresParaMijar(t *testing.T) {

	// lista de situações em que um mictório pode se encontrar
	var mictórios = []Mictório{
		{[]bool{false}, []bool{true}},
		{[]bool{true}, []bool{false}},

		{[]bool{false, true}, []bool{false, false}},
		{[]bool{true, false}, []bool{false, false}},

		{[]bool{false, false}, []bool{true, true}},
		{[]bool{true, true}, []bool{false, false}},

		{[]bool{false, false, false}, []bool{true, true, true}},
		{[]bool{false, false, true}, []bool{true, false, false}},
		{[]bool{false, true, false}, []bool{false, false, false}},
		{[]bool{false, true, true}, []bool{false, false, false}},

		{[]bool{true, false, false}, []bool{false, false, true}},
		{[]bool{true, false, true}, []bool{false, false, false}},
		{[]bool{true, true, false}, []bool{false, false, false}},
		{[]bool{true, true, true}, []bool{false, false, false}},

		{[]bool{false, false, false, false}, []bool{true, true, true, true}},
		{[]bool{false, false, false, true}, []bool{true, true, false, false}},
		{[]bool{false, false, true, false}, []bool{true, false, false, false}},
		{[]bool{false, false, true, true}, []bool{true, false, false, false}},
		{[]bool{false, true, false, false}, []bool{false, false, false, true}},
		{[]bool{false, true, false, true}, []bool{false, false, false, false}},
		{[]bool{false, true, true, false}, []bool{false, false, false, false}},
		{[]bool{false, true, true, true}, []bool{false, false, false, false}},

		{[]bool{true, false, false, false}, []bool{false, false, true, true}},
		{[]bool{true, false, false, true}, []bool{false, false, false, false}},
		{[]bool{true, false, true, false}, []bool{false, false, false, false}},
		{[]bool{true, false, true, true}, []bool{false, false, false, false}},
		{[]bool{true, true, false, false}, []bool{false, false, false, true}},
		{[]bool{true, true, false, true}, []bool{false, false, false, false}},
		{[]bool{true, true, true, false}, []bool{false, false, false, false}},
		{[]bool{true, true, true, true}, []bool{false, false, false, false}},
	}

	// itera sob as diversas situações comparando se a função retornos ou disponibilidade de forma correta
	for _, situação := range mictórios {
		resposta := lugaresParaMijar(situação.ocupados)
		if !reflect.DeepEqual(resposta, situação.disponibilidade) {
			t.Errorf("lugaresParaMijar falhou na situação: %v | Recebido: %v", situação, resposta)
		}
	}
}
