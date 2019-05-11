# Prometheus SandBox

## Description

A sandbox Docker Environment to experiment with Prometheus and Grafana

## How to get started

Clone the Repo: `git clone https://github.com/grahamplata/PrometheusSandbox.git`

Start the demo cluster: `./bin/start`

Stop the demo cluster: `./bin/stop`

## Accessing Exposed Services

- Grafana `http://localhost:3000`
- Prometheus `http://localhost:9090`
- Node Exporter `http://localhost:9100`
- Dummy Application Metrics `http://localhost:8000`

## Defaults

- Grafana credentials: `admin:admin`

## Todos

- Add default dashboards and charts
- Set internal cron on app_two
