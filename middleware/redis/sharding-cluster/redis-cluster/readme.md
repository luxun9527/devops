#执行命令
docker exec -it redis-node4 redis-cli -h 192.168.2.159 -p 7001  -a bingo --cluster create 192.168.2.159:7001 192.168.2.159:7002 192.168.2.159:7003 192.168.2.159:7004 192.168.2.159:7005 192.168.2.159:7006 --cluster-replicas 1
