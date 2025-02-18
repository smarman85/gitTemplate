appName := infra-lsp
command := bash
build:
	podman build -t $(appName) .
run:
	podman run -dith $(appName) -v$(shell pwd):/app --name $(appName) localhost/$(appName):latest $(command)
init: build run
attach:
	podman exec -it $(appName) $(command)
delete:
	podman rm -f $(appName)
reset: delete init