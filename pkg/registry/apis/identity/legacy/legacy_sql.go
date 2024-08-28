package legacy

import (
	"context"
	"database/sql"
	"fmt"
	"text/template"

	"github.com/grafana/authlib/claims"
	"github.com/grafana/grafana/pkg/services/team"
	"github.com/grafana/grafana/pkg/services/user"
	"github.com/grafana/grafana/pkg/storage/legacysql"
	"github.com/grafana/grafana/pkg/storage/unified/sql/sqltemplate"
)

var (
	_ LegacyIdentityStore = (*legacySQLStore)(nil)
)

func NewLegacySQLStores(sql legacysql.LegacyDatabaseProvider) LegacyIdentityStore {
	return &legacySQLStore{
		sql: sql,
	}
}

type legacySQLStore struct {
	sql legacysql.LegacyDatabaseProvider
}

// ListTeams implements LegacyIdentityStore.
func (s *legacySQLStore) ListTeams(ctx context.Context, ns claims.NamespaceInfo, query ListTeamQuery) (*ListTeamResult, error) {
	// for continue
	query.Pagination.Limit += 1
	query.OrgID = ns.OrgID
	if ns.OrgID == 0 {
		return nil, fmt.Errorf("expected non zero orgID")
	}

	sql, err := s.sql(ctx)
	if err != nil {
		return nil, err
	}

	req := newListTeams(sql, &query)
	q, err := sqltemplate.Execute(sqlQueryTeams, req)
	if err != nil {
		return nil, fmt.Errorf("execute template %q: %w", sqlQueryTeams.Name(), err)
	}

	rows, err := sql.DB.GetSqlxSession().Query(ctx, q, req.GetArgs()...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	res := &ListTeamResult{}
	if err != nil {
		return nil, err
	}

	var lastID int64
	for rows.Next() {
		t := team.Team{}
		err = rows.Scan(&t.ID, &t.UID, &t.Name, &t.Email, &t.Created, &t.Updated)
		if err != nil {
			return res, err
		}

		lastID = t.ID
		res.Teams = append(res.Teams, t)
		if len(res.Teams) > int(query.Pagination.Limit)-1 {
			res.Teams = res.Teams[0 : len(res.Teams)-1]
			res.Continue = lastID
			break
		}
	}

	if query.UID == "" {
		res.RV, err = sql.GetResourceVersion(ctx, "team", "updated")
	}

	return res, err
}

// ListUsers implements LegacyIdentityStore.
func (s *legacySQLStore) ListUsers(ctx context.Context, ns claims.NamespaceInfo, query ListUserQuery) (*ListUserResult, error) {
	// for continue
	limit := int(query.Pagination.Limit)
	query.Pagination.Limit += 1

	query.OrgID = ns.OrgID
	if ns.OrgID == 0 {
		return nil, fmt.Errorf("expected non zero orgID")
	}

	sql, err := s.sql(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.queryUsers(ctx, sql, sqlQueryUsers, newListUser(sql, &query), limit)
	if err == nil && query.UID != "" {
		res.RV, err = sql.GetResourceVersion(ctx, "user", "updated")
	}

	return res, err
}

// GetDisplay implements LegacyIdentityStore.
func (s *legacySQLStore) GetDisplay(ctx context.Context, ns claims.NamespaceInfo, query GetUserDisplayQuery) (*ListUserResult, error) {
	query.OrgID = ns.OrgID
	if ns.OrgID == 0 {
		return nil, fmt.Errorf("expected non zero orgID")
	}

	sql, err := s.sql(ctx)
	if err != nil {
		return nil, err
	}

	return s.queryUsers(ctx, sql, sqlQueryDisplay, newGetDisplay(sql, &query), 10000)
}

func (s *legacySQLStore) queryUsers(ctx context.Context, sql *legacysql.LegacyDatabaseHelper, t *template.Template, req sqltemplate.Args, limit int) (*ListUserResult, error) {
	q, err := sqltemplate.Execute(t, req)
	if err != nil {
		return nil, fmt.Errorf("execute template %q: %w", sqlQueryUsers.Name(), err)
	}

	res := &ListUserResult{}
	rows, err := sql.DB.GetSqlxSession().Query(ctx, q, req.GetArgs()...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if err == nil {
		var lastID int64
		for rows.Next() {
			u := user.User{}
			err = rows.Scan(&u.OrgID, &u.ID, &u.UID, &u.Login, &u.Email, &u.Name,
				&u.Created, &u.Updated, &u.IsServiceAccount, &u.IsDisabled, &u.IsAdmin,
			)
			if err != nil {
				return res, err
			}

			lastID = u.ID
			res.Users = append(res.Users, u)
			if len(res.Users) > limit {
				res.Users = res.Users[0 : len(res.Users)-1]
				res.Continue = lastID
				break
			}
		}
	}

	return res, err
}

// ListTeamsBindings implements LegacyIdentityStore.
func (s *legacySQLStore) ListTeamBindings(ctx context.Context, ns claims.NamespaceInfo, query ListTeamBindingsQuery) (*ListTeamBindingsResult, error) {
	// for continue
	query.Pagination.Limit += 1
	query.OrgID = ns.OrgID
	if query.OrgID == 0 {
		return nil, fmt.Errorf("expected non zero orgID")
	}

	sql, err := s.sql(ctx)
	if err != nil {
		return nil, err
	}

	req := newListTeamBindings(sql, &query)
	q, err := sqltemplate.Execute(sqlQueryTeamBindings, req)
	if err != nil {
		return nil, fmt.Errorf("execute template %q: %w", sqlQueryTeams.Name(), err)
	}

	rows, err := sql.DB.GetSqlxSession().Query(ctx, q, req.GetArgs()...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	res := &ListTeamBindingsResult{}
	grouped := map[string][]TeamMember{}

	var lastID int64
	var atTeamLimit bool

	for rows.Next() {
		m := TeamMember{}
		err = rows.Scan(&m.ID, &m.TeamUID, &m.TeamID, &m.UserUID, &m.Created, &m.Updated, &m.Permission)
		if err != nil {
			return res, err
		}

		lastID = m.TeamID
		members, ok := grouped[m.TeamUID]
		if ok {
			grouped[m.TeamUID] = append(members, m)
		} else if !atTeamLimit {
			grouped[m.TeamUID] = []TeamMember{m}
		}

		if len(grouped) >= int(query.Pagination.Limit)-1 {
			atTeamLimit = true
			res.Continue = lastID
		}
	}

	if query.UID == "" {
		res.RV, err = sql.GetResourceVersion(ctx, "team_member", "updated")
	}

	res.Bindings = make([]TeamBinding, 0, len(grouped))
	for uid, members := range grouped {
		res.Bindings = append(res.Bindings, TeamBinding{
			TeamUID: uid,
			Members: members,
		})
	}

	return res, err
}

// ListTeamMembers implements LegacyIdentityStore.
func (s *legacySQLStore) ListTeamMembers(ctx context.Context, ns claims.NamespaceInfo, query ListTeamMembersQuery) (*ListTeamMembersResult, error) {
	query.Pagination.Limit += 1
	query.OrgID = ns.OrgID
	if query.OrgID == 0 {
		return nil, fmt.Errorf("expected non zero org id")
	}

	sql, err := s.sql(ctx)
	if err != nil {
		return nil, err
	}

	req := newListTeamMembers(sql, &query)
	q, err := sqltemplate.Execute(sqlQueryTeamMembers, req)
	if err != nil {
		return nil, fmt.Errorf("execute template %q: %w", sqlQueryTeams.Name(), err)
	}

	rows, err := sql.DB.GetSqlxSession().Query(ctx, q, req.GetArgs()...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	res := &ListTeamMembersResult{}
	var lastID int64
	for rows.Next() {
		m, err := scanMember(rows)
		if err != nil {
			return nil, err
		}

		lastID = m.ID
		res.Members = append(res.Members, m)
		if len(res.Members) > int(query.Pagination.Limit)-1 {
			res.Continue = lastID
			res.Members = res.Members[0 : len(res.Members)-1]
			break
		}
	}

	return res, err
}

func scanMember(rows *sql.Rows) (TeamMember, error) {
	m := TeamMember{}
	err := rows.Scan(&m.ID, &m.TeamUID, &m.TeamID, &m.UserUID, &m.UserID, &m.Name, &m.Email, &m.Username, &m.External, &m.Created, &m.Updated, &m.Permission)
	return m, err
}

// GetUserTeams implements LegacyIdentityStore.
func (s *legacySQLStore) GetUserTeams(ctx context.Context, ns claims.NamespaceInfo, uid string) ([]team.Team, error) {
	panic("unimplemented")
}