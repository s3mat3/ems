cmake_minimum_required(VERSION 3.24)
set(TARGET_CS ems-cs)
project("${TARGET_CS}"
  VERSION ${EMS_VERSION}
  LANGUAGES NONE)
## create package.json from package.json.in
CONFIGURE_FILE ("${EMS_CS_ROOT}/package.json.in" "${EMS_CS_ROOT}/package.json")
## install node_modules for building client side code
add_custom_target(install-dep
  COMMAND ${EMS_SCRIPTS_DIRECTORY}/install-node.sh
  WORKING_DIRECTORY ${EMS_CS_ROOT}
  COMMENT "Installing js dependency"
  )
## build client side code (make ems-cs)
add_custom_target(${TARGET_CS}
  COMMAND npm run build
  WORKING_DIRECTORY ${EMS_CS_ROOT}
  DEPENDS "install-dep"
  COMMENT "Building ems client side"
  )
## (make cs-move)
add_custom_target(cs-move
  COMMAND mv ${EMS_CS_ROOT}/dist ${EMS_OUTPUT_ROOT}/cs
  DEPENDS ${TARGET_CS}
  COMMENT "Moving client side code"
  )
