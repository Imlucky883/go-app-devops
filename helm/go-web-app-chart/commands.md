# Helm Commands

```
helm install go-web-app ./go-web-app-chart

helm upgrade go-web-app --set replicaCount=2

helm history 

helm rollback chart-name revision

helm uninstall go-web-app
```

`NOTE :  You need to be inside the /helm directory to run the install chart command`