# client-server-sample

simple app to demonstrate client-server http communication  

deployment consists of three components deployed into separate namespaces. The calling chaining is following:  

```
service1 -> service2 -> service3 -> service1
```

every service has two endpoints:  

`/server`  
component receive incoming request and return static response


`/server-client`  
component receive incoming request and make outgoing request to another component. The outgoing request endpoint is configurable

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

it results in three independent PODs:

```a
app-namespace-one      service-one-8594786dfc-lml4f                           2/2     Running  
app-namespace-two      service-two-6d554d6799-xss2w                           2/2     Running  
app-namespace-three    service-three-5bdf45df95-j6cpk                         2/2     Running   
```

## example call

after successful deploy you can call the chanin from localhost:  

```
kubectl port-forward service/service-one 8080:8080 -n app-namespace-one
```

```
curl http://localhost:8080/server-client
```