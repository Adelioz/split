package handlers

import (
	"net/http"

	"github.com/Adelioz/split/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(l *zap.Logger, s *service.Service) *mux.Router {
	router := mux.NewRouter()

	router.Use(
		RecoveryMiddleware(l),
		LoggingMiddleware(l),
	)

	userRouter := router.PathPrefix("/user/").Subrouter()

	userRouter.Path("/").Methods(http.MethodPost).Handler(AddUser(s))          // curl http://localhost:65000/user/ -X POST -v -d '{"id": 1, "username": "Adel"}'
	userRouter.Path("/").Methods(http.MethodPut).Handler(UpdateUser(s))        // curl http://localhost:65000/user/ -X PUT -v -d '{"id": 1, "username": "Adel"}'
	userRouter.Path("/{id:\\d+}/").Methods(http.MethodGet).Handler(GetUser(s)) // curl http://localhost:65000/user/12/ -X GET -v

	roomRouter := router.PathPrefix("/room/").Subrouter()

	roomRouter.Path("/").Methods(http.MethodPost).Handler(AddRoom(s))          // curl http://localhost:65000/room/ -X POST -v -d '{"id": 1, "name": "Tara"}'
	roomRouter.Path("/").Methods(http.MethodPut).Handler(UpdateRoom(s))        // curl http://localhost:65000/room/ -X PUT -v -d '{"id": 1, "name": "Tara"}'
	roomRouter.Path("/{id:\\d+}/").Methods(http.MethodGet).Handler(GetRoom(s)) // curl http://localhost:65000/room/12/ -X GET -v

	expenseRouter := router.PathPrefix("/expense/").Subrouter()

	expenseRouter.Path("/").Methods(http.MethodPost).Handler(AddExpense(s))                // curl http://localhost:65000/expense/ -X POST -v -d '{"id":1,"roomId":1,"daddyId":1,"currency":"EUR","tag":"Groceries","createdAt":"0001-01-01T00:00:00Z","title":"Groceries","desc":""}'
	expenseRouter.Path("/").Methods(http.MethodPut).Handler(UpdateExpense(s))              // curl http://localhost:65000/expense/ -X PUT -v -d '{"id":1,"roomId":1,"daddyId":1,"currency":"EUR","tag":"Groceries","createdAt":"0001-01-01T00:00:00Z","title":"Groceries","desc":""}'
	expenseRouter.Path("/{id:\\d+}/").Methods(http.MethodGet).Handler(GetExpense(s))       // curl http://localhost:65000/expense/12/ -X GET -v
	expenseRouter.Path("/{id:\\d+}/").Methods(http.MethodDelete).Handler(DeleteExpense(s)) // curl http://localhost:65000/expense/12/ -X DELETE -v

	return router
}
