## Use env

```bash
export DAEMON_ENABLED=true
export DAEMON_INTERVAL=60
export HTTP_TIMEOUT=5

./blacklist

```



## Use flags

```bash
./blacklist \
  -config /etc/blacklist.yaml \
  -daemon \
  -interval 120 \
  -http-timeout 3 \
  -output /var/lib/routes
```



## Logs

```bash
[blacklist] Загрузка: remote_list
[blacklist] Ошибка: context deadline exceeded (Client.Timeout exceeded)
```