package main
import(
        "fmt"
        "net/http"
        "github.com/gorilla/mux"
        "time"
)
func main() {
        r := mux.NewRouter()
        r.HandleFunc("/time", GetTime)
        http.ListenAndServe(":8000", r)
}
func GetTime(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Current time is %s\n", time.Now())
}
