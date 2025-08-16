Este microserviçop que atua como consumidor de mensagens do kafka, processando payloads e encaminhando-ios para outro microserviço via rest api.

consumir mensagens de um tppico especifico do kafka.
processar o payload recebido (formato definido em 'payload.json')
3 encaminhar o payload para outro micro serviço via REST API (metodo patch)

configuração:
crie um arquivo '.env' na raiz do projeto com as seguintes variáveis:

KAFKA_BOOTSTRAP_SERVERS = localhost:9092
KAFKA_USERNAME=exemplo
KAFKA_PASSWORD=exemplo
KAFKA_SASL_MECHANISM=SCRAM-SHA-256
KAFKA_SECURITY_PROTOCOL="SASL_PLAINTEXT"
KAFKA_GROUP_ID=seu-grupo
KAFKA_TOPIC=seu-topico
KAFKA_TOPIC_DLQ=seu-topico-dlq
TARGET_SERVICE_URL=http://localhost:8080/api/v1


Como executar:
Crie um ambiente em docker com Kafka v3.6 para realizar seu desenvolvimento e testes.

Dependências principais:

github.com/confluentinc/confluent-kafka-go/v2
github.com/spf13/viper (p. variaveis de ambiente)

Formato do payload (OBS: encaminhado para o serviço destino via metodo PATCH):

{
    "ordemVenda": "order12345",
    "etapaAtual":"FATURADO"
}

Logs:
O serviço registra logs para: 
Conexao com o Kafka, Recebimento de mensagens, processamento do payload, chamadas ao serviço destino, erros durante o processamento.