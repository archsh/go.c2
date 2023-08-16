###########
all:

clean:
	@echo "Cleaning up ..."
	@rm -f goc2d

serve: goc2d
	@echo "Starting ..."
	@./goc2d --config c2d/goc2d.yml serve
init: goc2d
	@echo "Starting ..."
	@./goc2d --config c2d/goc2d.yml init
goc2d: Makefile *.go c2d/*.go
	@echo "Building $@ ..."
	@CGO_ENABLED=0 go build -o $@ ./c2d
