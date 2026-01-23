## Программа `blacklist`: техническое описание

### Общее назначение

Программа `blacklist` предназначена для обработки списков сетевых адресов (в формате CIDR) и генерации конфигурационных файлов для анонсирования маршрутов посредством протокола BGP.

### Способы конфигурирования

Программа поддерживает три механизма задания параметров конфигурации (в порядке возрастания приоритета):

1. Файл конфигурации (наименьший приоритет).  
2. Переменные окружения (средний приоритет).  
3. Параметры командной строки (наивысший приоритет).

### 1. Параметры через переменные окружения (ENV)

Допустимые переменные окружения:

```bash
export CONFIG_PATH="/path/config_file.yaml"
export OUTPUT_DIR="/etc/bird"
export MD5_FILE="md5sums.json"
export SERVICE_NAME_RELOADS="bird.service"
export DAEMON_ENABLED=true
export DAEMON_INTERVAL=60
export HTTP_TIMEOUT=5
```

### 2. Параметры через командную строку

При запуске программы допускаются следующие флаги:

```bash
./blacklist \
  -config /etc/blacklist.yaml \
  -daemon \
  -interval 120 \
  -http-timeout 3 \
  -output /etc/bird \
  -md5 /etc/md5sum.json \
  -service bird.service
```

### 3. Параметры через конфигурационный файл (YAML)

Использование конфигурационного файла **обязательно**. Список источников данных задаётся исключительно через этот файл.

Пример конфигурационного файла (`blacklist.yaml`):

```yaml
sources:
  - name: ipv4.txt
    type: url
    path: https://example.com/ipv4.txt
  - name: local_list
    type: file
    path: /example/local-ipv4.txt

md5_file: md5sums.json
output_dir: /etc/bird

daemon:
  enabled: false
  interval_seconds: 300

http:
  timeout_seconds: 10
  insecure_skip_verify: true
service:
  name: bird
```

**Описание полей конфигурационного файла:**

- `sources` — список источников данных:
  - `name` — логическое имя источника;
  - `type` — тип источника (`url` или `file`);
  - `path` — путь или URL к источнику данных.
- `md5_file` — путь к файлу с контрольными суммами MD5.
- `output_dir` — директория для сохранения выходных файлов.
- `daemon` — настройки режима демона:
  - `enabled` — флаг включения демона (`true`/`false`);
  - `interval_seconds` — интервал между запусками (в секундах).
- `http` — настройки HTTP-запросов:
  - `timeout_seconds` — тайм-аут запроса (в секундах).
  - `insecure_skip_verify` - игнорировать проверку ssl сертификата
- `service` — настройки сервиса BGP:
  - `name` — имя сервиса для перезапуска.

### Формат входных данных

Входные списки должны содержать записи в формате CIDR, каждая на отдельной строке. Пример:

```
1.1.1.1/32
192.168.0.0/24
```

### Логирование

Программа выводит сообщения в стандартный поток ошибок (stderr). Примеры сообщений:

```bash
[blacklist] Загрузка: remote_list
[blacklist] Ошибка: context deadline exceeded (Client.Timeout exceeded)
```

**Формат сообщений:**

- `[blacklist]` — префикс, указывающий на источник сообщения.  
- `Загрузка: <источник>` — информационное сообщение о начале обработки источника.  
- `Ошибка: <описание>` — сообщение об ошибке с детализацией причины.