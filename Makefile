build:
	go build -o k8s-master-ip-ha ./cmd/main.go
clean:
	rm -f k8s-master-ip-ha

