FROM golang:1.23-alpine
# --build-arg USERNAME=`id -un`
ARG USERNAME="default"
# --build-arg USER_ID=`id -u`
ARG USER_ID=1001

ARG USER_HOME=/home/${USERNAME}
ARG PROJECT_BASE=${USER_HOME}/projects/ems
ARG TODO_DIR=${PROJECT_BASE}/todo
ARG PROJECT_ROOT=${PROJECT_BASE}/ss
######
# Create user, working directory
######
#RUN addgroup -g ${USER_ID} ${USERNAME} \
RUN adduser -D -s /bin/sh -u ${USER_ID} ${USERNAME} \
    && mkdir -p ${TODO_DIR}
### set build context set ems in compose.yaml
COPY my-first-golang/todo/ ${TODO_DIR}/

WORKDIR ${PROJECT_BASE}/todo
RUN go mod download

RUN go install github.com/cespare/reflex@latest \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#####
# login user
#####
USER ${USERNAME}
WORKDIR ${PROJECT_ROOT}

