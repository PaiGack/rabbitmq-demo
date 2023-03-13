RabbitMQ Demo
===

```console
# 15672 WEB管理端端口
# 5672 消息接收端口
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management
```

