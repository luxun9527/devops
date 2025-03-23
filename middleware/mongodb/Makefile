
prev:
	touch ./replset/mongo1/logs/mongod.log
	touch ./replset/mongo2/logs/mongod.log
	touch ./replset/mongo3/logs/mongod.log
	chmod 666 ./replset/mongo1/logs/mongod.log
	chmod 666 ./replset/mongo2/logs/mongod.log
	chmod 666 ./replset/mongo3/logs/mongod.log
	chmod 0400 ./replset/mongo1/conf/rs_keyfile
	chmod 0400 ./replset/mongo2/conf/rs_keyfile
	chmod 0400 ./replset/mongo3/conf/rs_keyfile

up:
	docker-compose -f docker-compose-replset.yaml up -d

down:
	docker-compose -f docker-compose-replset.yaml down
clear:
	docker rm -f mongo1 mongo2 mongo3 setup
	rm -rf ./replset/mongo1/data ./replset/mongo2/data ./replset/mongo3/data
