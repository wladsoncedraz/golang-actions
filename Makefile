build:
	GOOS=linux GOARCH=amd64 go build -o dist/simple-http

send-bin:
	rsync dist/simple-http ubuntu@$(AWS_REMOTE_SERVER_IP):/home/ubuntu

send-server:
	rsync http-server.service ubuntu@$(AWS_REMOTE_SERVER_IP):/home/ubuntu

deploy: build send-bin send-server
	ssh -t ubuntu@$(AWS_REMOTE_SERVER_IP) '\
		sudo mv ~/http-server.service /etc/systemd/system/ \
		&& sudo systemctl enable http-server \
		&& sudo systemctl restart http-server \
	'