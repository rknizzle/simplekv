key_name=$1

curl -i -F upload=@./curl-requests/testData.txt http://localhost:8080/$key_name
