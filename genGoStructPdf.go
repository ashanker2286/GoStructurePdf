package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"models/objects"
	"os"
	"strings"
)

var codeHeader string = `
::

	import sys
	import os
	from flexswitchV2 import FlexSwitch

	if __name__ == '__main__':
		switchIP := "192.168.56.101"
		swtch = FlexSwitch (switchIP, 8080)  # Instantiate object to talk to flexSwitch
`

var codeTail string = `
		if error != None: #Error not being None implies there is some problem
			print error
		else :
			print 'Success'
`

var ObjMap map[string]bool = map[string]bool{
	"Acl":                               false,
	"AclRule":                           false,
	"AclRuleState":                      false,
	"AclState":                          false,
	"AlarmState":                        true,
	"ApiInfoState":                      true,
	"ArpEntryHwState":                   false,
	"ArpEntryState":                     false,
	"ArpGlobal":                         false,
	"ArpLinuxEntryState":                false,
	"AsicGlobalPM":                      true,
	"AsicGlobalPMState":                 true,
	"AsicGlobalState":                   true,
	"BGPCounters":                       false,
	"BGPGlobal":                         false,
	"BGPGlobalState":                    false,
	"BGPMessages":                       false,
	"BGPPolicyAction":                   false,
	"BGPPolicyActionState":              false,
	"BGPPolicyCondition":                false,
	"BGPPolicyConditionState":           false,
	"BGPPolicyDefinition":               false,
	"BGPPolicyDefinitionState":          false,
	"BGPPolicyDefinitionStmtPrecedence": false,
	"BGPPolicyStmt":                     false,
	"BGPPolicyStmtState":                false,
	"BGPQueues":                         false,
	"BGPv4Aggregate":                    false,
	"BGPv4Neighbor":                     false,
	"BGPv4NeighborState":                false,
	"BGPv4PeerGroup":                    false,
	"BGPv4RouteState":                   false,
	"BGPv6Aggregate":                    false,
	"BGPv6Neighbor":                     false,
	"BGPv6NeighborState":                false,
	"BGPv6PeerGroup":                    false,
	"BGPv6RouteState":                   false,
	"BfdGlobal":                         false,
	"BfdGlobalState":                    false,
	"BfdSession":                        false,
	"BfdSessionParam":                   false,
	"BfdSessionParamState":              false,
	"BfdSessionState":                   false,
	"BufferGlobalStatState":             false,
	"BufferPortStatState":               false,
	"ComponentLogging":                  true,
	"ConfigLogState":                    true,
	"CoppStatState":                     false,
	"DWDMModule":                        true,
	"DWDMModuleClntIntf":                true,
	"DWDMModuleClntIntfState":           true,
	"DWDMModuleNwIntf":                  true,
	"DWDMModuleNwIntfPMState":           true,
	"DWDMModuleNwIntfState":             true,
	"DWDMModulePMData":                  true,
	"DWDMModuleState":                   true,
	"DaemonState":                       true,
	"DhcpGlobalConfig":                  false,
	"DhcpIntfConfig":                    false,
	"DhcpRelayGlobal":                   false,
	"DhcpRelayHostDhcpState":            false,
	"DhcpRelayIntf":                     false,
	"DhcpRelayIntfServerState":          false,
	"DhcpRelayIntfState":                false,
	"DistributedRelay":                  false,
	"DistributedRelayState":             false,
	"EthernetPM":                        true,
	"EthernetPMState":                   true,
	"FMgrGlobal":                        true,
	"Fan":                               false,
	"FanSensor":                         true,
	"FanSensorPMData":                   true,
	"FanSensorPMDataState":              true,
	"FanSensorState":                    true,
	"FanState":                          false,
	"FaultState":                        true,
	"IPv4Intf":                          false,
	"IPv4IntfState":                     false,
	"IPv4Route":                         false,
	"IPv4RouteHwState":                  false,
	"IPv4RouteState":                    false,
	"IPv6Intf":                          false,
	"IPv6IntfState":                     false,
	"IPv6Route":                         false,
	"IPv6RouteHwState":                  false,
	"IPv6RouteState":                    false,
	"IpTableAcl":                        false,
	"IppLinkState":                      false,
	"IsisGlobal":                        false,
	"IsisGlobalState":                   false,
	"LLDPGlobal":                        true,
	"LLDPGlobalState":                   true,
	"LLDPIntf":                          true,
	"LLDPIntfState":                     true,
	"LaPortChannel":                     false,
	"LaPortChannelIntfRefListState":     false,
	"LaPortChannelState":                false,
	"LacpGlobal":                        false,
	"LacpGlobalState":                   false,
	"Led":                               false,
	"LedState":                          false,
	"LinkScopeIpState":                  false,
	"LogicalIntf":                       false,
	"LogicalIntfState":                  false,
	"MacTableEntryState":                true,
	"NDPEntryState":                     false,
	"NDPGlobal":                         false,
	"NDPGlobalState":                    false,
	"NDPIntfState":                      false,
	"NdpEntryHwState":                   false,
	"NeighborEntry":                     false,
	"NextBestRouteInfo":                 false,
	"NextHopInfo":                       false,
	"NotifierEnable":                    true,
	"OspfAreaEntry":                     false,
	"OspfAreaEntryState":                false,
	"OspfEventState":                    false,
	"OspfGlobal":                        false,
	"OspfGlobalState":                   false,
	"OspfIPv4RouteState":                false,
	"OspfIfEntry":                       false,
	"OspfIfEntryState":                  false,
	"OspfIfMetricEntry":                 false,
	"OspfLsaKey":                        false,
	"OspfLsdbEntryState":                false,
	"OspfNbrEntryState":                 false,
	"OspfNextHop":                       false,
	"OspfVirtIfEntry":                   false,
	"OspfVirtNbrEntryState":             false,
	"PMData":                            true,
	"PathInfo":                          false,
	"PerProtocolRouteCount":             false,
	"PlatformMgmtDeviceState":           true,
	"PlatformState":                     false,
	"PolicyCondition":                   false,
	"PolicyConditionState":              false,
	"PolicyDefinition":                  false,
	"PolicyDefinitionState":             false,
	"PolicyDefinitionStmtPriority":      false,
	"PolicyPrefix":                      false,
	"PolicyPrefixSet":                   false,
	"PolicyPrefixSetState":              false,
	"PolicyStmt":                        false,
	"PolicyStmtState":                   false,
	"Port":                              true,
	"PortState":                         true,
	"PowerConverterSensor":              true,
	"PowerConverterSensorPMData":        true,
	"PowerConverterSensorPMDataState":   true,
	"PowerConverterSensorState":         true,
	"Psu":                         false,
	"PsuState":                    false,
	"Qsfp":                        true,
	"QsfpChannel":                 true,
	"QsfpChannelPMData":           true,
	"QsfpChannelPMDataState":      true,
	"QsfpChannelState":            true,
	"QsfpPMData":                  true,
	"QsfpPMDataState":             true,
	"QsfpState":                   true,
	"RIBEventState":               false,
	"RepoInfo":                    true,
	"RouteDistanceState":          false,
	"RouteInfoSummary":            false,
	"RouteStatState":              false,
	"RouteStatsPerInterfaceState": false,
	"RouteStatsPerProtocolState":  false,
	"Sfp":                          false,
	"SfpState":                     false,
	"SourcePolicyList":             false,
	"StpBridgeInstance":            false,
	"StpBridgeInstanceState":       false,
	"StpGlobal":                    false,
	"StpPort":                      false,
	"StpPortState":                 false,
	"SubIPv4Intf":                  false,
	"SubIPv6Intf":                  false,
	"SystemLogging":                true,
	"SystemParam":                  true,
	"SystemParamState":             true,
	"SystemStatusState":            true,
	"SystemSwVersionState":         true,
	"TemperatureSensor":            true,
	"TemperatureSensorPMData":      true,
	"TemperatureSensorPMDataState": true,
	"TemperatureSensorState":       true,
	"ThermalState":                 false,
	"Vlan":                         true,
	"VlanState":                    true,
	"VoltageSensor":                true,
	"VoltageSensorPMData":          true,
	"VoltageSensorPMDataState":     true,
	"VoltageSensorState":           true,
	"VrrpIntf":                     false,
	"VrrpIntfState":                false,
	"VrrpVridState":                false,
	"VxlanInstance":                false,
	"VxlanVtepInstance":            false,
	"XponderGlobal":                true,
	"XponderGlobalState":           true,
}

