// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/masseelch/elk/internal/integration/pets/ent"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

// handler has some convenience methods used on node-handlers.
type handler struct{}

// Bitmask to configure which routes to register.
type Routes uint32

func (rs Routes) has(r Routes) bool { return rs&r != 0 }

func FastHttpHandler(h http.HandlerFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		fasthttpH := fasthttpadaptor.NewFastHTTPHandlerFunc(h)
		fasthttpH(ctx)
		return nil
	}
}

const (
	BadgeCreate Routes = 1 << iota
	BadgeRead
	BadgeUpdate
	BadgeDelete
	BadgeList
	BadgeWearer
	BadgeRoutes = 1<<iota - 1
)

// BadgeHandler handles http crud operations on ent.Badge.
type BadgeHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewBadgeHandler(c *ent.Client, l *zap.Logger) *BadgeHandler {
	return &BadgeHandler{
		client: c,
		log:    l.With(zap.String("handler", "BadgeHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *BadgeHandler) Mount(r fiber.Router, rs Routes) {
	if rs.has(BadgeCreate) {
		r.Post("/", FastHttpHandler(h.Create))
	}
	if rs.has(BadgeRead) {
		r.Get("/:id", FastHttpHandler(h.Read))
	}
	if rs.has(BadgeUpdate) {
		r.Patch("/id", FastHttpHandler(h.Update))
	}
	if rs.has(BadgeDelete) {
		r.Delete("/id", FastHttpHandler(h.Delete))
	}
	if rs.has(BadgeList) {
		r.Get("/", FastHttpHandler(h.List))
	}
	if rs.has(BadgeWearer) {
		r.Get("/:id/wearer", FastHttpHandler(h.Wearer))
	}
}

const (
	PetCreate Routes = 1 << iota
	PetRead
	PetUpdate
	PetDelete
	PetList
	PetBadge
	PetProtege
	PetMentor
	PetSpouse
	PetToys
	PetParent
	PetChildren
	PetPlayGroups
	PetFriends
	PetRoutes = 1<<iota - 1
)

// PetHandler handles http crud operations on ent.Pet.
type PetHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewPetHandler(c *ent.Client, l *zap.Logger) *PetHandler {
	return &PetHandler{
		client: c,
		log:    l.With(zap.String("handler", "PetHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *PetHandler) Mount(r fiber.Router, rs Routes) {
	if rs.has(PetCreate) {
		r.Post("/", FastHttpHandler(h.Create))
	}
	if rs.has(PetRead) {
		r.Get("/:id", FastHttpHandler(h.Read))
	}
	if rs.has(PetUpdate) {
		r.Patch("/id", FastHttpHandler(h.Update))
	}
	if rs.has(PetDelete) {
		r.Delete("/id", FastHttpHandler(h.Delete))
	}
	if rs.has(PetList) {
		r.Get("/", FastHttpHandler(h.List))
	}
	if rs.has(PetBadge) {
		r.Get("/:id/badge", FastHttpHandler(h.Badge))
	}
	if rs.has(PetProtege) {
		r.Get("/:id/protege", FastHttpHandler(h.Protege))
	}
	if rs.has(PetMentor) {
		r.Get("/:id/mentor", FastHttpHandler(h.Mentor))
	}
	if rs.has(PetSpouse) {
		r.Get("/:id/spouse", FastHttpHandler(h.Spouse))
	}
	if rs.has(PetToys) {
		r.Get("/:id/toys", FastHttpHandler(h.Toys))
	}
	if rs.has(PetParent) {
		r.Get("/:id/parent", FastHttpHandler(h.Parent))
	}
	if rs.has(PetChildren) {
		r.Get("/:id/children", FastHttpHandler(h.Children))
	}
	if rs.has(PetPlayGroups) {
		r.Get("/:id/play_groups", FastHttpHandler(h.PlayGroups))
	}
	if rs.has(PetFriends) {
		r.Get("/:id/friends", FastHttpHandler(h.Friends))
	}
}

const (
	PlayGroupCreate Routes = 1 << iota
	PlayGroupRead
	PlayGroupUpdate
	PlayGroupDelete
	PlayGroupList
	PlayGroupParticipants
	PlayGroupRoutes = 1<<iota - 1
)

// PlayGroupHandler handles http crud operations on ent.PlayGroup.
type PlayGroupHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewPlayGroupHandler(c *ent.Client, l *zap.Logger) *PlayGroupHandler {
	return &PlayGroupHandler{
		client: c,
		log:    l.With(zap.String("handler", "PlayGroupHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *PlayGroupHandler) Mount(r fiber.Router, rs Routes) {
	if rs.has(PlayGroupCreate) {
		r.Post("/", FastHttpHandler(h.Create))
	}
	if rs.has(PlayGroupRead) {
		r.Get("/:id", FastHttpHandler(h.Read))
	}
	if rs.has(PlayGroupUpdate) {
		r.Patch("/id", FastHttpHandler(h.Update))
	}
	if rs.has(PlayGroupDelete) {
		r.Delete("/id", FastHttpHandler(h.Delete))
	}
	if rs.has(PlayGroupList) {
		r.Get("/", FastHttpHandler(h.List))
	}
	if rs.has(PlayGroupParticipants) {
		r.Get("/:id/participants", FastHttpHandler(h.Participants))
	}
}

const (
	ToyCreate Routes = 1 << iota
	ToyRead
	ToyUpdate
	ToyDelete
	ToyList
	ToyOwner
	ToyRoutes = 1<<iota - 1
)

// ToyHandler handles http crud operations on ent.Toy.
type ToyHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewToyHandler(c *ent.Client, l *zap.Logger) *ToyHandler {
	return &ToyHandler{
		client: c,
		log:    l.With(zap.String("handler", "ToyHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *ToyHandler) Mount(r fiber.Router, rs Routes) {
	if rs.has(ToyCreate) {
		r.Post("/", FastHttpHandler(h.Create))
	}
	if rs.has(ToyRead) {
		r.Get("/:id", FastHttpHandler(h.Read))
	}
	if rs.has(ToyUpdate) {
		r.Patch("/id", FastHttpHandler(h.Update))
	}
	if rs.has(ToyDelete) {
		r.Delete("/id", FastHttpHandler(h.Delete))
	}
	if rs.has(ToyList) {
		r.Get("/", FastHttpHandler(h.List))
	}
	if rs.has(ToyOwner) {
		r.Get("/:id/owner", FastHttpHandler(h.Owner))
	}
}

func stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func zapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}
