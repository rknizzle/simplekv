key_name=$1

curl -i -X POST -H 'Content-type: text/plain' \
  --data-binary @./curl-requests/testData.txt \
  http://localhost:8080/$key_name
