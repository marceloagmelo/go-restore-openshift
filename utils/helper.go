package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/marceloagmelo/go-restore-openshift/logger"
)

// CheckErrFatal checar o erro
func CheckErrFatal(err error, msg string) {
	if err != nil {
		log.Printf("CheckErr(): %q\n", err)
		log.Fatalf("%s: %s", msg, err)
	}
}

// CheckErr checar o erro
func CheckErr(err error, msg string) string {
	mensagem := ""

	if err != nil {
		mensagem = fmt.Sprintf("CheckErr(): %s: %s", msg, err)
		log.Printf(mensagem)
	}

	return mensagem
}

//BytesToString converter bytes para string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//InBetween Intervalo de números
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

//IsEmpty verifica se esta vazio
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

//LoadFile carregar arquivo
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//SearchStringInFile
/*func SearchStringInFile(valor string, f *os.File) error {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), valor) {
			mensagem := fmt.Sprintf("%s", "Arquivo não encontrado no gitlab!")
			logger.Erro.Println(mensagem)
			err := errors.New(mensagem)
			return err
		}
	}
	return nil
}*/

//SearchStringInFile
func SearchStringInFile(valor, arquivo string) error {
	f, err := os.Open(arquivo)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao ler o arquivo [%s]: %s", arquivo, err.Error())
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), valor) {
			mensagem := fmt.Sprintf("%s", "Arquivo não encontrado no gitlab!")
			logger.Erro.Println(mensagem)
			err := errors.New(mensagem)
			return err
		}
	}
	return nil
}
