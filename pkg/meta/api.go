package meta

import (
	"database/sql"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/consts"
	"github.com/DitoAdriel99/go-monsterdex/pkg/timeutil"
)

// ErrInvalidMetadata is an error when metadata is invalid.
// This error usually returned by the implementation of Filter interface.
var ErrInvalidMetadata = errors.New("invalid metadata")

// Metadata represents a metadata for HTTP API.
type Metadata struct {
	Pagination
	Filtering
	*DateRange `json:"date_range,omitempty"`
}

// MetadataFromURL gets metadata from the given request url.
func MetadataFromURL(u url.Values) Metadata {
	return Metadata{
		Pagination: PaginationFromURL(u),
		Filtering:  FilterFromURL(u),
	}
}

// Pagination is a meta data for pagination.
type Pagination struct {
	PerPage   int   `json:"per_page"`
	Page      int   `json:"page"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"`
}

// PaginationFromURL gets pagination meta from request URL.
func PaginationFromURL(u url.Values) Pagination {
	p := Pagination{
		PerPage: consts.PaginationPerPageMinDefault,
		Page:    1,
	}

	pps := u.Get("per_page")
	if v, err := strconv.Atoi(pps); err == nil {
		if v <= 0 {
			v = consts.PaginationPerPageMinDefault
		}

		if v >= consts.PaginationPerPageMaxDefault {
			v = consts.PaginationPerPageMaxDefault
		}

		p.PerPage = v
	}

	ps := u.Get("page")
	if v, err := strconv.Atoi(ps); err == nil {
		if v < 1 {
			v = 1
		}

		p.Page = v
	}

	return p
}

// SortXXX are default values for order_type query params.
const (
	SortAscending  = "asc"
	SortDescending = "desc"
)

// Filtering represents a filterable fields.
type Filtering struct {
	OrderBy   string   `json:"order_by"`
	OrderType string   `json:"order_type"`
	Search    string   `json:"search,omitempty"`
	SearchBy  string   `json:"search_by,omitempty"`
	IsCatched string   `json:"is_catched,omitempty"`
	Type      []string `json:"type,omitempty"`
	Name      string   `json:"name,omitempty"`
	MonsterID int      `json:"monster_id,omitempty"`
}

// FilterFromURL gets filter values from query params.
func FilterFromURL(u url.Values) Filtering {
	f := Filtering{
		OrderBy:   "created_at",
		OrderType: SortDescending,
	}

	ob := u.Get("order_by")
	if len(ob) != 0 {
		f.OrderBy = strings.ToLower(strings.ToLower(ob))
	}

	nm := u.Get("name")
	if len(nm) != 0 {
		f.Name = strings.ToLower(strings.ToLower(nm))
	}

	ot := u.Get("order_type")
	if len(ot) != 0 {
		ot = strings.TrimSpace(strings.ToLower(ot))
		if ot == SortAscending {
			f.OrderType = SortAscending
		}
	}

	search := strings.TrimSpace(u.Get("search"))
	if len(search) == 0 {
		search = strings.TrimSpace(u.Get("keyword"))
	}

	if len(search) != 0 {
		f.Search = search
	}

	searchBy := strings.TrimSpace(u.Get("search_by"))
	if len(searchBy) != 0 {
		f.SearchBy = searchBy
	}

	ic := u.Get("is_catched")
	if len(ic) != 0 {
		f.IsCatched = strings.ToLower(strings.ToLower(ic))
	}

	mi := u.Get("monster_id")
	if len(mi) != 0 {
		i, _ := strconv.Atoi(mi)
		f.MonsterID = i
	}

	tp := u["type"]

	if len(tp) > 0 {
		var types []string
		types = append(types, tp...)
		f.Type = types
	}
	return f
}

