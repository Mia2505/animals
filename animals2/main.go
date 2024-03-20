package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Animal struct {
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

var animals []Animal

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// func init() {
// 	// Load data from data.json when starting the server
// 	data, err := ioutil.ReadFile("data.json")
// 	if err != nil {
// 		fmt.Println("Error reading data.json:", err)
// 		return
// 	}
// 	err = json.Unmarshal(data, &animals)
// 	if err != nil {
// 		fmt.Println("Error unmarshalling data:", err)
// 		return
// 	}
// }

// func GetAnimals() []Animal {
// 	return animals
// }

// func AddAnimal(newAnimal Animal) {
// 	animals = append(animals, newAnimal)

// 	// Save data to data.json after adding a new animal
// 	saveData()
// }

// func saveData() {
// 	data, err := json.MarshalIndent(animals, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshalling data:", err)
// 		return
// 	}
// 	err = ioutil.WriteFile("data.json", data, 0644)
// 	if err != nil {
// 		fmt.Println("Error writing data to data.json:", err)
// 		return
// 	}
// }

// func HandleListAnimals(w http.ResponseWriter, r *http.Request) {
// 	// Return list animals
// 	jsonData, err := json.Marshal(GetAnimals())
// 	if err != nil {
// 		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonData)
// }

func LoadAnimals() error {
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&animals)
	if err != nil {
		return err
	}

	return nil
}

func HandleListAnimals(w http.ResponseWriter, r *http.Request) {
	err := LoadAnimals()
	if err != nil {
		http.Error(w, "Error loading animals", http.StatusInternalServerError)
		return
	}

	// Return list animals
	for i := range animals {
		animals[i].Price *= 1.2 // Increase price by 20%
	}

	jsonData, err := json.Marshal(animals)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// func HandleAddAnimal(w http.ResponseWriter, r *http.Request) {
// 	// Parse request body
// 	var newAnimal Animal
// 	err := json.NewDecoder(r.Body).Decode(&newAnimal)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Add new animal to data source
// 	AddAnimal(newAnimal)

// 	// Return success response
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Animal added successfully"})
// }

func HandleAddAnimal(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var newAnimal Animal
	err := json.NewDecoder(r.Body).Decode(&newAnimal)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add new animal to data source
	animals = append(animals, newAnimal)

	// Save animals to file
	err = SaveAnimals()
	if err != nil {
		http.Error(w, "Error saving animals", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Animal added successfully"})
}


func SaveAnimals() error {
	file, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(animals)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/api/animals", HandleListAnimals)
	http.HandleFunc("/api/add-animal", HandleAddAnimal)
	http.Handle("/image", http.StripPrefix("/image", http.FileServer(http.Dir("image"))))
	fmt.Println("Server is running on http://localhost:6969")
	_ = http.ListenAndServe(":6969", nil)
}
