package api

import (
	"io"
	"strings"

	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

//ListarRecurso
func ListarRecurso(endPoint string) (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderOpenshift("application/json")
	apiRequest.EndPoint = endPoint
	apiRequest.Metodo = "GET"

	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

//CriarRecurso
func CriarRecurso(endPoint string, body io.Reader) (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderOpenshift("application/json")
	apiRequest.EndPoint = endPoint
	apiRequest.Body = body
	apiRequest.Metodo = "POST"
	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

//getHeader
func montarHeaderOpenshift(contentType string) []Header {
	var headers []Header
	header := Header{}

	var bearerAuth = "Bearer " + strings.TrimSuffix(variaveis.OpenshiftToken, "\n")
	header.Chave = "Authorization"
	header.Valor = bearerAuth
	headers = append(headers, header)

	header.Chave = "Accept"
	header.Valor = "application/json"
	headers = append(headers, header)

	header.Chave = "Content-Type"
	header.Valor = contentType
	headers = append(headers, header)

	return headers
}
