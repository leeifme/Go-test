package vm

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// LoginViewModelOp strutc
type LoginViewModelOp struct{}

// AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}
