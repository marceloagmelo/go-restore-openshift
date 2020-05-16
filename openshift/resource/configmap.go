package resource

import (
	"io"
	"net/http"

	"github.com/marceloagmelo/go-restore-openshift/api"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

// CriarConfigMap criar
func CriarConfigMap(namespace string, body io.Reader) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftUrl + variaveis.ApiV1 + "namespaces/" + namespace + "/configmaps"

	interf, statusCode, err := api.CriarRecurso(endpoint, body)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
