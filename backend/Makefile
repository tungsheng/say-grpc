build:
	GOOS=linux go build -o app
	docker build -t gcr.io/tonka-sh/say .
	rm -f app

push:
	gcloud docker -- push gcr.io/tonka-sh/say
