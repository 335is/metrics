# Introduction
This repo contains several different examples of client programs pushing metrics to a local StatsD daemon.

They assume there is a StatsD server running and listening on UDP 127.0.0.1:8125

All of the following commands have been tested in WSL2.

---

## command

```
cd command
source send.sh
```
#### Metrics Name
**command-client.main.forloop.increment**

---

## go/datadog

### StatsD Client Library
<https://github.com/datadog/datadog-go>

```
cd go/datadog
go run main.go
```
#### Metrics Name
**datadog-client.main.forloop.increment**

---

## go/cactus

### StatsD Client Library
<https://github.com/cactus/go-statsd-client>

```
cd go/cactus
go run main.go
```
#### Metrics Name
**cactus-client.main.forloop.increment**

---

## csharp/justeat-core

### StatsD Client Library
<https://github.com/justeat/JustEat.StatsD/>  
<https://www.nuget.org/packages/JustEat.StatsD/>

```
cd csharp/justeat-core
dotnet run
```
#### Metrics Name
**justeat-core-client.main.forloop.increment**

---

## cpp
### StatsD Client Library
<https://github.com/vthiery/cpp-statsd-client>  

```
cd cpp
mkdir build
cd build
cmake -S ../ -B .
make
./main
```

#### Metrics Name
**cpp-client.main.forloop.increment**

---

## StatsD Local Daemon (Linux)
The common implementation of StatsD is a Node.js server that listens to UDP messages on port 8125.

### Installation
I followed the steps in this article.
<https://www.sentinelone.com/blog/statsd-measure-anything-in-your-system/>

#### Install Node.js
```
sudo apt-get install -y nodejs
```

#### Clone the statsd repo
```
cd /mnt/c/test/statsd
git clone https://github.com/statsd/statsd
cd statsd
```

#### Modify the config file
```
cp exampleConfig.js localConfig.js
nano localConfig.js
# comment out everything at the bottom, add these instead
# this enables verbose output to the console so you can view stats as they accumulate
{
    debug: true
, port: 8125
, backends: [ "./backends/console" ]
}
```
#### Run StatsD
```
node stats.js localConfig.js
```

---

## StatsD Local Daemon (Windows)

### Installation
I followed the steps in this article.
<https://stackoverflow.com/questions/5436606/using-etsys-statsd-in-a-windows-environment>

#### Install Node.js
Download the MSI installer and run it to completion.
https://nodejs.org/en/download/
https://nodejs.org/dist/v14.18.0/node-v14.18.0-x64.msi


#### Install the statsd server
```
md c:\code\node.js\StatsD
cd c:\code\node.js\StatsD
npm install https://github.com/etsy/statsd.git
```

#### Create a config file
```
# configure StatsD with debug output and using console for the backend
cd c:\code\node.js\StatsD\node_modules\statsd
# create a config.json file containing the following text.
{
  debug: true
, backends: [ "./backends/console" ]
}
```

#### Run StatsD
```
node stats.js localConfig.js
```

### Sample Debug Output
```
Flushing stats at  Thu Sep 30 2021 08:45:34 GMT-0700 (Pacific Daylight Time)
{ counters:
   { 'statsd.bad_lines_seen': 0,
     'statsd.packets_received': 4,
     'statsd.metrics_received': 4,
     'command-client.main.forloop.increment': 0,
     'justeat-core-client.main.forloop.increment': 0,
     'datadog-client.main.forloop.increment': 0,
     'cactus-client.main.forloop.increment': 0,
     'cpp-client.main.forloop.increment': 4 },
  timers: {},
  gauges: { 'statsd.timestamp_lag': 0 },
  timer_data: {},
  counter_rates:
   { 'statsd.bad_lines_seen': 0,
     'statsd.packets_received': 0.4,
     'statsd.metrics_received': 0.4,
     'command-client.main.forloop.increment': 0,
     'justeat-core-client.main.forloop.increment': 0,
     'datadog-client.main.forloop.increment': 0,
     'cactus-client.main.forloop.increment': 0,
     'cpp-client.main.forloop.increment': 0.4 },
  sets: {},
  pctThreshold: [ 90 ] }
  ```
