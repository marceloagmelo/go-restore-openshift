package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marceloagmelo/go-git-cli/utils"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

var apiProjetos = "/api/v4/projects"

//GetProjeto recuperar os dados do projeto
func GetProjeto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idProjeto, _ := strconv.Atoi(vars["id"])
	resultado, projeto := utils.GetProjeto(variaveis.GitUrl, variaveis.GitToken, idProjeto)
	if resultado > 0 {
		fmt.Println("[GetProjeto] Projeto não encontrado")
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
	}
	respondJSON(w, http.StatusOK, projeto)
}

//Projetos lista dos projetos
func Projetos(w http.ResponseWriter, r *http.Request) {
	resultado, projetos := utils.Projetos(variaveis.GitUrl, variaveis.GitToken)
	if resultado > 0 {
		fmt.Println("[Projetos] Projetos não encontrados")
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
	}
	respondJSON(w, http.StatusOK, projetos)
}

//ListaBranch lista dos branches
func ListaBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idProjeto, _ := strconv.Atoi(vars["idProjeto"])
	resultado, branches := utils.Branches(variaveis.GitUrl, variaveis.GitToken, idProjeto)
	if resultado > 0 {
		fmt.Println("[ListaBranch] Branches não encontrados")
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
	}
	respondJSON(w, http.StatusOK, branches)
}

// GetRecursoDeploymentConfig recuperar o arquivo de recurso de deploymentconfig do git
func GetRecursoDeploymentConfig(url string, token string, idProjeto string, branch string, openshiftProjeto string, recurso string, nomeArquivo string) (resultado int, resposta string) {
	endpoint := url + apiProjetos + "/" + idProjeto + "/repository/files/" + openshiftProjeto + "%2F" + recurso + "%2F" + nomeArquivo + "/raw?ref=" + branch
	cmdRetirarCampos := "| grep resourceVersion | sed s/\\\"resourceVersion[^,]*,//g"
	cmd := "curl -s -k --request GET --header 'Private-Token: " + token + "' '" + endpoint + "' " + cmdRetirarCampos

	resultado, resposta = utils.ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[GetArquivo] Erro ao executar o CURL.")
	}
	return resultado, resposta
}

// DownloadRecursoDeploymentConfig download do arquivo de recurso de deploymentconfig do git
func DownloadRecursoDeploymentConfig(url string, token string, idProjeto string, branch string, openshiftProjeto string, recurso string, nomeArquivo string) (resultado int, resposta string) {
	endpoint := url + apiProjetos + "/" + idProjeto + "/repository/files/" + openshiftProjeto + "%2F" + recurso + "%2F" + nomeArquivo + "/raw?ref=" + branch
	arquivoSalvo := "/tmp/" + nomeArquivo
	cmdRetirarCampos := "| sed -i s/\\\"resourceVersion[^,]*,//g"
	cmd := "curl -s -k --request GET --header 'Private-Token: " + token + "' '" + endpoint + "' " + cmdRetirarCampos + " -o '" + arquivoSalvo

	resultado, resposta = utils.ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[DownloadArquivo] Erro ao executar o CURL.")
	}
	resposta = arquivoSalvo
	return resultado, resposta
}
