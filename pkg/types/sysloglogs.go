// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     NetDevice.avsc
 *     User.avsc
 *     campaign_finance.avsc
 *     clickstream.avsc
 *     clickstream_codes.avsc
 *     clickstream_users.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     orders.avsc
 *     pageviews.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customers.avsc
 *     shoe_orders.avsc
 *     siem_logs.avsc
 *     stockTrades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Sysloglogs struct {
	Name string `json:"name"`

	Type string `json:"type"`

	Message string `json:"message"`

	Host string `json:"host"`

	Version string `json:"version"`

	Tag string `json:"tag"`

	Level int32 `json:"level"`

	Facility string `json:"facility"`

	Severity int32 `json:"severity"`

	AppName string `json:"appName"`

	RemoteAddress string `json:"remoteAddress"`

	RawMessage string `json:"rawMessage"`

	ProcessId string `json:"processId"`

	MessageId string `json:"messageId"`

	DeviceVendor string `json:"deviceVendor"`

	DeviceProduct string `json:"deviceProduct"`

	DeviceVersion string `json:"deviceVersion"`

	Ts int64 `json:"ts"`
}

const SysloglogsAvroCRC64Fingerprint = "T\xcaIޚ\x8f\x92\xf4"

func NewSysloglogs() Sysloglogs {
	r := Sysloglogs{}
	return r
}

func DeserializeSysloglogs(r io.Reader) (Sysloglogs, error) {
	t := NewSysloglogs()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSysloglogsFromSchema(r io.Reader, schema string) (Sysloglogs, error) {
	t := NewSysloglogs()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSysloglogs(r Sysloglogs, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Message, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Host, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Version, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tag, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Level, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Facility, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Severity, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.AppName, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.RemoteAddress, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.RawMessage, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ProcessId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.MessageId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DeviceVendor, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DeviceProduct, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DeviceVersion, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	return err
}

func (r Sysloglogs) Serialize(w io.Writer) error {
	return writeSysloglogs(r, w)
}

func (r Sysloglogs) Schema() string {
	return "{\"fields\":[{\"name\":\"name\",\"type\":{\"arg.properties\":{\"regex\":\"[a-zA-Z]{5,8}\"},\"type\":\"string\"}},{\"name\":\"type\",\"type\":{\"arg.properties\":{\"options\":[\"RFC5424\",\"RFC3164\",\"CEF\",\"UNKNOWN\"]},\"type\":\"string\"}},{\"name\":\"message\",\"type\":{\"arg.properties\":{\"regex\":\"[a-z ]{10,15}\"},\"type\":\"string\"}},{\"name\":\"host\",\"type\":{\"arg.properties\":{\"options\":[\"121.46.66.201\",\"127.26.223.241\",\"185.167.200.143\",\"137.92.22.136\",\"191.84.20.142\",\"151.253.100.45\",\"76.115.125.169\",\"25.19.61.124\",\"140.136.224.11\",\"107.226.4.123\"]},\"type\":\"string\"}},{\"name\":\"version\",\"type\":{\"arg.properties\":{\"options\":[\"3.25.1\"]},\"type\":\"string\"}},{\"name\":\"tag\",\"type\":{\"arg.properties\":{\"options\":[\".source.s_src\"]},\"type\":\"string\"}},{\"name\":\"level\",\"type\":{\"arg.properties\":{\"range\":{\"max\":191,\"min\":0}},\"type\":\"int\"}},{\"name\":\"facility\",\"type\":{\"arg.properties\":{\"options\":[\"syslog\",\"authpriv\",\"cron\"]},\"type\":\"string\"}},{\"name\":\"severity\",\"type\":{\"arg.properties\":{\"range\":{\"max\":8,\"min\":0}},\"type\":\"int\"}},{\"name\":\"appName\",\"type\":{\"arg.properties\":{\"regex\":\"[A-Z]{5,8}\"},\"type\":\"string\"}},{\"name\":\"remoteAddress\",\"type\":{\"arg.properties\":{\"options\":[\"91.253.222.9\",\"48.92.8.255\",\"107.150.81.0\",\"254.11.108.139\",\"14.53.111.201\",\"215.70.232.123\",\"122.96.193.183\",\"185.128.89.135\",\"246.46.59.238\",\"238.158.147.172\"]},\"type\":\"string\"}},{\"name\":\"rawMessage\",\"type\":{\"arg.properties\":{\"regex\":\"[a-z ]{10,15}\"},\"type\":\"string\"}},{\"name\":\"processId\",\"type\":{\"arg.properties\":{\"regex\":\"[1-9]{1}[0-9]{6}\"},\"type\":\"string\"}},{\"name\":\"messageId\",\"type\":{\"arg.properties\":{\"regex\":\"[1-9]{1}[0-9]{2}\"},\"type\":\"string\"}},{\"name\":\"deviceVendor\",\"type\":{\"arg.properties\":{\"regex\":\"[a-z]{5}\"},\"type\":\"string\"}},{\"name\":\"deviceProduct\",\"type\":{\"arg.properties\":{\"regex\":\"[a-z]{5}\"},\"type\":\"string\"}},{\"name\":\"deviceVersion\",\"type\":{\"arg.properties\":{\"options\":[\"1.0\",\"2.0\",\"3.0\",\"4.0\",\"5.0\"]},\"type\":\"string\"}},{\"name\":\"ts\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1609459200000,\"step\":100}},\"type\":\"long\"}}],\"name\":\"syslogs.sysloglogs\",\"type\":\"record\"}"
}

func (r Sysloglogs) SchemaName() string {
	return "syslogs.sysloglogs"
}

func (_ Sysloglogs) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Sysloglogs) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Sysloglogs) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Sysloglogs) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Sysloglogs) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Sysloglogs) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Sysloglogs) SetString(v string)   { panic("Unsupported operation") }
func (_ Sysloglogs) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Sysloglogs) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Name}

		return w

	case 1:
		w := types.String{Target: &r.Type}

		return w

	case 2:
		w := types.String{Target: &r.Message}

		return w

	case 3:
		w := types.String{Target: &r.Host}

		return w

	case 4:
		w := types.String{Target: &r.Version}

		return w

	case 5:
		w := types.String{Target: &r.Tag}

		return w

	case 6:
		w := types.Int{Target: &r.Level}

		return w

	case 7:
		w := types.String{Target: &r.Facility}

		return w

	case 8:
		w := types.Int{Target: &r.Severity}

		return w

	case 9:
		w := types.String{Target: &r.AppName}

		return w

	case 10:
		w := types.String{Target: &r.RemoteAddress}

		return w

	case 11:
		w := types.String{Target: &r.RawMessage}

		return w

	case 12:
		w := types.String{Target: &r.ProcessId}

		return w

	case 13:
		w := types.String{Target: &r.MessageId}

		return w

	case 14:
		w := types.String{Target: &r.DeviceVendor}

		return w

	case 15:
		w := types.String{Target: &r.DeviceProduct}

		return w

	case 16:
		w := types.String{Target: &r.DeviceVersion}

		return w

	case 17:
		w := types.Long{Target: &r.Ts}

		return w

	}
	panic("Unknown field index")
}

