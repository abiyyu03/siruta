package register

// var userRepository = new(repository.UserRepository)

// func Register(ctx *fiber.Ctx) error {
// 	var user *model.User

// 	if err := ctx.BodyParser(&user); err != nil {
// 		return entity.Error(ctx, fiber.StatusUnprocessableEntity, fiber.ErrUnprocessableEntity.Message)
// 	}

// 	registeredUser, err := userRepository.RegisterUser(user, user.RoleID)

// 	if err != nil {
// 		return entity.Error(ctx, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message)
// 	}

// 	return entity.Success(ctx, &registeredUser, "Register user successfully")

// }
