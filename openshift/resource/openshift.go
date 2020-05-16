package resource

import (
	"io"
)

//Resource
type (
	OpenshiftInterf interface{}
	Openshift       struct {
		Metodo    string
		Namespace string
		Body      io.Reader
	}
)

//Executar
func Executar(openshift Openshift) (interface{}, int, error) {
	var resposta interface{}
	var statusCode int
	var err error

	if openshift.Metodo == "CriarTemplate" {
		resposta, statusCode, err = CriarTemplate(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarCronJob" {
		resposta, statusCode, err = CriarCronJob(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarService" {
		resposta, statusCode, err = CriarService(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarDeploymentConfig" {
		resposta, statusCode, err = CriarDeploymentConfig(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarRoleBinding" {
		resposta, statusCode, err = CriarRoleBinding(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarBuildConfig" {
		resposta, statusCode, err = CriarBuildConfig(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarConfigMap" {
		resposta, statusCode, err = CriarConfigMap(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarDaemonSet" {
		resposta, statusCode, err = CriarDaemonSet(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarDeployment" {
		resposta, statusCode, err = CriarDeployment(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarImageStream" {
		resposta, statusCode, err = CriarImageStream(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarLimitRange" {
		resposta, statusCode, err = CriarLimitRange(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarReplicaSet" {
		resposta, statusCode, err = CriarReplicaSet(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarResourceQuota" {
		resposta, statusCode, err = CriarResourceQuota(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarRole" {
		resposta, statusCode, err = CriarRole(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarRoute" {
		resposta, statusCode, err = CriarRoute(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarSecret" {
		resposta, statusCode, err = CriarSecret(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarServiceAccount" {
		resposta, statusCode, err = CriarServiceAccount(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "CriarStateFulSet" {
		resposta, statusCode, err = CriarStateFulSet(openshift.Namespace, openshift.Body)
	} else if openshift.Metodo == "ListarNamespaces" {
		resposta, statusCode, err = ListarNamespaces()
	} else if openshift.Metodo == "CriarNamespace" {
		resposta, statusCode, err = CriarNamespace(openshift.Body)
	}
	return resposta, statusCode, err
}
