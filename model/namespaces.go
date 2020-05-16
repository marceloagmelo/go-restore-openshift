package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/marceloagmelo/go-restore-openshift/logger"
	"github.com/marceloagmelo/go-restore-openshift/openshift/resource"
)

//Namespaces dados
type Namespaces struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Router string `json:"router"`
			} `json:"labels"`
			Annotations struct {
				OpenshiftIoDisplayName             string `json:"openshift.io/display-name"`
				OpenshiftIoNodeSelector            string `json:"openshift.io/node-selector"`
				OpenshiftIoSaSccMcs                string `json:"openshift.io/sa.scc.mcs"`
				OpenshiftIoSaSccSupplementalGroups string `json:"openshift.io/sa.scc.supplemental-groups"`
				OpenshiftIoSaSccUIDRange           string `json:"openshift.io/sa.scc.uid-range"`
			} `json:"annotations"`
		} `json:"metadata"`
		Spec struct {
			Finalizers []string `json:"finalizers"`
		} `json:"spec"`
		Status struct {
			Phase string `json:"phase"`
		} `json:"status"`
	} `json:"items"`
}

//GetNamespaces
func GetNamespaces() (Namespaces, int, error) {
	recursos := Namespaces{}
	// Recuperar recursos
	interf, statusCode, err := resource.ListarNamespaces()
	if err != nil {
		mensagemErro := fmt.Sprintf("%s", err)
		logger.Erro.Println(mensagemErro)
		return recursos, statusCode, err
	}
	if statusCode == http.StatusOK {
		reqBody, err := json.Marshal(interf)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao recuperar a lista de namespaces", err)
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return recursos, statusCode, err
		}

		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao converter o conteúdo para  o JSON de namespaces", err)
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return recursos, statusCode, err
		}
		if len(recursos.Items) <= 0 {
			mensagemErro := fmt.Sprintln("Não existem namespaces")
			logger.Erro.Println(mensagemErro)
			err := errors.New(mensagemErro)
			return recursos, statusCode, err
		}
	}
	return recursos, statusCode, nil
}
