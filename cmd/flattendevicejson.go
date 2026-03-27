/*
Copyright © 2026 Patrick Buick <patrick.buick@powerfleet.com>
This file is part of a CLI application for PowerFleet.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// From quicktype.io - would have saved a *LOT* of time and typing!
// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    items, err := UnmarshalItems(bytes)
//    bytes, err = items.Marshal()

type Items map[string]ItemValue

func UnmarshalItems(data []byte) (Items, error) {
	var r Items
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Items) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ItemValue struct {
	Device              Device     `json:"device"`
	Serial              string     `json:"serial"`
	Imei                string     `json:"imei"`
	Iccid               string     `json:"iccid"`
	Imsi                string     `json:"imsi"`
	VcmProtocol         string     `json:"vcmProtocol"`
	Vin                 string     `json:"vin"`
	Connected           bool       `json:"connected"`
	ConnectionDate      uint64     `json:"connectionDate"`
	CheckinDate         uint64     `json:"checkinDate"`
	TelemetryDate       uint64     `json:"telemetryDate"`
	PropertiesDate      uint64     `json:"propertiesDate"`
	CloudPropVer        uint64     `json:"cloudPropVer"`
	DevicePropVer       uint64     `json:"devicePropVer"`
	SynchedSettingsDate uint64     `json:"synchedSettingsDate"`
	SynchedSettingsVer  uint64     `json:"synchedSettingsVer"`
	CurrSettingsVer     uint64     `json:"currSettingsVer"`
	CurrSettingsDate    uint64     `json:"currSettingsDate"`
	BootFwVer           string     `json:"bootFwVer"`
	MainFwVer           string     `json:"mainFwVer"`
	BluetoothFwVer      string     `json:"bluetoothFwVer"`
	ModemFwVer          string     `json:"modemFwVer"`
	Scratchpad          Scratchpad `json:"scratchpad"`
	CAChain             string     `json:"caChain"`
	Make                string     `json:"make"`
	Model               string     `json:"model"`
	Year                string     `json:"year"`
	LastCmdDate         uint64     `json:"lastCmdDate"`
	BattPercent         uint64     `json:"battPercent"`
}

type Device struct {
	ID              string          `json:"id"`
	TenantID        string          `json:"tenantId"`
	ProductID       string          `json:"productId"`
	GroupID         string          `json:"groupId"`
	VcmProtocolID   string          `json:"vcmProtocolId"`
	VcmProtocol2ID  string          `json:"vcmProtocol2Id"`
	VcmVehicleID    string          `json:"vcmVehicleId"`
	VcmVehicle2ID   string          `json:"vcmVehicle2Id"`
	VcmStatic       bool            `json:"vcmStatic"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	State           string          `json:"state"`
	StateStartDate  uint64          `json:"stateStartDate"`
	ActivationDate  uint64          `json:"activationDate"`
	ModifiedDate    uint64          `json:"modifiedDate"`
	CreatedDate     uint64          `json:"createdDate"`
	Tags            Tags            `json:"tags"`
	PermissionIDS   []interface{}   `json:"permissionIds"`
	Commands        []interface{}   `json:"commands"`
	Settings        Settings        `json:"settings"`
	SettingsSynched SettingsSynched `json:"settingsSynched"`
	Properties      Properties      `json:"properties"`
}

type Properties struct {
	OwnerID      string         `json:"ownerId"`
	Version      uint64         `json:"version"`
	Data         PropertiesData `json:"data"`
	ModifiedDate uint64         `json:"modifiedDate"`
}

type PropertiesData struct {
	Apn                     string `json:"apn"`
	Ver                     string `json:"ver"`
	Vin                     string `json:"vin"`
	Imsi                    string `json:"imsi"`
	BTMAC                   string `json:"btMac"`
	GpsHw                   string `json:"gpsHw"`
	GpsSw                   string `json:"gpsSw"`
	Iccid                   string `json:"iccid"`
	BootVer                 string `json:"bootVer"`
	EcmOvrd                 string `json:"ecmOvrd"`
	MainVer                 string `json:"mainVer"`
	VcmType                 string `json:"vcmType"`
	WifiMAC                 string `json:"wifiMac"`
	BoardRev                string `json:"boardRev"`
	ModemVer                string `json:"modemVer"`
	CFGBitmask              string `json:"cfgBitmask"`
	ProductType             string `json:"productType"`
	VcmProtocol             string `json:"vcmProtocol"`
	BluetoothVer            string `json:"bluetoothVer"`
	HarshDriveOvrd          string `json:"harshDriveOvrd"`
	VehicleManufacturerInfo string `json:"vehicleManufacturerInfo"`
}

type Settings struct {
	OwnerID      string       `json:"ownerId"`
	Version      uint64       `json:"version"`
	Data         SettingsData `json:"data"`
	ModifiedDate uint64       `json:"modifiedDate"`
}

type SettingsData struct {
	Io             Io             `json:"io"`
	Eld            Eld            `json:"eld"`
	Ota            Ota            `json:"ota"`
	Ver            uint64         `json:"ver"`
	Wifi           Wifi           `json:"wifi"`
	Config         Config         `json:"config"`
	Motion         Motion         `json:"motion"`
	EXTTemp        EXTTemp        `json:"extTemp"`
	Vehicle        Vehicle        `json:"vehicle"`
	DriverID       DriverID       `json:"driverId"`
	Analytics      Analytics      `json:"analytics"`
	Bluetooth      Bluetooth      `json:"bluetooth"`
	Telemetry      Telemetry      `json:"telemetry"`
	DriverBehavior DriverBehavior `json:"driverBehavior"`
}

type Analytics struct {
	Feature uint64 `json:"feature"`
}

type Bluetooth struct {
	Scan         uint64 `json:"scan"`
	MACFilter    string `json:"macFilter"`
	NameFilter   string `json:"nameFilter"`
	BeaconFilter string `json:"beaconFilter"`
	MoveInterval uint64 `json:"moveInterval"`
	StopInterval uint64 `json:"stopInterval"`
}

type Config struct {
	Feature    uint64 `json:"feature"`
	MqttURL    string `json:"mqttUrl"`
	Interval   uint64 `json:"interval"`
	RootCAURL  string `json:"rootCAUrl"`
	ReportList string `json:"reportList"`
}

type DriverBehavior struct {
	SpdAdj              uint64 `json:"spdAdj"`
	HeavyImp            uint64 `json:"heavyImp"`
	LightImp            uint64 `json:"lightImp"`
	HeavyTurn           uint64 `json:"heavyTurn"`
	LightTurn           uint64 `json:"lightTurn"`
	HeavyAccel          uint64 `json:"heavyAccel"`
	HeavyBrake          uint64 `json:"heavyBrake"`
	LightAccel          uint64 `json:"lightAccel"`
	LightBrake          uint64 `json:"lightBrake"`
	HarshBuzzer         bool   `json:"harshBuzzer"`
	HarshConfig         uint64 `json:"harshConfig"`
	HeavyPeriod         uint64 `json:"heavyPeriod"`
	LightPeriod         uint64 `json:"lightPeriod"`
	DisableHarsh        bool   `json:"disableHarsh"`
	ImpactPeriod        uint64 `json:"impactPeriod"`
	SpdToTrigrWrnKmh    uint64 `json:"spdToTrigrWrnKmh"`
	SpdToTrigrWrnDurSEC uint64 `json:"spdToTrigrWrnDurSec"`
}

type DriverID struct {
	OffTime  uint64 `json:"offTime"`
	Timeout  uint64 `json:"timeout"`
	Reminder uint64 `json:"reminder"`
}

type EXTTemp struct {
	Always          bool    `json:"always"`
	DeltaRpt        float64 `json:"deltaRpt"`
	RptInterval     uint64  `json:"rptInterval"`
	CollectInterval uint64  `json:"collectInterval"`
}

type Eld struct {
	Enable  bool   `json:"enable"`
	MoveInt uint64 `json:"moveInt"`
	StopInt uint64 `json:"stopInt"`
}

type Io struct {
	Uart1CFG           string `json:"uart1Cfg"`
	Uart2CFG           string `json:"uart2Cfg"`
	DigIn1CFG          uint64 `json:"digIn1Cfg"`
	DigIn1Dbf          bool   `json:"digIn1Dbf"`
	DigIn2CFG          uint64 `json:"digIn2Cfg"`
	DigIn2Dbf          bool   `json:"digIn2Dbf"`
	DigIn3CFG          uint64 `json:"digIn3Cfg"`
	DigIn3Dbf          bool   `json:"digIn3Dbf"`
	DigIn4CFG          uint64 `json:"digIn4Cfg"`
	DigIn4Dbf          bool   `json:"digIn4Dbf"`
	SfoTxFreq          uint64 `json:"sfoTxFreq"`
	DigOut1CFG         uint64 `json:"digOut1Cfg"`
	DigOut2CFG         uint64 `json:"digOut2Cfg"`
	DigOut3CFG         uint64 `json:"digOut3Cfg"`
	DigOut4CFG         uint64 `json:"digOut4Cfg"`
	SfoCoverageTO      uint64 `json:"sfoCoverageTO"`
	DoorUnlockPulseCnt uint64 `json:"doorUnlockPulseCnt"`
	DoorUnlockPulseDur uint64 `json:"doorUnlockPulseDur"`
}

type Motion struct {
	CurvFit            bool   `json:"curvFit"`
	TargAcc            uint64 `json:"targAcc"`
	EnableHF           bool   `json:"enableHF"`
	TargAccTO          uint64 `json:"targAccTO"`
	CurvFitInt         uint64 `json:"curvFitInt"`
	BattLowMult        uint64 `json:"battLowMult"`
	BattMedMult        uint64 `json:"battMedMult"`
	CurvFitTrig        uint64 `json:"curvFitTrig"`
	BattRptCount       uint64 `json:"battRptCount"`
	StopInterval       uint64 `json:"stopInterval"`
	RecheckTmSecs      uint64 `json:"recheckTmSecs"`
	TelemInterval      uint64 `json:"telemInterval"`
	InterimStopInt     uint64 `json:"interimStopInt"`
	MoveDistThresh     uint64 `json:"moveDistThresh"`
	StopAccelWupThresh uint64 `json:"stopAccelWupThresh"`
}

type Ota struct {
	Feature           uint64 `json:"feature"`
	MainURL           string `json:"mainUrl"`
	MainVer           string `json:"mainVer"`
	ModemURL          string `json:"modemUrl"`
	ModemVer          string `json:"modemVer"`
	BluetoothURL      string `json:"bluetoothUrl"`
	BluetoothVer      string `json:"bluetoothVer"`
	EphemBaseURL      string `json:"ephemBaseUrl"`
	VcmVehicleLIBURL  string `json:"vcmVehicleLibUrl"`
	VcmProtocolLIBURL string `json:"vcmProtocolLibUrl"`
}

type Telemetry struct {
	Topic        string `json:"topic"`
	Feature      uint64 `json:"feature"`
	MqttURL      string `json:"mqttUrl"`
	RootCAURL    string `json:"rootCAUrl"`
	ReportList   string `json:"reportList"`
	ProvisionURL string `json:"provisionUrl"`
}

type Vehicle struct {
	VinOverride        string `json:"vinOverride"`
	CFGIgnThresh       uint64 `json:"cfgIgnThresh"`
	FuelTankSize       uint64 `json:"fuelTankSize"`
	LowVinCutoff       uint64 `json:"lowVinCutoff"`
	GpsOdomOffset      uint64 `json:"gpsOdomOffset"`
	DisableEVDetection bool   `json:"disableEVDetection"`
}

type Wifi struct {
	Apn        string `json:"apn"`
	Pwd        string `json:"pwd"`
	Auth       string `json:"auth"`
	SSID       string `json:"ssid"`
	ApnPwd     string `json:"apnPwd"`
	Encrypt    string `json:"encrypt"`
	ApnUname   string `json:"apnUname"`
	Broadcast  bool   `json:"broadcast"`
	EnableWifi bool   `json:"enableWifi"`
}

type SettingsSynched struct {
	OwnerID  string       `json:"ownerId"`
	Version  uint64       `json:"version"`
	Data     SettingsData `json:"data"`
	Errors   Tags         `json:"errors"`
	SyncDate uint64       `json:"syncDate"`
}

type Tags struct {
}

type Scratchpad struct {
	Ver                string  `json:"ver"`
	TotalFuel          uint64  `json:"totalFuel"`
	FeatureCFG         uint64  `json:"featureCfg"`
	GpsOdoCalc         float64 `json:"gpsOdoCalc"`
	EngineHrsCalc      float64 `json:"engineHrsCalc"`
	SwitchDisable      bool    `json:"switchDisable"`
	OdometerOffset     uint64  `json:"odometerOffset"`
	EngineHrsOffset    float64 `json:"engineHrsOffset"`
	OdometerAllowed    bool    `json:"odometerAllowed"`
	OdometerDerived    float64 `json:"odometerDerived"`
	IgnFollowsMotion   bool    `json:"ignFollowsMotion"`
	EnableHarshBuzzer  bool    `json:"enableHarshBuzzer"`
	TempConFreqExTime  uint64  `json:"tempConFreqExTime"`
	ShutdownExpiration uint64  `json:"shutdownExpiration"`
}

func printDevices(items Items) string {
	// For each "item" aka device parsed from the JSON
	count := 0
	last := len(items)
	fmt.Fprintf(os.Stderr, "I found %d items.", last)
	var IE = "\","
	var builder strings.Builder
	builder.WriteString("[\n")
	for _, value := range items {
		if count%20 == 0 {
			fmt.Fprintf(os.Stderr, "\rProcessing number %d", count)
		}
		// Start the "item"
		builder.WriteString("{\"device.id\": \"" + value.Device.ID + IE)
		builder.WriteString("\"device.tenantId\": \"" + value.Device.TenantID + IE)
		builder.WriteString("\"device.productId\": \"" + value.Device.ProductID + IE)
		builder.WriteString("\"device.groupId\": \"" + value.Device.GroupID + IE)
		builder.WriteString("\"device.vcmProtocolId\": \"" + value.Device.VcmProtocolID + IE)
		builder.WriteString("\"device.vcmProtocol2Id\": \"" + value.Device.VcmProtocol2ID + IE)
		builder.WriteString("\"device.vcmVehicleId\": \"" + value.Device.VcmVehicleID + IE)
		builder.WriteString("\"device.vcmVehicle2Id\": \"" + value.Device.VcmVehicle2ID + IE)
		builder.WriteString("\"device.vcmStatic\": " + strconv.FormatBool(value.Device.VcmStatic) + ",")
		builder.WriteString("\"device.name\": \"" + value.Device.Name + IE)
		builder.WriteString("\"device.description\": \"" + value.Device.Description + IE)
		builder.WriteString("\"device.state\": \"" + value.Device.State + IE)
		builder.WriteString("\"device.stateStartDate\": " + strconv.FormatUint(value.Device.StateStartDate, 10) + ",")
		builder.WriteString("\"device.activationDate\": " + strconv.FormatUint(value.Device.ActivationDate, 10) + ",")
		builder.WriteString("\"device.modifiedDate\": " + strconv.FormatUint(value.Device.ModifiedDate, 10) + ",")
		builder.WriteString("\"device.createdDate\": " + strconv.FormatUint(value.Device.CreatedDate, 10) + ",")
		builder.WriteString("\"device.permissionIds\": [")
		pcount := len(value.Device.PermissionIDS)
		for i, item := range value.Device.PermissionIDS {
			thing := fmt.Sprintf("%v", item)
			if i < pcount {
				builder.WriteString(thing + ",")
			} else {
				builder.WriteString(thing)
			}
		}
		builder.WriteString("],")
		builder.WriteString("\"device.commands\": [")
		ccount := len(value.Device.Commands)
		for i, item := range value.Device.Commands {
			thing := fmt.Sprintf("%v", item)
			if i < ccount {
				builder.WriteString(thing + ",")
			} else {
				builder.WriteString(thing)
			}
		}
		builder.WriteString("],")
		builder.WriteString("\"device.settings.ownerId\": \"" + value.Device.Settings.OwnerID + IE)
		builder.WriteString("\"device.settings.version\": " + strconv.FormatUint(value.Device.Settings.Version, 10) + ",")
		builder.WriteString("\"device.settings.data.io.uart1Cfg\": \"" + value.Device.Settings.Data.Io.Uart1CFG + IE)
		builder.WriteString("\"device.settings.data.io.uart2Cfg\": \"" + value.Device.Settings.Data.Io.Uart2CFG + IE)
		builder.WriteString("\"device.settings.data.io.digIn1Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigIn1CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digIn1Dbf\": " + strconv.FormatBool(value.Device.Settings.Data.Io.DigIn1Dbf) + ",")
		builder.WriteString("\"device.settings.data.io.digIn2Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigIn2CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digIn2Dbf\": " + strconv.FormatBool(value.Device.Settings.Data.Io.DigIn2Dbf) + ",")
		builder.WriteString("\"device.settings.data.io.digIn3Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigIn3CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digIn3Dbf\": " + strconv.FormatBool(value.Device.Settings.Data.Io.DigIn3Dbf) + ",")
		builder.WriteString("\"device.settings.data.io.digIn4Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigIn4CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digIn4Dbf\": " + strconv.FormatBool(value.Device.Settings.Data.Io.DigIn4Dbf) + ",")
		builder.WriteString("\"device.settings.data.io.sfoTxFreq\": " + strconv.FormatUint(value.Device.Settings.Data.Io.SfoTxFreq, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digOut1Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigOut1CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digOut2Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigOut2CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digOut3Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigOut3CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.digOut4Cfg\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DigOut4CFG, 10) + ",")
		builder.WriteString("\"device.settings.data.io.sfoCoverateTO\": " + strconv.FormatUint(value.Device.Settings.Data.Io.SfoCoverageTO, 10) + ",")
		builder.WriteString("\"device.settings.data.io.doorUnlockPulseCnt\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DoorUnlockPulseCnt, 10) + ",")
		builder.WriteString("\"device.settings.data.io.doorUnlockPulseDur\": " + strconv.FormatUint(value.Device.Settings.Data.Io.DoorUnlockPulseDur, 10) + ",")
		builder.WriteString("\"device.settings.data.eld.enable\": " + strconv.FormatBool(value.Device.Settings.Data.Eld.Enable) + ",")
		builder.WriteString("\"device.settings.data.eld.moveInt\": " + strconv.FormatUint(value.Device.Settings.Data.Eld.MoveInt, 10) + ",")
		builder.WriteString("\"device.settings.data.eld.stopInt\": " + strconv.FormatUint(value.Device.Settings.Data.Eld.StopInt, 10) + ",")
		builder.WriteString("\"device.settings.data.ota.feature\": " + strconv.FormatUint(value.Device.Settings.Data.Ota.Feature, 10) + ",")
		builder.WriteString("\"device.settings.data.ota.mainUrl\": \"" + value.Device.Settings.Data.Ota.MainURL + IE)
		builder.WriteString("\"device.settings.data.ota.mainVer\": \"" + value.Device.Settings.Data.Ota.MainVer + IE)
		builder.WriteString("\"device.settings.data.ota.modemUrl\": \"" + value.Device.Settings.Data.Ota.ModemURL + IE)
		builder.WriteString("\"device.settings.data.ota.modemVer\": \"" + value.Device.Settings.Data.Ota.ModemVer + IE)
		builder.WriteString("\"device.settings.data.ota.bluetoothUrl\": \"" + value.Device.Settings.Data.Ota.BluetoothURL + IE)
		builder.WriteString("\"device.settings.data.ota.bluetoothVer\": \"" + value.Device.Settings.Data.Ota.BluetoothVer + IE)
		builder.WriteString("\"device.settings.data.ota.ephemBaseUrl\": \"" + value.Device.Settings.Data.Ota.EphemBaseURL + IE)
		builder.WriteString("\"device.settings.data.ota.vcmVehicleLibUrl\": \"" + value.Device.Settings.Data.Ota.VcmVehicleLIBURL + IE)
		builder.WriteString("\"device.settings.data.ota.vcmProtocolLibUrl\": \"" + value.Device.Settings.Data.Ota.VcmProtocolLIBURL + IE)
		builder.WriteString("\"device.settings.data.ver\": " + strconv.FormatUint(value.Device.Settings.Data.Ver, 10) + ",")

		//TODO: Finish the rest of the output

		builder.WriteString("}")
		count++
		// End the "item"
		if count != last {
			builder.WriteString(",\n")
		}
	}
	builder.WriteString("\n]")
	fmt.Fprintf(os.Stderr, "\r                                                       \r")
	return builder.String()
}

// flattendevicejsonCmd represents the flattendevicejson command
var flattendevicejsonCmd = &cobra.Command{
	Use:   "flattendevicejson",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: I need to define the structures for the PUI JSON format coming out of the queries I do. *DONE*
		// TODO: Then I need to read in the file, the content is all on one line! I need to split it?
		// TODO:   Nope, just define a top-level struct that contains an undefined number of units as properties? It's not an array!
		// TODO:   Can't these people do "normal" JSON?
		// TODO: It starts with a { but after that each entry starts with a name, then the rest contained in {} and ending with a comma.
		// TODO: Then I need to parse the JSON to the nested structure and add it to the collection.
		// TODO: Then when I have enough units, write them out and empty the "collection".
		// TODO: When done, finish and let the user know.
		// TODO: An idea - can I catch if they use - as a parameter, perhaps the filename and
		//   collect from STDIN and dump to STDOUT vs specifying filenames?
		if filename, err := cmd.Flags().GetString("filename"); err == nil {
			if filename == "-" {
				fmt.Fprintf(os.Stderr, "flattendevicejson using STDIN and STDOUT")
				scanner := bufio.NewScanner(os.Stdin)
				// TODO: Remove this when it is production code or this gets inserted into the output stream
				//fmt.Println("Enter text. Use CTRL+Z and enter for windows or CTRL+D on Linux/Unix/MacOS or CTRL+C to stop:")

				for scanner.Scan() {
					line := scanner.Bytes()
					// TODO: Remove this for production code as we only want to output the processed JSON
					//fmt.Printf("Received: %s\n", line)
					// TODO: This is where my processing code is called / placed for handling STDIN to STDOUT
					items, _ := UnmarshalItems(line)
					//bytes, _ := items.Marshal()
					//fmt.Print(string(bytes))
					//fmt.Printf("%+v", items)
					for key, value := range items {
						fmt.Fprintf(os.Stderr, "%s has imei %s\n", key, value.Imei)
						//fmt.Printf("IMEI is %s", items)
					}
				}

				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "reading standard input:", err)
				}
			} else {
				fmt.Fprintf(os.Stderr, "flattendevicejson called with filename: %s\n", filename)
				// Filename is not "-", so presumably a filename.
				file, err := os.Open(filename)
				if errors.Is(err, os.ErrNotExist) {
					// Handle the case where the file does not exist
					fmt.Fprintln(os.Stderr, "File not found, handling the case")
				} else if err != nil {
					// Handle other potential errors
					fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
				} else {
					// File was opened successfully
					scanner := bufio.NewScanner(file)
					defer file.Close() // Remember to close the file
					fmt.Fprintln(os.Stderr, "File opened successfully")
					// Up the buffer size to 70MB
					const maxCapacity = 70 * 1024 * 1024 // 10 MB
					buf := make([]byte, maxCapacity)
					scanner.Buffer(buf, maxCapacity)
					for scanner.Scan() {
						line := scanner.Bytes()
						items, _ := UnmarshalItems(line)
						// TODO: I have to write my print function to print the standard "item" aka device JSON to string.
						// TODO: Then I can either spit the string out STDOUT or to file depending on which part of this code I am in.
						fmt.Print(printDevices(items))
						// for key, value := range items {
						// 	fmt.Printf("%s has imei %s\n", key, value.Imei)
						// }
					}
					fmt.Fprintln(os.Stderr, "Finished processing")
					if err := scanner.Err(); err != nil {
						fmt.Fprintln(os.Stderr, "reading file:", err)
					}
				}

				// TODO: Check if the file exists, is readable and open it.
				// TODO: Read the file line-by-line and process it exactly the same as for the STDIN - STDOUT mode.
				// TODO: Make the output file name just have "out" or "flat" somewhere in it.
			}
		} else {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(flattendevicejsonCmd)

	flattendevicejsonCmd.Flags().String("filename", "", "Provide a filename with JSON to flatten, - to use pipes.")

}