type ConfigObjJson struct {
	Owner         string   `json:"owner"`
	SrcFile       string   `json:"srcfile"`
	Access        string   `json:"access"`
	AutoCreate    bool     `json:"autoCreate"`
	AutoDiscover  bool     `json:"autoDiscover"`
	LinkedObjects []string `json:"linkedObjects"`
	Multiplicity  string   `json:"multiplicity"`
}

type StructDetails struct {
	FieldName    string
	Description  string
	Type         string
	IsKey        bool
	IsDefaultSet bool
	Default      string
	Selection    []string
	AutoDiscover bool
	AutoCreate   bool
	Multiplicity bool
}

type DaemonDetail map[string][]StructDetails

var ModelObj map[string]DaemonDetail

type ParameterDetails struct {
	Type         string   `json:"type"`
	IsKey        bool     `json:"isKey"`
	Description  string   `json:"description"`
	Default      string   `json:"default"`
	IsDefaultSet bool     `json:"isDefaultSet"`
	Selection    []string `json:"selections"`
	AutoDiscover bool     `json:"autoDiscover"`
	AutoCreate   bool     `json:"autoCreate"`
}

func generateParameterDetailList(structName string, owner string, Multiplicity string) {
	var paramDetailMap map[string]ParameterDetails
	paramDetailMap = make(map[string]ParameterDetails)
	fileName := "../../../reltools/codegentools/._genInfo/" + structName + "Members.json"
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
			AutoDiscover: val.AutoDiscover,
			AutoCreate:   val.AutoCreate,
		}
		if val.Selection != nil {
			structDetail.Selection = append(structDetail.Selection, val.Selection...)
		}
		if strings.ContainsAny(Multiplicity, "*") {
			structDetail.Multiplicity = true
		}
		structDetails = append(structDetails, structDetail)
	}
	ModelObjEnt[structName] = structDetails
	ModelObj[owner] = ModelObjEnt
}

