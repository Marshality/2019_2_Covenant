package album

import "2019_2_Covenant/internal/models"

type Repository interface {
	FindLike(name string, count uint64) ([]*models.Album, error)
}
