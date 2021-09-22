
url: pulumi.up 
	pulumi stack output bucket_endpoint

pulumi.up:
	pulumi up

init: init.python init.dotenv example.Pulumi.dev.yaml

init.python:
	@python3 -m venv venv && \
		. venv/bin/activate && \
		pip3 install -r requirements.txt

init.dotenv:
	@test ! -f .envrc && \
		echo 'source venv/bin/activate' >> .envrc && \
		direnv allow . && \
		cd . || \
		exit 0
examples=example*
.PHONY: $(examples)
$(examples):
	@export EXAMPLE_FILE=$(subst example.,,$@) && \
		test ! -f $${EXAMPLE_FILE} && \
		cp $@ $${EXAMPLE_FILE} && \
		echo You have to change the value in $${EXAMPLE_FILE} || \
		exit 0

.PHONY: url pulumi.% init init.% $(examples) 

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
