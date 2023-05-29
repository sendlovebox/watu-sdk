package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sendlovebox/watu-sdk/model"
)

// GetBillGroups gets all bill groups
func (c *Call) GetBillGroups(ctx context.Context) ([]*model.BillGroup, error) {
	endpoint := fmt.Sprintf("%s%s?country=%s", c.baseURL, "/watubill/groups", "NG")
	fL := c.logger.With().Str("func", "GetBillGroups").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "").Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorPayload{}
	response := struct {
		Data []*model.BillGroup `json:"data"`
	}{}

	resp, err := c.client.R().
		SetAuthToken(c.publicKey).
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return nil, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return nil, model.ErrNetworkError
	}

	return response.Data, nil
}

// GetBillsByGroup gets bills by group
func (c *Call) GetBillsByGroup(ctx context.Context, group string) ([]*model.Bill, error) {
	endpoint := fmt.Sprintf("%s%s?country=%s&group=%s", c.baseURL, "/watubill/channels", "NG", group)
	fL := c.logger.With().Str("func", "GetBillsByGroup").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, group).Msg("group")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorPayload{}
	response := struct {
		Data []*model.Bill `json:"data"`
	}{}

	resp, err := c.client.R().
		SetAuthToken(c.publicKey).
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return nil, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return nil, model.ErrNetworkError
	}

	return response.Data, nil
}

// GetBillInformation gets bills info
func (c *Call) GetBillInformation(ctx context.Context, billChannel string) (model.Bill, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, "/channel/info/", billChannel)
	fL := c.logger.With().Str("func", "GetBillsByGroup").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, billChannel).Msg("billChannel")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorPayload{}
	response := struct {
		Data model.Bill `json:"data"`
	}{}

	resp, err := c.client.R().
		SetAuthToken(c.publicKey).
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return model.Bill{}, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return model.Bill{}, model.ErrNetworkError
	}

	return response.Data, nil
}

// GetBillTypes gets bill types
func (c *Call) GetBillTypes(ctx context.Context, billChannel string) ([]*model.BillType, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "/watubill/bill-types")
	fL := c.logger.With().Str("func", "GetBillTypes").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, billChannel).Msg("billChannel")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorPayload{}
	response := struct {
		Data []*model.BillType `json:"data"`
	}{}
	body := struct {
		Channel string `json:"channel"`
	}{
		Channel: billChannel,
	}

	resp, err := c.client.R().
		SetAuthToken(c.publicKey).
		SetBody(body).
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return nil, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return nil, model.ErrNetworkError
	}

	return response.Data, nil
}

// ValidateBill validates bills
func (c *Call) ValidateBill(ctx context.Context, body interface{}) (interface{}, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "/watubill/validate")
	fL := c.logger.With().Str("func", "GetBillTypes").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, body).Msg("body")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorPayload{}
	response := struct {
		Data interface{} `json:"data"`
	}{}

	resp, err := c.client.R().
		SetAuthToken(c.publicKey).
		SetBody(body).
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return nil, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return nil, model.ErrNetworkError
	}

	return response.Data, nil
}

// VendBill actually vends the bill
func (c *Call) VendBill(ctx context.Context, body interface{}) (model.VendBillResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "/watubill/vend")
	fL := c.logger.With().Str("func", "VendBill").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, body).Msg("body")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorPayload{}
	response := struct {
		Data model.VendBillResponse `json:"data"`
	}{}

	resp, err := c.client.R().
		SetAuthToken(c.secretKey).
		SetBody(body).
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return model.VendBillResponse{}, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return model.VendBillResponse{}, model.ErrNetworkError
	}

	return response.Data, nil
}
