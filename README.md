pinbadissector
--------------

This utility reads [pinba](https://github.com/badoo/pinba2) packets from network or from pcap file and prints them in json format.

Usage
-----

How to collect pcap file?

```bash
sudo tcpdump -A -n "udp and host 10.20.210.161 and port 30002" -w pinba.pcap
```

Use [jq](https://stedolan.github.io/jq/) utility to simplify parsing.

```bash
$ ./pinbadissector -pcapfile pinba.pcap | jq .server_name | sort | uniq -c | sort -n -r
11243 "co-front0"
11188 "co-front0-shard0"
4106 "pe-front0"
4076 "pe-front0-shard0"
 754 "chappy_us-front0"
 710 "lumen_us-front0-shard0"
 710 "lumen_us-front0"
 225 "chappy_us-front0-shard3"
 219 "chappy_us-front0-shard0"
 187 "chappy_us-front0-shard1"
 126 "chappy_us-front0-shard2"
```