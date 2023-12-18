docker-latinoamerica:
	sudo docker build -f Dockerfile.ServerLatinoAmerica . -t containerized_latinoamerica:latest
	sudo docker run --rm --name Server-LatinoAmerica --network="host" containerized_latinoamerica:latest

docker-asia:
	sudo docker build -f Dockerfile.ServerAsia . -t containerized_asia:latest
	sudo docker run --rm --name Server-Asia --network="host" containerized_asia:latest

docker-europa:
	sudo docker build -f Dockerfile.ServerEuropa . -t containerized_europa:latest
	sudo docker run --rm --name Server-Europa --network="host" containerized_europa:latest

docker-australia:
	sudo docker build -f Dockerfile.ServerAustralia . -t containerized_australia:latest
	sudo docker run --rm --name Server-Australia --network="host" containerized_australia:latest

docker-onu:
	sudo docker build -f Dockerfile.Onu . -t containerized_onu:latest
	sudo docker run -i --rm --name Server-Onu --network="host" containerized_onu:latest

docker-datanode1:
	sudo docker build -f Dockerfile.Datanode1 . -t containerized_datanode1:latest
	sudo docker run --rm --name Server-Datanode1 -p 9005:9005 containerized_datanode1:latest

docker-datanode2:
	sudo docker build -f Dockerfile.Datanode2 . -t containerized_datanode2:latest
	sudo docker run --rm --name Server-Datanode2 -p 9006:9006 containerized_datanode2:latest

docker-oms:
	sudo docker build -f Dockerfile.Oms . -t containerized_oms:latest
	sudo docker run --rm --name Server-Oms -p 9000:9000 -p 9001:9001 -p 9002:9002 -p 9003:9003 -p 9004:9004 containerized_oms:latest
