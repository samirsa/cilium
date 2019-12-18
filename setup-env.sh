RUNTIME=containerd K8S=1 MEMORY=8192 CPUS=4 vagrant destroy
#!/usr/bin/env bash
kubectl create -f install/kubernetes/quick-install.yaml; kubectl create -f examples/minikube/http-sw-app.yaml
kubectl -n kube-system exec $(kubectl -n kube-system get pods -l k8s-app=cilium -o json | jq -r '.items[0].metadata.name') -- cilium endpoint list
kubectl exec xwing -- curl -s -XPOST deathstar.default.svc.cluster.local/v1/request-landing
kubectl exec tiefighter -- curl -s -XPOST deathstar.default.svc.cluster.local/v1/request-landing
kubectl -n kube-system logs --timestamps --follow $(kubectl -n kube-system get pods -l k8s-app=cilium -o json | jq -r '.items[0].metadata.name')
kubectl create -f examples/minikube/sw_l3_l4_policy.yaml
kubectl delete -f examples/minikube/sw_l3_l4_policy.yaml
kubectl delete -f examples/minikube/http-sw-app.yaml; kubectl delete -f install/kubernetes/quick-install.yaml

sudo make install
sudo mkdir -p /etc/sysconfig/
sudo cp contrib/systemd/cilium.service /etc/systemd/system/
sudo cp contrib/systemd/cilium  /etc/sysconfig/cilium
sudo usermod -a -G cilium vagrant
sudo systemctl enable cilium
sudo systemctl restart cilium

cilium policy trace --src-identity 61224 --dst-identity 43083 --dport 80/TCP -v
kubectl exec -ti $(kubectl -n kube-system get pods -l k8s-app=cilium -o json | jq -r '.items[0].metadata.name') -n kube-system -- cilium policy trace --src-k8s-pod default:xwing -d any:class=deathstar,k8s:org=empire,k8s:io.kubernetes.pod.namespace=default --dport 80
docker login -u ssaklika -p clotXuNRZhhLf5V90bzwb1L1v/Z7rnHBmDoVBzJrREgMnHszhujqLBP/BRJG7m4tiGo3fHL3jTka/iwILmxljz6FbSvL02MfFf1eHe5G0eA= containers.cisco.com
