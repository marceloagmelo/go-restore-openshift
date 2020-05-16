#!/usr/bin/env bash

source setenv.sh

echo "Subindo o ${APP_NAME}..."
docker run -d --name ${APP_NAME}  \
-p 7070:8080 \
-e OPENSHIFT_URL=${OPENSHIFT_URL} \
-e OPENSHIFT_USERNAME=${OPENSHIFT_USERNAME} \
-e OPENSHIFT_PASSWORD=${OPENSHIFT_PASSWORD} \
-e GIT_URL=${GIT_URL} \
-e GITLAB_PRIVATE_KEY=${GITLAB_PRIVATE_KEY} \
-e GITLAB_PROJECT_ID=${GITLAB_PROJECT_ID} \
-e RECURSOS_FILE=${RECURSOS_FILE} \
-e TZ=America/Sao_Paulo \
${DOCKER_REGISTRY}/${APP_NAME}:${DOCKER_TAG}

# Listando os containers
docker ps