package renderer

import (
	"encoding/json"
	"fmt"
	"github.com/zendern/getprs/models"
	"os"
)

func RenderJson(status []models.PRStatus){
	data, err := json.MarshalIndent(status, "", "\t")

	if err != nil {
		fmt.Println(">>> Failed to convert to json <<< : " + err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
}
