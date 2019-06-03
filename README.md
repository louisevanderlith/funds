# funds
Mango API: Funds holds payment information, Transaction and payment processing.

## Run with Docker
```
>$ docker build -t avosa/funds:dev .
>$ docker rm FundsDEV
>$ docker run -d -p 8092:8092 --network mango_net --name FundsDEV avosa/funds:dev
>$ docker logs FundsDEV
```