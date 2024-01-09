package models

type Atividade struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type QSA struct {
	Nome string `json:"nome"`
	Qual string `json:"qual"`
}

type ResponseCnpj struct {
	Abertura              string      `json:"abertura"`
	Situacao              string      `json:"situacao"`
	Tipo                  string      `json:"tipo"`
	Nome                  string      `json:"nome"`
	Porte                 string      `json:"porte"`
	NaturezaJuridica      string      `json:"natureza_juridica"`
	AtividadePrincipal    []Atividade `json:"atividade_principal"`
	AtividadesSecundarias []Atividade `json:"atividades_secundarias"`
	QSA                   []QSA       `json:"qsa"`
	Logradouro            string      `json:"logradouro"`
	Numero                string      `json:"numero"`
	Complemento           string      `json:"complemento"`
	Municipio             string      `json:"municipio"`
	Bairro                string      `json:"bairro"`
	UF                    string      `json:"uf"`
	CEP                   string      `json:"cep"`
	Email                 string      `json:"email"`
	Telefone              string      `json:"telefone"`
	DataSituacao          string      `json:"data_situacao"`
	CNPJ                  string      `json:"cnpj"`
	UltimaAtualizacao     string      `json:"ultima_atualizacao"`
	Status                string      `json:"status"`
	Fantasia              string      `json:"fantasia"`
	EFR                   string      `json:"efr"`
	MotivoSituacao        string      `json:"motivo_situacao"`
	SituacaoEspecial      string      `json:"situacao_especial"`
	DataSituacaoEspecial  string      `json:"data_situacao_especial"`
	CapitalSocial         string      `json:"capital_social"`
	Extra                 struct{}    `json:"extra"`
	Billing               struct {
		Free     bool `json:"free"`
		Database bool `json:"database"`
	} `json:"billing"`
}
