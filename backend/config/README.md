# 配置文件说明

文件名: `config.yaml`

## 文件内容示例:

```yaml
# cSpell: disable

server: # 服务器设置
  port: 8005 # 服务器端口
  encrypt_salt: "awa" # SH265 加密盐
  request_auth: "qwq" # 接口请求鉴权 (Authorization)

postgresql: # PostgreSQL 数据库设置
  host: "host" # 连接地址
  port: 5432 # 端口
  database: "name" # 数据库名
  user: "name" # 用户名
  password: "pwd" # 密码

openai: # OpenAI 相关设置
  token: "sk-***" # OpenAI API Token
  base_url: "https://awa.com" # OpenAI API 地址
```
