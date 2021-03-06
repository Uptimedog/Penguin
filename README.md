<p align="center">
    <img src="/assets/gopher1.png" width="230" />
    <h3 align="center">Penguin</h3>
    <p align="center">Daemon for fast and flexible stats aggregation and collection</p>
    <p align="center">
        <a href="https://github.com/Clivern/Penguin/actions"><img src="https://github.com/Clivern/Penguin/workflows/Build/badge.svg"></a>
        <a href="https://github.com/Clivern/Penguin/actions"><img src="https://github.com/Clivern/Penguin/workflows/Release/badge.svg"></a>
        <a href="https://github.com/Clivern/Penguin/releases"><img src="https://img.shields.io/badge/Version-0.0.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Penguin"><img src="https://goreportcard.com/badge/github.com/Clivern/Penguin?v=0.0.1"></a>
        <a href="https://github.com/Clivern/Penguin/blob/main/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>
<br/>

Penguin daemon listens for statistics like counters, gauges, histogram ... etc, sent over HTTP, UDP, TCP or observe log files and send them to a pluggable backend services like graphite or expose them to prometheus. Penguin is inspired by [statsd](https://github.com/statsd/statsd) and this [article](https://stripe.com/blog/canonical-log-lines) but still it is pretty different from [statsd](https://github.com/statsd/statsd).


## Documentation

Download [the latest penguin binary](https://github.com/Clivern/Penguin/releases). Make it executable from everywhere.

```bash
$ curl -sL https://github.com/Clivern/Penguin/releases/download/vx.x.x/penguin_x.x.x_OS.tar.gz | tar xz
```

Create a config file from `config.dist.yml`

```yaml
# Metrics Input
inputs:
    # HTTP endpoint for metrics collection
    http:
        enabled: on
        mode: prod
        port: 8080
        tls:
            status: off
            pemPath: cert/server.pem
            keyPath: cert/server.key
        path: /
        api_key: ""

    # Log files to watch
    log:
        enabled: on
        paths:
            - /app/logs/metrics_1.log
            - /app/logs/metrics_2.log

# Metrics Cache Driver
cache:
    type: memory
    enabled: on

    drivers:
        memory:
            buffer_size: 10

# Metrics Output
output:
    # Output metrics to console
    console:
        enabled: on

    # Expose to prometheus
    prometheus:
        enabled: on
        endpoint: /metrics

    # TODO: Support Graphite
    graphite:
        enabled: off

# Log configs
log:
    # Log level, it can be debug, info, warn, error, panic, fatal
    level: info
    # output can be stdout or abs path to log file /var/logs/penguin.log
    output: stdout
    # Format can be json
    format: json

```

Run Penguin

```bash
$ penguin run -c /absolute/path/to/config.yml
```

Send metrics to log files that penguin observes

```bash
for ((i=1;i<=100000;i++)); echo '{"type":"counter","name":"penguin_orders","help":"the amount of orders.","method":"inc","value":1,"labels":{"type":"shirts"}}' >> /app/logs/metrics_1.log

for ((i=1;i<=100000;i++)); echo '{"type":"counter","name":"penguin_orders","help":"the amount of orders.","method":"inc","value":1,"labels":{"type":"pants"}}' >> /app/logs/metrics_2.log
```

Send metrics to penguin HTTP endpoint

```bash
curl -X POST \
    -d '{"type":"counter","name":"penguin_orders","help":"the amount of orders.","method":"inc","value":1,"labels":{"type":"trousers"}}' \
    http://127.0.0.1:8080
```

Configure prometheus to scrape this URL `http://127.0.0.1:8080/metrics`


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Penguin is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/penguin/releases) for changelogs for each release version of Penguin. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/penguin/issues


## Security Issues

If you discover a security vulnerability within Penguin, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Penguin** is authored and maintained by [@clivern](http://github.com/clivern).
