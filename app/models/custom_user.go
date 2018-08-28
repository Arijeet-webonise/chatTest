package models

// CustomPortalUserImpl implements CustomPortalUser
type CustomPortalUserImpl struct {
	DB XODB
}

// CustomPortalUser handle custom Portal user function
type CustomPortalUser interface {
	PortalUserByEmailID(email string) (*PortalUser, error)
	PortalUserByID(userID int) (*PortalUser, error)
}

// PortalUserByEmailID retrieves a row from 'public.portal_user' as a PortalUser.
//
// Generated from index 'unq_ename'.
func (serviceImpl *CustomPortalUserImpl) PortalUserByEmailID(email string) (*PortalUser, error) {
	return PortalUserByEmailID(serviceImpl.DB, email)
}

// PortalUserByID retrieves a row from 'public.portal_user' as a PortalUser.
//
// Generated from index 'portal_user_pkey'.
func (serviceImpl *CustomPortalUserImpl) PortalUserByID(userID int) (*PortalUser, error) {
	return PortalUserByID(serviceImpl.DB, userID)
}
