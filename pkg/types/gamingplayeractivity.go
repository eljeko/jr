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
 *     shoe_customer.avsc
 *     shoe_order.avsc
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

type Gamingplayeractivity struct {
	Player_id int32 `json:"player_id"`

	Game_room_id int32 `json:"game_room_id"`

	Points int32 `json:"points"`

	Coordinates string `json:"coordinates"`
}

const GamingplayeractivityAvroCRC64Fingerprint = "\xc0\x14&ɿ\xa2\xfc\xb3"

func NewGamingplayeractivity() Gamingplayeractivity {
	r := Gamingplayeractivity{}
	return r
}

func DeserializeGamingplayeractivity(r io.Reader) (Gamingplayeractivity, error) {
	t := NewGamingplayeractivity()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeGamingplayeractivityFromSchema(r io.Reader, schema string) (Gamingplayeractivity, error) {
	t := NewGamingplayeractivity()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeGamingplayeractivity(r Gamingplayeractivity, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Player_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Game_room_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Points, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Coordinates, w)
	if err != nil {
		return err
	}
	return err
}

func (r Gamingplayeractivity) Serialize(w io.Writer) error {
	return writeGamingplayeractivity(r, w)
}

func (r Gamingplayeractivity) Schema() string {
	return "{\"fields\":[{\"name\":\"player_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1101,\"min\":1001}},\"type\":\"int\"}},{\"name\":\"game_room_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5000,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"points\",\"type\":{\"arg.properties\":{\"range\":{\"max\":500,\"min\":10}},\"type\":\"int\"}},{\"name\":\"coordinates\",\"type\":{\"arg.properties\":{\"regex\":\"\\\\[[0-9][0-9],[0-9][0-9]\\\\]\"},\"type\":\"string\"}}],\"name\":\"gaming.gamingplayeractivity\",\"type\":\"record\"}"
}

func (r Gamingplayeractivity) SchemaName() string {
	return "gaming.gamingplayeractivity"
}

func (_ Gamingplayeractivity) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetString(v string)   { panic("Unsupported operation") }
func (_ Gamingplayeractivity) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Gamingplayeractivity) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Player_id}

		return w

	case 1:
		w := types.Int{Target: &r.Game_room_id}

		return w

	case 2:
		w := types.Int{Target: &r.Points}

		return w

	case 3:
		w := types.String{Target: &r.Coordinates}

		return w

	}
	panic("Unknown field index")
}

func (r *Gamingplayeractivity) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Gamingplayeractivity) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Gamingplayeractivity) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Gamingplayeractivity) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Gamingplayeractivity) HintSize(int)                     { panic("Unsupported operation") }
func (_ Gamingplayeractivity) Finalize()                        {}

func (_ Gamingplayeractivity) AvroCRC64Fingerprint() []byte {
	return []byte(GamingplayeractivityAvroCRC64Fingerprint)
}

func (r Gamingplayeractivity) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["player_id"], err = json.Marshal(r.Player_id)
	if err != nil {
		return nil, err
	}
	output["game_room_id"], err = json.Marshal(r.Game_room_id)
	if err != nil {
		return nil, err
	}
	output["points"], err = json.Marshal(r.Points)
	if err != nil {
		return nil, err
	}
	output["coordinates"], err = json.Marshal(r.Coordinates)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Gamingplayeractivity) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["player_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Player_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for player_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["game_room_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Game_room_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for game_room_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["points"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Points); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for points")
	}
	val = func() json.RawMessage {
		if v, ok := fields["coordinates"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Coordinates); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for coordinates")
	}
	return nil
}
