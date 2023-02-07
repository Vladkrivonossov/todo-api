package main

import (
  "context"
  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/sirupsen/logrus"
  "github.com/spf13/viper"
  todo "go-api"
  "go-api/pkg/handler"
  "go-api/pkg/repository"
  "go-api/pkg/service"
  "log"
  "os"
  "os/signal"
  "syscall"
)

func main() {
  if err := initConfig(); err != nil {
    log.Fatal("error initializing configs")
  }

  err := godotenv.Load(`.env`)
  if err != nil {
    log.Fatal(err)
  }

  db, err := repository.NewPostgresDB(repository.Config{
    Host: os.Getenv("POSTGRES_HOST"),
    Port: "5432",
    Username: os.Getenv("POSTGRES_USER"),
    DBName: os.Getenv("POSTGRES_DB"),
    SSLMode: "disable",
    Password: os.Getenv("POSTGRES_PASSWORD"),
  })
  if err != nil {
    log.Fatal("failed to initialize db")
  }

  repos := repository.NewRepository(db)
  services := service.NewService(repos)
  handlers := handler.NewHandler(services)

  srv := new(todo.Server)
  go func () {
    if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
      log.Fatal(err)
    }
  }()

  logrus.Print("Started")

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
  <- quit

  logrus.Print("Started")

  if err := srv.Shutdown(context.Background()); err != nil {
    logrus.Error("error occured on shitting down")
  }
  if err := db.Close(); err != nil {
    logrus.Error("error occured on db connection close")
  }
}

func initConfig() error  {
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig()
}