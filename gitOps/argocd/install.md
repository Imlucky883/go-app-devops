
# Installation
```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
# Accessing the UI
```
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

kubectl get svc argocd-server -n argocd 

```
# Login Credentials
```
kg secrets/argocd-initial-admin-secret -o jsonpath='{.data.password}' | base64 --decode
```