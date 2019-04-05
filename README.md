# funds
Mango API: Funds

## Run with Docker
* $ docker build -t avosa/funds:dev .
* $ docker rm FundsDEV
* $ docker run -d --network host --name FundsDEV avosa/funds:dev 
* $ docker logs FundsDEV