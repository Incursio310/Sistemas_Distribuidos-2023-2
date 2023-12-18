docker-rabbit:
	sudo docker run -it --rm --name rabbitmq -p 5673:5672 -p 15673:15672 rabbitmq:3.9-management

docker-america:
	sudo docker build -f Dockerfile.ServerAmerica . -t containerized_america:latest
	sudo docker run --rm --name Server-America -p 9000:9000 --network="host" containerized_america:latest

docker-asia:
	sudo docker build -f Dockerfile.ServerAsia . -t containerized_asia:latest
	sudo docker run --rm --name Server-Asia -p 9001:9001 --network="host" containerized_asia:latest

docker-europa:
	sudo docker build -f Dockerfile.ServerEuropa . -t containerized_europa:latest
	sudo docker run --rm --name Server-Europa -p 9002:9002 --network="host" containerized_europa:latest

docker-oceania:
	sudo docker build -f Dockerfile.ServerOceania . -t containerized_oceania:latest
	sudo docker run --rm --name Server-Oceania -p 9003:9003 --network="host" containerized_oceania:latest

docker-central:
	sudo docker build -f Dockerfile.centralServer . -t containerized_central:latest
	sudo docker run --rm --name central-Server --network="host" containerized_central:latest