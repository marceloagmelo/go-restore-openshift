package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/marceloagmelo/go-restore-openshift/api"
	"github.com/marceloagmelo/go-restore-openshift/logger"
)

//Tags dados
type Tags []struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Target  string `json:"target"`
	Commit  struct {
		ID             string    `json:"id"`
		ShortID        string    `json:"short_id"`
		Title          string    `json:"title"`
		CreatedAt      time.Time `json:"created_at"`
		ParentIds      []string  `json:"parent_ids"`
		Message        string    `json:"message"`
		AuthorName     string    `json:"author_name"`
		AuthorEmail    string    `json:"author_email"`
		AuthoredDate   time.Time `json:"authored_date"`
		CommitterName  string    `json:"committer_name"`
		CommitterEmail string    `json:"committer_email"`
		CommittedDate  time.Time `json:"committed_date"`
	} `json:"commit"`
	Release interface{} `json:"release"`
}

// OrdenaTagsPorNameDesc
type OrdenaTagsPorNameDesc Tags

//Len
func (x OrdenaTagsPorNameDesc) Len() int {
	return len(x)
}

//Less
func (x OrdenaTagsPorNameDesc) Less(i, j int) bool {
	return x[i].Name > x[j].Name
}

//Swap
func (x OrdenaTagsPorNameDesc) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//GetTags
func GetTags() (Tags, int, error) {
	tags := Tags{}
	// Recuperar recursos
	interf, statusCode, err := api.ListarTags()
	if err != nil {
		mensagemErro := fmt.Sprintf("%s", err)
		logger.Erro.Println(mensagemErro)
		return tags, statusCode, err
	}
	if statusCode == http.StatusOK {
		reqBody, err := json.Marshal(interf)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao recuperar a lista de arquivos", err)
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return tags, statusCode, err
		}

		err = json.Unmarshal(reqBody, &tags)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao converter o conteúdo para  o JSON de namespaces", err)
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return tags, statusCode, err
		}
	}

	sort.Sort(OrdenaTagsPorNameDesc(tags))

	return tags, statusCode, nil

}
