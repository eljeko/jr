// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     clickstream.avsc
 *     clickstreamcodes.avsc
 *     clickstreamusers.avsc
 *     creditcards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     pageviews.avsc
 *     payrollbonus.avsc
 *     payrollemployee.avsc
 *     payrollemployeelocation.avsc
 *     pizzaorders.avsc
 *     pizzaorderscancelled.avsc
 *     pizzaorderscompleted.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
 *     stockTrades.avsc
 *     stores.avsc
 *     sysloglogs.avsc
 *     transactions.avsc
 *     user.avsc
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

type ShoeCustomer struct {
	Id string `json:"id"`

	First_name string `json:"first_name"`

	Last_name string `json:"last_name"`

	Email string `json:"email"`

	Phone_number string `json:"phone_number"`

	Street_address string `json:"street_address"`

	State string `json:"state"`

	Zip_code string `json:"zip_code"`

	Country string `json:"country"`

	Country_code string `json:"country_code"`
}

const ShoeCustomerAvroCRC64Fingerprint = "\xb9S<9+WU\x12"

func NewShoeCustomer() ShoeCustomer {
	r := ShoeCustomer{}
	return r
}

func DeserializeShoeCustomer(r io.Reader) (ShoeCustomer, error) {
	t := NewShoeCustomer()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoeCustomerFromSchema(r io.Reader, schema string) (ShoeCustomer, error) {
	t := NewShoeCustomer()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoeCustomer(r ShoeCustomer, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.First_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Last_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Phone_number, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Street_address, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.State, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Zip_code, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Country, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Country_code, w)
	if err != nil {
		return err
	}
	return err
}

func (r ShoeCustomer) Serialize(w io.Writer) error {
	return writeShoeCustomer(r, w)
}

func (r ShoeCustomer) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"first_name\",\"type\":\"string\"},{\"name\":\"last_name\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"phone_number\",\"type\":\"string\"},{\"name\":\"street_address\",\"type\":\"string\"},{\"name\":\"state\",\"type\":\"string\"},{\"name\":\"zip_code\",\"type\":\"string\"},{\"name\":\"country\",\"type\":\"string\"},{\"name\":\"country_code\",\"type\":\"string\"}],\"name\":\"shoes.ShoeCustomer\",\"type\":\"record\"}"
}

func (r ShoeCustomer) SchemaName() string {
	return "shoes.ShoeCustomer"
}

func (_ ShoeCustomer) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ShoeCustomer) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ShoeCustomer) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ShoeCustomer) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ShoeCustomer) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ShoeCustomer) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ShoeCustomer) SetString(v string)   { panic("Unsupported operation") }
func (_ ShoeCustomer) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ShoeCustomer) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.First_name}

		return w

	case 2:
		w := types.String{Target: &r.Last_name}

		return w

	case 3:
		w := types.String{Target: &r.Email}

		return w

	case 4:
		w := types.String{Target: &r.Phone_number}

		return w

	case 5:
		w := types.String{Target: &r.Street_address}

		return w

	case 6:
		w := types.String{Target: &r.State}

		return w

	case 7:
		w := types.String{Target: &r.Zip_code}

		return w

	case 8:
		w := types.String{Target: &r.Country}

		return w

	case 9:
		w := types.String{Target: &r.Country_code}

		return w

	}
	panic("Unknown field index")
}

func (r *ShoeCustomer) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ShoeCustomer) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ShoeCustomer) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ShoeCustomer) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ShoeCustomer) HintSize(int)                     { panic("Unsupported operation") }
func (_ ShoeCustomer) Finalize()                        {}

func (_ ShoeCustomer) AvroCRC64Fingerprint() []byte {
	return []byte(ShoeCustomerAvroCRC64Fingerprint)
}

func (r ShoeCustomer) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["first_name"], err = json.Marshal(r.First_name)
	if err != nil {
		return nil, err
	}
	output["last_name"], err = json.Marshal(r.Last_name)
	if err != nil {
		return nil, err
	}
	output["email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["phone_number"], err = json.Marshal(r.Phone_number)
	if err != nil {
		return nil, err
	}
	output["street_address"], err = json.Marshal(r.Street_address)
	if err != nil {
		return nil, err
	}
	output["state"], err = json.Marshal(r.State)
	if err != nil {
		return nil, err
	}
	output["zip_code"], err = json.Marshal(r.Zip_code)
	if err != nil {
		return nil, err
	}
	output["country"], err = json.Marshal(r.Country)
	if err != nil {
		return nil, err
	}
	output["country_code"], err = json.Marshal(r.Country_code)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ShoeCustomer) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["first_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.First_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for first_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["last_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Last_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for last_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["phone_number"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Phone_number); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for phone_number")
	}
	val = func() json.RawMessage {
		if v, ok := fields["street_address"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Street_address); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for street_address")
	}
	val = func() json.RawMessage {
		if v, ok := fields["state"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.State); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for state")
	}
	val = func() json.RawMessage {
		if v, ok := fields["zip_code"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Zip_code); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for zip_code")
	}
	val = func() json.RawMessage {
		if v, ok := fields["country"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Country); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for country")
	}
	val = func() json.RawMessage {
		if v, ok := fields["country_code"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Country_code); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for country_code")
	}
	return nil
}
