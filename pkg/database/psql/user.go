package psql

import (
	"log"

	"github.com/dsafanyuk/fetchr-go/pkg/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type userRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) Create(user *user.User) (*user.User, error) {
	rows, err := r.db.NamedQuery(`
	INSERT INTO users (
		email_address, password, room_num, first_name, last_name, phone_number
	) VALUES (
		:email_address, :password, :room_num, :first_name, :last_name, :phone_number
	) RETURNING *`,
		user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return user, nil
}

func (r *userRepository) FindByID(id string) (*user.User, error) {
	user := user.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE user_id = $1", id)
	if err != nil {
		panic(err)
	}
	return &user, nil
}

func (r *userRepository) FindAll() (users []*user.User, err error) {
	// todo refactor to be cleaner
	rows, err := r.db.Queryx("SELECT * FROM users")
	defer rows.Close()

	for rows.Next() {
		user := new(user.User)
		if err = rows.StructScan(&user); err != nil {
			log.Print(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Update(user *user.User) (*user.User, error) {

	rows, err := r.db.NamedQuery(`
	UPDATE users SET
		email_address = :email_address,
		password = :password,
		wallet = :wallet,
		is_active = :is_active,
		is_admin = :is_admin,
		room_num = :room_num,
		first_name = :first_name,
		last_name = :last_name,
		phone_number = :phone_number
	WHERE user_id = :user_id
		RETURNING *`, user)
	if err != nil {
		return nil, err
	}
	// Retreive newly created user_id
	if rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return user, nil

}

func (r *userRepository) Delete(id string) (*user.User, error) {
	user := user.User{}
	err := r.db.Get(&user, `
		UPDATE users
			SET is_active = false
		WHERE user_id = $1
			RETURNING *`,
		id)
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	return &user, nil
}
