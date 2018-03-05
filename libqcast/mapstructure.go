package libqcast

import (
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
)

func DecodeMap(inputMap, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			TagName: "json",
			Result:  result,
		},
	)
	if err != nil {
		glog.Errorf("Failed to create decoder: %s", err)
		return err
	}
	err = decoder.Decode(inputMap)
	if err != nil {
		glog.Errorf("Failed to decode from map: %s", err)
		return err
	}

	return nil
}
