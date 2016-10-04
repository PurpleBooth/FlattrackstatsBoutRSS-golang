package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type FlattrackstatRssServer struct {
	Port              int
	FeedUrl           string
	FlattrackstatsUrl string
}

func (s FlattrackstatRssServer) Listen() {
	router := httprouter.New()
	router.GET("/", s.makeIndexHandler)
	router.GET("/:id", s.makeRssHandler)

	router.
		log.Println(fmt.Sprintf("Listening on :%v", s.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), router))
}

func (s FlattrackstatRssServer) makeIndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Visit %v/teams and then visit the team page ", s.FlattrackstatsUrl)
	fmt.Fprintf(w, "%v/teams/$teamId, copy the teamId and ", s.FlattrackstatsUrl)
	fmt.Fprint(w, "put it put it at the end of this url e.g. ")
	fmt.Fprintf(w, "%v/13214", s.FeedUrl)
}

func (s FlattrackstatRssServer) makeRssHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, ps.ByName("id"))
}
