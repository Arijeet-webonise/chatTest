// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// AccessControlMapping represents a row from 'public.access_control_mapping'.
type AccessControlMapping struct {
	ID         int      `json:"id"`         // id
	AccessID   int      `json:"access_id"`  // access_id
	Role       Roletype `json:"role"`       // role
	Permission bool     `json:"permission"` // permission

	_exists, _deleted bool
}

type AccessControlMappingService interface {
	DoesAccessControlMappingExists(acm *AccessControlMapping) (bool, error)
	InsertAccessControlMapping(acm *AccessControlMapping) error
	UpdateAccessControlMapping(acm *AccessControlMapping) error
	UpsertAccessControlMapping(acm *AccessControlMapping) error
	DeleteAccessControlMapping(acm *AccessControlMapping) error
	GetAllAccessControlMappings() ([]*AccessControlMapping, error)
	GetChunkedAccessControlMappings(limit int, offset int) ([]*AccessControlMapping, error)
}

type AccessControlMappingServiceImpl struct {
	DB XODB
}

// Exists determines if the AccessControlMapping exists in the database.
func (serviceImpl *AccessControlMappingServiceImpl) DoesAccessControlMappingExists(acm *AccessControlMapping) (bool, error) {
	panic("not yet implemented")
}

// Insert inserts the AccessControlMapping to the database.
func (serviceImpl *AccessControlMappingServiceImpl) InsertAccessControlMapping(acm *AccessControlMapping) error {
	var err error

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.access_control_mapping (` +
		`access_id, role, permission` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, acm.AccessID, acm.Role, acm.Permission)
	err = serviceImpl.DB.QueryRow(sqlstr, acm.AccessID, acm.Role, acm.Permission).Scan(&acm.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the AccessControlMapping in the database.
func (serviceImpl *AccessControlMappingServiceImpl) UpdateAccessControlMapping(acm *AccessControlMapping) error {
	var err error

	// sql query
	const sqlstr = `UPDATE public.access_control_mapping SET (` +
		`access_id, role, permission` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, acm.AccessID, acm.Role, acm.Permission, acm.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, acm.AccessID, acm.Role, acm.Permission, acm.ID)
	return err
}

// Save saves the AccessControlMapping to the database.
/*
	func (acm *AccessControlMapping) Save(db XODB) error {
		if acm.Exists() {
			return acm.Update(db)
		}

		return acm.Insert(db)
	}
*/

// Upsert performs an upsert for AccessControlMapping.
//
// NOTE: PostgreSQL 9.5+ only
func (serviceImpl *AccessControlMappingServiceImpl) UpsertAccessControlMapping(acm *AccessControlMapping) error {
	var err error

	// sql query
	const sqlstr = `INSERT INTO public.access_control_mapping (` +
		`id, access_id, role, permission` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, access_id, role, permission` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.access_id, EXCLUDED.role, EXCLUDED.permission` +
		`)`

	// run query
	XOLog(sqlstr, acm.ID, acm.AccessID, acm.Role, acm.Permission)
	_, err = serviceImpl.DB.Exec(sqlstr, acm.ID, acm.AccessID, acm.Role, acm.Permission)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the AccessControlMapping from the database.
func (serviceImpl *AccessControlMappingServiceImpl) DeleteAccessControlMapping(acm *AccessControlMapping) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM public.access_control_mapping WHERE id = $1`

	// run query
	XOLog(sqlstr, acm.ID)
	_, err = serviceImpl.DB.Exec(sqlstr, acm.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllAccessControlMappings returns all rows from 'public.access_control_mapping',
// ordered by "created_at" in descending order.
func (serviceImpl *AccessControlMappingServiceImpl) GetAllAccessControlMappings() ([]*AccessControlMapping, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.access_control_mapping`

	q, err := serviceImpl.DB.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*AccessControlMapping
	for q.Next() {
		acm := AccessControlMapping{}

		// scan
		err = q.Scan(&acm.ID, &acm.AccessID, &acm.Role, &acm.Permission)
		if err != nil {
			return nil, err
		}

		res = append(res, &acm)
	}

	return res, nil
}

// GetChunkedAccessControlMappings returns pagingated rows from 'public.access_control_mapping',
// ordered by "created_at" in descending order.
func (serviceImpl *AccessControlMappingServiceImpl) GetChunkedAccessControlMappings(limit int, offset int) ([]*AccessControlMapping, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.access_control_mapping LIMIT $1 OFFSET $2`

	q, err := serviceImpl.DB.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*AccessControlMapping
	for q.Next() {
		acm := AccessControlMapping{}

		// scan
		err = q.Scan(&acm.ID, &acm.AccessID, &acm.Role, &acm.Permission)
		if err != nil {
			return nil, err
		}

		res = append(res, &acm)
	}

	return res, nil
}

// AccessControl returns the AccessControl associated with the AccessControlMapping's AccessID (access_id).
//
// Generated from foreign key 'access_control_mapping_access_id_fkey'.
func (acm *AccessControlMapping) AccessControl(db XODB) (*AccessControl, error) {
	return AccessControlByID(db, acm.AccessID)
}

// AccessControlMappingByID retrieves a row from 'public.access_control_mapping' as a AccessControlMapping.
//
// Generated from index 'access_control_mapping_pkey'.
func AccessControlMappingByID(db XODB, id int) (*AccessControlMapping, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, access_id, role, permission ` +
		`FROM public.access_control_mapping ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	acm := AccessControlMapping{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&acm.ID, &acm.AccessID, &acm.Role, &acm.Permission)
	if err != nil {
		return nil, err
	}

	return &acm, nil
}

// AccessControlMappingByAccessIDRole retrieves a row from 'public.access_control_mapping' as a AccessControlMapping.
//
// Generated from index 'accessid_role'.
func AccessControlMappingByAccessIDRole(db XODB, accessID int, role Roletype) (*AccessControlMapping, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, access_id, role, permission ` +
		`FROM public.access_control_mapping ` +
		`WHERE access_id = $1 AND role = $2`

	// run query
	XOLog(sqlstr, accessID, role)
	acm := AccessControlMapping{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, accessID, role).Scan(&acm.ID, &acm.AccessID, &acm.Role, &acm.Permission)
	if err != nil {
		return nil, err
	}

	return &acm, nil
}