package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/DATA-DOG/go-txdb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func ConnectDatabase() (webEngine *gin.Engine, db *sql.DB) {
	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	// Open inicia un pool de conexiones. Sólx abrir un vez
	db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")

	webEngine = gin.Default()

	return webEngine, db
}

func ConnectMockDatabase() (webEngine *gin.Engine, db *sql.DB, err error) {
	webEngine = gin.Default()

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	// Open inicia un pool de conexiones. Sólx abrir un vez
	txdb.Register("txdb", "mysql", configDB.FormatDSN())
	db, err = sql.Open("txdb", uuid.New().String())
	if err != nil {
		return webEngine, nil, err
	}
	if err = db.Ping(); err != nil {
		return webEngine, nil, err
	}
	log.Println("database configured")

	return webEngine, db, nil
}

func ConnectDynamoDB() (*dynamodb.DynamoDB, error) {
	region := "us-east-1"
	endpoint := "http://localhost:8000"
	credentials := credentials.NewStaticCredentials(os.Getenv("DYNAMOID"), os.Getenv("DYNAMOSECRET"), os.Getenv("DYNAMOTOKEN"))

	session, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(credentials))
	if err != nil {
		return nil, err
	}

	dynamodb := dynamodb.New(session)

	return dynamodb, nil
}
