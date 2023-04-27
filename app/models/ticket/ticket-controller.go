package ticket

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kajtuszd/cinema-api/app/models/seat"
	"github.com/kajtuszd/cinema-api/app/models/user"
	"github.com/kajtuszd/cinema-api/app/validators"
	"net/http"
)

type TicketController interface {
	GetTicket(ctx *gin.Context)
	GetTickets(ctx *gin.Context)
	GetAllTickets(ctx *gin.Context)
	CreateTicket(ctx *gin.Context)
	DeleteTicket(ctx *gin.Context)
	UpdateTicket(ctx *gin.Context)
	handleError(ctx *gin.Context, err error) error
}

type ticketController struct {
	ticketService TicketService
	userService   user.UserService
	seatService   seat.SeatService
	validator     *validator.Validate
}

func NewController(t TicketService, u user.UserService, s seat.SeatService) TicketController {
	v := validator.New()
	v.RegisterValidation("password", validators.PasswordValidator)
	return &ticketController{
		ticketService: t,
		userService:   u,
		seatService:   s,
		validator:     v,
	}
}

type TicketInput struct {
	SeatID uint `json:"seat_id"`
}

func (c *ticketController) handleError(ctx *gin.Context, err error) error {
	if err != nil {
		if errors.Is(err, ErrTicketNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": ErrTicketNotFound.Error()})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (c *ticketController) GetTicket(ctx *gin.Context) {
	id := ctx.Param("id")
	ticket, err := c.ticketService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": ticket})
}

func (c *ticketController) GetTickets(ctx *gin.Context) {
	u := ctx.MustGet("user").(user.User)
	tickets, err := c.ticketService.GetTicketsByUser(u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tickets})
}

func (c *ticketController) GetAllTickets(ctx *gin.Context) {
	tickets, err := c.ticketService.GetAllTickets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tickets})
}

func (c *ticketController) CreateTicket(ctx *gin.Context) {
	var input TicketInput
	u := ctx.MustGet("user").(user.User)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userObj, err := c.userService.GetByUsername(fmt.Sprintf("%s", u.Username))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	seatObj, err := c.seatService.GetByID(fmt.Sprintf("%d", input.SeatID))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if c.seatService.IsSeatReserved(seatObj) {
		ctx.JSON(http.StatusOK, gin.H{"message": "seat is already reserved"})
		return
	}
	ticket := &Ticket{
		ID:     uint(uuid.New().ID()),
		UserID: userObj.ID,
		SeatID: input.SeatID,
		Owner:  *userObj,
		Seat:   *seatObj,
		Price:  0,
	}
	if err := c.validator.Struct(ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ticketService.Create(ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.seatService.SetSeatReserved(seatObj); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"ticket": ticket})
}

func (c *ticketController) DeleteTicket(ctx *gin.Context) {
	id := ctx.Param("id")
	ticket, err := c.ticketService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	if err = c.ticketService.Delete(ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.seatService.SetSeatAvailable(&ticket.Seat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ticket deleted successfully"})
}

func (c *ticketController) UpdateTicket(ctx *gin.Context) {
	var input TicketInput
	id := ctx.Param("id")
	ticket, err := c.ticketService.GetByID(id)
	oldSeat := ticket.Seat
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.SeatID > 0 {
		seatObj, err := c.seatService.GetByID(fmt.Sprintf("%d", input.SeatID))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if c.seatService.IsSeatReserved(seatObj) {
			ctx.JSON(http.StatusOK, gin.H{"message": "seat is already reserved"})
			return
		}
		ticket.SeatID = input.SeatID
		ticket.Seat = *seatObj
	}
	if err := c.validator.Struct(ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.ticketService.Update(ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.seatService.SetSeatAvailable(&oldSeat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.seatService.SetSeatReserved(&ticket.Seat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ticket updated successfully"})
}
