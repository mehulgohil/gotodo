package utility

import (
	"encoding/json"
	"fmt"
	"github.com/mehulgohil/gotodo/models"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

func PrintOnOutputFormat(tasks models.Todos, outputFormat string) {
	switch outputFormat {
	case "json":
		marshal, err := json.Marshal(tasks)
		if err != nil {
			return
		}
		fmt.Println(string(marshal))
		break
	case "yaml":
		marshal, err := yaml.Marshal(tasks)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(marshal))
		break
	default:
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

		fs := "%s\t%s\t%s\t%s\n"
		fmt.Fprintf(w, fs, "ID", "Name", "Due Date", "Completed")
		for _, eachTask := range tasks.Todos {
			fmt.Fprintf(w, fs, strconv.Itoa(eachTask.ID), eachTask.Name, eachTask.DueDate, strconv.FormatBool(eachTask.Completed))
		}

		w.Flush()
	}
}
