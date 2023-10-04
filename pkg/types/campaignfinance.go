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

type Campaignfinance struct {
	Time int64 `json:"time"`

	Candidate_id string `json:"candidate_id"`

	Party_affiliation string `json:"party_affiliation"`

	Contribution int32 `json:"contribution"`
}

const CampaignfinanceAvroCRC64Fingerprint = "gx\xd3\xc5[\xcd2\xb1"

func NewCampaignfinance() Campaignfinance {
	r := Campaignfinance{}
	return r
}

func DeserializeCampaignfinance(r io.Reader) (Campaignfinance, error) {
	t := NewCampaignfinance()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCampaignfinanceFromSchema(r io.Reader, schema string) (Campaignfinance, error) {
	t := NewCampaignfinance()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCampaignfinance(r Campaignfinance, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Time, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Candidate_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Party_affiliation, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Contribution, w)
	if err != nil {
		return err
	}
	return err
}

func (r Campaignfinance) Serialize(w io.Writer) error {
	return writeCampaignfinance(r, w)
}

func (r Campaignfinance) Schema() string {
	return "{\"fields\":[{\"name\":\"time\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1619273364600,\"min\":1587715775521}},\"type\":\"long\"}},{\"name\":\"candidate_id\",\"type\":{\"arg.properties\":{\"regex\":\"[A-Z][1-9]{8,8}\"},\"type\":\"string\"}},{\"name\":\"party_affiliation\",\"type\":{\"arg.properties\":{\"options\":[\"REP\",\"DEM\",\"IND\"]},\"type\":\"string\"}},{\"name\":\"contribution\",\"type\":{\"arg.properties\":{\"range\":{\"max\":3500,\"min\":20}},\"type\":\"int\"}}],\"name\":\"datagen.campaignfinance\",\"type\":\"record\"}"
}

func (r Campaignfinance) SchemaName() string {
	return "datagen.campaignfinance"
}

func (_ Campaignfinance) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Campaignfinance) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Campaignfinance) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Campaignfinance) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Campaignfinance) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Campaignfinance) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Campaignfinance) SetString(v string)   { panic("Unsupported operation") }
func (_ Campaignfinance) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Campaignfinance) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Time}

		return w

	case 1:
		w := types.String{Target: &r.Candidate_id}

		return w

	case 2:
		w := types.String{Target: &r.Party_affiliation}

		return w

	case 3:
		w := types.Int{Target: &r.Contribution}

		return w

	}
	panic("Unknown field index")
}

func (r *Campaignfinance) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Campaignfinance) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Campaignfinance) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Campaignfinance) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Campaignfinance) HintSize(int)                     { panic("Unsupported operation") }
func (_ Campaignfinance) Finalize()                        {}

func (_ Campaignfinance) AvroCRC64Fingerprint() []byte {
	return []byte(CampaignfinanceAvroCRC64Fingerprint)
}

func (r Campaignfinance) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["time"], err = json.Marshal(r.Time)
	if err != nil {
		return nil, err
	}
	output["candidate_id"], err = json.Marshal(r.Candidate_id)
	if err != nil {
		return nil, err
	}
	output["party_affiliation"], err = json.Marshal(r.Party_affiliation)
	if err != nil {
		return nil, err
	}
	output["contribution"], err = json.Marshal(r.Contribution)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Campaignfinance) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for time")
	}
	val = func() json.RawMessage {
		if v, ok := fields["candidate_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Candidate_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for candidate_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["party_affiliation"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Party_affiliation); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for party_affiliation")
	}
	val = func() json.RawMessage {
		if v, ok := fields["contribution"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contribution); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contribution")
	}
	return nil
}
