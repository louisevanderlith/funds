# funds
Mango API: Funds

## Run with Docker
*$ go build
*$ docker build -t avosa/funds:dev .
*$ docker rm fundsDEV
*$ docker run -d --network host --name fundsDEV avosa/funds:dev 
*$ docker logs fundsDEV