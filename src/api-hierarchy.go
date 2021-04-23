// package apihierarchy

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/kyokomi/emoji"
// )

// func healthz(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	w.Write([]byte(`{"code": 200, "text": "Service is alive"}`))
// 	log.Print("Healthz check consume")
// }

// // Init function, set up the Hierachie app API
// func Init() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/api/v1/healthz", healthz)
// 	emoji.Println("[:white_check_mark:] Hierarchy API started")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
