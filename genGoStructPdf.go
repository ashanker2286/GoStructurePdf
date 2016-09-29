package main

import (
	"encoding/json"
	"fmt"
	//"github.com/jung-kurt/gofpdf"
	//"github.com/jung-kurt/gofpdf/internal/example"
	"io/ioutil"
	"models/objects"
	//"reflect"
	"os"
	"strings"
)

var listOfDaemon map[string]bool = map[string]bool{
	"opticd": true,
}

type ConfigObjJson struct {
	Owner         string   `json:"owner"`
	SrcFile       string   `json:"srcfile"`
	Access        string   `json:"access"`
	AutoCreate    bool     `json:"autoCreate"`
	AutoDiscover  bool     `json:"autoDiscover"`
	LinkedObjects []string `json:"linkedObjects"`
}

type StructDetails struct {
	FieldName    string
	Description  string
	Type         string
	IsKey        bool
	IsDefaultSet bool
	Default      string
}

type DaemonDetail map[string][]StructDetails

var ModelObj map[string]DaemonDetail

/*
func Example(pdf *gofpdf.Fpdf) {
	for daemonName, modelObjEnt := range ModelObj {
		for structName, structDetails := range modelObjEnt {
			pdf.SetFont("Arial", "B", 10)
			//pdf.Cell(20, 10, daemonName)
			//pdf.Ln(-1)
			pdf.Cell(20, 10, structName)
			pdf.SetFont("Arial", "", 8)
			for _, structDetail := range structDetails {
				fmt.Println("DaemonName:", daemonName, "StructName:", structName, "FieldName:", structDetail.FieldName, "DataType:", structDetail.Type)
				//fmt.Println(structDetail.FieldName)
				str := structDetail.FieldName + "( " + structDetail.Type + " )"
				pdf.Ln(-1)
				pdf.SetFont("Arial", "B", 8)
				pdf.Cell(20, 10, str)
				pdf.Ln(-1)
				//pdf.Cell(20, 10, structDetail.Type)
				//pdf.Cell(20, 10, steructDetail.Tag)
				pdf.SetFont("Arial", "", 8)
				pdf.MultiCell(0, 3, structDetail.Tag, "", "", false)
				//pdf.CellFormat(40, 6, structDetail.FieldName, "1", 0, "", false, 0, "")
				//pdf.CellFormat(60, 6, structDetail.Type, "1", 0, "", false, 0, "")
				//pdf.CellFormat(20, 6, structDetail.Type, "1", 0, "", false, 0, "")
			}

			pdf.Ln(-1)
			pdf.Ln(-1)
		}
	}
	//example.Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/basic.pdf
}
*/

type ParameterDetails struct {
	Type         string `json:"type"`
	IsKey        bool   `json:"isKey"`
	Description  string `json:"description"`
	Default      string `json:"default"`
	IsDefaultSet bool   `json:"isDefaultSet"`
}

/*
func reflectThis(intf interface{}, owner string, structName string) {
	//fmt.Println("Struct Name:", structName)
	val := reflect.TypeOf(intf).Elem()
	numOfField := val.NumField()
	//fmt.Println("Val: ", val)
	//fmt.Println("NumField: ", numOfField)
	//fmt.Println("Intf: ", intf)

	ModelObjEnt, _ := ModelObj[owner]
	structDetails, _ := ModelObjEnt[structName]
	for i := 1; i < numOfField; i++ {
		field := val.Field(i)
		fieldName := field.Name
		tag := field.Tag
		Type := field.Type
		//fmt.Println("Field Name:", fieldName, "Tag:", tag, "Type:", Type)
		structDetail := StructDetails{
			FieldName: fieldName,
			Tag:       string(tag),
			Type:      Type.String(),
		}
		structDetails = append(structDetails, structDetail)
	}
	ModelObjEnt[structName] = structDetails
	ModelObj[owner] = ModelObjEnt
}
*/

