package core

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/Ovsienko023/reporter/app/repository"
)

const (
	templateXlsxSheet = "Лист1"
	templateXlsxPath  = "store_fs/template.xlsx"
)

func (c *Core) ExportXLSXReports(ctx context.Context) (string, error) {
	//invokerId, err := c.authorize(msg.Token)
	//if err != nil {
	//	return err
	//}

	fileData, err := c.Fs.ReadFile(templateXlsxPath)
	if err != nil {
		return "", ErrInternal
	}

	file, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		return "", ErrInternal
	}

	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	reports := c.getReports(ctx)

	for i, row := range reports {
		dataRow := i + 2
		for j, col := range row {
			err := file.SetCellValue(templateXlsxSheet, fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
			if err != nil {
				return "", ErrInternal
			}
		}
	}

	if err := file.SaveAs("19a3d7a1-5fdd-4156-a246-ecde104f21fc.xlsx"); err != nil {
		return "", ErrInternal
	}

	return "19a3d7a1-5fdd-4156-a246-ecde104f21fc.xlsx", nil
}

func (c *Core) getReports(ctx context.Context) [][]interface{} {
	reportsDb, _, err := c.db.GetReports(ctx, &repository.GetReports{
		InvokerId: "19a3d7a1-5fdd-4156-a246-ecde104f21fc",
		DateFrom:  nil,
		DateTo:    nil,
	})

	if err != nil {
		panic(err.Error())
	}

	reports := make([][]interface{}, 0, len(reportsDb))
	for _, report := range reportsDb {
		startTime := time.Unix(*report.StartTime, 0).UTC().Format("15:04")
		endTime := time.Unix(*report.EndTime, 0).UTC().Format("15:04")
		breakTime := time.Unix(*report.BreakTime, 0).UTC().Format("15:04")

		item := []interface{}{*report.Date, startTime, endTime, breakTime}
		reports = append(reports, item)
	}

	fmt.Println(reports)

	return reports
}
