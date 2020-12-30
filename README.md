<p align="center">
    <img src="/static/logo.png" width="230" />
    <h3 align="center">Penguin</h3>
    <p align="center">Daemon for fast and flexible stats aggregation and collection</p>
    <p align="center">
        <a href="https://github.com/uptimedog/penguin/actions/workflows/build.yml">
            <img src="https://github.com/uptimedog/penguin/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/uptimedog/penguin/actions/workflows/release.yml">
            <img src="https://github.com/uptimedog/penguin/actions/workflows/release.yml/badge.svg">
        </a>
        <a href="https://github.com/uptimedog/penguin/releases">
            <img src="https://img.shields.io/badge/Version-1.0.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/uptimedog/penguin">
            <img src="https://goreportcard.com/badge/github.com/uptimedog/penguin?v=1.0.0">
        </a>
        <a href="https://hub.docker.com/r/clivern/penguin">
            <img src="https://img.shields.io/badge/Docker-Latest-green">
        </a>
        <a href="https://github.com/uptimedog/penguin/blob/main/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>

Penguin daemon listens for statistics like counters, gauges, histogram ... etc, sent over HTTP, UDP, TCP and expose them to prometheus. Penguin is inspired by [statsd](https://github.com/statsd/statsd) and this [article](https://stripe.com/blog/canonical-log-lines) but still it is pretty different from [statsd](https://github.com/statsd/statsd).


## Documentation

Download [the latest penguin binary](https://github.com/uptimedog/penguin/releases). Make it executable from everywhere.

```bash
$ curl -sL https://github.com/uptimedog/penguin/releases/download/vx.x.x/penguin_x.x.x_OS.tar.gz | tar xz
```

Create a config file from `config.dist.yml`

```yaml
# App configs
app:
  # App name
  name: ${PENGUIN_NAME:-penguin}

  # Env mode (dev or prod)
  mode: ${PENGUIN_MODE:-dev}

  # HTTP port
  port: ${PENGUIN_PORT:-8000}

  # Hostname
  hostname: ${PENGUIN_HOSTNAME:-127.0.0.1}

  # TLS configs
  tls:
    status: ${PENGUIN_TLS_STATUS:-off}
    crt_path: ${PENGUIN_TLS_PEMPATH:-cert/server.crt}
    key_path: ${PENGUIN_TLS_KEYPATH:-cert/server.key}

  # Global timeout
  timeout: ${PENGUIN_TIMEOUT:-50}

  # API Key
  api_key: ${PENGUIN_API_KEY:-xxxx-xxxx-xxxx-xxxx}

  # Log configs
  log:
    # Log level, it can be debug, info, warn, error, panic, fatal
    level: ${PENGUIN_LOG_LEVEL:-debug}
    # Output can be stdout or abs path to log file /var/logs/penguin.log
    output: ${PENGUIN_LOG_OUTPUT:-stdout}
    # Format can be json
    format: ${PENGUIN_LOG_FORMAT:-json}
```

Run Penguin

```bash
$ penguin server -c /absolute/path/to/config.yml
```

Send metrics to penguin HTTP endpoint

```bash
curl -X POST \
    -H "X-API-KEY: xxxx-xxxx-xxxx-xxxx" \
    -d '{"type":"counter","name":"penguin_orders","help":"the amount of orders.","method":"inc","value":1,"labels":{"type":"trousers"}}' \
    http://127.0.0.1:8000/_listen -v
```

Configure prometheus to scrape this URL `http://127.0.0.1:8000/metrics`


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Penguin is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/uptimedog/penguin/releases) for changelogs for each release version of Penguin. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/uptimedog/penguin/issues


## Security Issues

If you discover a security vulnerability within Penguin, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, Uptimedog. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Penguin** is authored and maintained by [@Uptimedog](http://github.com/uptimedog).
