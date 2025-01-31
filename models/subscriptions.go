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

// Subscription is an object representing the database table.
type Subscription struct {
	SubscriptionID int       `boil:"subscription_id" json:"subscription_id" toml:"subscription_id" yaml:"subscription_id"`
	FeedID         int       `boil:"feed_id" json:"feed_id" toml:"feed_id" yaml:"feed_id"`
	ChannelID      string    `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	CreatedAt      null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *subscriptionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L subscriptionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SubscriptionColumns = struct {
	SubscriptionID string
	FeedID         string
	ChannelID      string
	CreatedAt      string
}{
	SubscriptionID: "subscription_id",
	FeedID:         "feed_id",
	ChannelID:      "channel_id",
	CreatedAt:      "created_at",
}

// Generated where

var SubscriptionWhere = struct {
	SubscriptionID whereHelperint
	FeedID         whereHelperint
	ChannelID      whereHelperstring
	CreatedAt      whereHelpernull_Time
}{
	SubscriptionID: whereHelperint{field: "\"subscriptions\".\"subscription_id\""},
	FeedID:         whereHelperint{field: "\"subscriptions\".\"feed_id\""},
	ChannelID:      whereHelperstring{field: "\"subscriptions\".\"channel_id\""},
	CreatedAt:      whereHelpernull_Time{field: "\"subscriptions\".\"created_at\""},
}

// SubscriptionRels is where relationship names are stored.
var SubscriptionRels = struct {
	Feed string
}{
	Feed: "Feed",
}

// subscriptionR is where relationships are stored.
type subscriptionR struct {
	Feed *Feed
}

// NewStruct creates a new relationship struct
func (*subscriptionR) NewStruct() *subscriptionR {
	return &subscriptionR{}
}

// subscriptionL is where Load methods for each relationship are stored.
type subscriptionL struct{}

var (
	subscriptionAllColumns            = []string{"subscription_id", "feed_id", "channel_id", "created_at"}
	subscriptionColumnsWithoutDefault = []string{"feed_id", "channel_id"}
	subscriptionColumnsWithDefault    = []string{"subscription_id", "created_at"}
	subscriptionPrimaryKeyColumns     = []string{"subscription_id"}
)

type (
	// SubscriptionSlice is an alias for a slice of pointers to Subscription.
	// This should generally be used opposed to []Subscription.
	SubscriptionSlice []*Subscription

	subscriptionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	subscriptionType                 = reflect.TypeOf(&Subscription{})
	subscriptionMapping              = queries.MakeStructMapping(subscriptionType)
	subscriptionPrimaryKeyMapping, _ = queries.BindMapping(subscriptionType, subscriptionMapping, subscriptionPrimaryKeyColumns)
	subscriptionInsertCacheMut       sync.RWMutex
	subscriptionInsertCache          = make(map[string]insertCache)
	subscriptionUpdateCacheMut       sync.RWMutex
	subscriptionUpdateCache          = make(map[string]updateCache)
	subscriptionUpsertCacheMut       sync.RWMutex
	subscriptionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single subscription record from the query using the global executor.
func (q subscriptionQuery) OneG() (*Subscription, error) {
	return q.One(boil.GetDB())
}

// One returns a single subscription record from the query.
func (q subscriptionQuery) One(exec boil.Executor) (*Subscription, error) {
	o := &Subscription{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for subscriptions")
	}

	return o, nil
}

// AllG returns all Subscription records from the query using the global executor.
func (q subscriptionQuery) AllG() (SubscriptionSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Subscription records from the query.
func (q subscriptionQuery) All(exec boil.Executor) (SubscriptionSlice, error) {
	var o []*Subscription

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Subscription slice")
	}

	return o, nil
}

// CountG returns the count of all Subscription records in the query, and panics on error.
func (q subscriptionQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Subscription records in the query.
func (q subscriptionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count subscriptions rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q subscriptionQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q subscriptionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if subscriptions exists")
	}

	return count > 0, nil
}

// Feed pointed to by the foreign key.
func (o *Subscription) Feed(mods ...qm.QueryMod) feedQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"feed_id\" = ?", o.FeedID),
	}

	queryMods = append(queryMods, mods...)

	query := Feeds(queryMods...)
	queries.SetFrom(query.Query, "\"feeds\"")

	return query
}

// LoadFeed allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (subscriptionL) LoadFeed(e boil.Executor, singular bool, maybeSubscription interface{}, mods queries.Applicator) error {
	var slice []*Subscription
	var object *Subscription

	if singular {
		object = maybeSubscription.(*Subscription)
	} else {
		slice = *maybeSubscription.(*[]*Subscription)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subscriptionR{}
		}
		args = append(args, object.FeedID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subscriptionR{}
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

	query := NewQuery(qm.From(`feeds`), qm.WhereIn(`feeds.feed_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feed")
	}

	var resultSlice []*Feed
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feed")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for feeds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for feeds")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Feed = foreign
		if foreign.R == nil {
			foreign.R = &feedR{}
		}
		foreign.R.Subscriptions = append(foreign.R.Subscriptions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FeedID == foreign.FeedID {
				local.R.Feed = foreign
				if foreign.R == nil {
					foreign.R = &feedR{}
				}
				foreign.R.Subscriptions = append(foreign.R.Subscriptions, local)
				break
			}
		}
	}

	return nil
}

