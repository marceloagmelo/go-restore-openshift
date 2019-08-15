package controller

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo"
	"gitlab.produbanbr.corp/paas-brasil/go-git-cli/utils"
	"gitlab.produbanbr.corp/paas-brasil/go-restore-openshift/variaveis"
)

//Home página inicial
func Home(c echo.Context) error {
	branches := []string{}
	retBranches := []string{""}
	idProjeto, _ := strconv.Atoi(variaveis.GitIdProjeto)
	resultado, listaBranches := utils.Branches(variaveis.GitUrl, variaveis.GitToken, idProjeto)
	if resultado == 0 {
		for i := 0; i < len(listaBranches); i++ {
			branches = append(branches, listaBranches[i].Name)
		}
	}
	// Order o slice de forma descendente
	sort.Slice(branches, func(i, j int) bool {
		return branches[i] > branches[j]
	})

	for indice, branch := range branches {
		retBranches = append(retBranches, branch)
		if indice == 10 {
			break
		}
	}
	data := map[string]interface{}{
		"titulo":   "Restore de Recursos no Openshift",
		"branches": retBranches,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

/*
func Home(c echo.Context) error {
	branch := model.Branch{}
	branches := model.Branches{}
	resultado, listaBranches := utils.Branches(variaveis.GitUrl, variaveis.GitToken, variaveis.GitIdProjeto)
	if resultado == 0 {
		for i := 0; i < len(listaBranches); i++ {
			branch.Commit.ID = listaBranches[i].Commit.ID
			branch.Name = listaBranches[i].Name
			branch.Commit.AuthorName = listaBranches[i].Commit.AuthorName
			branches = append(branches, branch)

			if i == 10 {
				break
			}
		}
	}
	data := map[string]interface{}{
		"titulo":   "Titulo",
		"branches": branches,
	}

	return c.Render(http.StatusOK, "index.html", data)
}
*/
