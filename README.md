mustgather-report-operator
====================

Operator which deploy various operating systems on top of Kubernetes.

# Build Operator
```bash
export GO111MODULE=on
operator-sdk generate k8s # if changes made to *_types.go
go mod vendor
operator-sdk build quay.io/$USER/must-gather-operator:v0.0.1
sed -i 's|REPLACE_IMAGE|quay.io/$USER/must-gather-operator:v0.0.1|g' deploy/operator.yaml
docker push qquay.io/$USER/must-gather-operator:v0.0.1
```

# Installation
```bash
kubectl create -f deploy/crds/mustgather_v1alpha1_mustgatherreport_crd.yaml
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/operator.yaml
```

# Create Fedora virtual machine
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

TBD

# Development
After cloning the repository, run the operator locally using:
```bash
export GO111MODULE=on
go mod vendor
operator-sdk up local --namespace=default
```

After changes to types file run:
```bash
operator-sdk generate k8s
```

In order to debug the operator locally using 'dlv', start the operator locally:
```bash
operator-sdk build quay.io/$USER/must-gather-operator:v0.0.1
OPERATOR_NAME=must-gather-operator WATCH_NAMESPACE=default ./build/_output/bin/must-gather-operator
```

Kubernetes cluster should be avaiable and pointed by `~/.kube/config`.
The CRDs of `./deploy/crds/` should be applied on it.

From a second terminal window run:
```bash
dlv attach --headless --api-version=2 --listen=:2345 $(pgrep -f must-gather-operator) ./build/_output/bin/must-gather-operator
```

Connect to the debug session, i.e. if using vscode, create launch.json as:

```yaml
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Connect to must-gather-operator",
            "type": "go",
            "request": "launch",
            "mode": "remote",
            "remotePath": "${workspaceFolder}",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceFolder}",
            "env": {},
            "args": []
        }
    ]
}
```