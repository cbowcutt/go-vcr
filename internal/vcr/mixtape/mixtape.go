package mixtape

import (
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape/track"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"time"
)

type Mixtape struct {
	Intro  *track.Track   `yaml:"intro"`
	Tracks []*track.Track `'yaml:"tracks"`
}

var mixtapeInstance *Mixtape

func GetMixtape() *Mixtape {
	if mixtapeInstance == nil {
		mixtapeInstance = &Mixtape{}
	}
	return mixtapeInstance
}
func (m *Mixtape) FromYaml(yamlBytes []byte) error {
	yaml.Unmarshal(yamlBytes, m)
	err := yaml.Unmarshal(yamlBytes, m)
	if err != nil {
		return err
	}
	err = m.Intro.Decode()
	if err != nil {
		return err
	}
	for _, track := range m.Tracks {
		err = track.Decode()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Mixtape) ToYaml() ([]byte, error) {
	return yaml.Marshal(m)
}

func (m *Mixtape) ToFile() error {
	filepath := ""
	if m.Intro.Kind == "grpc" {
		second := time.Now().Unix()
		//segment := m.Intro.Data.(*track.GrpcTrackData).Method
		filepath = "./" + strconv.Itoa(int(second)) + ".yaml"
	}
	bytes, err := m.ToYaml()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(bytes)
	return err
}
