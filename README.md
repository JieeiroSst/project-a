mockgen -destination=mocks/mock_doer.go -package=mocks github.com/JieeiroSst/itjob/users/internal/repository UserRepository
 
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/JieeiroSst/itjob/users/internal/usecase UserUsecase
 
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/JieeiroSst/itjob/users/internal/http UserHttp
 
mockgen -destination=mocks/mock_doer.go -package=mocks  github.com/JieeiroSst/itjob/users/internal/delivery/http DeliveryHttp
 
mockgen -destination=mocks/mock_doer.go -package=mocks  github.com/JieeiroSst/itjob/users/internal/db UserDB
 
github.com/JieeiroSst/itjob/users/internal
 