package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang-database-user/model"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repo *userRepositoryImpl) EmailExists(ctx context.Context, email string) (bool, error) {
	query := "SELECT COUNT(1) FROM mst_user WHERE email = $1"

	var count int
	err := repo.DB.QueryRowContext(ctx, query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertUser : fungsi untuk melakukan query ke dalam database. ( contoh di bawah ini adalah fungsi untuk membuat data user )
func (repo *userRepositoryImpl) InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error) {

	query := "INSERT INTO mst_user(id_user, name, email, password, phone_number, role_id) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := repo.DB.ExecContext(ctx, query, user.IdUser, user.Name, user.Email, user.Password, user.PhoneNumber, user.Role.IdRole)
	if err != nil {
		return model.MstUser{}, err
	}

	return user, nil
}

func (repo *userRepositoryImpl) UpdateUser(ctx context.Context, user model.MstUser, userId string) (model.MstUser, error) {
	if userId == "" {
		return model.MstUser{}, fmt.Errorf("invalid user ID")
	}
	query := "UPDATE mst_user SET name = $1, email = $2, password = $3, phone_number = $4 WHERE id_user = $5"

	_, err := repo.DB.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.PhoneNumber, userId)
	if err != nil {
		return model.MstUser{}, err
	}
	return user, nil
}

func (repo *userRepositoryImpl) ReadUser(ctx context.Context) ([]model.MstUser, error) {
	query := "SELECT id_user, name, email, phone_number FROM mst_user"

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err) // Tambahkan error handling yang jelas
	}
	defer rows.Close()

	var users []model.MstUser
	for rows.Next() {
		var user model.MstUser

		err := rows.Scan(&user.IdUser, &user.Name, &user.Email, &user.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err) // Error handling saat scan
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("no user found") // Mengembalikan error jika tidak ada user
	}

	return users, nil
}


func (repo *userRepositoryImpl) DeleteUser(ctx context.Context, userId string) error {
	if userId == "" {
        return fmt.Errorf("ID user tidak boleh kosong")
    }
    
	query := "DELETE FROM mst_user WHERE id_user = $1"
	_, err := repo.DB.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}
	return nil
}

