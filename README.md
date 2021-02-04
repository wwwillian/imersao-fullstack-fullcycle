#Imersão Full Stack && Full Cycle


***Desafio 2***
***

**Para execultar:**

docker-compose up -d
***
###Escreva sua mensagens

exec -it imersao-full-cycle_kafka_1 bash

[appuser@c3689cb7de42 ~]$ kafka-console-producer --topic=kafka --bootstrap-server=localhost:9092
***
###Receba sua mensagens

exec -it imersao-full-cycle_kafka_1 bash

[appuser@c3689cb7de42 ~]$ kafka-console-consumer --topic=kafka --bootstrap-server=localhost:9092
***
#####ou receba pelo confluent
navegador:
localhost:9021