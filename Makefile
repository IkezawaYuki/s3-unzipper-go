STACK_NAME := stack-unzipper-lambda
TEMPLATE_FILE := template.yml
SAM_FILE := sam.yml

build:
	GOARCH=amd64 GOOS=linux go build -o artifact/unzipper
.PHONY: build

deploy: build
	sam package \
		--template-file $(TEMPLATE_FILE) \
		--s3-bucket $(STACK_BUCKET) \
		--output-template-file $(SAM_FILE)
	