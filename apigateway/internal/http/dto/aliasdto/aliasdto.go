package aliasdto

type URLSaveRequest struct {
	URL string `json:"url"`
}

type URLSaveResponse struct {
	Message string `json:"message"`
	Alias   string `json:"alias"`
}

type URLGetResponse struct {
	Message string `json:"message"`
	URL     string `json:"url"`
	Alias   string `json:"alias"`
}
