package yaygo

import (
	"encoding/json"
	"net/http"
)

type ReviewApi struct {
	s *Session
}

func newReviewApi(s *Session) *ReviewApi {
	return &ReviewApi{
		s: s,
	}
}



// 名前かぶったからとりまV1, V2...
type CreateReviewV2Params struct {
	ID         int    `json:"id,omitempty"`
	Comment    string `json:"comment,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	ApiKey     string `json:"api_key,omitempty"`
	Timestamp  int    `json:"timestamp"`
	SignedInfo string `json:"signed_info"`
	Context    string `json:"context"`
}

func (r *ReviewApi) CreateReviewV2(params *CreateReviewV2Params) (st *Response, err error) {
	resp, err := r.s.request(http.MethodPost, EndpointChatRoomsV2 + "create_review_v2", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type CreateReviewV1Params struct {
	UserIDs    []int  `json:"user_ids[],omitempty"`
	Comment    string `json:"comment,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	ApiKey     string `json:"api_key,omitempty"`
	Timestamp  int    `json:"timestamp"`
	SignedInfo string `json:"signed_info"`
}

func (r *ReviewApi) CreateReviewV1(params *CreateReviewV1Params) (st *Response, err error) {
	resp, err := r.s.request(http.MethodPost, EndpointChatRoomsV1 + "create_review_v1", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type DeleteReviewsParams struct {
	ReviewIDs []int `json:"review_ids[],omitempty"`
}

func (r *ReviewApi) DeleteReviews(params *DeleteReviewsParams) (st *Response, err error) {
	resp, err := r.s.request(http.MethodDelete, EndpointChatRoomsV1 + "delete_reviews", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetMyReviewsParams struct {
	FromID int `json:"from_id,omitempty"`
}

func (r *ReviewApi) GetMyReviews(params *GetMyReviewsParams) (st *ReviewsResponse, err error) {
	resp, err := r.s.request(http.MethodGet, EndpointChatRoomsV1 + "get_my_reviews", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type GetReviewsParams struct {
	ID     int `json:"id,omitempty"`
	FromID int `json:"from_id,omitempty"`
}

func (r *ReviewApi) GetReviews(params *GetReviewsParams) (st *ReviewsResponse, err error) {
	resp, err := r.s.request(http.MethodGet, EndpointChatRoomsV2 + "get_reviews", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type PinReviewParams struct {
	ID int `json:"id,omitempty"`
}

func (r *ReviewApi) PinReview(params *PinReviewParams) (st *Response, err error) {
	resp, err := r.s.request(http.MethodPost, EndpointChatRoomsV1 + "pin_review", nil, params, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}

type UnpinReviewParams struct {
	ID int `json:"id,omitempty"`
}

func (r *ReviewApi) UnpinReview(params *UnpinReviewParams) (st *Response, err error) {
	resp, err := r.s.request(http.MethodDelete, EndpointChatRoomsV1 + "unpin_review", params, nil, false)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &st)

	return
}
