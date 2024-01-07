package apperrors

import "net/http"

var (
	MongoCollectorRepositoryCreateError = AppError{
		Message:  "Failed to create collector",
		Code:     "MONGO_COLLECTOR_REPOSITORY_CREATE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	MongoCollectorRepositoryUpdateError = AppError{
		Message:  "Failed to update collector",
		Code:     "MONGO__REPOSITORY_UPDATE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	MongoCollectorRepositoryUpdateMarshalError = AppError{
		Message:  "Failed to marshal collector",
		Code:     "MONGO_COLLECTOR_REPOSITORY_UPDATE_MARSHAL_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	MongoCollectorUpdateUnmarshalError = AppError{
		Message:  "Failed to unmarshal collector",
		Code:     "MONGO_COLLECTOR_REPOSITORY_UPDATE_UNMARSHAL_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	MongoCollectorRepositoryDeleteError = AppError{
		Message:  "Failed to delete collector",
		Code:     "MONGO_COLLECTOR_REPOSITORY_DELETE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	MongoCollectorRepositoryGetByIdErrNoDocuments = AppError{
		Message:  "Failed to get collector by id",
		Code:     "MONGO_COLLECTOR_REPOSITORY_GET_BY_ID_ERR_NO_DOCUMENTS",
		HTTPCode: http.StatusInternalServerError,
	}
	MongoCollectorRepositoryGetByIdError = AppError{
		Message:  "Failed to get collector by id",
		Code:     "MONGO_COLLECTOR_REPOSITORY_GET_BY_ID_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)
