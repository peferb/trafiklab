package realtime_departures_v4

type Metro struct {
	GroupOfLine          string      `json:"GroupOfLine"`
	DisplayTime          string      `json:"DisplayTime"`
	TransportMode        string      `json:"TransportMode"`
	LineNumber           string      `json:"LineNumber"`
	Destination          string      `json:"Destination"`
	JourneyDirection     int         `json:"JourneyDirection"`
	StopAreaName         string      `json:"StopAreaName"`
	StopAreaNumber       int         `json:"StopAreaNumber"`
	StopPointNumber      int         `json:"StopPointNumber"`
	StopPointDesignation string      `json:"StopPointDesignation"`
	TimeTabledDateTime   string      `json:"TimeTabledDateTime"`
	ExpectedDateTime     string      `json:"ExpectedDateTime"`
	JourneyNumber        int         `json:"JourneyNumber"`
	Deviations           []Deviation `json:"Deviations"`
}

type Bus struct {
	GroupOfLine          string      `json:"GroupOfLine"`
	TransportMode        string      `json:"TransportMode"`
	LineNumber           string      `json:"LineNumber"`
	Destination          string      `json:"Destination"`
	JourneyDirection     int         `json:"JourneyDirection"`
	StopAreaName         string      `json:"StopAreaName"`
	StopAreaNumber       int         `json:"StopAreaNumber"`
	StopPointNumber      int         `json:"StopPointNumber"`
	StopPointDesignation string      `json:"StopPointDesignation"`
	TimeTabledDateTime   string      `json:"TimeTabledDateTime"`
	ExpectedDateTime     string      `json:"ExpectedDateTime"`
	DisplayTime          string      `json:"DisplayTime"`
	JourneyNumber        int         `json:"JourneyNumber"`
	Deviations           []Deviation `json:"Deviations"`
}

type Train struct {
	SecondaryDestinationName string      `json:"SecondaryDestinationName"`
	GroupOfLine              string      `json:"GroupOfLine"`
	TransportMode            string      `json:"TransportMode"`
	LineNumber               string      `json:"LineNumber"`
	Destination              string      `json:"Destination"`
	JourneyDirection         int         `json:"JourneyDirection"`
	StopAreaName             string      `json:"StopAreaName"`
	StopAreaNumber           int         `json:"StopAreaNumber"`
	StopPointNumber          int         `json:"StopPointNumber"`
	StopPointDesignation     string      `json:"StopPointDesignation"`
	TimeTabledDateTime       string      `json:"TimeTabledDateTime"`
	ExpectedDateTime         string      `json:"ExpectedDateTime"`
	DisplayTime              string      `json:"DisplayTime"`
	JourneyNumber            int         `json:"JourneyNumber"`
	Deviations               []Deviation `json:"Deviations"`
}

type Trams struct {
	TransportMode        string      `json:"TransportMode"`
	LineNumber           string      `json:"LineNumber"`
	Destination          string      `json:"Destination"`
	JourneyDirection     int         `json:"JourneyDirection"`
	GroupOfLine          string      `json:"GroupOfLine"`
	StopAreaName         string      `json:"StopAreaName"`
	StopAreaNumber       int         `json:"StopAreaNumber"`
	StopPointNumber      int         `json:"StopPointNumber"`
	StopPointDesignation string      `json:"StopPointDesignation"`
	TimeTabledDateTime   string      `json:"TimeTabledDateTime"`
	ExpectedDateTime     string      `json:"ExpectedDateTime"`
	DisplayTime          string      `json:"DisplayTime"`
	JourneyNumber        int         `json:"JourneyNumber"`
	Deviations           []Deviation `json:"Deviations"`
}

type Ship struct {
	TransportMode        string      `json:"TransportMode"`
	LineNumber           string      `json:"LineNumber"`
	Destination          string      `json:"Destination"`
	JourneyDirection     int         `json:"JourneyDirection"`
	GroupOfLine          string      `json:"GroupOfLine"`
	StopAreaName         string      `json:"StopAreaName"`
	StopAreaNumber       int         `json:"StopAreaNumber"`
	StopPointNumber      int         `json:"StopPointNumber"`
	StopPointDesignation string      `json:"StopPointDesignation"`
	TimeTabledDateTime   string      `json:"TimeTabledDateTime"`
	ExpectedDateTime     string      `json:"ExpectedDateTime"`
	DisplayTime          string      `json:"DisplayTime"`
	JourneyNumber        int         `json:"JourneyNumber"`
	Deviations           []Deviation `json:"Deviations"`
}

type Deviation struct {
	Text            string `json:"Text"`
	Consequence     string `json:"Consequence"`
	ImportanceLevel int    `json:"ImportanceLevel"`
}

type StopInfo struct {
	StopAreaNumber int    `json:"StopAreaNumber"`
	StopAreaName   string `json:"StopAreaName"`
	TransportMode  string `json:"TransportMode"`
	GroupOfLine    string `json:"GroupOfLine"`
}

type Response struct {
	StatusCode    int    `json:"StatusCode"`
	Message       string `json:"Message"`
	ExecutionTime int32  `json:"ExecutionTime"`
	ResponseData  struct {
		LatestUpdate        string  `json:"LatestUpdate"`
		DataAge             int     `json:"DataAge"`
		Metros              []Metro `json:"Metros"`
		Buses               []Bus   `json:"Buses"`
		Trains              []Train `json:"Trains"`
		Trams               []Trams `json:"Trams"`
		Ships               []Ship  `json:"Ships"`
		StopPointDeviations []struct {
			StopInfo  StopInfo `json:"StopInfo"`
			Deviation `json:"Deviation"`
		} `json:"StopPointDeviations"`
	} `json:"ResponseData"`
}
