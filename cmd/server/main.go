package main

import (
        "context"
        "log"
        "net/http"
        "os"

        "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
        dbURL := os.Getenv("DATABASE_URL")
        if dbURL == "" {
                log.Fatal("DATABASE_URL is not set")
        }

        ctx := context.Background()

        pool, err := pgxpool.New(ctx, dbURL)
        if err != nil {
                log.Fatalf("cannot create database pool: %v", err)
        }
        defer pool.Close()

        err = pool.Ping(ctx)
        if err != nil {
                log.Fatalf("cannot reach database: %v", err)
        }
        log.Println("connected to postgres")

        mux := http.NewServeMux()
        mux.HandleFunc("GET /healthz", healthz)

        log.Println("listening on :8080")
        err = http.ListenAndServe(":8080", mux)
        log.Fatal(err)
}

func healthz(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
}