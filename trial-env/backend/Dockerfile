FROM golang:1.23-alpine
# --build-arg USERNAME=`id -un`
ARG USERNAME="default"
# --build-arg USER_ID=`id -u`
ARG USER_ID=1001

ARG USER_HOME=/home/${USERNAME}
ARG PROJECT_BASE=${USER_HOME}/projects/ems
ARG PROJECT_ROOT=${PROJECT_BASE}/ss
ARG SCRIPTS=${PROJECT_BASE}/scripts
#COPY ss/go.mod ${PROJECT_ROOT}/go.mod
######
# Create user, working directory
######
#RUN addgroup -g ${USER_ID} ${USERNAME} \
RUN adduser -D -s /bin/sh -u ${USER_ID} ${USERNAME}
#RUN go install github.com/cespare/reflex@latest

# RUN --mount=type=cache,target=/go/pkg/mod \
#     --mount=type=cache,target=${USER_HOME}/.cache/go-build \
#     --mount=type=bind,source=ss,target=${PROJECT_ROOT},readwrite \
#     go mod tidy
    # go build ${PROJECT_ROOT}/cmd/manage/main.go
    # go mod download -x
# WORKDIR ${PROJECT_ROOT}
RUN chown -R ${USERNAME}:${USERNAME} /go
#RUN go mod download
#####
# login user
#####
USER ${USERNAME}
