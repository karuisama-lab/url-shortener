package aliasdto

// тут сохраняем то, что приходит от клиента на сервер и наоборот

type URLSaveRequest struct { //TODO: в рф чаще первым идет существительное в названии или глагол, URLSaveRequest или SaveURLRequest?
	UserID string `json:"user_id"`
	URL    string `json:"url"`
}

type URLSaveResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
