```
docker run -it --rm --entrypoint /bin/bash docker.zhai.cm/linuxserver/bookstack:latest appkey 
```

执行此命令生成密钥，修改docker-compose文件配置。

```
- APP_KEY=base64:f1p8OZZjzHk6Oa4Nkr5TF1lfQ5jyYznTT70Z909yUFM=
```

```
admin@admin.com password
```

默认的账号密码