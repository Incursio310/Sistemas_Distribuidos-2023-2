Mauricio Cortés - 202004529-1
Claudio Espinoza - 202004539-9
Gianfranco Dissi - 202004514-3

A CONSIDERAR:
~ El Makefile contiene un make por cada servidor. Ej: make docker-america

~ La ejecución del programa del servidor central finaliza de manera automática, mientras que la ejecución de los servidores regionales y cola rabbit deben detenerse manualmente mediante la combinación de teclas "Ctrl + C".

Máquinas Virtuales (y sus contraseñas) asignadas a cada servidor:

	Máquina - Contraseña

	dist005 - YGrLddgmdvYV ---> centralServer y ServerAmerica

	dist006 - Xfsdn4fFeVuE ---> ServerAsia

	dist007 - ZTAtXuA59qDM ---> ServerEuropa

	dist008 - jrZWunU4AUuy ---> ServerOceania y rabbitmq

~ Se requiere ejecutar en el siguiente orden los make's (siguiendo con la idea de que en cada máquina se ejecuta algo en específico, tal como se indica arriba):
	make docker-rabbit  ---> En la dist 008 (Primer comando a ejecutar)
	make docker-america ---> En la dist 005
	make docker-asia    ---> En la dist 006
	make docker-europa  ---> En la dist 007
	make docker-oceania ---> En la dist 008
	make docker-central ---> En la dist 005 (Último comando a ejecutar)

	Disclaimer: Los make docker de los servidores regionales pueden variar su orden. Lo importante es ejecutar primero la cola rabbit y último el servidor central.

~ La única forma que nos resultó para importar el archivo proto fue crear un repositorio nuevo exclusivo para estos archivos, y referenciarlos desde este nuevo repositorio al repositorio del Laboratorio.
Para más entendimiento, en caso de ser necesario, revisar, por ejemplo, la línea 16 de cualquier main.go de algún servidor regional.
