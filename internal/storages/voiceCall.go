package storages

import (
	"github.com/esvas/FinalProject/internal/voiceCall"
	"github.com/esvas/FinalProject/pkg/pars"
)

type VCStorage []*voiceCall.VoiceCall

func (vcs *VCStorage) Add(obj *voiceCall.VoiceCall) {
	*vcs = append(*vcs, obj)
}

func NewVCStorage(filename string) (*VCStorage, error) {
	return createVCStorage(filename)
}

func createVCStorage(filename string) (*VCStorage, error) {
	smsStr, err := pars.FileToStr(filename)
	if err != nil {
		return nil, err
	}
	ss := VCStorage{}
	for _, s := range smsStr {
		res := voiceCall.FromSTR(s)
		if res == nil {
			continue
		}
		ss.Add(res)
	}
	return &ss, nil
}