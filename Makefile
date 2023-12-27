# Makefile

.PHONY: run clean

run:
	CompileDaemon --build="GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o ./cityvibe ./" --command="./cityvibe"

clean:
	rm -f cityvibe
