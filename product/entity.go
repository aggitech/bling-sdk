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

type ResponseCreatorModel struct {
	Response struct {
		bling.ResponseErrorModel
		Products []struct {
			ProductCreator `json:"produto,omitempty"`
		} `json:"produtos,omitempty"`
	} `json:"retorno"`
}

type Category struct {
	ID        string `json:"id"`
	Descricao string `json:"descricao"`
}

type Warehouse struct {
	ID            string `json:"id" xml:"id,omitempty"`
	Nome          string `json:"nome"`
	Saldo         int    `json:"saldo" xml:"estoque,omitempty"`
	Desconsiderar string `json:"desconsiderar"`
	SaldoVirtual  int    `json:"saldoVirtual"`
}

type Variant struct {
	//XMLName      xml.Name `xml:"variacao"`
	Nome           string `json:"nome,omitempty" xml:"nome,omitempty"`
	Codigo         string `json:"codigo,omitempty" xml:"codigo,omitempty"`
	VlrUnit        string `json:"vlr_unit,omitempty" xml:"vlr_unit,omitempty"`
	EstoqueAtual   int    `json:"estoqueAtual,omitempty" xml:"estoque,omitempty"`
	ClonarDadosPai string `json:"clonarDadosPai,omitempty" xml:"clonarDadosPai,omitempty"`
	Depositos      []struct {
		Deposito Warehouse `json:"depositos,omitempty" xml:"depositos,omitempty"`
	} `json:"depositos,omitempty" xml:"depositos,omitempty"`
}

type Images struct {
	Url string `json:"link" xml:"url"`
}

type Product struct {
	XMLName               xml.Name `xml:"produto"`
	Codigo                string   `json:"codigo" xml:"codigo,omitempty"`
	Descricao             string   `json:"descricao" xml:"descricao,omitempty"`
	Tipo                  string   `json:"tipo" xml:"tipo,omitempty"`
	Situacao              string   `json:"situacao" xml:"situacao,omitempty"`
	Unidade               string   `json:"unidade" xml:"unidade,omitempty"`
	Preco                 string   `json:"preco" xml:"preco,omitempty"`
	PrecoCusto            string   `json:"precoCusto" xml:"preco_custo,omitempty"`
	VlrUnit               string   `xml:"vlr_unit,omitempty"`
	DescricaoCurta        string   `json:"descricaoCurta" xml:"descricaoCurta,omitempty"`
	DescricaoComplementar string   `json:"descricaoComplementar" xml:"descricaoComplementar,omitempty"`
	DataInclusao          string   `json:"dataInclusao"  xml:",omitempty"`
	DataAlteracao         string   `json:"dataAlteracao" xml:",omitempty"`
	ImageThumbnail        string   `json:"imageThumbnail"  xml:",omitempty"`
	URLVideo              string   `json:"urlVideo" xml:"urlVideo,omitempty"`
	NomeFornecedor        string   `json:"nomeFornecedor" xml:"nomeFornecedor,omitempty"`
	CodigoFabricante      string   `json:"codigoFabricante" xml:"codigoFabricante,omitempty"`
	Marca                 string   `json:"marca" xml:"marca,omitempty"`
	ClassFiscal           string   `json:"class_fiscal" xml:"class_fiscal,omitempty"`
	Cest                  string   `json:"cest" xml:"cest,omitempty"`
	Origem                string   `json:"origem" xml:"origem,omitempty"`
	IDGrupoProduto        string   `json:"idGrupoProduto" xml:"idGrupoProduto,omitempty"`
	LinkExterno           string   `json:"linkExterno" xml:"linkExterno,omitempty"`
	Observacoes           string   `json:"observacoes" xml:"observacoes,omitempty"`
	Garantia              string   `json:"garantia" xml:"garantia,omitempty"`
	DescricaoFornecedor   string   `json:"descricaoFornecedor" xml:"descricaoFornecedor,omitempty"`
	IDFabricante          string   `json:"idFabricante" xml:"idFabricante,omitempty"`
	Categoria             Category `json:"categoria"  xml:",omitempty"`
	IDCategoria           string   `json:"idCategoria" xml:"idCategoria,omitempty"`
	PesoLiq               string   `json:"pesoLiq" xml:"peso_liq,omitempty"`
	PesoBruto             string   `json:"pesoBruto" xml:"peso_bruto,omitempty"`
	EstoqueMinimo         string   `json:"estoqueMinimo" xml:"estoqueMinimo,omitempty"`
	EstoqueMaximo         string   `json:"estoqueMaximo" xml:"estoqueMaximo,omitempty"`
	Gtin                  string   `json:"gtin" xml:"gtin,omitempty"`
	GtinEmbalagem         string   `json:"gtinEmbalagem" xml:"gtinEmbalagem,omitempty"`
	LarguraProduto        string   `json:"larguraProduto" xml:"largura,omitempty"`
	AlturaProduto         string   `json:"alturaProduto" xml:"altura,omitempty"`
	ProfundidadeProduto   string   `json:"profundidadeProduto" xml:"profundidade,omitempty"`
	UnidadeMedida         string   `json:"unidadeMedida" xml:"unidadeMedida,omitempty"`
	ItensPorCaixa         int      `json:"itensPorCaixa" xml:"itensPorCaixa,omitempty"`
	Volumes               int      `json:"volumes" xml:"volumes,omitempty"`
	Localizacao           string   `json:"localizacao" xml:"localizacao,omitempty"`
	Crossdocking          string   `json:"crossdocking" xml:"crossdocking,omitempty"`
	Condicao              string   `json:"condicao" xml:"condicao,omitempty"`
	FreteGratis           string   `json:"freteGratis" xml:"freteGratis,omitempty"`
	Producao              string   `json:"producao" xml:"producao,omitempty"`
	DataValidade          string   `json:"dataValidade" xml:"dataValidade,omitempty"`
	SpedTipoItem          string   `json:"spedTipoItem" xml:"spedTipoItem,omitempty"`
	ClonarDadosPai        string   `json:"clonarDadosPai" xml:"clonarDadosPai,omitempty"`
	CodigoPai             string   `json:"codigoPai" xml:"codigoPai,omitempty"`
	EstoqueAtual          int      `json:"estoqueAtual" xml:"estoqueAtual,omitempty"`
	Depositos             []struct {
		Warehouse `json:"deposito"`
	} `json:"depositos"`
	Deposito  Warehouse `json:"deposito.omitempty" xml:"deposito,omitempty"`
	Variacoes []struct {
		Variant `json:"variacao,omitempty" xml:"variacao,omitempty"`
	} `json:"variacoes,omitempty" xml:"variacoes>variacao,omitempty"`
	Imagens []Images `json:"imagens,omitempty" xml:"imagens,omitempty"`
	Image   []Images `json:"imagem,omitempty" xml:"image,omitempty"`
}

type ProductCreator struct {
	Codigo string `json:"codigo,omitempty"`
}
