mustgather-report-operator
====================

Operator which deploy various operating systems on top of Kubernetes.

# Build Operator
```bash
export GO111MODULE=on
operator-sdk generate k8s # if changes made to *_types.go
go mod vendor
operator-sdk build quay.io/$USER/must-gather-operator:v0.0.1
sed -i "s|REPLACE_IMAGE|quay.io/$USER/must-gather-operator:v0.0.1|g" deploy/operator.yaml
docker push quay.io/$USER/must-gather-operator:v0.0.1
```

# Installation
```bash
kubectl create -f deploy/crds/mustgather_v1alpha1_mustgatherreport_crd.yaml
kubectl create -f deploy/02-namespace.yaml
kubectl create -f deploy/03-cluster-operator.yaml
kubectl create -f deploy/04-service_account.yaml
kubectl create -f deploy/05-role_binding.yaml
kubectl create -f deploy/06-operator.yaml
```

# Create Must Gather Report
```bash
cat <<EOF | kubectl create -f -
apiVersion: mustgather.openshift.io/v1alpha1
kind: MustGatherReport
metadata:
  name: example-mustgatherreport
spec:
  images:
  - quay.io/kubevirt/must-gather
EOF
```

Verify MustGatherReport was created:

```bash
kubectl get mustgatherreport example-mustgatherreport
```

# Troubleshooting

The Pods for must-gather are being created with a new persistent volume claim as their volume using the default storage class.
If there is no storage-class defined, the must-gather will fail. In order to define a storage class as the default follow:
```bash
  kubectl patch storageclass <your-class-name> -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
```

Verify default storage class is set by:
```bash
kubectl get storageclass
NAME              PROVISIONER                    AGE
local (default)   kubernetes.io/no-provisioner   142m
```

# Development
After cloning the repository, run the operator locally using:
```bash
export GO111MODULE=on
go mod vendor
operator-sdk up local --namespace=operator-must-gather
```

After changes to types file run:
```bash
operator-sdk generate k8s
```

In order to debug the operator locally using 'dlv', start the operator locally by running (assuming namespace is 'openshift-must-gather'):
```bash
operator-sdk build quay.io/$USER/must-gather-operator:v0.0.1
operator-sdk up local  --enable-delve --namespace=openshift-must-gather
```
Kubernetes cluster should be avaiable and pointed by `~/.kube/config`.
The CRDs of `./deploy/crds/` should be applied on it.

Connect to the debug session, i.e. if using vscode, create launch.json as:

```yaml
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Must-Gather Operator",
      "type": "go",
      "request": "attach",
      "mode": "remote",
      "port": 2345,
      "host": "127.0.0.1"
    }
  ]
}
```
