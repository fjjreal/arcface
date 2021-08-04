
# step

- cp example.arc.json arc.json

- edit arc.json ARC_APPID & ARC_APPKEY, apply from [arc-soft](https://ai.arcsoft.com.cn/product/arcface.html)

- docker build -t "hrface:latest" -f dev.Dockerfile .

- docker-compose up -d

- curl http://127.0.0.1:8080/active

- curl http://127.0.0.1:8080/compare?name=ty

- curl http://127.0.0.1:8080/compare?name=zzy