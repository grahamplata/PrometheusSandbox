# Get python 3 image from dockerhub
FROM python:3

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64
RUN chmod +x /usr/local/bin/dumb-init

# Copy repo contents into working directory
COPY . /app
WORKDIR /app

# Install prometheus client library
RUN pip install prometheus_client

# Dumb init
ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]

# Run the application
CMD ["python", "app.py"]
