# gg

go command tool library

```sh
  go install github.com/bobacgo/gg@latest
```

### tree

- gg tree
- gg tree /path/to/dir -l 2

### time t

- gg time                # 输出当前时间、秒级时间戳、毫秒级时间戳
- gg time 2024-06-01     # 输出该日期的起止时间戳
- gg time "2024-06-01 12:00:00" # 输出该时间的时间戳
- gg time 1717219200     # 输出时间戳对应的时间

### uuid

- gg uuid 3
- gg uuid 3 -f -

### http

- gg http http://localhost:8080/api/v1/cfg --debug -e=test
- gg http /api/v1/cfg -b dev=http://localhost:8080
- gg http post http://localhost:8080/api/v1/user -r "{\"name\": \"bobacgo\"}"
- gg http post http://www.imooc.com/search/hotwords -H token=234 -H app=1

### json

- gg json -f data.json                   # 格式化 JSON 文件内容并覆盖原文件
- gg json -f '{"name":"Alice","age":30}' # 格式化 JSON 字符串
- gg json -e data.txt                    # 将文件内容序列化为 JSON 字符串
- gg json -d data.json                   # 将 JSON 字符串反序列化为原始字符串

### cron

- gg cron "0 0 * * *"

```sh
  近5次执行时间（未来）:
  2025-06-26 00:00:00
  2025-06-27 00:00:00
  2025-06-28 00:00:00
  2025-06-29 00:00:00
  2025-06-30 00:00:00
```

### md5

- gg md5 <string|filepath>

### token jwt

- gg token `<your-jwt-token>`

### pwd kv

- gg pwd                # List all passwords
- gg pwd key value      # Add or update a password with key and value
- gg pwd -d key         # Delete the password with the specified key

### base64 b64

- gg base64 encode "hello"
- gg base64 decode "aGVsbG8="
