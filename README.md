# energy-consumption-instrumentation

## Goal

Instrumentation of server code to relate energy consumption to server activity.

The activity is server deal with certain afford activity, the energy consumption is reprsative by the CPU cycle number, the IO from disk, the IO from network.

write a gemini protocl server both python and go.

## Technical details

### Gemini Protocl

<https://gemini.circumlunar.space/docs/specification.gmi>

<https://gemini.circumlunar.space/software/>

in Golang

<https://github.com/makeworld-the-better-one/go-gemini>

in Python

<https://pypi.org/project/gemeaux/>

### How to measure data transfers

Use client that can **control the size** of the messages.

To get the actual size of the data being transferred for every message send by the client or server:

* observe by a program like **tcpdump**, or a library like **libpcap**

* **calculate based on the protocol specification** (for example, every ActivityPub message has a fixed format, so you can count the size in bytes. This is sent using https, so the JSON data is wrapped with an http header and encrypted, then this is sent over tcp/ip. The tcp/ip headers are of a fixed size, depending on whether the protocol is IPv4 or IPv6. There is also the overhead for authentication: to establish a session between the client and the server, a series of packets are exchanged, depending on the protocol. For both Gemini and ActivityPub over https, the first handshake will be the TLS (Transport Layer Security, which does authentication and encryption). Then for ActivityPub there will be an additional authentication.)

### Stress tests

Control messages size and number per second. Check CPU time, memory utilisation and I/O.

### Data support

file system support. database support is a extension.

### TLS

use reverse proxy (e.g. nginx)

## Dissertation Keywords

HLS

<https://www.dacast.com/blog/video-streaming-protocol/>

<https://www.sandvine.com/global-internet-phenomena-report-2019>

## Develop plan

### Requirement

* python server and both python and go sever
* The server side measures the cpu clock, network usage, and memory usage when providing this protocol service in multiple threads, provided that the protocol functionality is implemented.
* The client side enables the definition of packet sizes and bumber of data sent per second
* Before developing the function, first develop Unit testing, Integration testing or stress testing(for measuring resource consumption under different pressures)
* The data used in the project should first be provided by the documentation and then transferred to the database provide

### Schedule

#### day1

Read opensource client code, implement a python language client to communicate with a opensource server.

#### day2

Implement the client that read a given speed, size and content from file, then send them.

#### day3

Find methods of measuring and displaying resource consumption in both python and go. Find how to measure the resource consumption in multi-threaded programs over a fixed period of time based on the threading models of both languages. If it can't find why it is difficult.

#### day4

Implement go server-side with reference to opensource server-side code and protocol specification.

#### day5

Implement python server-side with reference to opensource server-side code and protocol specification.

#### weekend1

Complete the first five days work if not already done. Read and find improvement points.

#### day6

Complete resource consumption measurement functionality in both languages.

#### day7

Adapt server-side and resource consumption measurement for both languages in multi-threaded.

#### day8

Design formal experimental data and run examples, get some data.

#### day9

Migrate to database provide data and read code

#### day10

Fix the code to make it more elegant and add more comments

#### weekend2

Complete the first five days if not already done. Read and look for points I can improve
