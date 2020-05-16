package resource

import (
	"net/http"

	"github.com/marceloagmelo/go-restore-openshift/api"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

// ListarTags listar
func ListarTags() (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.GitlabApiURL + variaveis.GitlabApiProjetos + "/" + variaveis.GitlabProjectID + "/repository/tags"

	interf, statusCode, err := api.ListarRecurso(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
