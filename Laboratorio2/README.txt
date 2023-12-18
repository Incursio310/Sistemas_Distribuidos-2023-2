Mauricio Cortés - 202004529-1
Claudio Espinoza - 202004539-9
Gianfranco Dissi - 202004514-3

A CONSIDERAR:
~ El Makefile contiene un make por cada servidor. Ej: make docker-oms

~ La ejecución de los DataNodes deben detenerse manualmente mediante la combinación de teclas "Ctrl + C". El resto de entidades termina de manera automática al realizar lo anterior.


DISTRIBUCIÓN DE MÁQUINAS

dist005 - YGrLddgmdvYV ---> Australia - ONU
dist006 - Xfsdn4fFeVuE ---> Asia - NameNode (OMS)
dist007 - ZTAtXuA59qDM ---> Europa - DataNode1
dist008 - jrZWunU4AUuy ---> Latinoamerica - DataNode2

ORDEN EJECUCIÓN

~ Se requiere ejecutar en el siguiente orden los makes (siguiendo con la idea de que en cada máquina se ejecuta algo en específico, tal como se indica arriba):
	make docker-datanode1     ---> En la dist 007 (Primer comando a ejecutar)
	make docker-datanode2     ---> En la dist 008 (Segundo comando a ejecutar)
   	make docker-oms           ---> En la dist 006 (Tercero comando a ejecutar)
	make docker-asia          ---> En la dist 006
	make docker-australia     ---> En la dist 005
	make docker-europa        ---> En la dist 007
	make docker-latinoamerica ---> En la dist 008 
   	make docker-onu           ---> En la dist 005

	Disclaimer: Los primeros 3 comando ingresados importan su orden. Los demás pueden variar en orden.


DISTRIBUCIÓN DE PUERTOS

NameNode:
    (Escucha)
    9000: Asia
    9001: Australia
    9002: Europa
    9003: Latinoamerica
    9004: ONU

    (Comunica)
    9005: DataNode1
    9006: DataNode2

DataNode1
    (Escucha)
    9005: NameNode

DataNode2
    (Escucha)
    9006: NameNode
