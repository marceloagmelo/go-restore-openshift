package controller

import (
	"net/http"
	"sort"

	"github.com/labstack/echo"
	"github.com/marceloagmelo/go-git-cli/utils"
	"github.com/marceloagmelo/go-restore-openshift/variaveis"
)

//Home página inicial
func Home(c echo.Context) error {
	branches := []string{}
	retBranches := []string{""}

	resultado, listaBranches := utils.Branches(variaveis.GitUrl, variaveis.GitToken, variaveis.GitRepositorio)
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
