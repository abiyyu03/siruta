package http

import (
	"github.com/abiyyu03/siruta/entity"
	"github.com/abiyyu03/siruta/entity/constant"
	"github.com/abiyyu03/siruta/entity/model"
	"github.com/abiyyu03/siruta/usecase/letter"
	"github.com/gofiber/fiber/v2"
)

type LetterReqHttp struct {
	letterReqUsecase letter.LetterReqUsecaseInterface
}

type CombinedRequest struct {
	OutcomingLetter *model.OutcomingLetter      `json:"outcoming_letter"`
	CheckResident   *entity.CheckResidentMember `json:"check_resident"`
	Member          *model.Member               `json:"member"`
}

func NewLetterReqHttp(letterReqUC letter.LetterReqUsecaseInterface) *LetterReqHttp {
	return &LetterReqHttp{
		letterReqUsecase: letterReqUC,
	}
}

func (l *LetterReqHttp) CreateData(ctx *fiber.Ctx) error {
	var request *CombinedRequest

	if err := ctx.BodyParser(&request); err != nil {
		return entity.Error(ctx, fiber.StatusUnprocessableEntity, constant.Errors["UnprocessableEntity"].Message, constant.Errors["UnprocessableEntity"].Clue)
	}

	return l.letterReqUsecase.StoreOutcomingLetter(
		ctx,
		request.Member,
		request.OutcomingLetter,
		request.CheckResident.MemberStatus,
		request.CheckResident.BirthDate,
		request.CheckResident.NikNumber,
	)
}

func (l *LetterReqHttp) UpdateApprovalStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("letter_req_id")

	return l.letterReqUsecase.UpdateApprovalStatusByRT(ctx, id)
}
