// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Feed is an object representing the database table.
type Feed struct {
	FeedID              int       `boil:"feed_id" json:"feed_id" toml:"feed_id" yaml:"feed_id"`
	Title               string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	URL                 string    `boil:"url" json:"url" toml:"url" yaml:"url"`
	RequestFailureCount int       `boil:"request_failure_count" json:"request_failure_count" toml:"request_failure_count" yaml:"request_failure_count"`
	CreatedAt           null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *feedR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L feedL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FeedColumns = struct {
	FeedID              string
	Title               string
	URL                 string
	RequestFailureCount string
	CreatedAt           string
}{
	FeedID:              "feed_id",
	Title:               "title",
	URL:                 "url",
	RequestFailureCount: "request_failure_count",
	CreatedAt:           "created_at",
}

// Generated where

var FeedWhere = struct {
	FeedID              whereHelperint
	Title               whereHelperstring
	URL                 whereHelperstring
	RequestFailureCount whereHelperint
	CreatedAt           whereHelpernull_Time
}{
	FeedID:              whereHelperint{field: "\"feeds\".\"feed_id\""},
	Title:               whereHelperstring{field: "\"feeds\".\"title\""},
	URL:                 whereHelperstring{field: "\"feeds\".\"url\""},
	RequestFailureCount: whereHelperint{field: "\"feeds\".\"request_failure_count\""},
	CreatedAt:           whereHelpernull_Time{field: "\"feeds\".\"created_at\""},
}

// FeedRels is where relationship names are stored.
var FeedRels = struct {
	Articles      string
	Subscriptions string
}{
	Articles:      "Articles",
	Subscriptions: "Subscriptions",
}

// feedR is where relationships are stored.
type feedR struct {
	Articles      ArticleSlice
	Subscriptions SubscriptionSlice
}

// NewStruct creates a new relationship struct
func (*feedR) NewStruct() *feedR {
	return &feedR{}
}

// feedL is where Load methods for each relationship are stored.
type feedL struct{}

var (
	feedAllColumns            = []string{"feed_id", "title", "url", "request_failure_count", "created_at"}
	feedColumnsWithoutDefault = []string{"title", "url"}
	feedColumnsWithDefault    = []string{"feed_id", "request_failure_count", "created_at"}
	feedPrimaryKeyColumns     = []string{"feed_id"}
)

type (
	// FeedSlice is an alias for a slice of pointers to Feed.
	// This should generally be used opposed to []Feed.
	FeedSlice []*Feed

	feedQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	feedType                 = reflect.TypeOf(&Feed{})
	feedMapping              = queries.MakeStructMapping(feedType)
	feedPrimaryKeyMapping, _ = queries.BindMapping(feedType, feedMapping, feedPrimaryKeyColumns)
	feedInsertCacheMut       sync.RWMutex
	feedInsertCache          = make(map[string]insertCache)
	feedUpdateCacheMut       sync.RWMutex
	feedUpdateCache          = make(map[string]updateCache)
	feedUpsertCacheMut       sync.RWMutex
	feedUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single feed record from the query using the global executor.
func (q feedQuery) OneG() (*Feed, error) {
	return q.One(boil.GetDB())
}

// One returns a single feed record from the query.
func (q feedQuery) One(exec boil.Executor) (*Feed, error) {
	o := &Feed{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feeds")
	}

	return o, nil
}

// AllG returns all Feed records from the query using the global executor.
func (q feedQuery) AllG() (FeedSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Feed records from the query.
func (q feedQuery) All(exec boil.Executor) (FeedSlice, error) {
	var o []*Feed

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Feed slice")
	}

	return o, nil
}

// CountG returns the count of all Feed records in the query, and panics on error.
func (q feedQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Feed records in the query.
func (q feedQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feeds rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q feedQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q feedQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feeds exists")
	}

	return count > 0, nil
}

// Articles retrieves all the article's Articles with an executor.
func (o *Feed) Articles(mods ...qm.QueryMod) articleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"articles\".\"feed_id\"=?", o.FeedID),
	)

	query := Articles(queryMods...)
	queries.SetFrom(query.Query, "\"articles\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"articles\".*"})
	}

	return query
}

