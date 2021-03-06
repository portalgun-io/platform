package spectests

import (
	"time"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/functions/inputs"
	"github.com/influxdata/flux/functions/transformations"
)

func init() {
	RegisterFixture(
		NewFixture(
			`SHOW TAG VALUES ON "db0" WITH KEY = "host"`,
			&flux.Spec{
				Operations: []*flux.Operation{
					{
						ID: "from0",
						Spec: &inputs.FromOpSpec{
							BucketID: bucketID.String(),
						},
					},
					{
						ID: "range0",
						Spec: &transformations.RangeOpSpec{
							Start: flux.Time{
								Relative:   -time.Hour,
								IsRelative: true,
							},
							Stop: flux.Now,
						},
					},
					{
						ID: "keyValues0",
						Spec: &transformations.KeyValuesOpSpec{
							KeyColumns: []string{"host"},
						},
					},
					{
						ID: "group0",
						Spec: &transformations.GroupOpSpec{
							Columns: []string{"_measurement", "_key"},
							Mode:    "by",
						},
					},
					{
						ID: "distinct0",
						Spec: &transformations.DistinctOpSpec{
							Column: execute.DefaultValueColLabel,
						},
					},
					{
						ID: "group1",
						Spec: &transformations.GroupOpSpec{
							Columns: []string{"_measurement"},
							Mode:    "by",
						},
					},
					{
						ID: "rename0",
						Spec: &transformations.RenameOpSpec{
							Columns: map[string]string{
								"_key":   "key",
								"_value": "value",
							},
						},
					},
					{
						ID: "yield0",
						Spec: &transformations.YieldOpSpec{
							Name: "0",
						},
					},
				},
				Edges: []flux.Edge{
					{Parent: "from0", Child: "range0"},
					{Parent: "range0", Child: "keyValues0"},
					{Parent: "keyValues0", Child: "group0"},
					{Parent: "group0", Child: "distinct0"},
					{Parent: "distinct0", Child: "group1"},
					{Parent: "group1", Child: "rename0"},
					{Parent: "rename0", Child: "yield0"},
				},
				Now: Now(),
			},
		),
	)
}
