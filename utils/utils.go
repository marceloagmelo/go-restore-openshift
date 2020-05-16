package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/marceloagmelo/go-restore-openshift/logger"
	"github.com/marceloagmelo/go-restore-openshift/openshift/resource"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

//GetToken recuperar Token do usuário.
func GetToken(url string, username string, password string) (string, error) {
	resposta := ""
	endpoint := variaveis.OpenshiftUrl + "/oauth/authorize?client_id=openshift-challenging-client&response_type=token"

	cmdCurl := "curl -s -u " + username + ":" + password + " -kI '" + endpoint + "' | grep -oP 'access_token=\\K[^&]*'"

	resposta, err := ExecCmd(cmdCurl)

	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o CURL", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return resposta, err
	}
	return resposta, nil
}

//ExecCmd execuctar comando no OS.
func ExecCmd(strCurl string) (string, error) {
	resposta := ""
	cmd := exec.Command("/bin/bash", "-c", strCurl)

	out, err := cmd.CombinedOutput()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o comando no OS", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return resposta, err
	}
	resposta = string(out)

	return resposta, nil
}

//FindSlices procurar palavra no slice
func FindSlices(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

//CriarRecurso
func CriarRecurso(namespace string, arquivo string, recurso string) error {
	var openshift resource.Openshift
	openshift.Namespace = namespace
	openshift.Metodo = getMetodoRecurso(recurso)
	f, err := os.Open(arquivo)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao ler o arquivo [%s]: %s", arquivo, err.Error())
		logger.Erro.Println(mensagem)
		return err
	}
	defer f.Close()
	openshift.Body = f

	/*err = SearchStringInFile("File Not Found", f)
	if err != nil {
		return http.StatusNotFound, err
	}*/

	_, _, err = resource.Executar(openshift)
	if err != nil {
		return err
	}

	return nil
}

//getMetodoRecurso
func getMetodoRecurso(recurso string) string {

	switch recurso {
	case "service":
		return "CriarService"
	case "deploymentconfig":
		return "CriarDeploymentConfig"
	case "secret":
		return "CriarSecret"
	case "configmap":
		return "CriarConfigMap"
	case "pvc":
		return "CriarPVC"
	case "rolebinding":
		return "CriarRoleBinding"
	case "clusterrolebinding":
		return "CriarClusterRoleBinding"
	case "role":
		return "CriarRole"
	case "clusterrole":
		return "CriarClusterRole"
	case "route":
		return "CriarRoute"
	case "statefulset":
		return "CriarStateFulSet"
	case "buildconfig":
		return "CriarBuildConfig"
	case "serviceaccount":
		return "CriarServiceAccount"
	case "replicaset":
		return "CriarReplicaSet"
	}
	return ""
}

// DownloadArquivo recuperar o arquivo do git
func DownloadArquivo(tagName, openshiftProjeto, recurso, nomeArquivo string) (string, error) {
	endpoint := variaveis.GitlabApiURL + variaveis.GitlabApiProjetos + "/" + variaveis.GitlabProjectID + "/repository/files/" + openshiftProjeto + "%2F" + recurso + "%2F" + nomeArquivo + "/raw?ref=" + tagName
	arquivoSalvo := "/tmp/" + nomeArquivo
	cmd := "curl -s -k --request GET --header 'Private-Token: " + variaveis.GitlabToken + "' -o '" + arquivoSalvo + "' '" + endpoint + "'"

	resposta, err := ExecCmd(cmd)
	if err != nil {
		fmt.Println("[DownloadArquivo] Erro ao executar o CURL.")
		mensagem := fmt.Sprintf("%s [%s]: %s", "Erro ao fazer o download do arquivo", nomeArquivo, err)
		logger.Erro.Println(mensagem)
		err = errors.New(mensagem)
		return resposta, nil
	}

	resposta = arquivoSalvo

	return resposta, nil
}
