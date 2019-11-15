mackerel-plugin-jitsi-videobridge
================================================================================

[Jitsi Videobridge][] custom metrics plugin for [mackerel agent][].

[mackerel agent]: https://github.com/mackerelio/mackerel-agent
[Jitsi Videobridge]: https://jitsi.org/jitsi-videobridge/


Description
--------------------------------------------------------------------------------

The plugin to posts Jitsi Videobridge statistics as custom metrics to Mackerel.


Requirements
--------------------------------------------------------------------------------

- Jitsi Videobridge (1124-1 or higher)


Installation
--------------------------------------------------------------------------------

Latest release:

```
$ mkr plugin install tomohiro/mackerel-plugin-jitsi-videobridge
```

Specified release version:

```
$ mkr plugin install tomohiro/mackerel-plugin-jitsi-videobridge@v0.0.1
```


Usage
--------------------------------------------------------------------------------

```
mackerel-plugin-jitsi-videobridge [-host=<host>] [-port=<port>] [-tempfile=<tempfile>]
```

### Example mackerel-agent.conf

```
[plugin.metrics.jitsi-videobridge]
command = "/usr/bin/mackerel-plugin-jitsi-videobridge -host=127.0.0.1 -port=8080
```


References
--------------------------------------------------------------------------------

- [jitsi-videobridge/rest.md at master · jitsi/jitsi-videobridge](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/rest.md)
- [jitsi-videobridge/tcp.md at master · jitsi/jitsi-videobridge](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/tcp.md)
- [jitsi-videobridge/statistics.md at master · jitsi/jitsi-videobridge](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md)


LICENSE
--------------------------------------------------------------------------------

© 2019 Tomohiro Taira.

This project is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
