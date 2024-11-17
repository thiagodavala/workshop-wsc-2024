package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	dynamoTable string
	port        string
)

func init() {
	// Leitura das variáveis de ambiente
	dynamoTable = os.Getenv("DYNAMO_TABLE")
	port = os.Getenv("PORT")
	if dynamoTable == "" || port == "" {
		log.Fatal("As variáveis de ambiente DYNAMO_TABLE e PORT devem ser definidas.")
	}
}

func main() {
	// Parse de argumentos para o comando --help
	flag.Parse()
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Println("Variáveis de ambiente necessárias:")
		fmt.Println("  DYNAMO_TABLE: Nome da tabela DynamoDB")
		fmt.Println("  PORT: Porta onde o servidor vai rodar")
		return
	}

	// Inicialização do cliente do DynamoDB
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := dynamodb.NewFromConfig(cfg)

	// Definindo os endpoints
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		handleCalc(svc, w, r)
	})
	http.HandleFunc("/health", handleHealth)

	// Iniciando o servidor HTTP
	log.Printf("Iniciando servidor na porta %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}

// handleHealth retorna a mensagem de saúde
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Healthy!")
}

// handleCalc processa o parâmetro 'input' e grava no DynamoDB
func handleCalc(svc *dynamodb.Client, w http.ResponseWriter, r *http.Request) {
	// Obtém o valor do parâmetro 'input'
	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Parâmetro 'input' é obrigatório", http.StatusBadRequest)
		return
	}

	// Gera o hash SHA256 do valor de entrada
	hash := sha256.Sum256([]byte(input))

	// Converte o hash para base64
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])

	// Persistindo no DynamoDB
	err := persistHashInDynamoDB(svc, base64Hash)
	if err != nil {
		http.Error(w, "Erro ao persistir no DynamoDB", http.StatusInternalServerError)
		return
	}

	// Retorna o valor convertido
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"base64_hash": "%s"}`, base64Hash)
}

// persistHashInDynamoDB salva o valor base64 na tabela DynamoDB
func persistHashInDynamoDB(svc *dynamodb.Client, base64Hash string) error {
	// Prepara o item para inserir na tabela
	item := map[string]types.AttributeValue{
		"hash": &types.AttributeValueMemberS{Value: base64Hash},
	}

	// Insere o item na tabela do DynamoDB
	_, err := svc.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(dynamoTable),
		Item:      item,
	})

	return err
}
