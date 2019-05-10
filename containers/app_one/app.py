from prometheus_client import start_http_server, Gauge
import random
import time

g = Gauge('demo_gauge', 'Description of demo gauge')


def emit_data(t):
    """Emit fake data"""
    time.sleep(t)
    g.set(t)


if __name__ == '__main__':
    start_http_server(8000)
    while True:
        emit_data(random.random())
