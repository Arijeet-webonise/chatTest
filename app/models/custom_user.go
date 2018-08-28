package models

// CustomPortalUserImpl implements CustomPortalUser
type CustomPortalUserImpl struct {
	DB XODB
}

// CustomPortalUser handle custom Portal user function
type CustomPortalUser interface {
	PortalUserByEmailID(email string) (*PortalUser, error)
}

// PortalUserByEmailID retrieves a row from 'public.portal_user' as a PortalUser.
//
// Generated from index 'unq_ename'.
func (serviceImpl *CustomPortalUserImpl) PortalUserByEmailID(email string) (*PortalUser, error) {
	return PortalUserByEmailID(serviceImpl.DB, email)
}
