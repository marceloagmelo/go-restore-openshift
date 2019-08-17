package variaveis

import (
	"log"
	"os"
	"time"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

//DataFormat formato da data
var DataFormat = "02/01/2006 15:04:05"

//DataFormatArquivo formato da data para arquivos
var DataFormatArquivo = "20060102-150405"

//DataHoraAtual a data e hora tual
var DataHoraAtual = time.Now()

//GitUrl
var GitUrl = os.Getenv("GIT_URL")

//GitUrlDownload
var GitUrlDownload = os.Getenv("GIT_URL_DOWNLOAD")

//GitToken
var GitToken = os.Getenv("GIT_TOKEN")

//GitRepositorio
var GitRepositorio = os.Getenv("GIT_REPOSITORIO")

//OpenshiftUrl
var OpenshiftUrl = os.Getenv("OPENSHIFT_URL")

//OpenshiftUsername
var OpenshiftUsername = os.Getenv("OPENSHIFT_USERNAME")

//OpenshiftPassword
var OpenshiftPassword = os.Getenv("OPENSHIFT_PASSWORD")
