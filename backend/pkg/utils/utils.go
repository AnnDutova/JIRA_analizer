package utils

import (
	"Backend/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Href struct {
	Href string `json:"href"`
}

type BaseMessage struct {
	status  string
	message string
	name    string
	links   map[string]Href
}

func Message(status bool, message string, name string, url string) map[string]interface{} {
	links := make(map[string]Href)
	links["self"] = Href{Href: "http://localhost:8000" + url}

	return map[string]interface{}{"status": status, "message": message, "name": name, "_links": links}
}

func RespondAny(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func RespondVariable(w http.ResponseWriter, status int, data ...any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func SortCategories(count map[string]interface{}) []string {
	result := make([]string, 0, len(count))
	for h := 0; h < 24; h++ {
		str := fmt.Sprintf("%d hours", h)
		if _, ok := count[str]; ok {
			result = append(result, str)
		}
	}
	for d := 0; d < 31; d++ {
		str := fmt.Sprintf("%d day", d)
		if _, ok := count[str]; ok {
			result = append(result, str)
		}
	}
	str := "1+year"
	if _, ok := count[str]; ok {
		result = append(result, str)
	} else {
		for y := 0; y < 8; y++ {
			str := fmt.Sprintf("%d year", y)
			if _, ok := count[str]; ok {
				result = append(result, str)
			}
		}
		str = "8+years"
		if _, ok := count[str]; ok {
			result = append(result, str)
		}
	}
	return result
}

func SortMinutesCategories(category map[string]interface{}) []string {
	keys := make([]int, 0, len(category))
	for k := range category {
		m := strings.Split(k, "m")
		minutes, err := strconv.Atoi(m[0])
		if err != nil {
			return nil
		}
		keys = append(keys, minutes)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	minutes := make([]string, 0, len(category))
	for _, el := range keys {
		minutes = append(minutes, fmt.Sprintf("%dm", el))
	}
	return minutes
}

func SortDatesForActivityGraph(data map[string]interface{}) ([]string, error) {
	keys := make([]string, 0, len(data))
	dates := make([]time.Time, 0)
	for k := range data {
		sp := strings.Split(k, ".")
		year, err := strconv.Atoi(sp[2])
		if err != nil {
			return nil, err
		}
		month, err := strconv.Atoi(sp[1])
		if err != nil {
			return nil, err
		}
		day, err := strconv.Atoi(sp[0])
		if err != nil {
			return nil, err
		}
		t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		dates = append(dates, t)
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	for _, el := range dates {
		keys = append(keys, el.Format("2.1.2006"))
	}
	return keys, nil
}

func CreateResultMap(projectCount, ipr int, data []models.GraphOutput, result map[string]interface{}) map[string]interface{} {
	for _, el := range data {
		if val, ok := result[el.Title]; ok {
			val.([]int)[ipr] = el.Count
			result[el.Title] = val
		} else {
			arr := make([]int, projectCount)
			arr[ipr] = el.Count
			result[el.Title] = arr
		}
	}
	return result
}

func JoinToMap(data1, data2 []models.GraphOutput, result map[string]interface{}) map[string]interface{} {
	for _, val := range data1 {
		if _, ok := result[val.Title]; !ok {
			result[val.Title] = 1
		}
	}
	for _, val := range data2 {
		if _, ok := result[val.Title]; !ok {
			result[val.Title] = 1
		}
	}
	return result
}
