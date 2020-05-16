package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/marceloagmelo/go-restore-openshift/api"
	"github.com/marceloagmelo/go-restore-openshift/logger"
)

//Arquivos
type Arquivos []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Mode string `json:"mode"`
}

// OrdenaArquivosPorNameDesc
type OrdenaArquivosPorNameDesc Arquivos

//Len
func (x OrdenaArquivosPorNameDesc) Len() int {
	return len(x)
}

//Less
func (x OrdenaArquivosPorNameDesc) Less(i, j int) bool {
	return x[i].Name < x[j].Name
}

//Swap
func (x OrdenaArquivosPorNameDesc) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//GetArquivosGit
func GetArquivosGit(tag, path string) (Arquivos, int, error) {
	arquivos := Arquivos{}
	qtdePorPagina := "500"
	// Recuperar recursos
	interf, statusCode, err := api.ListarArquivos(tag, path, qtdePorPagina)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s", err)
		logger.Erro.Println(mensagemErro)
		return arquivos, statusCode, err
	}
	if statusCode == http.StatusOK {
		reqBody, err := json.Marshal(interf)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao recuperar a lista de arquivos", err)
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return arquivos, statusCode, err
		}

		err = json.Unmarshal(reqBody, &arquivos)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao converter o conteúdo para  o JSON de namespaces", err)
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return arquivos, statusCode, err
		}
	}

	sort.Sort(OrdenaArquivosPorNameDesc(arquivos))

	return arquivos, statusCode, nil
}
