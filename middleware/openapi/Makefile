api:
	docker build -t apiredoc .


run:
	docker run -it --rm -p 8888:8888 \
      --name apiredoc \
      -v ./docs:/usr/share/nginx/html/apifiles \
      -e PORT=8888 \
      apiredoc:latest