package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/marceloagmelo/go-restore-openshift/logger"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

//RecursosValidos
type RecursosValidos struct {
	Recursos []struct {
		Nome string `json:"nome"`
	} `json:"recursos"`
}

//GetRecursosValidos
func GetRecursosValidos() (RecursosValidos, error) {
	recursosValidos := RecursosValidos{}

	jsonFile, err := os.Open(variaveis.RecursosFile)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao ler o arquivo de recursos válidos [%s]: %s", variaveis.RecursosFile, err.Error())
		logger.Erro.Println(mensagem)
		return recursosValidos, err
	}
	defer jsonFile.Close()

	// Ler o json como um array de bytes
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao converter o arquivo [%s] para bytes: %s", variaveis.RecursosFile, err.Error())
		logger.Erro.Println(mensagem)
		return recursosValidos, err
	}

	err = json.Unmarshal(byteValue, &recursosValidos)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao converter o arquivo [%s] para o struct de recursos válidos: %s", variaveis.RecursosFile, err.Error())
		logger.Erro.Println(mensagem)
		return recursosValidos, err
	}

	return recursosValidos, nil
}
