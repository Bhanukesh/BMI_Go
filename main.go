package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func main() {
    http.HandleFunc("/bmi", calculateBMI) // Route to handle BMI calculation
    fmt.Println("Server starting on port 8080...")
    err := http.ListenAndServe(":8080", nil) // Start the HTTP server on port 8080
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}

func calculateBMI(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Received request for BMI calculation")

    heightStr := r.URL.Query().Get("height")
    weightStr := r.URL.Query().Get("weight")

    fmt.Printf("Height: %s, Weight: %s\n", heightStr, weightStr)

    if heightStr == "" || weightStr == "" {
        http.Error(w, "Please provide height and weight parameters.", http.StatusBadRequest)
        return
    }

    height, err := strconv.ParseFloat(heightStr, 64)
    if err != nil {
        http.Error(w, "Invalid height.", http.StatusBadRequest)
        return
    }

    weight, err := strconv.ParseFloat(weightStr, 64)
    if err != nil {
        http.Error(w, "Invalid weight.", http.StatusBadRequest)
        return
    }

    bmi := weight / (height * height)
    fmt.Printf("Calculated BMI: %.2f\n", bmi)

    var category, recommendation, yogaAsanas string
    switch {
    case bmi < 18.5:
        category = "Underweight"
        recommendation = "Increase calorie intake with nutritious foods, consult a dietitian."
        yogaAsanas = "Vajrasana, Bhujangasana, and Pawanmuktasana."
    case bmi <= 24.9:
        category = "Normal weight"
        recommendation = "Maintain your current diet and exercise regimen."
        yogaAsanas = "Tadasana, Virabhadrasana, and Trikonasana."
    case bmi <= 29.9:
        category = "Overweight"
        recommendation = "Consider reducing calorie intake and increasing physical activity, consult a dietitian."
        yogaAsanas = "Dhanurasana, Ardha Matsyendrasana, and Surya Namaskar."
    default:
        category = "Obesity"
        recommendation = "Significantly reduce calorie intake, increase physical activity, and consult a healthcare provider."
        yogaAsanas = "Kapalbhati, Setu Bandhasana, and Ardhachakrasana."
    }

    response := fmt.Sprintf("Your BMI is: %.2f, Category: %s. Recommendations: %s Yoga Asanas: %s", bmi, category, recommendation, yogaAsanas)
    fmt.Println("Response: ", response)
    fmt.Fprintf(w, response)
}
