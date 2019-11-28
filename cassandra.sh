just cass setup
kubectl create -f install/kubernetes/quick-install.yaml; kubectl create -f examples/kubernetes-cassandra/cass-sw-app.yaml
examples/kubernetes-cassandra/cass-populate-tables.sh | bash
kubectl create -f examples/kubernetes-cassandra/cass-sw-security-policy.yaml

create -
kubectl create -f install/kubernetes/quick-install.yaml; kubectl create -f examples/kubernetes-cassandra/cass-sw-app.yaml
examples/kubernetes-cassandra/cass-populate-tables.sh | bash
kubectl create -f examples/kubernetes-cassandra/pdp.yaml ; kubectl create -f examples/kubernetes-cassandra/starwars.yaml
kubectl create -f examples/kubernetes-cassandra/pdp-sw-security-policy.yaml

starwars ui -
-f install/kubernetes/quick-install.yaml -f examples/kubernetes-cassandra/just-cass-svc.yaml -f examples/kubernetes-cassandra/pdp.yaml -f examples/kubernetes-cassandra/starwars.yaml
examples/kubernetes-cassandra/cass-populate-tables.sh | bash
kubectl create -f examples/kubernetes-cassandra/pdp-just-sw-ui-policy.yaml



CASS_SERVER=$(kubectl get pods -l app=cass-server -o jsonpath='{.items[0].metadata.name}')
HQ_POD=$(kubectl get pods -l app=empire-hq -o jsonpath='{.items[0].metadata.name}')
OUTPOST_POD=$(kubectl get pods -l app=empire-outpost -o jsonpath='{.items[0].metadata.name}')
kubectl exec -it $OUTPOST_POD -- cqlsh cassandra-svc
kubectl exec -it $CASS_SERVER -- cqlsh -e "SELECT * FROM attendance.daily_records;"
kubectl exec -it $OUTPOST_POD -- cqlsh cassandra-svc -e "SELECT * FROM attendance.daily_records;"
kubectl exec -it $OUTPOST_POD -- cqlsh cassandra-svc -e "SELECT * FROM deathstar.scrum_notes;"
kubectl create -f examples/kubernetes-cassandra/cass-sw-security-policy.yaml
CILIUM_POD=$(kubectl get pods -n kube-system -l k8s-app=cilium -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n kube-system -it $CILIUM_POD -- bash
kubectl -n kube-system logs --timestamps --follow $(kubectl -n kube-system get pods -l k8s-app=cilium -o json | jq -r '.items[0].metadata.name')
kubectl -n kube-system logs --timestamps $(kubectl -n kube-system get pods -l k8s-app=cilium -o json | jq -r '.items[0].metadata.name') > pdplogs.txt
kubectl delete -f examples/kubernetes-cassandra/cass-sw-security-policy.yaml; kubectl delete -f examples/kubernetes-cassandra/cass-sw-app.yaml; kubectl delete -f install/kubernetes/quick-install.yaml


cleanup - 
kubectl delete -f examples/kubernetes-cassandra/pdp.yaml ; kubectl delete -f examples/kubernetes-cassandra/starwars.yaml
kubectl delete -f examples/kubernetes-cassandra/pdp-sw-security-policy.yaml; kubectl delete -f examples/kubernetes-cassandra/cass-sw-app.yaml; kubectl delete -f install/kubernetes/quick-install.yaml

/etc/resolv.conf
nameserver 172.20.0.10
search default.svc.cluster.local svc.cluster.local cluster.local
