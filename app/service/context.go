package service

// Context 上下文管理服务
var Context = contextService{}

type contextService struct{}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
//func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
//	r.SetCtxVar(model.ContextKey, customCtx)
//}
//
//// Get 获得上下文变量，如果没有设置，那么返回nil
//func (s *contextService) Get(ctx context.Context) *model.Context {
//	value := ctx.Value(model.ContextKey)
//	if value == nil {
//		return nil
//	}
//	localCtx, ok := value.(*model.Context)
//	if ok {
//		return localCtx
//	}
//	return nil
//}
//
//// SetUserInfo 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserInfo(ctx context.Context, ctxUser *model.UserInfo) {
//	s.Get(ctx).User = &model.ContextUser{}
//	s.Get(ctx).User.UserInfo = ctxUser
//}
//
//// SetUserEmployee 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserEmployee(ctx context.Context, ctxUser *model.Employee) {
//	s.Get(ctx).User.EmployeeInfo = ctxUser
//}
//
//// SetUserJob 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserJob(ctx context.Context, ctxUser []entity.Job) {
//	s.Get(ctx).User.JobInfo = ctxUser
//}
//
//// SetUserDepartment 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserDepartment(ctx context.Context, ctxUser []entity.Department) {
//	s.Get(ctx).User.DepartmentInfo = ctxUser
//}
//
//// SetUserProduct 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserProduct(ctx context.Context, ctxUser []*model.ProductMember) {
//	s.Get(ctx).User.ProductMemberList = ctxUser
//}
//
//// SetUserProductIds 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserProductIds(ctx context.Context, productIds []uint) {
//	s.Get(ctx).User.ProductIds = productIds
//}
//
//// SetUserRoleLevel 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserRoleLevel(ctx context.Context, roleLevel int) {
//	s.Get(ctx).User.RoleLevel = roleLevel
//}
//
//// SetUserProductRole 将上下文信息设置到上下文请求中，注意是完整覆盖
//func (s *contextService) SetUserProductRole(ctx context.Context, roleName int) {
//	s.Get(ctx).User.ProductRole = roleName
//}
