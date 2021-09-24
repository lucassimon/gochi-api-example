compile:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(PWD)/build/worker_comunicados_sem_ticker .

compress:
	@upx $(PWD)/build/worker_comunicados_sem_ticker

build: compile compress
