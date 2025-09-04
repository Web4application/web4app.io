cmake_minimum_required(VERSION 3.15)
project(nodenv-realpath-lib C)

# Set position independent code for shared libraries
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# Define the shared library target
add_library(nodenv_realpath SHARED src/realpath.c)

# Set output directory for the shared lib
set_target_properties(nodenv_realpath PROPERTIES
    LIBRARY_OUTPUT_DIRECTORY ${PROJECT_SOURCE_DIR}/libexec
    PREFIX ""  # optional: remove the 'lib' prefix if you want (e.g. nodenv-realpath.dylib not libnodenv_realpath.dylib)
    OUTPUT_NAME "nodenv-realpath"
)

# Optional: add compile definitions or flags if needed
target_compile_options(nodenv_realpath PRIVATE -Wall -O2)
