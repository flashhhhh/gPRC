package handler

import (
	"net/http"
	"strconv"

	"flash/server/rest/service"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	firstNumString := r.URL.Query().Get("first_num")
	firstNum, err := strconv.Atoi(firstNumString)
	if err != nil {
		http.Error(w, "first_num query param must be an integer", http.StatusBadRequest)
		return
	}

	secondNumString := r.URL.Query().Get("second_num")
	secondNum, err := strconv.Atoi(secondNumString)
	if err != nil {
		http.Error(w, "second_num query param must be an integer", http.StatusBadRequest)
		return
	}

	result, err := service.Add(firstNum, secondNum)
	if err != nil {
		http.Error(w, "Failed to add", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(result)))
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	firstNumString := r.URL.Query().Get("first_num")
	firstNum, err := strconv.Atoi(firstNumString)
	if err != nil {
		http.Error(w, "first_num query param must be an integer", http.StatusBadRequest)
		return
	}

	secondNumString := r.URL.Query().Get("second_num")
	secondNum, err := strconv.Atoi(secondNumString)
	if err != nil {
		http.Error(w, "second_num query param must be an integer", http.StatusBadRequest)
		return
	}

	result, err := service.Subtract(firstNum, secondNum)
	if err != nil {
		http.Error(w, "Failed to subtract", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(result)))
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	firstNumString := r.URL.Query().Get("first_num")
	firstNum, err := strconv.Atoi(firstNumString)
	if err != nil {
		http.Error(w, "first_num query param must be an integer", http.StatusBadRequest)
		return
	}

	secondNumString := r.URL.Query().Get("second_num")
	secondNum, err := strconv.Atoi(secondNumString)
	if err != nil {
		http.Error(w, "second_num query param must be an integer", http.StatusBadRequest)
		return
	}

	result, err := service.Multiply(firstNum, secondNum)
	if err != nil {
		http.Error(w, "Failed to multiply", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(result)))
}

func DivisionHandler(w http.ResponseWriter, r *http.Request) {
	firstNumString := r.URL.Query().Get("first_num")
	firstNum, err := strconv.Atoi(firstNumString)
	if err != nil {
		http.Error(w, "first_num query param must be an integer", http.StatusBadRequest)
		return
	}

	secondNumString := r.URL.Query().Get("second_num")
	secondNum, err := strconv.Atoi(secondNumString)
	if err != nil {
		http.Error(w, "second_num query param must be an integer", http.StatusBadRequest)
		return
	}

	result, err := service.Division(firstNum, secondNum)
	if err != nil {
		http.Error(w, "Failed to divide", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(result)))
}

func PowerHandler(w http.ResponseWriter, r *http.Request) {
	firstNumString := r.URL.Query().Get("first_num")
	firstNum, err := strconv.Atoi(firstNumString)
	if err != nil {
		http.Error(w, "first_num query param must be an integer", http.StatusBadRequest)
		return
	}

	secondNumString := r.URL.Query().Get("second_num")
	secondNum, err := strconv.Atoi(secondNumString)
	if err != nil {
		http.Error(w, "second_num query param must be an integer", http.StatusBadRequest)
		return
	}

	result, err := service.Power(firstNum, secondNum)
	if err != nil {
		http.Error(w, "Failed to power", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(result)))
}