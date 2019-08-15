package controller

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	gitutils "github.com/marceloagmelo/go-git-cli/utils"
	openshiftutils "github.com/marceloagmelo/go-openshift-cli/utils"
	"github.com/marceloagmelo/go-restore-openshift/utils"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

var recusosValidos = make([]string, 15)

func init() {
	recusosValidos[0] = "namespace"
	recusosValidos[1] = "template"
	recusosValidos[2] = "rolebinding"
	recusosValidos[3] = "service"
	recusosValidos[4] = "serviceaccount"
	recusosValidos[5] = "route"
	recusosValidos[6] = "secret"
	recusosValidos[7] = "pvc"
	recusosValidos[8] = "configmap"
	recusosValidos[9] = "imagestream"
	recusosValidos[10] = "buildconfig"
	recusosValidos[11] = "deploymentconfig"
	recusosValidos[12] = "statefulset"
	recusosValidos[13] = "daemonset"
	recusosValidos[14] = "replicaSet"
}

func criarRecurso(token string, url string, projeto string, conteudoJSON string, recurso string) (resultado int, erro string) {
	switch recurso {
	case "service":
		resultado, erro = openshiftutils.CriarService(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "deploymentconfig":
		resultado, erro = openshiftutils.CriarDeploymentConfig(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "secret":
		resultado, erro = openshiftutils.CriarSecret(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "configmap":
		resultado, erro = openshiftutils.CriarConfigMap(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "pvc":
		resultado, erro = openshiftutils.CriarPvc(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "rolebinding":
		resultado, erro = openshiftutils.CriarRoleBinding(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "route":
		resultado, erro = openshiftutils.CriarRoute(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "statefulset":
		resultado, erro = openshiftutils.CriarStateFulSet(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "buildconfig":
		resultado, erro = openshiftutils.CriarBuildConfig(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "serviceaccount":
		resultado, erro = openshiftutils.CriarServiceAccount(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	case "replicaset":
		resultado, erro = openshiftutils.CriarReplicaSet(token, variaveis.OpenshiftUrl, projeto, conteudoJSON)
	}
	return resultado, erro
}

//Restore executa o restore do recurso no openshift
func Restore(c echo.Context) error {
	branch := c.FormValue("branch")
	projeto := c.FormValue("projeto")
	recurso := c.FormValue("recurso")
	nomeArquivo := c.FormValue("nomeArquivo")

	mensagem := "Restore executado com sucesso!"
	htmlResposta := "restore-sucesso.html"

	resultado := 0
	token := ""
	arquivoSalvo := ""

	// Validar campos de entrada
	resultado, retorno := validarDadosEntrada(branch, projeto, recurso, nomeArquivo)
	if resultado > 0 {
		mensagem = retorno
		htmlResposta = "restore-erro.html"
	}

	if resultado == 0 {
		resultado, arquivoSalvo = gitutils.DownloadArquivo(variaveis.GitUrl, variaveis.GitToken, variaveis.GitIdProjeto, branch, projeto, recurso, nomeArquivo)
		if resultado > 0 {
			mensagem = "Erro ao executar o get do arquivo no git!"
			htmlResposta = "restore-erro.html"
		}
	}

	/*
		resultado, conteudoJSON := gitutils.GetArquivo(variaveis.GitUrl, variaveis.GitToken, variaveis.GitIdProjeto, branch, projeto, recurso, nomeArquivo)
		if resultado > 0 {
			mensagem = "Erro ao executar o get do arquivo no git!"
			htmlResposta = "restore-erro.html"
		}
	*/

	if resultado == 0 {
		resultado, token = openshiftutils.GetToken(variaveis.OpenshiftUrl, variaveis.OpenshiftUsername, variaveis.OpenshiftPassword)
		if resultado > 0 {
			mensagem = "Erro ao tentar recuperar o token no openshift!"
			htmlResposta = "restore-erro.html"
		}
	}

	if resultado == 0 {
		resultado, erro := criarRecurso(token, variaveis.OpenshiftUrl, projeto, arquivoSalvo, recurso)
		if resultado > 0 {
			mensagem = "Erro ao executar a criação do recurso no openshift: " + erro
			htmlResposta = "restore-erro.html"
		}
	}

	data := map[string]interface{}{
		"titulo":      "Restore de Recursos no Openshift",
		"mensagem":    mensagem,
		"branch":      branch,
		"projeto":     projeto,
		"recurso":     recurso,
		"nomeArquivo": nomeArquivo,
	}
	return c.Render(http.StatusOK, htmlResposta, data)
}

func validarDadosEntrada(branch string, projeto string, recurso string, nomeArquivo string) (retorno int, mensagem string) {
	retorno = 0

	// Validar branch
	if len(strings.TrimSpace(branch)) == 0 {
		mensagem = "Branch não pode ser vazia!"
		retorno = 1
	}
	// Validar projeto
	if len(strings.TrimSpace(projeto)) == 0 {
		mensagem = "Projeto não pode ser vazio!"
		retorno = 1
	}
	// Validar recurso
	if len(strings.TrimSpace(recurso)) == 0 {
		mensagem = "Recurso não pode ser vazio"
	} else {
		if utils.FindSlices(recusosValidos, recurso) < 0 {
			mensagem = "Recurso inválido!"
			retorno = 1
		}
	}
	// Validar nome do arquivo
	if len(strings.TrimSpace(nomeArquivo)) == 0 {
		mensagem = "Nome do arquivo não pode ser vazio!"
		retorno = 1
	}

	return retorno, mensagem
}
