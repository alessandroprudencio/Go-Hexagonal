package server
import (
 "github.com/alessandroprudencio/Go-Hexagonal/application"
 "github.com/alessandroprudencio/Go-Hexagonal/adapters/web/handler"
 "net/http"
 "github.com/gorilla/mux"
 "github.com/codegangsta/negroni"
 "log"
 "os"
 "time"
 )

type Webserver struct {
    Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
    return &Webserver{}
}

func (w Webserver) Serve(){
    r := mux.NewRouter()
    n := negroni.New(
        negroni.NewLogger(),
    )

    handler.MakeProductHandlers(r, n, w.Service)
    http.Handle("/", r)

    server := &http.Server{
        ReadHeaderTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        Addr: ":8080",
        Handler: http.DefaultServeMux,
        ErrorLog: log.New(os.Stderr, "log: ", log.  Lshortfile),
    }

    err := server.ListenAndServe()

    if  err != nil {
        log.Fatal(err)
    }
}
