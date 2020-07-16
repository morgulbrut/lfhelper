package cmd

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dannav/hhmmss"
	"github.com/morgulbrut/helferlein"
	"github.com/morgulbrut/lfhelper/trainingdata"
	"gopkg.in/yaml.v2"
)

// CsvScheme descripes the names of the fields because
// LifeFitness don't know how to CSV
type CsvScheme struct {
	ProgName  string `yaml:progname`
	RunTime   string `yaml:runtime`
	Cals      string `yaml:cals`
	Distance  string `yaml:distance`
	RiseDist  string `yaml:risedist`
	AvSpeed   string `yaml:avspeed`
	AvPace    string `yaml:avpace`
	AvPower   string `yaml:avpower`
	AvHeartR  string `yaml:avheartr`
	Product   string `yaml:product`
	Timestamp string `yaml:timestamp`
}

func (csv *CsvScheme) readScheme(fn string) *CsvScheme {

	yamlFile, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("ReadScheme read: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, csv)
	if err != nil {
		log.Fatalf("ReadScheme unmarshal: %v", err)
	}

	return csv
}

func ReadCVS(fn string) trainingdata.TrainingUnit {
	var scheme CsvScheme
	var lf trainingdata.TrainingUnit
	scheme.readScheme("scheme.yml")

	data, err := helferlein.ReadLines(fn)
	if err != nil {
		log.Fatalf("Read: %v", err)
	}

	for _, line := range data {

		// because LF can't do shit correctly, it seems
		line = strings.Replace(line, "\x00", "", -1)

		l := strings.Split(line, ",")
		if len(l) == 3 {
			key := l[1]
			val := l[2]
			key = strings.TrimSpace(key)
			switch key {
			case scheme.Product:
				lf.Product = val
			case scheme.ProgName:
				lf.ProgName = val
			case scheme.RunTime:
				lf.RunTime, _ = hhmmss.Parse(val)
			case scheme.AvPace:
				val = strings.Split(val, " ")[0]
				val = strings.TrimSpace(val)
				lf.AvPace, _ = hhmmss.Parse(val)
			case scheme.Timestamp:
				lf.Timestamp, err = time.Parse("01/02/2006 15:04:05", val)
			case scheme.Cals:
				lf.Cals, _ = strconv.Atoi(val)
			case scheme.AvPower:
				lf.AvPower, _ = strconv.Atoi(val)
			case scheme.AvHeartR:
				lf.AvHeartR, _ = strconv.Atoi(val)
			case scheme.RiseDist:
				val = strings.Split(val, " ")[0]
				val = strings.TrimSpace(val)
				lf.RiseDist, _ = strconv.Atoi(val)
			case scheme.Distance:
				val = strings.Split(val, " ")[0]
				val = strings.TrimSpace(val)
				lf.Distance, _ = strconv.ParseFloat(val, 64)
			case scheme.AvSpeed:
				val = strings.Split(val, " ")[0]
				val = strings.TrimSpace(val)
				lf.AvSpeed, _ = strconv.ParseFloat(val, 64)
			}
		}
	}
	return lf
}
