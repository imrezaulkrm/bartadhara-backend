package routes

import (
    "github.com/gorilla/mux"
    "news_backend/controllers"
)

func NewsRoutes(r *mux.Router) {
    r.HandleFunc("/news", controllers.GetAllNews).Methods("GET")       // সব নিউজ পেতে
    r.HandleFunc("/news/{id}", controllers.GetNewsByID).Methods("GET") // নির্দিষ্ট আইডি দ্বারা নিউজ পেতে
    r.HandleFunc("/news", controllers.CreateNews).Methods("POST")      // নতুন নিউজ তৈরি করতে
    r.HandleFunc("/news/{id}", controllers.UpdateNews).Methods("PUT")  // নতুন আপডেট রাউট
    r.HandleFunc("/news/{id}", controllers.DeleteNews).Methods("DELETE")

}
