
#to build and run the application

build:
	cd cmd/server && go build -o opti-collab . && ./opti-collab

run: 
	cd cmd/server && ./opti-collab

container-up:
	docker start python-container
	docker start go-container
	docker start js-container
	docker start gcc-container
	docker start java-container

container-down:
	docker stop python-container
	docker stop go-container
	docker stop js-container
	docker stop gcc-container
	docker stop java-container

#to automate the git push
push:
	git add .
	git commit -m "$(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))"
	git push

%:
	@: