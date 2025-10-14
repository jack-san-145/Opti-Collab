
#to build and run the application

build:
	cd cmd/server && go build -o opti-collab .

run: 
	cd cmd/server && ./opti-collab



#to automate the git push
push:
	git add .
	git commit -m "$(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))"
	git push

%:
	@: