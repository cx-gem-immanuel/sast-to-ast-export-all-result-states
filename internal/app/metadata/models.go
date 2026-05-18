package metadata

type (
	Record struct {
		ProjectID  int
		ScanID     string
		PathErrors []PathError
		Queries    []*RecordQuery `json:"queries"`
	}

	PathError struct {
		ProjectID int
		ScanID    string
		PathID    string
		Reason    string
	}

	RecordQuery struct {
		QueryID string          `json:"queryId"`
		Results []*RecordResult `json:"results"`
	}

	RecordResult struct {
		ResultID string        `json:"resultId"`
		Paths    []*RecordPath `json:"paths"`
	}

	RecordPath struct {
		PathID           string `json:"pathId"`
		SimilarityID     string `json:"similarityId"`
		ResultID         string `json:"-"`
		SASTSimilarityID string `json:"-"`
		DetectionDate    string `json:"-"`
	}

	Query struct {
		QueryID  string
		Language string
		Name     string
		Group    string
		Results  []*Result
	}

	Result struct {
		PathID        string
		ResultID      string
		SimilarityID  string
		DetectionDate string
		FirstNode     Node
		LastNode      Node
	}

	Node struct {
		FileName string
		Name     string
		Line     string
		Column   string
	}

	SimilarityCalculationJob struct {
		ResultID, PathID,
		Filename1, Name1, Line1, Column1, MethodLine1,
		Filename2, Name2, Line2, Column2, MethodLine2,
		QueryID       string
		SimIDVersion  int
		DetectionDate string
	}

	SimilarityCalculationResult struct {
		Err                            error
		ResultID, PathID, SimilarityID string
		DetectionDate                  string
	}
)
