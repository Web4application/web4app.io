cmake_minimum_required(VERSION 3.15)
project(nodenv-realpath-lib VERSION 1.0.0 LANGUAGES C)

set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# Honor BUILD_SHARED_LIBS (default: shared if ON, static if OFF)
add_library(nodenv_realpath src/realpath.c)

# Namespaced alias for consumers
add_library(nodenv::realpath ALIAS nodenv_realpath)

set_target_properties(nodenv_realpath PROPERTIES
    PREFIX ""                      # no 'lib' prefix
    OUTPUT_NAME "nodenv-realpath"  # final artifact name
)

target_compile_options(nodenv_realpath PRIVATE -Wall -O2)

target_include_directories(nodenv_realpath PUBLIC
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
    $<INSTALL_INTERFACE:include>
)

# Install library into libexec (static or shared)
install(TARGETS nodenv_realpath
    EXPORT nodenv-realpath-libTargets
    LIBRARY DESTINATION libexec    # .so / .dylib
    RUNTIME DESTINATION libexec    # Windows .dll
    ARCHIVE DESTINATION libexec    # .a
)

# Install headers
install(DIRECTORY include/ DESTINATION include
        FILES_MATCHING PATTERN "*.h")

# --- pkg-config support ---
set(PKGCONFIG_INSTALL_DIR "${CMAKE_INSTALL_PREFIX}/lib/pkgconfig")

configure_file(
    ${CMAKE_CURRENT_SOURCE_DIR}/nodenv-realpath.pc.in
    ${CMAKE_CURRENT_BINARY_DIR}/nodenv-realpath.pc
    @ONLY
)

install(FILES ${CMAKE_CURRENT_BINARY_DIR}/nodenv-realpath.pc
        DESTINATION lib/pkgconfig)

# --- CMake package export ---
include(CMakePackageConfigHelpers)

write_basic_package_version_file(
    "${CMAKE_CURRENT_BINARY_DIR}/nodenv-realpath-libConfigVersion.cmake"
    VERSION ${PROJECT_VERSION}
    COMPATIBILITY SameMajorVersion
)

configure_package_config_file(
    ${CMAKE_CURRENT_SOURCE_DIR}/cmake/nodenv-realpath-libConfig.cmake.in
    ${CMAKE_CURRENT_BINARY_DIR}/nodenv-realpath-libConfig.cmake
    INSTALL_DESTINATION lib/cmake/nodenv-realpath-lib
)

install(EXPORT nodenv-realpath-libTargets
    FILE nodenv-realpath-libTargets.cmake
    NAMESPACE nodenv::
    DESTINATION lib/cmake/nodenv-realpath-lib
)

install(FILES
    "${CMAKE_CURRENT_BINARY_DIR}/nodenv-realpath-libConfig.cmake"
    "${CMAKE_CURRENT_BINARY_DIR}/nodenv-realpath-libConfigVersion.cmake"
    DESTINATION lib/cmake/nodenv-realpath-lib
)
