#!/bin/sh
sed -i "s/graphite_host/$GRAPHITE_HOST/g" /etc/grafana/provisioning/datasources/graphite.yaml
exec /run.sh
