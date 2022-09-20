package bling

import (
	"time"
)

const DefaultTimeFormat = "02/01/2006"
const DefaultTimeout = 60 * time.Second
const DefaultNotFoundException = "A informacao desejada nao foi encontrada"
const DefaultResponseType = "/json"
const DefaultUrl = "https://bling.com.br/Api/v2"
const ProductsUrl = DefaultUrl + "/produtos"
const CategorysUrl = DefaultUrl + "/categorias"
