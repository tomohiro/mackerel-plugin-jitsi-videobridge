mackerel-plugin-jitsi-videobridge
================================================================================

[![GitHub Releases](https://img.shields.io/github/release/tomohiro/mackerel-plugin-jitsi-videobridge.svg?style=flat-square)](https://github.com/tomohiro/mackerel-plugin-jitsi-videobridge/releases)
[![Build Status](https://img.shields.io/travis/com/tomohiro/mackerel-plugin-jitsi-videobridge.svg?style=flat-square)](https://travis-ci.com/tomohiro/mackerel-plugin-jitsi-videobridge)
[![Coverage Status](https://img.shields.io/coveralls/tomohiro/mackerel-plugin-jitsi-videobridge.svg?style=flat-square)](https://coveralls.io/github/tomohiro/mackerel-plugin-jitsi-videobridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomohiro/mackerel-plugin-jitsi-videobridge?style=flat-square)](https://goreportcard.com/report/github.com/tomohiro/mackerel-plugin-jitsi-videobridge)
[![Apache License](http://img.shields.io/badge/license-Apache-blue.svg?style=flat-square)](https://github.com/tomohiro/mackerel-plugin-jitsi-videobridgeblob/master/LICENSE)


[Jitsi Videobridge][] custom metrics plugin for [mackerel agent][].

[mackerel agent]: https://github.com/mackerelio/mackerel-agent
[Jitsi Videobridge]: https://jitsi.org/jitsi-videobridge/


Description
--------------------------------------------------------------------------------

The plugin to posts Jitsi Videobridge statistics as custom metrics to Mackerel.

![Custom Metrics Example](https://user-images.githubusercontent.com/54254/69406552-a9268580-0d45-11ea-9701-0905b8fefa3e.png)


Supporting Jitsi Videobridge
--------------------------------------------------------------------------------

- Jitsi Videobridge (1124-1 or higher)


Installation
--------------------------------------------------------------------------------

Latest release:

```
$ mkr plugin install mackerel-plugin-jitsi-videobridge
```

Specified release version:

```
$ mkr plugin install mackerel-plugin-jitsi-videobridge@v0.0.1
```


Usage
--------------------------------------------------------------------------------

```
$ mackerel-plugin-jitsi-videobridge --help
Usage of mackerel-plugin-jitsi-videobridge:
  -host string
        Hostname or IP address of Jitsi Videobridge Colibri REST interface (default "127.0.0.1")
  -metric-key-prefix string
        Metric key prefix (default "jitsi-videobridge")
  -metric-label-prefix string
        Metric label prefix (default "JVB")
  -port string
        Port of Jitsi Videobridge Colibri REST interface (default "80")
  -tempfile string
        Temp file name
```

### Example mackerel-agent.conf

```
[plugin.metrics.jitsi-videobridge]
command = "/usr/bin/mackerel-plugin-jitsi-videobridge -host=127.0.0.1 -port=8080
```


Development
--------------------------------------------------------------------------------

### Requirements

- Go 1.13 or higher


### Release by manually

- Install goxz and ghr by `make setup`
- Edit CHANGELOG.md, git commit, git push
- git tag vx.y.z (Semantic Versioning)
- `make dist` to build and ship package as release assets
- `GITHUB_TOKEN=... make release` to publish assets to GitHub

### Contributing

Please check out the [CONTIRBUTING](CONTRIBUTING.md) guideline.


References
--------------------------------------------------------------------------------

- [jitsi-videobridge/rest.md at master · jitsi/jitsi-videobridge](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/rest.md)
- [jitsi-videobridge/tcp.md at master · jitsi/jitsi-videobridge](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/tcp.md)
- [jitsi-videobridge/statistics.md at master · jitsi/jitsi-videobridge](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md)


LICENSE
--------------------------------------------------------------------------------

© 2019 Tomohiro Taira.

This project is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
