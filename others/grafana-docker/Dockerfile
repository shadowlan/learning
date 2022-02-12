FROM grafana/grafana
USER grafana
COPY entrypoint.sh /entrypoint.sh
COPY grafana.ini /etc/grafana/grafana.ini
COPY dashboards /var/lib/grafana/dashboards
COPY graphite.yaml /etc/grafana/provisioning/datasources/graphite.yaml
COPY default-provider.yaml /etc/grafana/provisioning/dashboards/default-provider.yaml
RUN grafana-cli plugins install grafana-piechart-panel
ENTRYPOINT [ "/entrypoint.sh" ]