// Subscriptions retrieves all the subscription's Subscriptions with an executor.
func (o *Feed) Subscriptions(mods ...qm.QueryMod) subscriptionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"subscriptions\".\"feed_id\"=?", o.FeedID),
	)

	query := Subscriptions(queryMods...)
	queries.SetFrom(query.Query, "\"subscriptions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"subscriptions\".*"})
	}

	return query
}

// LoadArticles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (feedL) LoadArticles(e boil.Executor, singular bool, maybeFeed interface{}, mods queries.Applicator) error {
	var slice []*Feed
	var object *Feed

	if singular {
		object = maybeFeed.(*Feed)
	} else {
		slice = *maybeFeed.(*[]*Feed)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &feedR{}
		}
		args = append(args, object.FeedID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &feedR{}
			}

			for _, a := range args {
				if a == obj.FeedID {
					continue Outer
				}
			}

			args = append(args, obj.FeedID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`articles`), qm.WhereIn(`articles.feed_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load articles")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice articles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on articles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for articles")
	}

	if singular {
		object.R.Articles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &articleR{}
			}
			foreign.R.Feed = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeedID == foreign.FeedID {
				local.R.Articles = append(local.R.Articles, foreign)
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.Feed = local
				break
			}
		}
	}

	return nil
}

// LoadSubscriptions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (feedL) LoadSubscriptions(e boil.Executor, singular bool, maybeFeed interface{}, mods queries.Applicator) error {
	var slice []*Feed
	var object *Feed

	if singular {
		object = maybeFeed.(*Feed)
	} else {
		slice = *maybeFeed.(*[]*Feed)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &feedR{}
		}
		args = append(args, object.FeedID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &feedR{}
			}

			for _, a := range args {
				if a == obj.FeedID {
					continue Outer
				}
			}

			args = append(args, obj.FeedID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`subscriptions`), qm.WhereIn(`subscriptions.feed_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load subscriptions")
	}

	var resultSlice []*Subscription
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice subscriptions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on subscriptions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subscriptions")
	}

	if singular {
		object.R.Subscriptions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &subscriptionR{}
			}
			foreign.R.Feed = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeedID == foreign.FeedID {
				local.R.Subscriptions = append(local.R.Subscriptions, foreign)
				if foreign.R == nil {
					foreign.R = &subscriptionR{}
				}
				foreign.R.Feed = local
				break
			}
		}
	}

	return nil
}

// AddArticlesG adds the given related objects to the existing relationships
// of the feed, optionally inserting them as new records.
// Appends related to o.R.Articles.
// Sets related.R.Feed appropriately.
// Uses the global database handle.
func (o *Feed) AddArticlesG(insert bool, related ...*Article) error {
	return o.AddArticles(boil.GetDB(), insert, related...)
}

// AddArticles adds the given related objects to the existing relationships
// of the feed, optionally inserting them as new records.
// Appends related to o.R.Articles.
// Sets related.R.Feed appropriately.
func (o *Feed) AddArticles(exec boil.Executor, insert bool, related ...*Article) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FeedID = o.FeedID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"articles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"feed_id"}),
				strmangle.WhereClause("\"", "\"", 2, articlePrimaryKeyColumns),
			)
			values := []interface{}{o.FeedID, rel.ArticleID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FeedID = o.FeedID
		}
	}

	if o.R == nil {
		o.R = &feedR{
			Articles: related,
		}
	} else {
		o.R.Articles = append(o.R.Articles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &articleR{
				Feed: o,
			}
		} else {
			rel.R.Feed = o
		}
	}
	return nil
}

// AddSubscriptionsG adds the given related objects to the existing relationships
// of the feed, optionally inserting them as new records.
// Appends related to o.R.Subscriptions.
// Sets related.R.Feed appropriately.
// Uses the global database handle.
func (o *Feed) AddSubscriptionsG(insert bool, related ...*Subscription) error {
	return o.AddSubscriptions(boil.GetDB(), insert, related...)
}

