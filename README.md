# prometheus-uptimerobot-exporter

Prometheus Exporter for UptimeRobot written in Go.

## Usage

```shell
go install github.com/systemli/prometheus-uptimerobot-exporter@latest
export UPTIMEROBOT_API_KEY=your-api-key
$GOPATH/bin/prometheus-uptimerobot-exporter
```

### Commandline Options

```text
  -web.listen-address string
        Address on which to expose metrics and web interface. (default ":13121")
```

## Metrics

```text
# HELP uptimerobot_monitor_up Status of the UptimeRobot monitor
# TYPE uptimerobot_monitor_up gauge
uptimerobot_monitor_up{friendly_name="Google",id="1",status="2",type="1",url="https://www.google.com"} 2
```

## Docker

```shell
docker run -p 13121:13121 -e UPTIMEROBOT_API_KEY=your-api-key systemli/prometheus-uptimerobot-exporter:latest
```

## License

GPLv3
