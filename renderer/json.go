package renderer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zendern/getprs/models"
	"os"
)

func RenderJson(status []models.PRStatus){
	data, err := json.Marshal(status)

	if err != nil {
		fmt.Println(">>> Failed to convert to josn <<< : " + err.Error())
		os.Exit(1)
	}
	fmt.Println(jsonPrettyPrint(string(data)))
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}