func generateReqdObjMap(reqdObjMap *map[string]bool) {
	for key, val := range ObjMap {
		if val {
			(*reqdObjMap)[key] = true
		}
	}

}

func generateListOfDaemon(listOfDaemon, reqdObjMap *map[string]bool, objMap *map[string]ConfigObjJson) {
	for key, val := range *objMap {
		if (*reqdObjMap)[key] {
			(*listOfDaemon)[val.Owner] = true
		}
	}
}

func allocateModelObj(ModelObj *map[string]DaemonDetail, listOfDaemon *map[string]bool) {
	for key, val := range *listOfDaemon {
		if val == true {
			ModelObjEnt, _ := (*ModelObj)[key]
			ModelObjEnt = make(map[string][]StructDetails)
			(*ModelObj)[key] = ModelObjEnt
		}
	}

}

func generateModelObjectRstFile(listOfDaemon *map[string]bool) {
	f, err := os.Create("modelObjects.rst")
	check(err)
	f.WriteString("FlexSwitch Model Objects\n")
	f.WriteString("================================================================\n\n\n")
	f.WriteString(".. toctree::\n")
	f.WriteString("   :maxdepth: 1\n\n")
	for _, modelObjEnt := range ModelObj {
		for structName, _ := range modelObjEnt {
			rstFile := structName + "Objects.rst"
			f.WriteString("   " + structName + "  <" + rstFile + ">\n")
		}
	}
	f.Sync()
	f.Close()
}

