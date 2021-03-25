istioctl x precheck


```sh
# Automatic sidecar injection
kubectl label namespace default istio-injection=enabled --overwrite
```

Problems that we have in Rapiscan deployment:
- TLS communication with each pod
- DNS problem where we need to create a DNS resolution where we need to send a request from inside the pod to the kubernetes ingress
- Attach Rapiscan's wildcard certificate for mTLS of Istio
- Blue-green deployment

Things to demo:
- [x] Show that communication is via TLS
- [x] Check if we can use istio ingress gateway to create an ingress domain so we don't have to port-forward the service
- [ ] Check if we can use istio ingress gateway so we don't have to add DNS resolution to send request from malibustream to Gluu
- [x] Use rapiscan wildcard certificate for mTLS - CANNOT (need root certificate)
- [ ] Show blue-green deployment for Istio

Gateway is like your ingress and must be bounded to a specific virtualservice
VirtualService is like your buffed-up service object. You can add rules on how to port specific users depending on the rules to the pod

Show that you can restrict binding of this gateway to only specific namespace
```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: my-gateway
  namespace: some-config-namespace
spec:
  selector:
    app: my-gateway-controller
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "ns1/*"
    - "ns2/foo.bar.com"
```