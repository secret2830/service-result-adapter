package main

import (
	"fmt"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

type Request struct {
	FieldName string      `json:"field_name"`
	Result    interface{} `json:"result"`
}

func handle(req Request) (interface{}, error) {
	logger.Infof("Request received: %+v", req)

	json := models.JSON{}

	json, err := json.Add(req.FieldName, fmt.Sprintf("%v", req.Result))
	if err != nil {
		return nil, err
	}

	return json, nil
}
