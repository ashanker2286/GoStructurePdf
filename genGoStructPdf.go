package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"models/objects"
	"os"
	"strings"
)

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
	"LLDPGlobal":                        false,
	"LLDPGlobalState":                   false,
	"LLDPIntf":                          false,
	"LLDPIntfState":                     false,
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
	f.WriteString("============================================\n\n\n")
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

func constructRstFile() {
	autoCreateFlag := false
	autoDiscoverFlag := false
	for _, modelObjEnt := range ModelObj {
		for structName, structDetails := range modelObjEnt {
			f, err := os.Create(structName + "Objects.rst")
			check(err)
			f.WriteString(structName + " Model Objects\n")
			f.WriteString("============================================\n\n")
			if strings.Contains(structName, "State") {
				str := strings.TrimSuffix(structName, "State")
				f.WriteString("*state/" + str + "*\n")
			} else {
				f.WriteString("*config/" + structName + "*\n")
			}
			f.WriteString("------------------------------------\n\n")
			if structDetails[0].Multiplicity {
				f.WriteString("- Multiple of these objects can exist in a system.\n")
			} else {
				f.WriteString("- Only one of these object can exist in a system.\n")
			}

			for _, structDetail := range structDetails {
				if structDetail.IsKey {
					f.WriteString("- **" + structDetail.FieldName + "**\n")
					f.WriteString("\t- **Data Type**: " + structDetail.Type + "\n")
					f.WriteString("\t- **Description**: " + structDetail.Description + ".\n")
					if structDetail.IsDefaultSet {
						f.WriteString("\t- **Default**: " + structDetail.Default + "\n")
					}
					if structDetail.Selection != nil {
						f.WriteString("\t- **Possible Values**: ")
						for idx, val := range structDetail.Selection {
							f.WriteString(val)
							if idx != len(structDetail.Selection)-1 {
								f.WriteString(", ")
							} else {
								f.WriteString("\n")
							}
						}
					}
					f.WriteString("\t- This parameter is key element.\n")
					if autoDiscoverFlag == false && structDetail.AutoDiscover {
						autoDiscoverFlag = true
					}
					if autoCreateFlag == false && structDetail.AutoCreate {
						autoCreateFlag = true
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
			f.WriteString("**Flexswitch API Supported:**\n")
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
			f.WriteString("\n\n")
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
