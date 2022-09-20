package category

import "github.com/integrmais/bling"

type Category struct {
	Id             int64  `json:"id"`
	Descricao      string `json:"descricao"`
	IdCategoriaPai int64  `json:"idCategoriaPai"`
}

type ResponseModel struct {
	Response struct {
		bling.ResponseErrorModel
		Categories []struct {
			Category `json:"categoria,omitempty"`
		} `json:"categorias,omitempty"`
	} `json:"retorno"`
}