// SetFeedG of the subscription to the related item.
// Sets o.R.Feed to related.
// Adds o to related.R.Subscriptions.
// Uses the global database handle.
func (o *Subscription) SetFeedG(insert bool, related *Feed) error {
	return o.SetFeed(boil.GetDB(), insert, related)
}

// SetFeed of the subscription to the related item.
// Sets o.R.Feed to related.
// Adds o to related.R.Subscriptions.
func (o *Subscription) SetFeed(exec boil.Executor, insert bool, related *Feed) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"subscriptions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feed_id"}),
		strmangle.WhereClause("\"", "\"", 2, subscriptionPrimaryKeyColumns),
	)
	values := []interface{}{related.FeedID, o.SubscriptionID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeedID = related.FeedID
	if o.R == nil {
		o.R = &subscriptionR{
			Feed: related,
		}
	} else {
		o.R.Feed = related
	}

	if related.R == nil {
		related.R = &feedR{
			Subscriptions: SubscriptionSlice{o},
		}
	} else {
		related.R.Subscriptions = append(related.R.Subscriptions, o)
	}

	return nil
}

// Subscriptions retrieves all the records using an executor.
func Subscriptions(mods ...qm.QueryMod) subscriptionQuery {
	mods = append(mods, qm.From("\"subscriptions\""))
	return subscriptionQuery{NewQuery(mods...)}
}

// FindSubscriptionG retrieves a single record by ID.
func FindSubscriptionG(subscriptionID int, selectCols ...string) (*Subscription, error) {
	return FindSubscription(boil.GetDB(), subscriptionID, selectCols...)
}

// FindSubscription retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSubscription(exec boil.Executor, subscriptionID int, selectCols ...string) (*Subscription, error) {
	subscriptionObj := &Subscription{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"subscriptions\" where \"subscription_id\"=$1", sel,
	)

	q := queries.Raw(query, subscriptionID)

	err := q.Bind(nil, exec, subscriptionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from subscriptions")
	}

	return subscriptionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Subscription) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Subscription) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no subscriptions provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(subscriptionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	subscriptionInsertCacheMut.RLock()
	cache, cached := subscriptionInsertCache[key]
	subscriptionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			subscriptionAllColumns,
			subscriptionColumnsWithDefault,
			subscriptionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"subscriptions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"subscriptions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into subscriptions")
	}

	if !cached {
		subscriptionInsertCacheMut.Lock()
		subscriptionInsertCache[key] = cache
		subscriptionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Subscription record using the global executor.
// See Update for more documentation.
func (o *Subscription) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Subscription.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Subscription) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	subscriptionUpdateCacheMut.RLock()
	cache, cached := subscriptionUpdateCache[key]
	subscriptionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			subscriptionAllColumns,
			subscriptionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update subscriptions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"subscriptions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, subscriptionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, append(wl, subscriptionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update subscriptions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for subscriptions")
	}

	if !cached {
		subscriptionUpdateCacheMut.Lock()
		subscriptionUpdateCache[key] = cache
		subscriptionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q subscriptionQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q subscriptionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for subscriptions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for subscriptions")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SubscriptionSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SubscriptionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subscriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"subscriptions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, subscriptionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in subscription slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all subscription")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Subscription) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Subscription) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no subscriptions provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(subscriptionColumnsWithDefault, o)

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

	subscriptionUpsertCacheMut.RLock()
	cache, cached := subscriptionUpsertCache[key]
	subscriptionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			subscriptionAllColumns,
			subscriptionColumnsWithDefault,
			subscriptionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			subscriptionAllColumns,
			subscriptionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert subscriptions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(subscriptionPrimaryKeyColumns))
			copy(conflict, subscriptionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"subscriptions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert subscriptions")
	}

	if !cached {
		subscriptionUpsertCacheMut.Lock()
		subscriptionUpsertCache[key] = cache
		subscriptionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Subscription record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Subscription) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Subscription record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Subscription) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Subscription provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), subscriptionPrimaryKeyMapping)
	sql := "DELETE FROM \"subscriptions\" WHERE \"subscription_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from subscriptions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for subscriptions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q subscriptionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no subscriptionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from subscriptions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for subscriptions")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SubscriptionSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SubscriptionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subscriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"subscriptions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, subscriptionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from subscription slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for subscriptions")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Subscription) ReloadG() error {
	if o == nil {
		return errors.New("models: no Subscription provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Subscription) Reload(exec boil.Executor) error {
	ret, err := FindSubscription(exec, o.SubscriptionID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubscriptionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty SubscriptionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubscriptionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SubscriptionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subscriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"subscriptions\".* FROM \"subscriptions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, subscriptionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SubscriptionSlice")
	}

	*o = slice

	return nil
}

// SubscriptionExistsG checks if the Subscription row exists.
func SubscriptionExistsG(subscriptionID int) (bool, error) {
	return SubscriptionExists(boil.GetDB(), subscriptionID)
}

// SubscriptionExists checks if the Subscription row exists.
func SubscriptionExists(exec boil.Executor, subscriptionID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"subscriptions\" where \"subscription_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, subscriptionID)
	}
	row := exec.QueryRow(sql, subscriptionID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if subscriptions exists")
	}

	return exists, nil
}