type DateRange struct {
	Field string    `json:"field"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func DateRangeFromURL(u url.Values, field string, startQuery, endQuery string) (*DateRange, error) {
	ts := u.Get(startQuery)
	te := u.Get(endQuery)
	if len(ts) == 0 || len(te) == 0 {
		return nil, nil
	}

	dr := DateRange{
		Field: "created_at",
		Start: time.Time{},
		End:   time.Time{},
	}

	if v := u.Get(field); len(v) != 0 {
		dr.Field = strings.TrimSpace(strings.ToLower(v))
	}

	t, err := time.Parse("2006-01-02", ts)
	if err != nil {
		return nil, ErrInvalidMetadata
	}

	dr.Start = t

	t, err = time.Parse("2006-01-02", te)
	if err != nil {
		return nil, ErrInvalidMetadata
	}

	dr.End = t

	return &dr, nil
}

// Filter knows how to validate filterable fields.
// This Filter usually implemented by Repository.
type MetaFilter interface {
	// Sortable returns true if a given field is allowed for sorting.
	Sortable(field string) bool
}

func Filter(metadata *Metadata) *Query {
	var form, end sql.NullTime
	if metadata.DateRange != nil {

		form = sql.NullTime{
			Time:  timeutil.BeginOfDay(metadata.DateRange.Start),
			Valid: !metadata.DateRange.Start.IsZero(),
		}

		end = sql.NullTime{
			Time:  timeutil.BeginOfNextDay(metadata.DateRange.End),
			Valid: !metadata.DateRange.End.IsZero(),
		}
	}

	limit := metadata.PerPage
	offset := (metadata.Page - 1) * limit
	search := "%" + strings.ToLower(metadata.Search) + "%"

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
		DateFrom:       form,
		DateEnd:        end,
	}

	return &q
}

type Query struct {
	OrderBy        string
	OrderDirection string
	Search         string
	SearchBy       string
	Status         string
	Limit          int
	Offset         int
	DateFrom       sql.NullTime
	DateEnd        sql.NullTime
	IsCatched      string
	Type           []string
	Name           string
	MonsterID      int
}

func FromMetadata(metadata *Metadata, filter MetaFilter) (*Query, error) {
	if !filter.Sortable(metadata.OrderBy) {
		return nil, ErrInvalidMetadata
	}

	var form, end sql.NullTime
	if metadata.DateRange != nil {
		if !filter.Sortable(metadata.DateRange.Field) {
			return nil, ErrInvalidMetadata
		}

		form = sql.NullTime{
			Time:  timeutil.BeginOfDay(metadata.DateRange.Start),
			Valid: !metadata.DateRange.Start.IsZero(),
		}

		end = sql.NullTime{
			Time:  timeutil.BeginOfNextDay(metadata.DateRange.End),
			Valid: !metadata.DateRange.End.IsZero(),
		}
	}

	limit := metadata.PerPage
	offset := (metadata.Page - 1) * limit
	search := "%" + strings.ToLower(metadata.Search) + "%"

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
		DateFrom:       form,
		DateEnd:        end,
		SearchBy:       strings.ToLower(metadata.SearchBy),
		MonsterID:      metadata.MonsterID,
	}

	return &q, nil
}

func ParseMetaData(metadata *Metadata, filter MetaFilter) (*Query, error) {
	if !filter.Sortable(metadata.OrderBy) {
		metadata.OrderBy = "created_at"
	}

	var form, end sql.NullTime
	if metadata.DateRange != nil {
		if !filter.Sortable(metadata.DateRange.Field) {
			return nil, ErrInvalidMetadata
		}

		form = sql.NullTime{
			Time:  timeutil.BeginOfDay(metadata.DateRange.Start),
			Valid: !metadata.DateRange.Start.IsZero(),
		}

		end = sql.NullTime{
			Time:  timeutil.BeginOfNextDay(metadata.DateRange.End),
			Valid: !metadata.DateRange.End.IsZero(),
		}
	}

	limit := metadata.PerPage
	offset := (metadata.Page - 1) * limit
	search := "%" + strings.ToLower(metadata.Search) + "%"

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
		DateFrom:       form,
		DateEnd:        end,
		SearchBy:       strings.ToLower(metadata.SearchBy),
		Name:           strings.ToLower(metadata.Name),
	}

	q.IsCatched = metadata.IsCatched
	q.Type = metadata.Type
	return &q, nil
}

func SortableDownloadCFC(field string) bool {
	switch field {
	case "created_at":
		return true
	default:
		return false
	}
}

func ParseMetaDataDownloadCFC(metadata *Metadata, filter MetaFilter) (*Query, error) {
	if !SortableDownloadCFC(metadata.OrderBy) {
		metadata.OrderBy = "created_at"
	}

	var form, end sql.NullTime
	if metadata.DateRange != nil {
		if !filter.Sortable(metadata.DateRange.Field) {
			return nil, ErrInvalidMetadata
		}

		form = sql.NullTime{
			Time:  timeutil.BeginOfDay(metadata.DateRange.Start),
			Valid: !metadata.DateRange.Start.IsZero(),
		}

		end = sql.NullTime{
			Time:  timeutil.BeginOfNextDay(metadata.DateRange.End),
			Valid: !metadata.DateRange.End.IsZero(),
		}
	}

	limit := metadata.PerPage
	offset := (metadata.Page - 1) * limit
	search := "%" + strings.ToLower(metadata.Search) + "%"

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Search:         search,
		Limit:          limit,
		Offset:         offset,
		DateFrom:       form,
		DateEnd:        end,
		SearchBy:       strings.ToLower(metadata.SearchBy),
	}

	return &q, nil
}
