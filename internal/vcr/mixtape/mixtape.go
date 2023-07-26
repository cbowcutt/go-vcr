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

func (m *Mixtape) ToYaml() ([]byte, error) {
	//introYaml, err := m.Intro.ToYaml()
	//if err != nil {
	//	return nil, err
	//}
	////trackYamls := make([][]byte, 0)
	////
	////for _, t := range m.tracks {
	////	trackYaml, err := t.ToYaml()
	////	if err != nil {
	////		return nil, err
	////	}
	////	trackYamls = append(trackYamls, trackYaml)
	////}
	//yamlStruct := mixtapeYaml{
	//	Intro: introYaml,
	//}
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
