package resultsmapping

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/checkmarxDev/ast-sast-export/internal/app/metadata"
)

func GenerateCSV(records []*metadata.Record) [][]string {
	var items [][]string
	items = append(items, []string{
		"result_id",
		"cxone_similarity_id",
		"sast_similarity_id",
	})
	if records == nil {
		return items
	}
	for _, record := range records {
		for _, query := range record.Queries {
			for _, result := range query.Results {
				for _, path := range result.Paths {
					items = append(items, []string{
						path.ResultID,
						path.SimilarityID,
						path.SASTSimilarityID,
					})
				}
			}
		}
	}

	return items
}

// GenerateCSVWithDetectionDate is like GenerateCSV but includes a project_id and detection_date column.
// Used exclusively for the consolidated --simid-mapping-file output.
func GenerateCSVWithDetectionDate(records []*metadata.Record) [][]string {
	var items [][]string
	items = append(items, []string{
		"sast_project_id",
		"sast_scan_id",
		"sast_similarity_id",
		"cxone_similarity_id",
		"detection_date",
	})
	if records == nil {
		return items
	}
	for _, record := range records {
		projectID := strconv.Itoa(record.ProjectID)
		for _, query := range record.Queries {
			for _, result := range query.Results {
				for _, path := range result.Paths {
					items = append(items, []string{
						projectID,
						path.ResultID,
						path.SASTSimilarityID,
						path.SimilarityID,
						path.DetectionDate,
					})
				}
			}
		}
	}

	return items
}

func WriteAllToSanitizedCsv(records [][]string) []byte {
	for _, row := range records {
		for i := range row {
			row[i] = sanitize(row[i])
		}
	}

	rows := make([]string, len(records))
	for i, record := range records {
		rows[i] = strings.Join(record, ",") + "\n"
	}

	data := strings.Join(rows, "")

	return []byte(data)
}

// GenerateCSVWithDetectionDateForRecord returns data rows (no header) for a single record.
// Used by simIDMappingWriter to stream rows as each record becomes available.
func GenerateCSVWithDetectionDateForRecord(record *metadata.Record) [][]string {
	var rows [][]string
	projectID := strconv.Itoa(record.ProjectID)
	for _, query := range record.Queries {
		for _, result := range query.Results {
			for _, path := range result.Paths {
				rows = append(rows, []string{
					projectID,
					path.ResultID,
					path.SASTSimilarityID,
					path.SimilarityID,
					path.DetectionDate,
				})
			}
		}
	}
	return rows
}
// Used for the similarity mapping file where the apostrophe marker is not wanted.
func WriteAllToCsv(records [][]string) []byte {
	rows := make([]string, len(records))
	for i, record := range records {
		rows[i] = strings.Join(record, ",") + "\n"
	}
	return []byte(strings.Join(rows, ""))
}

func sanitize(cell string) string {
	escapedCell := strings.ReplaceAll(cell, `"`, `""`)
	return fmt.Sprintf(`"'%s"`, escapedCell)
}
