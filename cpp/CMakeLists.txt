# CMakeList.txt : CMake project for test, include source and define
# project specific logic here.
#
cmake_minimum_required (VERSION 3.16)

project ("main")

# add the cpp-statsd-client library
# https://github.com/vthiery/cpp-statsd-client
set(CMAKE_THREAD_PREFER_PTHREAD ON)                                                                                                                                                                                                           
set(THREADS_PREFER_PTHREAD_FLAG ON)
# pull in pthreads for cpp-statsd-client library
find_package(Threads REQUIRED)

# source files
set(StatsdClient "StatsdClient.hpp" "UDPSender.hpp")
add_executable (main "main.cpp" ${StatsdClient})

# CMake could not find the correct threading libraries without these explicit statements (Windows and Linux)
if(WIN32)
  target_link_libraries(main wsock32 ws2_32)
else()
  target_link_libraries(main PRIVATE Threads::Threads)
endif()
