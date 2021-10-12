# energy-consumption-instrumentation

## Goal

Instrumentation of server code to relate energy consumption to server activity.

The activity is server deal with certain afford activity, the energy consumption is reprsative by the CPU cycle number, the IO from disk, the IO from network.

write a gemini protocl server both python and go.

## Technical details

## Gemini Protocl

<https://gemini.circumlunar.space/docs/specification.gmi>

<https://gemini.circumlunar.space/software/>

in Golang

<https://github.com/makeworld-the-better-one/go-gemini>

in Python

<https://pypi.org/project/gemeaux/>

### How to measure data transfers

Use client that can **control the size** of the messages.

To get the actual size of the data being transferred for every message send bu the client or server:

* observe by a program like **tcpdump**, or a library like **libpcap** 

* **calculate based on the protocol specification** (for example, every ActivityPub message has a fixed format, so you can count the size in bytes. This is sent using https, so the JSON data is wrapped with an http header and encrypted, then this is sent over tcp/ip. The tcp/ip headers are of a fixed size, depending on whether the protocol is IPv4 or IPv6. There is also the overhead for authentication: to establish a session between the client and the server, a series of packets are exchanged, depending on the protocol. For both Gemini and ActivityPub over https, the first handshake will be the TLS (Transport Layer Security, which does authentication and encryption). Then for ActivityPub there will be an additional authentication.)

### stress tests

Control messages size and number per second. Check CPU time, memory utilisation and I/O.

### data support

file system support. database support is a extension.

### TLS

use reverse proxy (e.g. nginx)
