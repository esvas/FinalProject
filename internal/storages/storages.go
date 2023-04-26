package storages

import (
	"github.com/esvas/FinalProject/configs"
	"github.com/esvas/FinalProject/internal/billing"
)

type ResultSetT struct {
	SMS       []SMSStorage                 `json:"sms"`
	MMS       []MMSStorage                 `json:"mms"`
	VoiceCall VCStorage                    `json:"voice_call"`
	Email     map[string][]providerStorage `json:"email"`
	Billing   billing.BillingData          `json:"billing"`
	Support   []int                        `json:"support"`
	Incidents IncidentStorage              `json:"incident"`
}

type ResultT struct {
	Status bool        `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  []string    `json:"errors"` // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

var config = configs.GetConfig()

func GetResultData() ResultT {
	data := ResultSetT{}
	status := ResultT{
		Status: true,
		Data:   &data,
		Error:  nil,
	}
	errs := []string{}
	smsData, err := smsDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.SMS = smsData
	mmsData, err := mmsDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.MMS = mmsData
	vcData, err := vcDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.VoiceCall = vcData
	emailData, err := emailDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Email = emailData
	billingData, err := billingDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Billing = billingData
	supportData, err := supportDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Support = supportData
	incidentData, err := incidentDataF()
	if err != nil {
		errs = append(errs, err.Error())
	}
	data.Incidents = incidentData
	if len(errs) > 0 {
		status.Error = errs
		status.Status = false
	}
	return status
}

func smsDataF() ([]SMSStorage, error) {
	path := config.SMS
	smsStorage, err := NewSMSStorage(path)
	if err != nil {
		return nil, err
	}
	sortedProvider := make(SMSStorage, len(*smsStorage))
	smsStorage.SortProvider()
	copy(sortedProvider, *smsStorage)
	smsStorage.SortCountry()
	sortedCountry := make(SMSStorage, len(*smsStorage))
	copy(sortedCountry, *smsStorage)
	return []SMSStorage{sortedProvider, sortedCountry}, nil
}

func mmsDataF() ([]MMSStorage, error) {
	path := config.MMS
	mmsStorage, err := NewMMSStorage(path)
	if err != nil {
		return nil, err
	}
	sortedProvider := make(MMSStorage, len(*mmsStorage))
	mmsStorage.SortProvider()
	copy(sortedProvider, *mmsStorage)
	sortedCountry := make(MMSStorage, len(*mmsStorage))
	mmsStorage.SortCountry()
	copy(sortedCountry, *mmsStorage)
	return []MMSStorage{sortedProvider, sortedCountry}, nil
}

func vcDataF() (VCStorage, error) {
	path := config.VoiceCall
	vcData, err := NewVCStorage(path)
	return *vcData, err
}

func emailDataF() (map[string][]providerStorage, error) {
	path := config.Email
	emailStorage, err := NewEmailStorage(path)
	if err != nil {
		return nil, err
	}
	catalogEmailByCountry := emailStorage.catalogingByCountry()
	result := map[string][]providerStorage{}
	for country, emails := range catalogEmailByCountry {
		providers := emails.createStatisticProviders()
		providers.sort()
		topsProviders := providers.BestAndWorst()
		result[country] = topsProviders
	}
	return result, nil
}

func billingDataF() (billing.BillingData, error) {
	path := config.Billing
	billingData, err := billing.New(path)
	return *billingData, err
}

func supportDataF() ([]int, error) {
	path := config.Support
	supportData, err := NewSupportStorage(path)
	if err != nil {
		return nil, err
	}
	loadStatus, waitTime := supportData.CurrentLoad()
	return []int{loadStatus, waitTime}, nil
}

func incidentDataF() (IncidentStorage, error) {
	path := config.Incident
	incidentData, err := NewIncidentStorage(path)
	if err != nil {
		return nil, err
	}
	incidentData.sort()
	return *incidentData, nil
}