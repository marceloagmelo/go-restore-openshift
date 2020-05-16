package variaveis

import (
	"os"
	"time"
)

//DataFormat formato da data
var DataFormat = "02/01/2006 15:04:05"

//DataFormatArquivo formato da data para arquivos
var DataFormatArquivo = "20060102-150405"

//DataHoraAtual a data e hora tual
var DataHoraAtual = time.Now()

//Contexto
var Contexto = os.Getenv("CONTEXTO")

//GitlabApiURL
var GitlabApiURL = os.Getenv("GIT_URL")

//GitlabToken
var GitlabToken = os.Getenv("GITLAB_PRIVATE_KEY")

//GitlabProjectID
var GitlabProjectID = os.Getenv("GITLAB_PROJECT_ID")

//OpenshiftUrl
var OpenshiftUrl = os.Getenv("OPENSHIFT_URL")

//OpenshiftUsername
var OpenshiftUsername = os.Getenv("OPENSHIFT_USERNAME")

//OpenshiftPassword
var OpenshiftPassword = os.Getenv("OPENSHIFT_PASSWORD")

//OpenshiftToken token usuário do openshift
var OpenshiftToken string

//ApiProjeto
var ApiProject = "/apis/project.openshift.io/v1/projects/"

//ApiApps
var ApiApps = "/apis/apps.openshift.io/v1/"

//ApiV1
var ApiV1 = "/api/v1/"

//ApiRoutes
var ApiRoutes = "/apis/route.openshift.io/v1/"

//ApisAppsv1beta1
var ApisAppsv1beta1 = "/apis/apps/v1beta1/"

//ApisImageV1
var ApisImageV1 = "/apis/image.openshift.io/v1/"

//ApisAuthorizationOpenshiftV1
var ApisAuthorizationOpenshiftV1 = "/apis/authorization.openshift.io/v1/"

// /ApisExtensionsV1beta1
var ApisExtensionsV1beta1 = "/apis/extensions/v1beta1/"

//ApiBuilds
var ApiBuilds = "/apis/build.openshift.io/v1/"

//ApiTemplates
var ApiTemplates = "/apis/template.openshift.io/v1/"

//ApiUsers
var ApiUsers = "/apis/user.openshift.io/v1"

//RecursosFile
var RecursosFile = os.Getenv("RECURSOS_FILE")

//GitlabApiProjetos
var GitlabApiProjetos = "/api/v4/projects"
