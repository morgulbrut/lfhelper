package cmd

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dannav/hhmmss"
	"github.com/davecgh/go-spew/spew"
	"github.com/morgulbrut/helferlein"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		csv := ReadCVS(args[0])
		spew.Dump(csv)
	},
}

func init() {
	readCmd.AddCommand(csvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

type CSV struct {
	ProgName  string
	RunTime   time.Duration
	Cals      int
	Distance  float64
	RiseDist  int
	AvSpeed   float64
	AvPace    time.Duration
	AvPower   int
	AvHeartR  int
	Product   string
	Timestamp time.Time
}

func ReadCVS(fn string) CSV {
	var scheme CsvScheme
	var lf CSV
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
