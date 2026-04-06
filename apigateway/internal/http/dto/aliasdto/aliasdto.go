package aliasdto

// тут сохраняем то, что приходит от клиента на сервер и наоборот

type URLSaveRequest struct { //TODO: в рф чаще первым идет существительное в названии или глагол, URLSaveRequest или SaveURLRequest?
	URL string `json:"url"`
}

type URLSaveResponse struct {
	Message string `json:"message"`
	Alias   string `json:"alias"`
}

type URLGetResponse struct {
	URL     string `json:"url"`
	Message string `json:"message"`
	Alias   string `json:"alias"`
}
