package setdata

import (
	"encoding/csv"
	"fmt"
	"gin/database"
	ep "gin/model/emp_model"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
)

func PostCSVData(c *gin.Context) {

	file, err := c.FormFile("file") // Get the uploaded file from the request
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// Check the file extension
	extension := filepath.Ext(file.Filename)
	tokenid := c.GetFloat64("id")
	fmt.Println(tokenid)

	file1, err := file.Open()
	if err != nil {
		panic(err)
	}
	if extension == ".csv" {

		reader := csv.NewReader(file1)
		record, _ := reader.ReadAll()
		fmt.Println("jcbhjdsbvvbdsvbhjvbdsvhjbvhjbvircfrf", record)
		dataMap := []ep.Employee{}
		data := make(map[string]string)
		var email string
		for firstrow, row := range record {
			if firstrow == 0 {
				continue
			}
			var emp ep.Employee
			uuid := uuid.New().String()
			emp.ID = uuid
			name := row[0]
			email = row[2]
			fmt.Println("sdhbvd", email)
			phone := row[3]
			age := row[1]
			i, _ := strconv.Atoi(row[1])
			data["ID"] = uuid
			data["Name"] = name
			data["Age"] = age
			data["Email"] = email
			data["Phone_No"] = phone
			dataMap = append(dataMap, ep.Employee{ID: uuid, Name: name, Age: i, Email: email, Phone_No: phone, Userid: int(tokenid)})
		}
		database.Database.Create(dataMap)
	} else if extension == ".xlsx" {
		xlFile, err := xlsx.OpenReaderAt(file1, file.Size)
		if err != nil {
			fmt.Println("Error opening XLSX file:", err)
			return
		}
		dataMap := []ep.Employee{}
		data := make(map[string]string)
		// Iterate over the sheets and rows
		for _, sheet := range xlFile.Sheets {
			for firstrow, row := range sheet.Rows {
				if firstrow == 0 {
					continue
				}
				name := row.Cells[0].Value
				age := row.Cells[1].Value
				email := row.Cells[2].Value
				phone := row.Cells[3].Value
				floatValue, err := strconv.ParseFloat(age, 64)
				if err != nil {
					fmt.Println("Error:", err)
				}
				intValue := int(floatValue)
				uuid := uuid.New().String()
				data["ID"] = uuid
				data["Name"] = name
				data["Age"] = age
				data["Email"] = email
				data["Phone_No"] = phone
				dataMap = append(dataMap, ep.Employee{ID: uuid, Name: name, Age: intValue, Email: email, Phone_No: phone, Userid: int(tokenid)})
			}
		}
		fmt.Println(dataMap)
		database.Database.Create(dataMap)
	} else {
		c.String(http.StatusBadRequest, "unsupported file type")
		return
	}

	c.String(http.StatusOK, "file processed successfully")

}
