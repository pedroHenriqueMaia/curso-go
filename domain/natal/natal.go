package natal

import (
	"encoding/json"
	"fmt"

	"github.com/pedroHenriqueMaia/curso-go/domain"
)

type Natal struct {
	TotalCarne           int `json:"total-carne"`
	TotalPessoas         int `json:"total-pessoas"`
	TotalAcompanhamentos int `json:"total-acompanhamentos"`
	NaoAlcoolicas        int `json:"nao-alcoolicas"`
	Vinho                int `json:"vinho"`
	Peru                 int `json:"peru"`
}

func (b Natal) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

type Service struct{}

func NewNatal() Service {
	return Service{}
}

func (s Service) Calcula(p domain.Parametros) (domain.Resultado, error) {
	b := Natal{}
	if p.Mulheres <= 0 || p.Homens <= 0 || p.Criancas <= 0 {
		return b, fmt.Errorf("Festa fantasma?")
	}

	b.TotalCarne = p.Mulheres*150 + p.Homens*150 + p.Criancas*100
	b.TotalPessoas = p.Mulheres + p.Homens + p.Criancas
	b.TotalAcompanhamentos = b.TotalPessoas * 50
	b.NaoAlcoolicas = 400 * b.TotalPessoas
	b.Vinho = b.TotalPessoas / 2
	b.Peru = p.Mulheres*200 + p.Homens*250 + p.Criancas*100

	return b, nil
}
