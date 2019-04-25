# Pull latest image from dockerhub
FROM prom/prometheus:latest
# Add Configurations to the image
ADD ./config.yml /etc/prometheus/config.yml
# Pass in the configurations
CMD [ "--config.file=/etc/prometheus/config.yml" ] 
