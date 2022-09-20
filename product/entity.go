package product

import (
	"encoding/xml"

	"github.com/integrmais/bling"
)

type ResponseModel struct {
	Response struct {
		bling.ResponseErrorModel
		Products []struct {
			Product `json:"produto,omitempty"`
		} `json:"produtos,omitempty"`
	} `json:"retorno"`
}

type Category struct {
	ID        string `json:"id"`
	Descricao string `json:"descricao"`
}

type Warehouse struct {
	ID            string `json:"id"`
	Nome          string `json:"nome"`
	Saldo         int    `json:"saldo"`
	Desconsiderar string `json:"desconsiderar"`
	SaldoVirtual  int    `json:"saldoVirtual"`
}

type Images struct {
	Url string `json:"url" xml:"url"`
}

type Product struct {
	XMLName               xml.Name `xml:"produto"`
	ID                    string   `json:"id" xml:"id"`
	Codigo                string   `json:"codigo" xml:"codigo"`
	Descricao             string   `json:"descricao" xml:"descricao"`
	Tipo                  string   `json:"tipo" xml:"tipo"`
	Situacao              string   `json:"situacao" xml:"situacao"`
	Unidade               string   `json:"unidade" xml:"unidade"`
	Preco                 string   `json:"preco" xml:"preco"`
	PrecoCusto            string   `json:"precoCusto" xml:"precoCusto"`
	DescricaoCurta        string   `json:"descricaoCurta" xml:"descricaoCurta"`
	DescricaoComplementar string   `json:"descricaoComplementar" xml:"descricaoComplementar"`
	DataInclusao          string   `json:"dataInclusao"  xml:",omitempty"`
	DataAlteracao         string   `json:"dataAlteracao" xml:",omitempty"`
	ImageThumbnail        string   `json:"imageThumbnail"  xml:",omitempty"`
	URLVideo              string   `json:"urlVideo" xml:"urlVideo"`
	NomeFornecedor        string   `json:"nomeFornecedor" xml:"nomeFornecedor"`
	CodigoFabricante      string   `json:"codigoFabricante" xml:"codigoFabricante"`
	Marca                 string   `json:"marca" xml:"marca"`
	ClassFiscal           string   `json:"class_fiscal" xml:"class_fiscal"`
	Cest                  string   `json:"cest" xml:"cest"`
	Origem                string   `json:"origem" xml:"origem"`
	IDGrupoProduto        string   `json:"idGrupoProduto" xml:"idGrupoProduto"`
	LinkExterno           string   `json:"linkExterno" xml:"linkExterno"`
	Observacoes           string   `json:"observacoes" xml:"observacoes"`
	Garantia              string   `json:"garantia" xml:"garantia"`
	DescricaoFornecedor   string   `json:"descricaoFornecedor" xml:"descricaoFornecedor"`
	IDFabricante          string   `json:"idFabricante" xml:"idFabricante"`
	Categoria             Category `json:"categoria"  xml:",omitempty"`
	IDCategoria           string   `json:"idCategoria" xml:"idCategoria"`
	PesoLiq               string   `json:"pesoLiq" xml:"pesoLiq"`
	PesoBruto             string   `json:"pesoBruto" xml:"pesoBruto"`
	EstoqueMinimo         string   `json:"estoqueMinimo" xml:"estoqueMinimo"`
	EstoqueMaximo         string   `json:"estoqueMaximo" xml:"estoqueMaximo"`
	Gtin                  string   `json:"gtin" xml:"gtin"`
	GtinEmbalagem         string   `json:"gtinEmbalagem" xml:"gtinEmbalagem"`
	LarguraProduto        string   `json:"larguraProduto" xml:"larguraProduto"`
	AlturaProduto         string   `json:"alturaProduto" xml:"alturaProduto"`
	ProfundidadeProduto   string   `json:"profundidadeProduto" xml:"profundidadeProduto"`
	UnidadeMedida         string   `json:"unidadeMedida" xml:"unidadeMedida"`
	ItensPorCaixa         int      `json:"itensPorCaixa" xml:"itensPorCaixa"`
	Volumes               int      `json:"volumes" xml:"volumes"`
	Localizacao           string   `json:"localizacao" xml:"localizacao"`
	Crossdocking          string   `json:"crossdocking" xml:"crossdocking"`
	Condicao              string   `json:"condicao" xml:"condicao"`
	FreteGratis           string   `json:"freteGratis" xml:"freteGratis"`
	Producao              string   `json:"producao" xml:"producao"`
	DataValidade          string   `json:"dataValidade" xml:"dataValidade"`
	SpedTipoItem          string   `json:"spedTipoItem" xml:"spedTipoItem"`
	ClonarDadosPai        string   `json:"clonarDadosPai" xml:"clonarDadosPai"`
	CodigoPai             string   `json:"codigoPai" xml:"codigoPai"`
	EstoqueAtual          int      `json:"estoqueAtual" xml:"estoqueAtual"`
	Depositos             []struct {
		Warehouse `json:"deposito" xml:"deposito"`
	} `json:"depositos"`
	Deposito Warehouse `json:"deposito" xml:"deposito"`
	Imagens  Images    `json:"imagens" xml:"imagens"`
}
