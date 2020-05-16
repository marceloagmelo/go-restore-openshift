package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/marceloagmelo/go-restore-openshift/logger"
	"github.com/marceloagmelo/go-restore-openshift/model"
	"github.com/marceloagmelo/go-restore-openshift/utils"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

var view = template.Must(template.ParseGlob("views/*.html"))

//Home primeira página
func Home(w http.ResponseWriter, r *http.Request) {
	mensagemErro := ""
	namespaces := model.Namespaces{}

	cacheData := make(map[string]interface{}, 0)
	cacheData["titulo"] = "Restore de Recursos no Openshift"
	cacheData["mensagemErro"] = mensagemErro
	contexto := variaveis.Contexto
	if utils.IsEmpty(contexto) {
		contexto = "/"
	} else {
		contexto = "/" + contexto + "/"
	}
	cacheData["contexto"] = contexto

	var tag, namespace, recurso string

	tags, _, err := model.GetTags()
	if err != nil {
		mensagemErro = fmt.Sprintf("%s", err)
		cacheData["mensagemErro"] = mensagemErro

		view.ExecuteTemplate(w, "Index", cacheData)
		return
	}
	cacheData["tags"] = tags

	for _, v := range tags {
		tag = v.Name
		break
	}
	cacheData["tag"] = tag

	// Recuperar token
	variaveis.OpenshiftToken, err = utils.GetToken(variaveis.OpenshiftUrl, variaveis.OpenshiftUsername, variaveis.OpenshiftPassword)
	if err != nil {
		mensagemErro = fmt.Sprintf("%s: %s", "Erro ao tentar recuperar o token no openshift", err)
		cacheData["mensagemErro"] = mensagemErro

		view.ExecuteTemplate(w, "Index", cacheData)
		return
	}

	// Recuperar namespaces
	namespaces, _, err = model.GetNamespaces()
	if err != nil {
		mensagemErro = fmt.Sprintf("%s", err)
		cacheData["mensagemErro"] = mensagemErro

		view.ExecuteTemplate(w, "Index", cacheData)
		return
	}
	cacheData["namespaces"] = namespaces.Items

	for _, v := range namespaces.Items {
		namespace = v.Metadata.Name
		break
	}
	cacheData["namespace"] = namespace

	// Recuperar recursos válidos
	recursosValidos, err := model.GetRecursosValidos()
	if err != nil {
		mensagemErro = fmt.Sprintf("%s", err)
		cacheData["mensagemErro"] = mensagemErro

		view.ExecuteTemplate(w, "Index", cacheData)
		return
	}
	cacheData["recursos"] = recursosValidos.Recursos

	for _, v := range recursosValidos.Recursos {
		recurso = v.Nome
		break
	}
	cacheData["recurso"] = recurso

	arquivos, _, err := model.GetArquivosGit(tag, namespace+"/"+recurso)
	if err != nil {
		mensagemErro = fmt.Sprintf("%s", err)
		cacheData["mensagemErro"] = mensagemErro

		view.ExecuteTemplate(w, "Index", cacheData)
		return
	}
	cacheData["arquivos"] = arquivos

	view.ExecuteTemplate(w, "Index", cacheData)
}