func generateParameterDetailList(structName string, owner string) {
	var paramDetailMap map[string]ParameterDetails
	paramDetailMap = make(map[string]ParameterDetails)
	fileName := "../../../reltools/codegentools/._genInfo/" + structName + "Members.json"
	fmt.Println("FileName: ", fileName)
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error in reading Object Configuration File", fileName)
		return
	}
	err = json.Unmarshal(bytes, &paramDetailMap)
	if err != nil {
		fmt.Println("Error Unmarshalling Object Configuration file", fileName)
		return
	}

	ModelObjEnt, _ := ModelObj[owner]
	structDetails, _ := ModelObjEnt[structName]
	for key, val := range paramDetailMap {
		structDetail := StructDetails{
			FieldName:    key,
			Type:         val.Type,
			Description:  val.Description,
			IsKey:        val.IsKey,
			IsDefaultSet: val.IsDefaultSet,
			Default:      val.Default,
		}
		structDetails = append(structDetails, structDetail)
	}
	ModelObjEnt[structName] = structDetails
	ModelObj[owner] = ModelObjEnt
}

func main() {
	var objMap map[string]ConfigObjJson

	objMap = make(map[string]ConfigObjJson)
	ModelObj = make(map[string]DaemonDetail)
	for key, val := range listOfDaemon {
		if val == true {
			ModelObjEnt, _ := ModelObj[key]
			ModelObjEnt = make(map[string][]StructDetails)
			ModelObj[key] = ModelObjEnt
		}
	}
	objConfigFile := "genObjectConfig.json"

	bytes, err := ioutil.ReadFile(objConfigFile)
	if err != nil {
		fmt.Println("Error in reading Object Configuration File", objConfigFile)
		return
	}
	err = json.Unmarshal(bytes, &objMap)
	if err != nil {
		fmt.Println("Error Unmarshalling Object Configuration file", objConfigFile)
		return
	}

	for key, val := range objMap {
		_, exist := listOfDaemon[val.Owner]
		if !exist {
			continue
		}

		_, exist = objects.GenConfigObjectMap[strings.ToLower(key)]
		if !exist {
			fmt.Println("Error finding given Object in GenConfigObjectMap")
			continue
		}
		generateParameterDetailList(key, val.Owner)
	}

	for daemonName, modelObjEnt := range ModelObj {
		f, err := os.Create(daemonName + "Objects.rst")
		check(err)
		f.WriteString(strings.ToUpper(daemonName) + " Structure\n")
		f.WriteString("============================================\n\n")
		f.WriteString("Objects\n")
		f.WriteString("---------------------------------------------------------\n\n")
		for structName, structDetails := range modelObjEnt {
			if strings.Contains(structName, "State") {
				str := strings.Trim(structName, "State")
				f.WriteString("*state/" + str + "*\n")
			} else {
				f.WriteString("*config/" + structName + "*\n")
			}
			f.WriteString("\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\n\n")
			for _, structDetail := range structDetails {
				if structDetail.IsKey {
					f.WriteString("- **" + structDetail.FieldName + "**\n")
					f.WriteString("\t- **Data Type**: " + structDetail.Type + "\n")
					f.WriteString("\t- **Description**: " + structDetail.Description + ".\n")
					f.WriteString("\t- This parameter is the key element.\n")
					if structDetail.IsDefaultSet {
						f.WriteString("\t- **Default**: " + structDetail.Default + "\n")
					}
				}
			}
			for _, structDetail := range structDetails {
				if !structDetail.IsKey {
					f.WriteString("- **" + structDetail.FieldName + "**\n")
					f.WriteString("\t- **Data Type**: " + structDetail.Type + "\n")
					f.WriteString("\t- **Description**: " + structDetail.Description + ".\n")
					if structDetail.IsDefaultSet {
						f.WriteString("\t- **Default**: " + structDetail.Default + "\n")
					}
				}
			}
			f.WriteString("\n\n")
		}
		f.Sync()
		f.Close()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
