package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func newServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to this newServer")
}

//start a new http server and return error type
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/new", newServer)
	err := srv.ListenAndServe()
	return err
}

func main() {
	cancelCtx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(cancelCtx)

	srv := &http.Server{Addr: ":8080"}

	//use g.Go to start http server
	g.Go(func() error {
		return StartHttpServer(srv)
	})

	g.Go(func() error {
		<-ctx.Done()
		//close http server when error happened
		return srv.Shutdown(ctx)
	})
	//use notify to listen signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case term := <-ch:
				cancel()
				return errors.New(term.String())
			}
		}
		return nil
	})

	//wait until all groups stop
	err := g.Wait()

	if err != nil {
		fmt.Println("some group error happened: ", err)
		return
	}
	fmt.Println("all groups end normally")
}