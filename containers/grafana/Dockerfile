# Pull latest image from dockerhub
FROM grafana/grafana:latest
# Add Configurations to the image
ADD ./provisioning /etc/grafana/provisioning
ADD ./config.ini /etc/grafana/config.ini
ADD ./dashboards /var/lib/grafana/dashboards
