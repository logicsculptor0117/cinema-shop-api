// Code generated by goctl. DO NOT EDIT.
package types

type FilmListInfo struct {
	FilmId   int    `json:"film_id"`
	Type     int    `json:"type"`
	FilmName string `json:"film_name"`
	CoverPic string `json:"cover_pic"`
}

type FilmListInfoResp struct {
	Data  []FilmListInfo `json:"data"`
	Count int            `json:"count"`
	Page  int            `json:"page"`
}

type FilmListReq struct {
	Status int `path:"status"`
	Page   int `path:"page"`
	Limit  int `path:"limit"`
}

type FilmListResp struct {
	Data    FilmListInfoResp `json:"data"`
	ErrCode int              `json:"err_code"`
	ErrMsg  string           `json:"err_msg"`
}

type FilmDetailInfo struct {
	FilmId   int    `json:"film_id"`
	Cate     int    `json:"cate"`
	FilmName string `json:"film_name"`
	CoverPic string `json:"cover_pic"`
	FilmDesc string `json:"film_desc"`
}

type FilmDetailReq struct {
	FilmId int `path:"film_id"`
}

type FilmDetailResp struct {
	Data    FilmDetailInfo `json:"data"`
	ErrCode int            `json:"err_code"`
	ErrMsg  string         `json:"err_msg"`
}
