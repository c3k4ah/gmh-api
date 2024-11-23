package services

import (
	"github.com/c3k4ah/gmh-api/internal/configs"
	"github.com/c3k4ah/gmh-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type AppartementResponse struct {
	Appartement models.Appartement `json:"appartement"`
}

type AppartementsResponse struct {
	Appartements []models.Appartement `json:"appartements"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}

// GetAppartement godoc
// @Summary Récupérer un appartement par ID
// @Description Récupère les détails d'un appartement spécifique en utilisant son ID
// @Tags appartements
// @Param id path string true "ID de l'appartement"
// @Produce json
// @Success 200 {object} AppartementResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /appartements/{id} [get]
func GetAppartement(ginCxt *gin.Context) {
	var (
		appartement models.Appartement
		result      interface{}
	)
	id := ginCxt.Param("id")
	error := configs.Database.Where("id = ?", id).Preload("Localisation").Find(&appartement).Error

	if error != nil {
		if gorm.ErrRecordNotFound == error {
			result = ErrorResponse{Error: "Appartement not found"}
			ginCxt.JSON(404, result)
			return
		}
		result = ErrorResponse{Error: "Error while fetching appartement"}
		ginCxt.JSON(500, result)
		return
	}

	result = AppartementResponse{Appartement: appartement}
	ginCxt.JSON(200, result)
}

// CreateAppartement godoc
// @Summary Créer un nouvel appartement
// @Description Crée un appartement avec les informations fournies dans le corps de la requête
// @Tags appartements
// @Accept json
// @Produce json
// @Param appartement body models.Appartement true "Appartement à créer"
// @Success 200 {object} AppartementResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /appartements [post]
func CreateAppartement(ginCxt *gin.Context) {
	var (
		appartement models.Appartement
		result      interface{}
	)
	if error := ginCxt.ShouldBindJSON(&appartement); error != nil {
		result = ErrorResponse{Error: "Invalid input"}
		ginCxt.JSON(400, result)
		return
	}

	error := configs.Database.Create(&appartement).Error
	if error != nil {
		result = ErrorResponse{Error: "Error while creating appartement"}
		ginCxt.JSON(500, result)
		return
	}

	result = AppartementResponse{Appartement: appartement}
	ginCxt.JSON(200, result)
}

// UpdateAppartement godoc
// @Summary Mettre à jour un appartement existant
// @Description Met à jour les informations d'un appartement existant en utilisant son ID
// @Tags appartements
// @Param id path string true "ID de l'appartement"
// @Accept json
// @Produce json
// @Param appartement body models.Appartement true "Données de l'appartement à mettre à jour"
// @Success 200 {object} AppartementResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /appartements/{id} [put]
func UpdateAppartement(ginCxt *gin.Context) {
	var (
		appartement models.Appartement
		result      interface{}
	)
	id := ginCxt.Param("id")

	// Récupérer l'appartement existant
	if error := configs.Database.Where("id = ?", id).First(&appartement).Error; error != nil {
		if error == gorm.ErrRecordNotFound {
			result = ErrorResponse{Error: "Appartement not found"}
			ginCxt.JSON(404, result)
			return
		}
		result = ErrorResponse{Error: "Error while fetching appartement"}
		ginCxt.JSON(500, result)
		return
	}

	if error := ginCxt.ShouldBindJSON(&appartement); error != nil {
		result = ErrorResponse{Error: "Invalid input"}
		ginCxt.JSON(400, result)
		return
	}

	// Mettre à jour l'appartement
	if error := configs.Database.Save(&appartement).Error; error != nil {
		result = ErrorResponse{Error: "Error while updating appartement"}
		ginCxt.JSON(500, result)
		return
	}

	result = AppartementResponse{Appartement: appartement}
	ginCxt.JSON(200, result)
}

// DeleteAppartement godoc
// @Summary Supprimer un appartement
// @Description Supprime un appartement en utilisant son ID
// @Tags appartements
// @Param id path string true "ID de l'appartement"
// @Produce json
// @Success 200 {object} SuccessMessage
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /appartements/{id} [delete]
func DeleteAppartement(ginCxt *gin.Context) {
	var (
		appartement models.Appartement
		result      interface{}
	)
	id := ginCxt.Param("id")

	// Récupérer l'appartement existant
	if error := configs.Database.Where("id = ?", id).First(&appartement).Error; error != nil {
		if error == gorm.ErrRecordNotFound {
			result = ErrorResponse{Error: "Appartement not found"}
			ginCxt.JSON(404, result)
			return
		}
		result = ErrorResponse{Error: "Error while fetching appartement"}
		ginCxt.JSON(500, result)
		return
	}

	// Supprimer l'appartement
	if error := configs.Database.Delete(&appartement).Error; error != nil {
		result = ErrorResponse{Error: "Error while deleting appartement"}
		ginCxt.JSON(500, result)
		return
	}

	result = SuccessMessage{Message: "Appartement deleted successfully"}
	ginCxt.JSON(200, result)
}

// GetAllAppartements godoc
// @Summary Récupérer tous les appartements
// @Description Récupère la liste de tous les appartements disponibles
// @Tags appartements
// @Produce json
// @Success 200 {object} AppartementsResponse
// @Failure 500 {object} ErrorResponse
// @Router /appartements [get]
func GetAllAppartements(ginCxt *gin.Context) {
	var (
		appartements []models.Appartement
		result       interface{}
	)

	// Récupérer tous les appartements, en préchargeant la localisation
	if error := configs.Database.Preload("Localisation").Find(&appartements).Error; error != nil {
		result = ErrorResponse{Error: "Error while fetching appartements"}
		ginCxt.JSON(500, result)
		return
	}

	result = AppartementsResponse{Appartements: appartements}
	ginCxt.JSON(200, result)
}
