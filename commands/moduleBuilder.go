package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/urfave/cli"
)

var moduleFields []string
var modelName string
var moduleName string
var data []string

func ModuleBuilderInit(c *cli.Context) {
	pluralize := pluralize.NewClient()
	moduleName = c.String("moduleName")
	moduleFields = strings.Split(c.String("data"), ",")
	modelName = strings.Title(pluralize.Singular(strings.ToLower(c.String("moduleName"))))
	data = strings.Split(c.String("data"), ",")
	createModel()
	fmt.Println("Creating Module Model")
	createRoute()
	fmt.Println("Creating Module Route")
	createController()
	fmt.Println("Creating Module Controller")
	createDocs()
	fmt.Println("Creating Module Docs")
	fmt.Println("Your module is ready!")
}

func createModel() {
	os.MkdirAll("modules"+string(filepath.Separator)+moduleName, os.ModePerm)
	oringinaleContent := readFileContent("module" + string(filepath.Separator) + "models" + string(filepath.Separator) + "model.go")
	fields := "\n\t"
	// Title  string `gorm:"type:varchar(100);" json:"title" binding:required"`
	for _, field := range data {
		fields = fields + strings.Title(field) + " string " + "`gorm:\"type:varchar(100);" + "\"" + " json:" + "\"" + strings.ToLower(field) + "\"" + "\"" + "`" + "\n\t"
	}
	replacementsStrings := strings.NewReplacer(
		"{modelName}", modelName,
		"// Set the model fields", fields,
	)
	actualContent := replacementsStrings.Replace(oringinaleContent)
	os.MkdirAll("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"models", os.ModePerm)
	ioutil.WriteFile("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"models"+string(filepath.Separator)+modelName+".go", []byte(actualContent), os.ModePerm)
}

func createRoute() {
	oringinaleContent := readFileContent("module" + string(filepath.Separator) + "routes" + string(filepath.Separator) + "api.go")
	actualContent := strings.ReplaceAll(oringinaleContent, "{moduleName}", strings.ToLower(moduleName))
	os.MkdirAll("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"routes", os.ModePerm)
	ioutil.WriteFile("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"routes"+string(filepath.Separator)+"api.go", []byte(actualContent), os.ModePerm)
}

func createController() {
	pluralize := pluralize.NewClient()
	oringinaleContent := readFileContent("module" + string(filepath.Separator) + "controllers" + string(filepath.Separator) + "controller.go")
	replacementsStrings := strings.NewReplacer(
		"{moduleInPlural}", moduleName,
		"{modelName}", modelName,
		"{moduleInSingular}", pluralize.Singular(strings.ToLower(modelName)),
		"{modelNameInSmallCase}", strings.ToLower(modelName),
	)
	actualContent := replacementsStrings.Replace(oringinaleContent)
	os.MkdirAll("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"controllers", os.ModePerm)
	ioutil.WriteFile("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"controllers"+string(filepath.Separator)+moduleName+"Controller.go", []byte(actualContent), os.ModePerm)
}

func createDocs() {
	oringinaleContent := readFileContent("module" + string(filepath.Separator) + "docs" + string(filepath.Separator) + "module.postman.json")

	fields := "[\n\t\t\t\t\t\t"

	for _, field := range data {
		fields = fields + "{\n\t\t\t\t\t\t\t" + "\"key\": " + "\"" + field + "\"\n\t\t\t\t\t\t" + "\n\t\t\t\t\t\t}," + "\n\t\t\t\t\t\t"
	}
	// Remove last "," in fields

	fields = fields[:strings.LastIndex(fields, ",")] + fields[strings.LastIndex(fields, ",")+1:] + "]"
	replacementsStrings := strings.NewReplacer(
		"{moduleName}", moduleName,
		"{singleModuleName}", modelName,
		"\"{formData}\"", strings.Trim(fields, "\""),
	)

	actualContent := replacementsStrings.Replace(oringinaleContent)
	os.MkdirAll("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"docs", os.ModePerm)
	ioutil.WriteFile("modules"+string(filepath.Separator)+moduleName+string(filepath.Separator)+"docs"+string(filepath.Separator)+moduleName+".postman.json", []byte(actualContent), os.ModePerm)

}

func readFileContent(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
