# graphite

wrapper to pull graphite image and configure udp/tcp listener port.

## how to run it

1. docker build . -t graphite-statsd:local
2. docker run -d\
 --name graphite\
 --restart=always\
 -p 80:80\
 -p 2003-2004:2003-2004\
 -p 2023-2024:2023-2024\
 -p 8130:8130/udp\
 -p 8126:8126\
 graphite-statsd:local