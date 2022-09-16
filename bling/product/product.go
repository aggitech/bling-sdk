package product

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/integrmais/bling/bling"
)

type ProductService struct {
	AppKey  string
	Client  *http.Client
	Timeout time.Duration
}

func NewProductService(appKey string, c *http.Client, t time.Duration) *ProductService {
	return &ProductService{
		AppKey:  appKey,
		Client:  c,
		Timeout: t,
	}
}

func HandlerError(req ResponseModel) error {
	if len(req.Response.Errors) == 0 {
		return nil
	}

	reqErrors := []string{}
	for _, e := range req.Response.Errors {
		reqErrors = append(reqErrors, e.Error.Message)
	}

	return errors.New(
		strings.Join(reqErrors, ""),
	)
}

func HandlerProduct(response *http.Response) (ResponseModel, error) {
	defer response.Body.Close()

	var products ResponseModel
	err := json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		return ResponseModel{}, err
	}

	err = HandlerError(products)
	if err != nil {
		return ResponseModel{}, err
	}

	return products, nil
}

func (s *ProductService) GetProductById(ctx context.Context, productID int) (Product, error) {
	url := bling.ProductUrl + bling.DefaultResponseType

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	q := req.URL.Query()

	q.Add("apikey", s.AppKey)
	q.Add("imagem", "S")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		return Product{}, err
	}

	res, err := s.Client.Do(req)
	if err != nil {
		return Product{}, err
	}

	p, err := HandlerProduct(res)
	if err != nil {
		return Product{}, err
	}

	return p.Response.Products[0].Product, nil
}
