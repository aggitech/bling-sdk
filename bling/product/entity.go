package product

import "github.com/integrmais/bling/bling"

type ResponseModel struct {
	Response struct {
		bling.ResponseErrorModel
		Products []struct {
			Product `json:"produto"`
		} `json:"produtos"`
	} `json:"retorno"`
}

type Product struct {
	ID                    string      `json:"id"`
	Codigo                string      `json:"codigo"`
	Descricao             string      `json:"descricao"`
	Tipo                  string      `json:"tipo"`
	Situacao              string      `json:"situacao"`
	Unidade               string      `json:"unidade"`
	Preco                 string      `json:"preco"`
	PrecoCusto            string      `json:"precoCusto"`
	DescricaoCurta        string      `json:"descricaoCurta"`
	DescricaoComplementar interface{} `json:"descricaoComplementar"`
	DataInclusao          string      `json:"dataInclusao"`
	DataAlteracao         string      `json:"dataAlteracao"`
	ImageThumbnail        interface{} `json:"imageThumbnail"`
	URLVideo              string      `json:"urlVideo"`
	NomeFornecedor        string      `json:"nomeFornecedor"`
	CodigoFabricante      string      `json:"codigoFabricante"`
	Marca                 string      `json:"marca"`
	ClassFiscal           string      `json:"class_fiscal"`
	Cest                  string      `json:"cest"`
	Origem                string      `json:"origem"`
	IDGrupoProduto        string      `json:"idGrupoProduto"`
	LinkExterno           interface{} `json:"linkExterno"`
	Observacoes           string      `json:"observacoes"`
	GrupoProduto          interface{} `json:"grupoProduto"`
	Garantia              string      `json:"garantia"`
	DescricaoFornecedor   string      `json:"descricaoFornecedor"`
	IDFabricante          string      `json:"idFabricante"`
	Categoria             struct {
		ID        string `json:"id"`
		Descricao string `json:"descricao"`
	} `json:"categoria"`
	PesoLiq             string      `json:"pesoLiq"`
	PesoBruto           string      `json:"pesoBruto"`
	EstoqueMinimo       string      `json:"estoqueMinimo"`
	EstoqueMaximo       string      `json:"estoqueMaximo"`
	Gtin                string      `json:"gtin"`
	GtinEmbalagem       string      `json:"gtinEmbalagem"`
	LarguraProduto      string      `json:"larguraProduto"`
	AlturaProduto       string      `json:"alturaProduto"`
	ProfundidadeProduto string      `json:"profundidadeProduto"`
	UnidadeMedida       string      `json:"unidadeMedida"`
	ItensPorCaixa       int         `json:"itensPorCaixa"`
	Volumes             int         `json:"volumes"`
	Localizacao         string      `json:"localizacao"`
	Crossdocking        string      `json:"crossdocking"`
	Condicao            string      `json:"condicao"`
	FreteGratis         string      `json:"freteGratis"`
	Producao            string      `json:"producao"`
	DataValidade        interface{} `json:"dataValidade"`
	SpedTipoItem        string      `json:"spedTipoItem"`
	ClonarDadosPai      string      `json:"clonarDadosPai"`
	CodigoPai           string      `json:"codigoPai"`
	EstoqueAtual        int         `json:"estoqueAtual"`
	Depositos           []struct {
		Deposito struct {
			ID            string `json:"id"`
			Nome          string `json:"nome"`
			Saldo         int    `json:"saldo"`
			Desconsiderar string `json:"desconsiderar"`
			SaldoVirtual  int    `json:"saldoVirtual"`
		} `json:"deposito"`
	} `json:"depositos"`
}
