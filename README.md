# Restore dos Recursos do Openshift

Aplicação que realiza o restore de recursos do **Openshift**, este recursos estão armazenado no repositório do [GITLAB](https://gitlab.com/marceloagmelo/openshift-backup.git) **ID 18811050**, feitos pelo processo de backup, os seguintes recursos são permitidos:

- buildconfig
- configmap
- cronjob
- daemonset
- deployment
- deploymentconfig
- namespace
- imagestream
- replicaSet
- rolebinding
- route
- secret
- service
- serviceaccount
- statefulset
- role
- rolebinding
- template

----

# Instalação

```
go get -v github.com/marceloagmelo/go-restore-openshift
```
```
cd go-restore-openshift
```

## Build da Aplicação

```
./image-build.sh
```

## Iniciar a Aplicação
```
./start.sh
```

## Finalizar a Aplicação
```
./stop.sh
```

# Instalação no Openshift


Importar o [Template](https://github.com/marceloagmelo/go-restore-openshift/blob/master/openshift/template/go-restore-openshift-template.json) no projeto do openshift e preencher as seguintes informações:

```
Application Name: apagar-templates-default-openshift
Openshift URL: https://console.openshift.lab:8443
Openshift Username: 
Openshift Password:
Gitlab URL: https://gitlab.com
Gitlab Private Key:
Gitlab Project ID: 18811050
Caminho do arquivo de recursos: /go/bin/recursos.json
Contexto: go-restore-openshift
```