// AddSubscriptions adds the given related objects to the existing relationships
// of the feed, optionally inserting them as new records.
// Appends related to o.R.Subscriptions.
// Sets related.R.Feed appropriately.
func (o *Feed) AddSubscriptions(exec boil.Executor, insert bool, related ...*Subscription) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FeedID = o.FeedID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"subscriptions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"feed_id"}),
				strmangle.WhereClause("\"", "\"", 2, subscriptionPrimaryKeyColumns),
			)
			values := []interface{}{o.FeedID, rel.SubscriptionID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FeedID = o.FeedID
		}
	}

	if o.R == nil {
		o.R = &feedR{
			Subscriptions: related,
		}
	} else {
		o.R.Subscriptions = append(o.R.Subscriptions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &subscriptionR{
				Feed: o,
			}
		} else {
			rel.R.Feed = o
		}
	}
	return nil
}

// Feeds retrieves all the records using an executor.
func Feeds(mods ...qm.QueryMod) feedQuery {
	mods = append(mods, qm.From("\"feeds\""))
	return feedQuery{NewQuery(mods...)}
}

// FindFeedG retrieves a single record by ID.
func FindFeedG(feedID int, selectCols ...string) (*Feed, error) {
	return FindFeed(boil.GetDB(), feedID, selectCols...)
}

// FindFeed retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeed(exec boil.Executor, feedID int, selectCols ...string) (*Feed, error) {
	feedObj := &Feed{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feeds\" where \"feed_id\"=$1", sel,
	)

	q := queries.Raw(query, feedID)

	err := q.Bind(nil, exec, feedObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feeds")
	}

	return feedObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Feed) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Feed) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no feeds provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(feedColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	feedInsertCacheMut.RLock()
	cache, cached := feedInsertCache[key]
	feedInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			feedAllColumns,
			feedColumnsWithDefault,
			feedColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(feedType, feedMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(feedType, feedMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"feeds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"feeds\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into feeds")
	}

	if !cached {
		feedInsertCacheMut.Lock()
		feedInsertCache[key] = cache
		feedInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Feed record using the global executor.
// See Update for more documentation.
func (o *Feed) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Feed.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Feed) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	feedUpdateCacheMut.RLock()
	cache, cached := feedUpdateCache[key]
	feedUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			feedAllColumns,
			feedPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update feeds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feeds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, feedPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(feedType, feedMapping, append(wl, feedPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update feeds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for feeds")
	}

	if !cached {
		feedUpdateCacheMut.Lock()
		feedUpdateCache[key] = cache
		feedUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q feedQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q feedQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for feeds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for feeds")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeedSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeedSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), feedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"feeds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, feedPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in feed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all feed")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Feed) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Feed) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no feeds provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(feedColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	feedUpsertCacheMut.RLock()
	cache, cached := feedUpsertCache[key]
	feedUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			feedAllColumns,
			feedColumnsWithDefault,
			feedColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			feedAllColumns,
			feedPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert feeds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(feedPrimaryKeyColumns))
			copy(conflict, feedPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"feeds\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(feedType, feedMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(feedType, feedMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert feeds")
	}

	if !cached {
		feedUpsertCacheMut.Lock()
		feedUpsertCache[key] = cache
		feedUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Feed record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Feed) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Feed record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Feed) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Feed provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), feedPrimaryKeyMapping)
	sql := "DELETE FROM \"feeds\" WHERE \"feed_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from feeds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for feeds")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q feedQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no feedQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from feeds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for feeds")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o FeedSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeedSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), feedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"feeds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, feedPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from feed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for feeds")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Feed) ReloadG() error {
	if o == nil {
		return errors.New("models: no Feed provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Feed) Reload(exec boil.Executor) error {
	ret, err := FindFeed(exec, o.FeedID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeedSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeedSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeedSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FeedSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), feedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"feeds\".* FROM \"feeds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, feedPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeedSlice")
	}

	*o = slice

	return nil
}

// FeedExistsG checks if the Feed row exists.
func FeedExistsG(feedID int) (bool, error) {
	return FeedExists(boil.GetDB(), feedID)
}

// FeedExists checks if the Feed row exists.
func FeedExists(exec boil.Executor, feedID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"feeds\" where \"feed_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, feedID)
	}
	row := exec.QueryRow(sql, feedID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feeds exists")
	}

	return exists, nil
}
