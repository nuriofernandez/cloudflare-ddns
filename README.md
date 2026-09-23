<p align="center">
  <img src="https://i.imgur.com/KPGtP0d.png">
</p>

# Cloudflare DDNS

A lightweight Dynamic DNS updater that keeps a Cloudflare DNS record in sync with your current public IP address.

## How it works

On startup the binary fetches your public IP via [ipinfo.io](https://ipinfo.io) and updates the target Cloudflare DNS record. If `DDNS_REFRESH_INTERVAL_MINUTES` is set it continues running and repeats the update on that interval.

## Configuration

| Variable | Description |
|---|---|
| `IPINFO_API_TOKEN` | ipinfo.io API token |
| `CLOUDFLARE_API_TOKEN` | Cloudflare API token with DNS edit permissions |
| `CLOUDFLARE_ZONE_ID` | Cloudflare Zone ID |
| `CLOUDFLARE_DNS_RECORD_ID` | ID of the DNS record to update |
| `CLOUDFLARE_RECORD_NAME` | DNS record name (e.g. `home.example.com`) |
| `DDNS_REFRESH_INTERVAL_MINUTES` | _(optional)_ Repeat interval in minutes. Omit to run once and exit. |


## Usage

### Docker (recommended)

```yaml
version: '3.8'

services:
  cloudflare-ddns:
    image: nuriofernandez/cloudflare-ddns:latest
    container_name: cloudflare-ddns
    restart: always
    environment:
      - IPINFO_API_TOKEN=<token-here>
      - CLOUDFLARE_API_TOKEN=<token-here>
      - CLOUDFLARE_ZONE_ID=<zone-id-here>
      - CLOUDFLARE_DNS_RECORD_ID=<record-id-here>
      - CLOUDFLARE_RECORD_NAME=your.expected.dns.record.com
      - DDNS_REFRESH_INTERVAL_MINUTES=5
```

### Kubernetes (homelab)

```yml
apiVersion: v1
kind: Namespace
metadata:
  name: cloudflare
---
apiVersion: v1
kind: Secret
metadata:
  name: cloudflare-ddns-secrets
  namespace: cloudflare
type: Opaque
stringData:
  IPINFO_API_TOKEN: "<token-here>"
  CLOUDFLARE_API_TOKEN: "<token-here>"
  CLOUDFLARE_ZONE_ID: "<zone-id-here>"
  CLOUDFLARE_DNS_RECORD_ID: "<record-id-here>"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cloudflare-ddns
  namespace: cloudflare
  labels:
    app: cloudflare-ddns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cloudflare-ddns
  template:
    metadata:
      labels:
        app: cloudflare-ddns
    spec:
      containers:
      - name: cloudflare-ddns
        image: nuriofernandez/cloudflare-ddns:latest
        imagePullPolicy: Always
        env:
        - name: IPINFO_API_TOKEN
          valueFrom:
            secretKeyRef:
              name: cloudflare-ddns-secrets
              key: IPINFO_API_TOKEN
        - name: CLOUDFLARE_API_TOKEN
          valueFrom:
            secretKeyRef:
              name: cloudflare-ddns-secrets
              key: CLOUDFLARE_API_TOKEN
        - name: CLOUDFLARE_ZONE_ID
          valueFrom:
            secretKeyRef:
              name: cloudflare-ddns-secrets
              key: CLOUDFLARE_ZONE_ID
        - name: CLOUDFLARE_DNS_RECORD_ID
          valueFrom:
            secretKeyRef:
              name: cloudflare-ddns-secrets
              key: CLOUDFLARE_DNS_RECORD_ID
        - name: CLOUDFLARE_RECORD_NAME
          value: "your.expected.dns.record.com"
        - name: DDNS_REFRESH_INTERVAL_MINUTES
          value: "5"
```

### Cron job (binary)

Install the binary with `go install`:

```sh
go install github.com/nuriofernandez/cloudflare-ddns@latest
```
**Note:** `$GOPATH` must be on your `$PATH` in order to work. ```export PATH=${PATH}:`go env GOPATH`/bin```

Then add a crontab entry to run it on your desired schedule:

```
*/5 * * * * cloudflare-ddns
```

Set the required environment variables in your shell profile or pass them inline before the command.
