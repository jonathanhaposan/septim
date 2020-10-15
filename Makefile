build-backend:
	cd ./septim-backend && \
	go build -o backend-app && \
	./backend-app