# golang-argocd-kustomize

Start Minikube
---------------------
```bash
minikube start
```

Install Argo CD
---------------------
```bash
brew install argocd
````

Create namespace and install argo cd
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Run Argo CD
```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

Check Password Secret argo cd 
--------------------
(user: admin / password: ${secret})
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}"| base64 -d;echo
```