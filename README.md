# Introduction
Example of several client programs pushing metrics to a local StatsD daemon.

Assumes there is a StatsD server running and listening on UDP 127.0.0.1:8125

All of the following commands shall be run in WSL.

## command
```
cd command
source send.sh
```
Metrics name: command-client.main.forloop.increment

## go/datadog
<https://github.com/datadog/datadog-go>
```
cd go/datadog
go run main.go
```
Metrics name: "DataDog-client.main.forloop.increment"

## go/cactus
<https://github.com/cactus/go-statsd-client>
```
cd go/cactus
go run main.go
```
Metrics name: "cactus-client.main.forloop.increment"

## csharp/justeat
<https://github.com/justeat/JustEat.StatsD/>  
<https://www.nuget.org/packages/JustEat.StatsD/>
```
cd csharp/justeat-core
dotnet run
```
Metrics name: "justeat-client.main.forloop.increment"

## CPP
<https://github.com/justeat/JustEat.StatsD/>  
<https://www.nuget.org/packages/JustEat.StatsD/>

Metrics name: "CPP-client.main.forloop.increment"

### CMake Build and Run in WSL
```
cd cpp
mkdir build
cd build
cmake -S ../ -B .
make
./main
```
