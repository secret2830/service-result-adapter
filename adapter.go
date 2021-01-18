package main

import (
	"fmt"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

type Request struct {
	ResultField string      `json:"result_field"`
	Result      interface{} `json:"result"`
}

func handle(req Request) (interface{}, error) {
	logger.Infof("Request received: %+v", req)

	json := models.JSON{}

	json, err := json.Add(req.ResultField, fmt.Sprintf("%v", req.Result))
	if err != nil {
		return nil, err
	}

	return json.String(), nil
}
