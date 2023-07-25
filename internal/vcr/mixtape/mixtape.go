package mixtape

import (
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape/track"
	"gopkg.in/yaml.v3"
)

type Mixtape struct {
	Intro  track.Track   `yaml:"intro"`
	tracks []track.Track `'yaml:"tracks"`
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
