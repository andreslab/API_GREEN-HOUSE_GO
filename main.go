package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var db_name = "greenhousedb"

/*type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}*/

type Sensor struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date     string             `json:"date,omitempty" bson:"date,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Category string             `json:"category,omitempty" bson:"category,omitempty"`
	Value    float64            `json:"value,omitempty" bson:"value,omitempty"`
}

/*

Humedad (en ambiente, en suelo a distintas profundidades, en hoja)
Temperatura
Luminosidad y radiación solar y ultravioleta.
Contaminación y Gases (CO2, NH2…)
Crecimiento de tallo.
Diametro de fruto
Diametro de tronco
Presión atmosférica
Consumo de agua

*/

//humedad y temperatura
func SaveWetAndTemperature(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensor Sensor
	_ = json.NewDecoder(request.Body).Decode(&sensor)
	collection := client.Database(db_name).Collection("wet_and_temperature")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, sensor)
	json.NewEncoder(response).Encode(result)
}

//brillo
func SaveBrightness(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensor Sensor
	_ = json.NewDecoder(request.Body).Decode(&sensor)
	collection := client.Database(db_name).Collection("brightness")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, sensor)
	json.NewEncoder(response).Encode(result)
}

//radiación uv
func SaveUVRadiation(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensor Sensor
	_ = json.NewDecoder(request.Body).Decode(&sensor)
	collection := client.Database(db_name).Collection("uv_radiation")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, sensor)
	json.NewEncoder(response).Encode(result)
}

//gases
func SaveGas(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensor Sensor
	_ = json.NewDecoder(request.Body).Decode(&sensor)
	collection := client.Database(db_name).Collection("gas")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, sensor)
	json.NewEncoder(response).Encode(result)
}

//presion atmosferica
func SaveAtmosphericPressure(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensor Sensor
	_ = json.NewDecoder(request.Body).Decode(&sensor)
	collection := client.Database(db_name).Collection("atmospheric_pressure")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, sensor)
	json.NewEncoder(response).Encode(result)
}

//consumo de agua
func SaveWaterConsumption(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensor Sensor
	_ = json.NewDecoder(request.Body).Decode(&sensor)
	collection := client.Database(db_name).Collection("water_consumption")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, sensor)
	json.NewEncoder(response).Encode(result)
}

func GetSensors(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var sensors []string
	collection := client.Database(db_name).Collection("sensors")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var sensor string
		cursor.Decode(&sensor)
		sensors = append(sensors, sensor)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(sensors)
}

/*
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database(db_name).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}*/

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/sensors", GetSensors).Methods("GET")
	router.HandleFunc("/sensor/atmospheric-pressure", SaveAtmosphericPressure).Methods("POST")
	router.HandleFunc("/sensor/brightness", SaveBrightness).Methods("POST")
	router.HandleFunc("/sensor/gas", SaveGas).Methods("POST")
	router.HandleFunc("/sensor/uv-radiation", SaveUVRadiation).Methods("POST")
	router.HandleFunc("/sensor/water-consumption", SaveWaterConsumption).Methods("POST")
	router.HandleFunc("/sensor/wet-and-temperature", SaveWetAndTemperature).Methods("POST")

	/*router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")*/
	http.ListenAndServe(":12345", router)
}
