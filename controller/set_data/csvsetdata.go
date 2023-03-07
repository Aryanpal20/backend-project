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

	fil, err := c.FormFile("file") // Get the uploaded file from the request
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// Check the file extension to determine whether it is a CSV or an XLSX file
	extension := filepath.Ext(fil.Filename)
	tokenid := c.GetFloat64("id")
	fmt.Println(tokenid)
	var emp ep.Employee
	if extension == ".csv" {
		file, err := fil.Open()
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(file)
		emp.Userid = int(tokenid)
		record, _ := reader.ReadAll()

		fmt.Println("jcbhjdsbvvbdsvbhjvbdsvhjbvhjbvircfrf", record)
		for _, r := range record {
			uuid := uuid.New().String()
			id, _ := strconv.Atoi(uuid)
			emp.ID = id
			name := r[0]
			i, _ := strconv.Atoi(r[1])
			age := i
			emp.Name = name
			emp.Age = age
			database.Database.Create(&emp)
			fmt.Println("scjkdsnjkvbskvbsfv", emp)
		}
	} else if extension == ".xlsx" {
		f, err := fil.Open()
		if err != nil {
			c.String(500, "Internal server error")
			return
		}
		xlFile, err := xlsx.OpenReaderAt(f, fil.Size)
		if err != nil {
			fmt.Println("Error opening XLSX file:", err)
			return
		}

		arr1 := [][]string{}
		emp.Userid = int(tokenid)
		for _, sheet := range xlFile.Sheets {
			fmt.Println("Sheet name:", sheet.Name)
			// Loop through all the rows in the sheet.
			for _, row := range sheet.Rows {
				fmt.Println("hjdvhv", row)
				arr := []string{}
				// Loop through all the cells in the row.
				for _, cell := range row.Cells {
					text := cell.String()

					// Print the value of the cell.
					arr = append(arr, text)

				}
				fmt.Println("dbvhebfehifbibf", arr)
				arr1 = append(arr1, arr)
				uuid := uuid.New().String()
				id, _ := strconv.Atoi(uuid)
				emp.ID = id
				name := arr[0]
				i, _ := strconv.Atoi(arr[1])
				age := i
				emp.Name = name
				emp.Age = age
				database.Database.Create(&emp)
				fmt.Println()
			}
		}
		fmt.Println("hbdhbd", arr1)
	} else {
		c.String(http.StatusBadRequest, "unsupported file type")
		return
	}

	c.String(http.StatusOK, "file processed successfully")

}
