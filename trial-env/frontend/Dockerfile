FROM node:lts-alpine

ARG USERNAME="default"
ARG USER_ID=1001

ARG USER_HOME=/home/${USERNAME}
ARG PROJECT_ROOT=${USER_HOME}/projects/ems/cs
######
# Create user, working directory
######

RUN adduser -D -s /bin/sh -u ${USER_ID} ${USERNAME}

#####
# login user
#####
USER ${USERNAME}
WORKDIR ${PROJECT_ROOT}

