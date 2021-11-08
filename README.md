# go-shopping-cart

## technology
* go 1.13
* gin
* mongodb in aws
* redis
* swagger
* docker

## services

|  service name   | desc  | port |
|  ----  | ----  | ----  |
| user-service  | auth user | 8808 |
| shopping-cart-service  | product caches and shopping records | 8809 |

## install
```
move to service folder
$ docker build -t {service-name}:{version} .
```

## run
```
docker run -it {service-name} -p {service-port}:{bind-port} --name="{service-name}"
```

## use kubenates
    https://github.com/bruce770405/shopping-cart-deploy
