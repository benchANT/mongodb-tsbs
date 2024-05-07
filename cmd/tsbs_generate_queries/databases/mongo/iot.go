package mongo

import (
	"fmt"
//	"strings"
	"time"
	"encoding/gob"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/databases"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/iot"
	"github.com/timescale/tsbs/pkg/query"
)

func init() {
	// needed for serializing the mongo query to gob
	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})
	gob.Register([]map[string]interface{}{})
	gob.Register(bson.M{})
	gob.Register(bson.D{})
	gob.Register([]bson.M{})
}

func (i *IoT) getTrucksFilterArray(nTrucks int) []string {
	names, err := i.GetRandomTrucks(nTrucks)
	panicIfErr(err)
	return names
}

// IoT produces Mongo-specific queries for all the iot query types.
type IoT struct {
	*iot.Core
	*BaseGenerator
}

// NewIoT makes an IoT object ready to generate Queries.
func NewIoT(start, end time.Time, scale int, g *BaseGenerator) *IoT {
	c, err := iot.NewCore(start, end, scale)
	databases.PanicIfErr(err)
	return &IoT{
		Core:          c,
		BaseGenerator: g,
	}
}

func (i *IoT) LastLocByTruck(qi query.Query, nTrucks int) {
	trucks := i.getTrucksFilterArray(nTrucks)
	pipelineQuery := mongo.Pipeline{
		{{
			"$match", bson.M{
				"tags.name": bson.M{
					"$in": trucks,
				},
			},

		}},
		{{
			"$group", bson.M{
				"_id": "$tags.name",
				"output": bson.M{
					"$top": bson.M{
						"sortBy": bson.M{ "time" : -1},
						"output": bson.M{
							"longitude": "$longitude",
							"latitude": "$latitude",
							"time": "$time",
						},
					},
				},
			},
		}},
	}

	humanLabel := "MongoDB last location by specific truck(s)"
	humanDesc := fmt.Sprintf("%s: random %4d trucks (%v)", humanLabel, nTrucks, trucks)

	q := qi.(*query.Mongo)
	q.HumanLabel = []byte(humanLabel)
	q.Pipeline = pipelineQuery
	q.CollectionName = []byte("point_data")
	q.HumanDescription = []byte(humanDesc)
}
