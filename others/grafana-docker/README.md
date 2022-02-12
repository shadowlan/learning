# Grafana

grafana wrapper to pull grafana/grafana image, config it with graphite data source.

## how to run it

1. docker build . -t grafana-graphite:local
2. docker run -e GRAPHITE_HOST=localhost grafana-graphite:local
