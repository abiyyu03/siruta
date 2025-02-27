package register

// var userRepository = new(repository.UserRepository)

// func Register(ctx *fiber.Ctx) error {
// 	var user *model.User

// 	if err := ctx.BodyParser(&user); err != nil {
// 		return entity.Error(
// 	ctx,
// 	fiber.ErrBadRequest.Code,
// 	constant.Errors["UnprocessableEntity"].Message,
// 	constant.Errors["UnprocessableEntity"].Clue,
// )
// 	}

// 	registeredUser, err := userRepository.RegisterUser(user, user.RoleID)

// 	if err != nil {
// 		return entity.Error( ctx, fiber.StatusInternalServerError, constant.Errors["InternalError"].Message, constant.Errors["InternalError"].Clue)
// 	}

// 	return entity.Success(ctx, &registeredUser, "Registrasi Pengguna/User Berhasil")

// }
