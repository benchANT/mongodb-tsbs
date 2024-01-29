// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package mongo

import "strconv"

type MongoTagValue byte

const (
	MongoTagValueNONE            MongoTagValue = 0
	MongoTagValueMongoStringTag  MongoTagValue = 1
	MongoTagValueMongoFloat32Tag MongoTagValue = 2
)

var EnumNamesMongoTagValue = map[MongoTagValue]string{
	MongoTagValueNONE:            "NONE",
	MongoTagValueMongoStringTag:  "MongoStringTag",
	MongoTagValueMongoFloat32Tag: "MongoFloat32Tag",
}

var EnumValuesMongoTagValue = map[string]MongoTagValue{
	"NONE":            MongoTagValueNONE,
	"MongoStringTag":  MongoTagValueMongoStringTag,
	"MongoFloat32Tag": MongoTagValueMongoFloat32Tag,
}

func (v MongoTagValue) String() string {
	if s, ok := EnumNamesMongoTagValue[v]; ok {
		return s
	}
	return "MongoTagValue(" + strconv.FormatInt(int64(v), 10) + ")"
}
