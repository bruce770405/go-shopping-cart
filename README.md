# go-shopping-cart

## services
* user-service
  ```
    port 8808
  ```
* shopping-cart-service
   ```
    port 8809
  ```

## technology
* go 1.13
* gin
* mongodb in aws
* swagger
* docker

## install
```
  move to service folder
  $ docker build -t {service-name}:{version} .
```

## run
```
  docker run -it {service-name} -p {service-port}:{bind-port} --name="{service-name}"
```
