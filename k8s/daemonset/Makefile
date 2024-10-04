buildDocker:
	docker build -t 192.168.2.200:5000/daemonsetdemo:latest .
push:
	 docker push 192.168.2.200:5000/daemonsetdemo:latest
apply:
	kubectl apply -f daemonset.yaml
delete:
	kubectl delete -f deployment
	docker rmi 192.168.2.200:5000/daemonsetdemo:latest
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -buildvcs=false -o daemonsetdemo
dockerRun:
	docker run  --rm -p 12345:10086 --name daemonsetdemo -it 192.168.2.200:5000/daemonsetdemo:latest
start:
	make build
	make buildDocker
	make push
	make apply