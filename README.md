# CMPE273-Lab2

Routing example using HTTPRouter .


## Usage
Use HttpRouter to route URL's to appropriate destination
### Install

If the package("github.com/julienschmidt/httprouter") is not installed go to this file's folder and do
```
go get 
```
go build
```

Start the  server:

```
go run httpPostRouter.go
```
### Make POST calls from Client to the URL using cURL

```
curl -X POST -d "{\"name\":\"Hanumant\"}" http://localhost:8088/hello
```
Response:
{"greeting":"Hello,Hanumant!"}
```
### Make GET calls From Client to the URL using cURL
```
curl http://localhost:8088/hello/Hanu
```

```
Response:

Hello, Hanu!
```

Comments: Run the Server.go file first using "go run httpPostRouter.go" (without double quotes) and in separate command prompt run the cURL command as specified above
The same can be achieved using postman(application) where you can post raw values to URL and retrieve data
