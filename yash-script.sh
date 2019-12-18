curl -X POST -k http://localhost:20000/v1/abpolicyif/add -H "Content-Type: text/plain" -d '@test_data/policy/policy1.json'
echo
curl -X POST -k http://localhost:20000/v1/abpolicyif/add -H "Content-Type: text/plain" -d '@test_data/policy/policy2.json'
echo
curl -X POST -k http://localhost:20000/v1/abpolicyif/add -H "Content-Type: text/plain" -d '@test_data/policy/policy3.json'
echo
curl -X POST -k http://localhost:20000/v1/abpolicyif/add -H "Content-Type: text/plain" -d '@test_data/policy/policy4.json'
echo
curl -X POST -k http://localhost:10000/v1/pdp/query -H "Content-Type: text/plain" -d '@test_data/input.json'
