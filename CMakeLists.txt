#[[
 usage cmake -D CMAKE_BUILD_TYPE=(Debug | Release | '') -DCMAKE_EXPORT_COMPILE_COMMANDS=on -DCMAKE_VERBOSE_MAKEFILE:BOOL=TRUE
 -DMAKE_TEST=1 we will compile test and sample code
]]
cmake_minimum_required(VERSION 3.24)
#[[

]]
file(STRINGS VERSION EMS_VERSION_LINE)
string(REGEX MATCH "([0-9]+\\.?)+" EMS_VERSION ${EMS_VERSION_LINE})
set(TARGET ems)
message(${EMS_VERSION})
project("${TARGET}"
  VERSION ${EMS_VERSION}
  LANGUAGES NONE)

set(EMS_ROOT_NAME ${TARGET})
set(EMS_ROOT_DIRECTORY ${PROJECT_SOURCE_DIR})
set(EMS_BUILD_DIRECTORY ${PROJECT_BINARY_DIR})
set(EMS_SCRIPTS_DIRECTORY ${EMS_ROOT_DIRECTORY}/scripts)
set(EMS_CS_ROOT ${EMS_ROOT_DIRECTORY}/cs)
set(EMS_SS_ROOT ${EMS_ROOT_DIRECTORY}/ss)
set(EMS_STUDY_ROOT ${EMS_ROOT_DIRECTORY}/my-first-golang)
## output directories
set(EMS_OUTPUT_ROOT ${EMS_ROOT_DIRECTORY}/public)
set(EMS_BIN_OUTPUT ${EMS_OUTPUT_ROOT}/site)
set(RUNTIME_OUTPUT_DIRECTORY ${EMS_BIN_OUTPUT})
set(RUNTIME_OUTPUT_DIRECTORY_DEBUG ${EMS_BIN_OUTPUT}/debug)
set(RUNTIME_OUTPUT_DIRECTORY_RELEASE ${EMS_BIN_OUTPUT}/release)
set(EXECUTABLE_OUTPUT_PATH ${EMS_BIN_OUTPUT})
#
option(EMS_API_COMPILE "Server side API compile." OFF)
option(EMS_CS_COMPILE "Client side compile." ON)
option(EMS_API_MAKE_TEST "Test for server side" OFF)
option(EMS_CS_MAKE_TEST "Test for client side" OFF)


# create runtime directory
if(NOT IS_DIRECTORY ${EMS_OUTPUT_ROOT})
  file(MAKE_DIRECTORY ${EMS_OUTPUT_ROOT})
else()
  message("exists")
endif()



# if(EMS_API_MAKE_TEST OR EMS_CS_MAKE_TEST)
#   enable_testing()
# endif()

if(EMS_API_COMPILE)
  add_subdirectory(api)
endif()

if(EMS_CS_COMPILE)
  add_subdirectory(cs)
endif()

add_custom_target(run
  COMMAND ems-api
  WORKING_DIRECTORY ${EMS_BIN_OUTPUT}
  COMMENT "Running test server..."
  )

add_custom_target(compile-commands
  COMMAND ln -s ${EMS_BUILD_DIRECTORY}/compile_commands.json ${EMS_ROOT_DIRECTORY}/compile_commands.json
  )
