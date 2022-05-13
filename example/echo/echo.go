package echo

import (
	"context"
	"errors"

	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"

	iface2 "github.com/amjadjibon/encoding/iface"
	registry2 "github.com/amjadjibon/encoding/registry"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Echo struct {
	mValues       model.ConfigMap
	mEncodingName string
	mEncoding     iface2.IEncoding
}

func (e *Echo) Name() string {
	return "abesh_example_echo"
}

func (e *Echo) Version() string {
	return "0.0.1"
}

func (e *Echo) Category() string {
	return string(constant.CategoryService)
}

func (e *Echo) ContractId() string {
	return "abesh:ex_echo"
}

func (e *Echo) GetConfigMap() model.ConfigMap {
	return e.mValues
}

func (e *Echo) Setup() error {
	e.mEncoding = registry2.EncodingRegistry().GetEncoding(e.mEncodingName)
	if e.mEncoding == nil {
		return errors.New("encoder not found")
	}
	return nil
}

func (e *Echo) SetConfigMap(values model.ConfigMap) error {
	e.mValues = values
	e.mEncodingName = values.String("encoding_name", "json")
	return nil
}

func (e *Echo) New() iface.ICapability {
	return &Echo{}
}

func (e *Echo) Serve(_ context.Context, input *model.Event) (*model.Event, error) {
	var response = &Response{
		Code:    200,
		Message: "OK",
		Data: map[string]string{
			"message": "Hello World",
		},
	}

	data, err := e.mEncoding.Marshal(response)
	if err != nil {
		return nil, err
	}

	return model.GenerateOutputEvent(
		input.Metadata,
		e.ContractId(),
		"OK",
		200,
		"application/json",
		data,
	), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&Echo{})
}