func createParameterDescTable(structName string, structDetails []StructDetails, f *os.File) (bool, bool) {
	autoCreateFlag := false
	autoDiscoverFlag := false
	table := tablewriter.NewWriter(f)
	table.SetAutoFormatHeaders(true)
	table.SetHeader([]string{"**Parameter Name**", "**Data Type**", "**Description**", "**Default**", "**Valid Values**"})
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetHeaderLine(true)
	table.SetColWidth(30)
	data := [][]string{}
	for _, structDetail := range structDetails {
		val := []string{}
		if structDetail.IsKey {
			str := structDetail.FieldName
			str += " **[KEY]**"
			val = append(val, str)
			val = append(val, structDetail.Type)
			val = append(val, structDetail.Description)
			if structDetail.IsDefaultSet {
				val = append(val, structDetail.Default)
			} else {
				val = append(val, "N/A")
			}
			if structDetail.Selection != nil {
				var str string
				for idx := 0; idx < len(structDetail.Selection); idx++ {
					str = str + structDetail.Selection[idx]
					if idx != len(structDetail.Selection)-1 {
						str += ", "
					}
				}
				val = append(val, str)
			} else {
				val = append(val, "N/A")
			}
			data = append(data, val)
			if autoDiscoverFlag == false && structDetail.AutoDiscover {
				autoDiscoverFlag = true
			}
			if autoCreateFlag == false && structDetail.AutoCreate {
				autoCreateFlag = true
			}
		}
	}
	for _, structDetail := range structDetails {
		val := []string{}
		if !structDetail.IsKey {
			val = append(val, structDetail.FieldName)
			val = append(val, structDetail.Type)
			val = append(val, structDetail.Description)
			if structDetail.IsDefaultSet {
				val = append(val, structDetail.Default)
			} else {
				val = append(val, "N/A")
			}
			if structDetail.Selection != nil {
				var str string
				for idx := 0; idx < len(structDetail.Selection); idx++ {
					str = str + structDetail.Selection[idx]
					if idx != len(structDetail.Selection)-1 {
						str += ", "
					}
				}
				val = append(val, str)
			} else {
				val = append(val, "N/A")
			}
			data = append(data, val)
		}
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return autoDiscoverFlag, autoCreateFlag
}

func WriteCurlCommands(structName string, structDetails []StructDetails, f *os.File, autoDiscoverFlag bool, autoCreateFlag bool) {
	if strings.Contains(structName, "State") {
		str := strings.TrimSuffix(structName, "State")
		f.WriteString("\t- GET By Key\n")
		f.WriteString("\t\t curl -X GET -H 'Content-Type: application/json' --header 'Accept: application/json' -d '{<Model Object as json-Data>}' http://device-management-IP:8080/public/v1/state/" + str + "\n")
		if structDetails[0].Multiplicity {
			f.WriteString("\t- GET ALL\n")
			f.WriteString("\t\t curl -X GET http://device-management-IP:8080/public/v1/state/" + str + "?CurrentMarker=<x>&Count=<y>\n")
		}
	} else {
		f.WriteString("\t- GET By Key\n")
		f.WriteString("\t\t curl -X GET -H 'Content-Type: application/json' --header 'Accept: application/json' -d '{<Model Object as json-Data>}' http://device-management-IP:8080/public/v1/config/" + structName + "\n")
		f.WriteString("\t- GET By ID\n")
		f.WriteString("\t\t curl -X GET http://device-management-IP:8080/public/v1/config/" + structName + "/<uuid>\n")
		if structDetails[0].Multiplicity {
			f.WriteString("\t- GET ALL\n")
			f.WriteString("\t\t curl -X GET http://device-management-IP:8080/public/v1/config/" + structName + "?CurrentMarker=<x>&Count=<y>\n")
		}
		fmt.Println("StructName:", structName, "AutoCreate:", structDetails[0].AutoCreate, "AutoDiscover:", structDetails[0].AutoDiscover)
		if !autoCreateFlag && autoDiscoverFlag {
			f.WriteString("\t- CREATE(POST)\n")
			f.WriteString("\t\t curl -X POST -H 'Content-Type: application/json' --header 'Accept: application/json' -d '{<Model Object as json-Data>}' http://device-management-IP:8080/public/v1/config/" + structName + "\n")
			f.WriteString("\t- DELETE By Key\n")
			f.WriteString("\t\t curl -X DELETE -i -H 'Accept:application/json' -d '{<Model Object as json data>}' http://device-management-IP:8080/public/v1/config/" + structName + "\n")
			f.WriteString("\t- DELETE By ID\n")
			f.WriteString("\t\t curl -X DELETE http://device-management-IP:8080/public/v1/config/" + structName + "<uuid>\n")
		}
		f.WriteString("\t- UPDATE(PATCH) By Key\n")
		f.WriteString("\t\t curl -X PATCH -H 'Content-Type: application/json' -d '{<Model Object as json data>}'  http://device-management-IP:8080/public/v1/config/" + structName + "\n")
		f.WriteString("\t- UPDATE(PATCH) By ID\n")
		f.WriteString("\t\t curl -X PATCH -H 'Content-Type: application/json' -d '{<Model Object as json data>}'  http://device-management-IP:8080/public/v1/config/" + structName + "<uuid>\n")
	}

}

func WriteCodeExample(structName string, structDetails []StructDetails, f *os.File, autoDiscoverFlag bool, autoCreateFlag bool) {
	f.WriteString("- **GET**\n\n")
	f.WriteString(codeHeader)
	f.WriteString("\t\tresponse, error = swtch.get" + structName + "(")
	var arg []string
	for _, structDetail := range structDetails {
		if structDetail.IsKey {
			arg = append(arg, structDetail.FieldName)
		}
	}
	for idx, data := range arg {
		f.WriteString(data + "=" + strings.ToLower(data))
		if idx != len(arg)-1 {
			f.WriteString(", ")
		} else {
			f.WriteString(")")
		}
	}
	f.WriteString("\n")
	f.WriteString(codeTail)

	f.WriteString("\n\n")
	f.WriteString("- **GET By ID**\n\n")
	f.WriteString(codeHeader)
	//str = strings.TrimSuffix(structName, "State")
	f.WriteString("\t\tresponse, error = swtch.get" + structName + "ById(ObjectId=objectid)")
	f.WriteString("\n")
	f.WriteString(codeTail)
	f.WriteString("\n\n")

	f.WriteString("\n\n")
	f.WriteString("- **GET ALL**\n\n")
	f.WriteString(codeHeader)
	//str = strings.TrimSuffix(structName, "State")
	f.WriteString("\t\tresponse, error = swtch.getAll" + structName + "s()")
	f.WriteString("\n")
	f.WriteString(codeTail)
	f.WriteString("\n\n")
	if !strings.Contains(structName, "State") {
		if !autoCreateFlag && autoDiscoverFlag {
			f.WriteString("- **CREATE**\n")
			f.WriteString(codeHeader)
			f.WriteString("\t\tresponse, error = swtch.create" + structName + "(")
			var arg []string
			for _, structDetail := range structDetails {
				if structDetail.IsKey {
					arg = append(arg, structDetail.FieldName)
				}
			}
			for _, structDetail := range structDetails {
				if !structDetail.IsKey {
					arg = append(arg, structDetail.FieldName)
				}
			}
			for idx, data := range arg {
				f.WriteString(data + "=" + strings.ToLower(data))
				if idx != len(arg)-1 {
					f.WriteString(", ")
				} else {
					f.WriteString(")")
				}
			}
			f.WriteString("\n")
			f.WriteString(codeTail)

			f.WriteString("\n\n")
			f.WriteString("- **DELETE**\n")
			f.WriteString(codeHeader)
			f.WriteString("\t\tresponse, error = swtch.delete" + structName + "(")
			arg = []string{}
			for _, structDetail := range structDetails {
				if structDetail.IsKey {
					arg = append(arg, structDetail.FieldName)
				}
			}
			for idx, data := range arg {
				f.WriteString(data + "=" + strings.ToLower(data))
				if idx != len(arg)-1 {
					f.WriteString(", ")
				} else {
					f.WriteString(")")
				}
			}
			f.WriteString("\n")
			f.WriteString(codeTail)

			f.WriteString("\n\n")
			f.WriteString("- **DELETE By ID**\n")
			f.WriteString(codeHeader)
			f.WriteString("\t\tresponse, error = swtch.delete" + structName + "ById(ObjectId=objectid")
			f.WriteString("\n")
			f.WriteString(codeTail)

		}

		f.WriteString("\n\n")
		f.WriteString("- **UPDATE**\n")
		f.WriteString(codeHeader)
		f.WriteString("\t\tresponse, error = swtch.update" + structName + "(")
		var arg []string
		for _, structDetail := range structDetails {
			if structDetail.IsKey {
				arg = append(arg, structDetail.FieldName)
			}
		}
		for _, structDetail := range structDetails {
			if !structDetail.IsKey {
				arg = append(arg, structDetail.FieldName)
			}
		}
		for idx, data := range arg {
			f.WriteString(data + "=" + strings.ToLower(data))
			if idx != len(arg)-1 {
				f.WriteString(", ")
			} else {
				f.WriteString(")")
			}
		}
		f.WriteString("\n")
		f.WriteString(codeTail)

		f.WriteString("\n\n")
		f.WriteString("- **UPDATE By ID**\n")
		f.WriteString(codeHeader)
		f.WriteString("\t\tresponse, error = swtch.update" + structName + "ById(ObjectId=objectid")
		arg = []string{}
		for _, structDetail := range structDetails {
			if !structDetail.IsKey {
				arg = append(arg, structDetail.FieldName)
			}
		}
		for idx, data := range arg {
			f.WriteString(data + "=" + strings.ToLower(data))
			if idx != len(arg)-1 {
				f.WriteString(", ")
			} else {
				f.WriteString(")")
			}
		}
		f.WriteString("\n")
		f.WriteString(codeTail)
	}
}

func constructRstFile() {
	for _, modelObjEnt := range ModelObj {
		for structName, structDetails := range modelObjEnt {
			f, err := os.Create(structName + "Objects.rst")
			check(err)
			f.WriteString(structName + " Model Objects\n")
			f.WriteString("=============================================================\n\n")
			if strings.Contains(structName, "State") {
				str := strings.TrimSuffix(structName, "State")
				f.WriteString("*state/" + str + "*\n")
			} else {
				f.WriteString("*config/" + structName + "*\n")
			}
			f.WriteString("------------------------------------\n\n")
			if structDetails[0].Multiplicity {
				f.WriteString("- Multiple objects of this type can exist in a system.\n\n")
			} else {
				f.WriteString("- Only one object of this type can exist in a system.\n\n")
			}

			autoDiscoverFlag, autoCreateFlag := createParameterDescTable(structName, structDetails, f)

			f.WriteString("\n")

			f.WriteString("\n\n")

			f.WriteString("**FlexSwitch CURL API Supported:**\n")
			WriteCurlCommands(structName, structDetails, f, autoDiscoverFlag, autoCreateFlag)
			f.WriteString("\n\n")
			f.WriteString("**FlexSwitch SDK API Supported:**\n")
			f.WriteString("\n\n")

			WriteCodeExample(structName, structDetails, f, autoDiscoverFlag, autoCreateFlag)
			f.Sync()
			f.Close()
		}
	}
}

func main() {
	var objMap map[string]ConfigObjJson

	objMap = make(map[string]ConfigObjJson)
	objConfigFile := "../models/objects/genObjectConfig.json"

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

	reqdObjMap := make(map[string]bool)
	generateReqdObjMap(&reqdObjMap)

	listOfDaemon := make(map[string]bool)
	generateListOfDaemon(&listOfDaemon, &reqdObjMap, &objMap)

	ModelObj = make(map[string]DaemonDetail)
	allocateModelObj(&ModelObj, &listOfDaemon)

	for key, val := range objMap {
		_, exist := reqdObjMap[key]
		if !exist {
			continue
		}

		_, exist = objects.GenConfigObjectMap[strings.ToLower(key)]
		if !exist {
			fmt.Println("Error finding given Object in GenConfigObjectMap")
			continue
		}
		generateParameterDetailList(key, val.Owner, val.Multiplicity)
	}

	generateModelObjectRstFile(&listOfDaemon)
	constructRstFile()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
