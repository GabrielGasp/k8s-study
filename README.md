1. Create KIND cluster
```bash
kind create cluster --config kind.yaml
```

2. Install nginx controller
```bash
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

3. Create namspace, service account, configmap and secret
```bash
kubectl apply -f k8s/namespace.yaml && \
kubectl apply -f k8s/service-account.yaml && \
kubectl apply -f k8s/configmap.yaml && \
kubectl apply -f k8s/secret.yaml
```

4. Apply the rest
```bash
kubectl apply -f k8s/
```

5. Forward ingress port
```bash
kubectl port-forward -n ingress-nginx svc/ingress-nginx-controller 8000:80
```

6. Open browser and go to http://demo.127.0.0.1.nip.io:8000/

7. Cleanup
```bash
kind delete cluster
```