//ExecutarRestore de recurso
func ExecutarRestore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		namespaces := model.Namespaces{}

		mensagemErro := ""

		cacheData := make(map[string]interface{}, 0)
		cacheData["titulo"] = "Restore de Recursos no Openshift"
		cacheData["mensagemErro"] = mensagemErro
		contexto := variaveis.Contexto
		if utils.IsEmpty(contexto) {
			contexto = "/"
		} else {
			contexto = "/" + contexto + "/"
		}
		cacheData["contexto"] = contexto

		err := r.ParseForm()
		if err != nil {
			mensagemErro = fmt.Sprintf("%s: %s", "Erro no parse do formulário", err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		tag := r.FormValue("tag")
		namespace := r.FormValue("namespace")
		recurso := r.FormValue("recurso")
		nomeArquivo := r.FormValue("nomeArquivo")

		cacheData["tag"] = tag
		cacheData["namespace"] = namespace
		cacheData["recurso"] = recurso
		cacheData["nomeArquivo"] = nomeArquivo

		tags, _, err := model.GetTags()
		if err != nil {
			mensagemErro = fmt.Sprintf("%s", err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}
		cacheData["tags"] = tags

		variaveis.OpenshiftToken, err = utils.GetToken(variaveis.OpenshiftUrl, variaveis.OpenshiftUsername, variaveis.OpenshiftPassword)
		if err != nil {
			mensagemErro = fmt.Sprintf("%s: %s", "Erro ao tentar recuperar o token no openshift", err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		// Recuperar namespaces
		namespaces, _, err = model.GetNamespaces()
		if err != nil {
			mensagemErro = fmt.Sprintf("%s", err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}
		cacheData["namespaces"] = namespaces.Items

		// Recuperar recursos válidos
		recursosValidos, err := model.GetRecursosValidos()
		if err != nil {
			mensagemErro = fmt.Sprintf("%s", err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}
		cacheData["recursos"] = recursosValidos.Recursos

		// Validar campos de entrada
		ok, mensagem := validarDadosEntrada(tag, namespace, recurso, nomeArquivo)
		if !ok {
			mensagemErro = fmt.Sprintf("%s", mensagem)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		arquivoSalvo, err := utils.DownloadArquivo(tag, namespace, recurso, nomeArquivo)
		if err != nil {
			mensagemErro = fmt.Sprintf("%s", "Erro ao executar o get do arquivo no git!")
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		err = utils.SearchStringInFile("File Not Found", arquivoSalvo)
		if err != nil {
			mensagemErro = fmt.Sprintf("Não foi possível criar o recurso %s no openshift: %s", recurso, err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		//Retirar atributos do arquivo JSON
		err = retirarAtributosInvalidosNoJSON(arquivoSalvo)
		if err != nil {
			mensagemErro = fmt.Sprintf("Não foi possível criar o recurso %s no openshift: %s", recurso, err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		// Criar o recurso
		err = utils.CriarRecurso(namespace, arquivoSalvo, recurso)
		if err != nil {
			mensagemErro = fmt.Sprintf("Não foi possível criar o recurso %s no openshift: %s", recurso, err)
			cacheData["mensagemErro"] = mensagemErro

			view.ExecuteTemplate(w, "Index", cacheData)
			return
		}

		os.RemoveAll(arquivoSalvo)

		mensagem = fmt.Sprintf("Recurso %v criado com sucesso!", recurso)
		logger.Info.Println(mensagem)
		cacheData["mensagem"] = mensagem

		view.ExecuteTemplate(w, "Index", cacheData)
	}
}

//ListarArquivosDoGit listar os arquivos do git
func ListarArquivosDoGit(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("ref")
	path := r.URL.Query().Get("path")

	arquivos, _, err := model.GetArquivosGit(tag, path)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s", err)
		respondError(w, http.StatusInternalServerError, mensagemErro)
		return

	}

	respondJSON(w, http.StatusOK, arquivos)
}

//retirarAtributosInvalidosNoJSON
func retirarAtributosInvalidosNoJSON(arquivo string) error {
	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + arquivo
	_, err := utils.ExecCmd(cmd)
	if err != nil {
		mensagem := fmt.Sprintf("%s %s: %s", "Erro ao retirar atributo [resourceVersion] do arquivo", arquivo, err.Error())
		logger.Erro.Println(mensagem)
		return err
	}
	cmd = "sed -i s/\\\"clusterIP[^,]*,//g " + arquivo
	_, err = utils.ExecCmd(cmd)
	if err != nil {
		mensagem := fmt.Sprintf("%s %s: %s", "Erro ao retirar atributo [clusterIP] do arquivo", arquivo, err.Error())
		logger.Erro.Println(mensagem)
		return err
	}

	return nil
}

func validarDadosEntrada(tag string, namespace string, recurso string, nomeArquivo string) (bool, string) {
	// Validar tag
	if len(strings.TrimSpace(tag)) == 0 {
		mensagem := "Tag não pode ser vazia!"
		return false, mensagem
	}
	// Validar namespace
	if len(strings.TrimSpace(namespace)) == 0 {
		mensagem := "Projeto não pode ser vazio!"
		return false, mensagem
	}
	// Validar recurso
	if len(strings.TrimSpace(recurso)) == 0 {
		mensagem := "Recurso não pode ser vazio"
		return false, mensagem
	}
	// Validar nome do arquivo
	if len(strings.TrimSpace(nomeArquivo)) == 0 {
		mensagem := "Nome do arquivo não pode ser vazio!"
		return false, mensagem
	}

	return true, ""
}
