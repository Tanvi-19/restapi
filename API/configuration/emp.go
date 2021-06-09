package configuration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	Emp_ID      string `json:"id" bson:"id"`
	Emp_Name    string `json:"name" bson:"name"`
	Emp_Age     int    `json:"age" bson:"age"`
	Emp_email   string `json:"email" bson:"email"`
	Designation string `json:"designation" bson:"designation"`
	Salary      int    `json:"salary" bson:"salary"`
}

type bug struct {
	Statuscode int    `json:"statuscode"`
	Err        string `json:"err"`
}

func print(write io.Writer, result interface{}) {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprint(write, err)
	}
	fmt.Fprint(write, string(data))
}

func ViewAllEmp(res http.ResponseWriter, req *http.Request) {
		col:=ConnectToDB()
		var employees []Employee
		c := context.TODO()
		i, err := col.Find(c, bson.D{})
		if err != nil {
			panic(err)
		}
		for i.Next(c) {
			var employee Employee
			i.Decode(&employee)
			employees = append(employees, employee)
		}

		if err := i.Err(); err != nil {
			panic(err)
		}
		print(res, employees)

}

func ViewEmpByID(res http.ResponseWriter, req *http.Request) {
		col := ConnectToDB()
		vars := mux.Vars(req)
		id := vars["id"]
		filter := bson.D{{"id", id}}
		var employee Employee
		err := col.FindOne(context.TODO(), filter).Decode(&employee)
		if err != nil {
			print(res, bug{Statuscode: http.StatusNotFound, Err: err.Error()})
		} else {
			res.WriteHeader(http.StatusFound)
			print(res, employee)
		}

}


func CreateEmp(res http.ResponseWriter, req *http.Request) {
		var employee Employee
		col:=ConnectToDB()
		if err := json.NewDecoder(req.Body).Decode(&employee); err != nil {
			panic(err)
		}
		col.InsertOne(context.TODO(), employee)
		res.WriteHeader(http.StatusCreated)
		print(res, employee)

}

func UpdateEmp(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		filter := bson.D{{"id", id}}
		var employee Employee
		if err := json.NewDecoder(req.Body).Decode(&employee); err != nil {
			panic(err)
		}
		col:=ConnectToDB()
		result := col.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": employee}, options.FindOneAndUpdate().SetReturnDocument(1))
		decoded := Employee{}
		if err := result.Decode(&decoded); err != nil {
			panic(err)
		}

		print(res, decoded)
}

func DeleteEmp(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		filter := bson.D{{"id", id}}
		col:=ConnectToDB()
		var employee Employee
		col.FindOneAndDelete(context.TODO(), filter).Decode(&employee)
		print(res, employee)
}
