run:
	
runfront: 
	cd ./frontend && \
	npm install && \
	npm run build && \
	npm run start && \
	@echo "Running ${PROJECT} frontend.....http://localhost:4200 to view website. Use Ctrl + C TWICE to close."; 

viewfront:
	netstat -a -n -o &&\
	@echo "Showing ports..... To close website, use taskkill -f /pid [PID of website using port 4200]" 

cleanfront:
	rm -rf node_modules

buildback:
	@echo "Building ${PROJECT} backend..... not needed to run go project. "
	go build -o server/bin OverHere/server/main/. 

testback:
	@echo "Running backend tests..... in main. Use go test within server/main to see debug results"; 
	go test OverHere/server/main/.

runback:
	@echo "Running ${PROJECT} backend..... use Postman to http://localhost:8000 to test routing. Ctrl + C TWICE to close."; 
	go run OverHere/server/main/.