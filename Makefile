.PHONY: docker_run
CMD = docker run --rm -p 5432:5432  --name postgres  -v ./data:/var/lib/postgresql/data  -e POSTGRES_DB=larstest  -e POSTGRES_USER=isinov  -e POSTGRES_PASSWORD=qweasd postgres
docker_run:
		${CMD}