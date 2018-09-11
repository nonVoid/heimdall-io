package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/hubidu/e2e-backend/report-lib/db"
	"github.com/hubidu/e2e-backend/report-lib/model"
	"github.com/jasonlvhit/gocron"
)

func init() {
	db.Connect()
}

func deleteOldReports() {
	CleanupOldReports(7)
}

func importJob(baseDir string, removeReportFiles bool) {
	reports := model.GetReportFiles(baseDir)

	if len(reports) > 0 {
		fmt.Println("Importing reports from directory " + baseDir + " ...")

		// TODO Should not insert duplicate reports
		InsertReportsIntoDB(reports)

		fmt.Println("Inserted " + strconv.Itoa(len(reports)) + " report files into database ...")

		if removeReportFiles {
			// fmt.Println("Renaming report files ...")
			for _, report := range reports {
				reportFileName := path.Join(baseDir, report.ReportDir, report.ReportFileName)
				os.Rename(reportFileName, path.Join(baseDir, report.ReportDir, "report_imported.json"))
			}
		}
	}
}

func main() {
	removeReportFiles := flag.Bool("rm", false, "set to true to remove report files from disk")
	flag.Parse()

	args := flag.Args()

	var baseDir string
	if len(args) >= 1 {
		baseDir = args[0]
	} else {
		baseDir = "./fixtures"
	}

	importJobTask := func() {
		importJob(baseDir, *removeReportFiles)
	}

	var intervalInSeconds uint64
	intervalInSeconds = 3
	if len(os.Getenv("JOB_INTERVAL")) != 0 {
		interval, _ := strconv.Atoi(os.Getenv("JOB_INTERVAL"))
		intervalInSeconds = uint64(interval)
	}

	gocron.Every(intervalInSeconds).Seconds().Do(importJobTask)
	gocron.Every(1).Day().At("05:00").Do(deleteOldReports)

	<-gocron.Start()
}
