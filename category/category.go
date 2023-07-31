package category

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/aggitech/bling-sdk"
)

type CategoryService struct {
	AppKey string
	Client *http.Client
}

func NewBlingCategoryService(appKey string, c *http.Client) *CategoryService {
	return &CategoryService{
		AppKey: appKey,
		Client: c,
	}
}

func HandlerResponse(res *http.Response) (ResponseModel, error) {
	defer res.Body.Close()

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

func HandlerError(req ResponseModel) error {
	if req.Response.Errors[0].Error.Code == 0 {
		return nil
	}

	var errMessages []string
	for _, err := range req.Response.Errors {
		errMessages = append(errMessages, err.Error.Message)
	}

	return errors.New(strings.Join(errMessages, "\n"))
}

func (s *CategoryService) Get(ctx context.Context) (ResponseModel, error) {
	url := bling.CategorysUrl + bling.DefaultResponseType

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ResponseModel{}, err
	}

	q := req.URL.Query()
	q.Add("apikey", s.AppKey)

	req.URL.RawQuery = q.Encode()

	res, err := s.Client.Do(req)
	if err != nil {
		return ResponseModel{}, err
	}

	r, err := HandlerResponse(res)
	if err != nil {
		return ResponseModel{}, err
	}

	return r, nil
}