func (r *Sysloglogs) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Sysloglogs) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Sysloglogs) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Sysloglogs) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Sysloglogs) HintSize(int)                     { panic("Unsupported operation") }
func (_ Sysloglogs) Finalize()                        {}

func (_ Sysloglogs) AvroCRC64Fingerprint() []byte {
	return []byte(SysloglogsAvroCRC64Fingerprint)
}

func (r Sysloglogs) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["message"], err = json.Marshal(r.Message)
	if err != nil {
		return nil, err
	}
	output["host"], err = json.Marshal(r.Host)
	if err != nil {
		return nil, err
	}
	output["version"], err = json.Marshal(r.Version)
	if err != nil {
		return nil, err
	}
	output["tag"], err = json.Marshal(r.Tag)
	if err != nil {
		return nil, err
	}
	output["level"], err = json.Marshal(r.Level)
	if err != nil {
		return nil, err
	}
	output["facility"], err = json.Marshal(r.Facility)
	if err != nil {
		return nil, err
	}
	output["severity"], err = json.Marshal(r.Severity)
	if err != nil {
		return nil, err
	}
	output["appName"], err = json.Marshal(r.AppName)
	if err != nil {
		return nil, err
	}
	output["remoteAddress"], err = json.Marshal(r.RemoteAddress)
	if err != nil {
		return nil, err
	}
	output["rawMessage"], err = json.Marshal(r.RawMessage)
	if err != nil {
		return nil, err
	}
	output["processId"], err = json.Marshal(r.ProcessId)
	if err != nil {
		return nil, err
	}
	output["messageId"], err = json.Marshal(r.MessageId)
	if err != nil {
		return nil, err
	}
	output["deviceVendor"], err = json.Marshal(r.DeviceVendor)
	if err != nil {
		return nil, err
	}
	output["deviceProduct"], err = json.Marshal(r.DeviceProduct)
	if err != nil {
		return nil, err
	}
	output["deviceVersion"], err = json.Marshal(r.DeviceVersion)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Sysloglogs) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["message"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Message); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for message")
	}
	val = func() json.RawMessage {
		if v, ok := fields["host"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Host); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for host")
	}
	val = func() json.RawMessage {
		if v, ok := fields["version"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Version); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for version")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tag"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tag); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tag")
	}
	val = func() json.RawMessage {
		if v, ok := fields["level"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Level); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for level")
	}
	val = func() json.RawMessage {
		if v, ok := fields["facility"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Facility); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for facility")
	}
	val = func() json.RawMessage {
		if v, ok := fields["severity"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Severity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for severity")
	}
	val = func() json.RawMessage {
		if v, ok := fields["appName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AppName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for appName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["remoteAddress"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemoteAddress); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for remoteAddress")
	}
	val = func() json.RawMessage {
		if v, ok := fields["rawMessage"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RawMessage); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rawMessage")
	}
	val = func() json.RawMessage {
		if v, ok := fields["processId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProcessId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for processId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["messageId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MessageId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for messageId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["deviceVendor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeviceVendor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for deviceVendor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["deviceProduct"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeviceProduct); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for deviceProduct")
	}
	val = func() json.RawMessage {
		if v, ok := fields["deviceVersion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeviceVersion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for deviceVersion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ts")
	}
	return nil
}
