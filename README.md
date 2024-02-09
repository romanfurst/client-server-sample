# client-server-sample

simple app to demonstrate client-server http communication  

depllyment consists of three components deployed into separate namespaces. The calling chaing is following:  

```
service1 -> service2 -> service3 -> service2
```

## native build

```
go build 
```

## docker build 

```
docker build -t client-server-sample:1.0.0 . -f ./Dockerfile
```


## k8s deployment

change current directory to helm/client-server-sample

following command deploy all three components into separate namespaces

```
helm install sample-services  . 
```
