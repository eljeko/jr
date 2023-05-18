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
 *     shoe_clickstream.avsc
 *     shoe_customers.avsc
 *     shoe_orders.avsc
 *     shoes_product.avsc
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

type MyMapTestRecord struct {
	IntField map[string]int32 `json:"IntField"`

	LongField map[string]int64 `json:"LongField"`

	StringField map[string]string `json:"StringField"`

	FloatField map[string]float32 `json:"FloatField"`

	BoolField map[string]bool `json:"BoolField"`

	BytesField map[string]Bytes `json:"BytesField"`
}

const MyMapTestRecordAvroCRC64Fingerprint = ",%\xe7\x99:\xa1\xef\xbb"

func NewMyMapTestRecord() MyMapTestRecord {
	r := MyMapTestRecord{}
	r.IntField = make(map[string]int32)

	r.IntField["default"] = 1

	r.LongField = make(map[string]int64)

	r.LongField["default"] = 2

	r.StringField = make(map[string]string)

	r.StringField["default"] = "defaultstring"

	r.FloatField = make(map[string]float32)

	r.FloatField["default"] = 236

	r.BoolField = make(map[string]bool)

	r.BoolField["default"] = true

	r.BytesField = make(map[string]Bytes)

	r.BytesField["default"] = []byte("\x03\x0f\xde")

	return r
}

func DeserializeMyMapTestRecord(r io.Reader) (MyMapTestRecord, error) {
	t := NewMyMapTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMyMapTestRecordFromSchema(r io.Reader, schema string) (MyMapTestRecord, error) {
	t := NewMyMapTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMyMapTestRecord(r MyMapTestRecord, w io.Writer) error {
	var err error
	err = writeMapInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = writeMapLong(r.LongField, w)
	if err != nil {
		return err
	}
	err = writeMapString(r.StringField, w)
	if err != nil {
		return err
	}
	err = writeMapFloat(r.FloatField, w)
	if err != nil {
		return err
	}
	err = writeMapBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = writeMapBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r MyMapTestRecord) Serialize(w io.Writer) error {
	return writeMyMapTestRecord(r, w)
}

func (r MyMapTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":{\"default\":1},\"name\":\"IntField\",\"type\":{\"type\":\"map\",\"values\":\"int\"}},{\"default\":{\"default\":2},\"name\":\"LongField\",\"type\":{\"type\":\"map\",\"values\":\"long\"}},{\"default\":{\"default\":\"defaultstring\"},\"name\":\"StringField\",\"type\":{\"type\":\"map\",\"values\":\"string\"}},{\"default\":{\"default\":236},\"name\":\"FloatField\",\"type\":{\"type\":\"map\",\"values\":\"float\"}},{\"default\":{\"default\":true},\"name\":\"BoolField\",\"type\":{\"type\":\"map\",\"values\":\"boolean\"}},{\"default\":{\"default\":\"\\u0003\\u000fÞ\"},\"name\":\"BytesField\",\"type\":{\"type\":\"map\",\"values\":\"bytes\"}}],\"name\":\"MyMapTestRecord\",\"type\":\"record\"}"
}

func (r MyMapTestRecord) SchemaName() string {
	return "MyMapTestRecord"
}

func (_ MyMapTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ MyMapTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MyMapTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.IntField = make(map[string]int32)

		w := MapIntWrapper{Target: &r.IntField}

		return &w

	case 1:
		r.LongField = make(map[string]int64)

		w := MapLongWrapper{Target: &r.LongField}

		return &w

	case 2:
		r.StringField = make(map[string]string)

		w := MapStringWrapper{Target: &r.StringField}

		return &w

	case 3:
		r.FloatField = make(map[string]float32)

		w := MapFloatWrapper{Target: &r.FloatField}

		return &w

	case 4:
		r.BoolField = make(map[string]bool)

		w := MapBoolWrapper{Target: &r.BoolField}

		return &w

	case 5:
		r.BytesField = make(map[string]Bytes)

		w := MapBytesWrapper{Target: &r.BytesField}

		return &w

	}
	panic("Unknown field index")
}

func (r *MyMapTestRecord) SetDefault(i int) {
	switch i {
	case 0:
		r.IntField["default"] = 1

		return
	case 1:
		r.LongField["default"] = 2

		return
	case 2:
		r.StringField["default"] = "defaultstring"

		return
	case 3:
		r.FloatField["default"] = 236

		return
	case 4:
		r.BoolField["default"] = true

		return
	case 5:
		r.BytesField["default"] = []byte("\x03\x0f\xde")

		return
	}
	panic("Unknown field index")
}

func (r *MyMapTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MyMapTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MyMapTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MyMapTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ MyMapTestRecord) Finalize()                        {}

func (_ MyMapTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(MyMapTestRecordAvroCRC64Fingerprint)
}

func (r MyMapTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	output["LongField"], err = json.Marshal(r.LongField)
	if err != nil {
		return nil, err
	}
	output["StringField"], err = json.Marshal(r.StringField)
	if err != nil {
		return nil, err
	}
	output["FloatField"], err = json.Marshal(r.FloatField)
	if err != nil {
		return nil, err
	}
	output["BoolField"], err = json.Marshal(r.BoolField)
	if err != nil {
		return nil, err
	}
	output["BytesField"], err = json.Marshal(r.BytesField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MyMapTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IntField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IntField); err != nil {
			return err
		}
	} else {
		r.IntField = make(map[string]int32)

		r.IntField["default"] = 1

	}
	val = func() json.RawMessage {
		if v, ok := fields["LongField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LongField); err != nil {
			return err
		}
	} else {
		r.LongField = make(map[string]int64)

		r.LongField["default"] = 2

	}
	val = func() json.RawMessage {
		if v, ok := fields["StringField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StringField); err != nil {
			return err
		}
	} else {
		r.StringField = make(map[string]string)

		r.StringField["default"] = "defaultstring"

	}
	val = func() json.RawMessage {
		if v, ok := fields["FloatField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FloatField); err != nil {
			return err
		}
	} else {
		r.FloatField = make(map[string]float32)

		r.FloatField["default"] = 236

	}
	val = func() json.RawMessage {
		if v, ok := fields["BoolField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BoolField); err != nil {
			return err
		}
	} else {
		r.BoolField = make(map[string]bool)

		r.BoolField["default"] = true

	}
	val = func() json.RawMessage {
		if v, ok := fields["BytesField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BytesField); err != nil {
			return err
		}
	} else {
		r.BytesField = make(map[string]Bytes)

		r.BytesField["default"] = []byte("\x03\x0f\xde")

	}
	return nil
}
