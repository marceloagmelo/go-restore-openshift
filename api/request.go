package api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marceloagmelo/go-restore-openshift/logger"
)

//OpenshiftErro retorno de erro
type OpenshiftErro struct {
	APIVersion string `json:"apiVersion"`
	Code       int    `json:"code"`
	Details    struct {
	} `json:"details"`
	Kind     string `json:"kind"`
	Message  string `json:"message"`
	Metadata struct {
	} `json:"metadata"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}

//ApiRequest estrutura ApiRequest
type (
	Requisicao interface {
		cmdRequest() (interface{}, int, error)
	}
	Header struct {
		Chave string
		Valor string
	}
	ApiRequest struct {
		Headers  []Header
		EndPoint string
		Metodo   string
		Body     io.Reader
	}
)

//ExecutarRequest request
func ExecutarRequest(req Requisicao) (interface{}, int, error) {
	resposta, statusCode, err := req.cmdRequest()

	return resposta, statusCode, err
}

//cmdRequest request
func (apiR ApiRequest) cmdRequest() (interface{}, int, error) {
	var interf interface{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest(apiR.Metodo, apiR.EndPoint, apiR.Body)
	if err != nil {
		usuario := fmt.Sprintf("%s: %s", "Erro ao criar um request", err.Error())
		logger.Erro.Println(usuario)
		return nil, http.StatusInternalServerError, err
	}

	for _, header := range apiR.Headers {
		request.Header.Set(header.Chave, header.Valor)
	}

	resposta, err := cliente.Do(request)
	if err != nil {
		usuario := fmt.Sprintf("%s: %s", "Erro ao abrir o request", err.Error())
		logger.Erro.Println(usuario)
		return nil, http.StatusInternalServerError, err
	}
	defer resposta.Body.Close()

	corpo, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo da pagina", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return nil, http.StatusInternalServerError, err
	}
	err = json.Unmarshal(corpo, &interf)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON do Servidor", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return nil, http.StatusInternalServerError, err
	}

	statusCode, err := tratarRetornoRequest(interf, resposta.StatusCode, apiR.Metodo, err)
	if err != nil {
		return interf, statusCode, err
	}

	return interf, statusCode, nil
}

//tratarRetornoRequest o retorno do quest
func tratarRetornoRequest(interf interface{}, statusCode int, metodo string, err error) (int, error) {
	statusCodeValido := http.StatusOK
	if metodo == "POST" {
		statusCodeValido = http.StatusCreated
	}
	reqBody, err := json.Marshal(interf)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para []Byte", err)
		logger.Erro.Println(mensagem)
		return http.StatusInternalServerError, err
	}

	if statusCode != statusCodeValido {
		openshisftErro := OpenshiftErro{}
		json.Unmarshal(reqBody, &openshisftErro)

		mensagem := fmt.Sprintf("Erro [%v]: %v - %s", statusCode, openshisftErro.Code, openshisftErro.Message)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return statusCode, err
	}

	return statusCode, nil
}
