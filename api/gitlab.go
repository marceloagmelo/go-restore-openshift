package api

import (
	"strings"

	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

// ListarTags listar
func ListarTags() (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderGitlab()
	apiRequest.EndPoint = variaveis.GitlabApiURL + variaveis.GitlabApiProjetos + "/" + variaveis.GitlabProjectID + "/repository/tags?order_by=name&sort=desc"
	apiRequest.Metodo = "GET"

	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

// ListarArquivos listar
func ListarArquivos(ref, path, qtdePorPagina string) (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderGitlab()
	apiRequest.EndPoint = variaveis.GitlabApiURL + variaveis.GitlabApiProjetos + "/" + variaveis.GitlabProjectID + "/repository/tree?ref=" + ref + "&path=" + path + "&recursive=true&per_page=" + qtdePorPagina
	apiRequest.Metodo = "GET"

	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

//getHeader
func montarHeaderGitlab() []Header {
	var headers []Header
	header := Header{}

	header.Chave = "PRIVATE-TOKEN"
	header.Valor = strings.TrimSuffix(variaveis.GitlabToken, "\n")
	headers = append(headers, header)

	return headers
}
