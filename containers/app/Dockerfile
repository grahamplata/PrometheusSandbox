# Get python 3 image from dockerhub
FROM python:3
# Copy repo contents into working directory
COPY . /app
WORKDIR /app
# Install prometheus client library
RUN pip install prometheus_client
# Run the application
CMD ["python", "app.py"]