package main

import (
	"fmt"
	"log"
	"p1-reciever/broker"
	dbrepo "p1-reciever/dbRepo"
	"p1-reciever/dbRepo/repository"
)

func main() {
	// database config
	var (
		user     = "alireza"
		password = "1234"
		host     = "postgres"
		dbName   = "db1"
		port     = "5432"
	)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := dbrepo.NewPostgresDB(connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	postgreRepo := repository.NewPostgresRepo(db)
	err = postgreRepo.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("database is ready to be used")
	// rabbitMQ config
	var (
		Ruser      = "alireza"
		Rpassword  = "1234"
		Rhost      = "rabbitmq"
		rabbitPort = "5672"
		queueName  = ""
	)

	conn := fmt.Sprintf("amqp://%s:%s@%s:%s/", Ruser, Rpassword, Rhost, rabbitPort)
	err = broker.InitRabbit(conn, queueName)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("ready to serve")
	err = broker.Listen(postgreRepo)
	if err != nil {
		log.Fatalln(err)
	}

}
