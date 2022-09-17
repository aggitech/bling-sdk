package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/integrmais/bling"
	"github.com/integrmais/bling/internal"
)

type ProductService struct {
	AppKey string
	Client *http.Client
}

func NewProductService(appKey string, c *http.Client) *ProductService {
	return &ProductService{
		AppKey: appKey,
		Client: c,
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

func HandlerResponse(res *http.Response) (ResponseModel, error) {
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	var response ResponseModel
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ResponseModel{}, err
	}

	err = HandlerError(response)
	if err != nil {
		return ResponseModel{}, err
	}

	return response, nil
}

func (s *ProductService) GetProductById(ctx context.Context, productID string) (Product, error) {
	url := fmt.Sprintf(
		"%s/produto/%s/%s",
		bling.DefaultUrl,
		productID,
		bling.DefaultResponseType,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Product{}, err
	}

	q := req.URL.Query()
	q.Add("apikey", s.AppKey)
	q.Add("imagem", "S")
	q.Add("situacao", "A")
	q.Add("tipo", "P")

	req.URL.RawQuery = q.Encode()

	res, err := s.Client.Do(req)
	if err != nil {
		return Product{}, err
	}

	p, err := HandlerResponse(res)
	if err != nil {
		return Product{}, err
	}

	return p.Response.Products[0].Product, nil
}

func (s *ProductService) GetByRange(ctx context.Context, startAt time.Time) ([]Product, error) {
	url := bling.ProductsUrl + bling.DefaultResponseType

	by := internal.NormalizeDate(startAt)
	at := internal.NormalizeDate(time.Now())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", s.AppKey)
	q.Add("imagem", "S")
	q.Add("situacao", "A")
	q.Add("tipo", "P")
	q.Add("filters", fmt.Sprintf("dataAlteracao[%s TO %s]", by, at))

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.RawQuery, q)

	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	p, err := HandlerResponse(res)
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, product := range p.Response.Products {
		products = append(products, product.Product)
	}

	return products, nil
}
