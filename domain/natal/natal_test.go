package natal

import (
	"encoding/json"
	"testing"

	"github.com/pedroHenriqueMaia/curso-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculaNatalFraquinho(t *testing.T) {
	s := NewNatal()
	_, err := s.Calcula(domain.Parametros{
		Homens:          2,
		Mulheres:        1,
		Criancas:        2,
		Acompanhamentos: true,
	})
	assert.Equal(t, err.Error(), "Homens e mulheres devem ser igual ou maior que cinco")
}

func TestCalculaNatalBombando(t *testing.T) {
	s := NewNatal()
	ch, err := s.Calcula(domain.Parametros{
		Homens:          4,
		Mulheres:        4,
		Criancas:        2,
		Acompanhamentos: true,
	})
	assert.Nil(t, err)

	esperado := Natal{
		TotalCarne:           2000,
		TotalPessoas:         10,
		TotalAcompanhamentos: 1000,
		NaoAlcoolicas:        10,
		Vinho:                5,
		Peru:                 3000,
	}

	j, err := ch.ToJSON()
	jEsperado, err := json.Marshal(esperado)

	assert.Equal(t, ch, esperado)
	assert.Equal(t, j, jEsperado)
}
