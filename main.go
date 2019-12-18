package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "someusers"
)
const (
	CommandSave = iota + 1
	CommandListAll
	CommandGetByID
	CommandGetByPhone
	CommandGetByEmail
	CommandSearchByName
	CommandDelete
)

type ContactsRepositoryPostgreSQL struct {
	db *sql.DB
}

func NewContactsRepositoryPostgreSQL(db *sql.DB) *ContactsRepositoryPostgreSQL {
	return &ContactsRepositoryPostgreSQL{
		db: db,
	}
}

type ContactsRepository interface {
	Save(Contact) (Contact, error)
	ListAll() ([]Contact, error)
	GetByID(uint) (Contact, error)
	GetByPhone(string) (Contact, error)
	GetByEmail(string) (Contact, error)
	SearchByName(string) ([]Contact, error)
	Delete(uint) error
}

var db *sql.DB

type Contact struct {
	ID uint `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (c Contact) String() string {
	return fmt.Sprintf("Contact:\n\tID - %d\n\tFirst Name - %q\n\tLast Name - %q\n\tPhone - %q\n\tEmail - %q\n", c.ID, c.FirstName, c.LastName, c.Phone, c.Email)
}
func main() {
	fmt.Println("Start!")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	repository := NewContactsRepositoryPostgreSQL(db)

	for {
		fmt.Print(menu)
		var command int
		if _, err := fmt.Scanf("%d", &command); err != nil {
			log.Println(err)
		}

		switch command {
		case CommandSave:
			if err := Save(repository); err != nil {
				log.Println(err)
			}
		case CommandListAll:
			if err := ListAll(repository); err != nil {
				log.Println(err)
			}
				case CommandGetByID:
					if err := GetByID(repository); err != nil {
						log.Println(err)
					}
				case CommandGetByPhone:
					if err := GetByPhone(repository); err != nil {
						log.Println(err)
					}
				case CommandGetByEmail:
					if err := GetByEmail(repository); err != nil {
						log.Println(err)
					}
				case CommandSearchByName:
					if err := SearchByName(repository); err != nil {
						log.Println(err)
					}
				case CommandDelete:
					if err := Delete(repository); err != nil {
						log.Println(err)
					}
		default:
			log.Printf("command not foumd for value %d\n", command)
		}

		printSeparator()
	}
}

func ListAll(rep ContactsRepository) error {
	records, err := rep.ListAll()
	if err != nil {
		return fmt.Errorf("error in ListAll: %q", err.Error())
	}

	fmt.Println("ListAll:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

func GetByID(rep ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	record, err := rep.GetByID(id)
	if err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Println("GetByID")
	fmt.Println(record)

	return nil
}

func GetByPhone(rep ContactsRepository) error {
	phone := readString("Please enter an 'Phone' field and press Enter")

	record, err := rep.GetByPhone(phone)
	if err != nil {
		return fmt.Errorf("error in GetByPhone: %q", err.Error())
	}

	fmt.Println("GetByPhone:")
	fmt.Println(record)

	return nil
}

func GetByEmail(rep ContactsRepository) error {
	email := readString("Please enter an 'Email' field and press Enter")

	record, err := rep.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("error in GetByEmail: %q", err.Error())
	}

	fmt.Println("GetByEmail:")
	fmt.Println(record)

	return nil
}

func SearchByName(rep ContactsRepository) error {
	email := readString("Please enter prefix for 'Name' field and press Enter")

	records, err := rep.SearchByName(email)
	if err != nil {
		return fmt.Errorf("error in SearchByName: %q", err.Error())
	}

	fmt.Println("SearchByName:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

func Delete(rep ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	if err := rep.Delete(id); err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Printf("Delete:\nRecord with ID %d successfylly deleted\n", id)
	return nil
}

func Save(rep ContactsRepository) error {
	contact := Contact{
		FirstName: readString("Please enter an 'FirstName' field and press Enter"),
		LastName:  readString("Please enter an 'LastName' field and press Enter"),

		Phone: readString("Please enter an 'Phone' field and press Enter"),
		Email: readString("Please enter an 'Email' field and press Enter"),
	}

	result, err := rep.Save(contact)
	if err != nil {
		return err
	}

	fmt.Println("Save Contact:")
	fmt.Println(result)

	return nil
}

const menu = `
Please enter operation number:
  * 1 - Save
  * 2 - ListAll
  * 3 - GetByID
  * 4 - GetByPhone
  * 5 - GetByEmail
  * 6 - SearchByName
  * 7 - Delete 
  * Control + C - to exit 
`

func readString(message string) string {
	var r string

	for r == "" {
		fmt.Println(message)
		if _, err := fmt.Scanf("%s", &r); err != nil {
			fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
			printSeparator()
		}

	}

	return r
}

func readUint(message string) uint {
	var r uint

	for {
		fmt.Println(message)
		_, err := fmt.Scanf("%d", &r)
		if err == nil {
			break
		}

		fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
		printSeparator()
	}

	return r
}

func printSeparator() {
	for i := 0; i < 50; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}
func (r *ContactsRepositoryPostgreSQL) Save(contact Contact) (Contact, error) {
	query := "INSERT INTO users(first_name,last_name,phone,email) VALUES($1,$2,$3,$4) returning id;"
	err := r.db.QueryRow(query, contact.FirstName, contact.LastName, contact.Phone, contact.Email).Scan(&contact.ID)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"contacts_phone_idx\"") {
			return Contact{}, fmt.Errorf("contact with phone %q already exists", contact.Phone)
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"contacts_email_idx\"") {
			return Contact{}, fmt.Errorf("contact with email %q already exists", contact.Email)
		}

		return Contact{}, err
	}

	return contact, nil
}

func (r *ContactsRepositoryPostgreSQL) ListAll() (users []Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM users;"
	rows, err := r.db.Query(query)
	if err != nil {
		return users, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, contact)
	}

	return users, nil
}

func (r *ContactsRepositoryPostgreSQL) GetByID(id uint) (contact Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM users WHERE id = $1;"
	row := r.db.QueryRow(query, id)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return Contact{}, err
		}

		return Contact{}, fmt.Errorf("record not found")
	}

	return contact, nil
}

func (r *ContactsRepositoryPostgreSQL) GetByPhone(phone string) (contact Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM users WHERE phone = $1;"
	row := r.db.QueryRow(query, phone)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return Contact{}, err
		}

		return Contact{}, fmt.Errorf("record not found")
	}

	return contact, nil
}

func (r *ContactsRepositoryPostgreSQL) GetByEmail(email string) (contact Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM users WHERE email = $1;"
	row := r.db.QueryRow(query, email)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err == sql.ErrNoRows {
		return Contact{}, fmt.Errorf("record not found")
	} else if err != nil {
		return Contact{}, err
	}
	return contact, nil
}

func (r *ContactsRepositoryPostgreSQL) SearchByName(n string) (contacts []Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM users WHERE first_name = $1;"
	rows, err := r.db.Query(query, n)
	if err != nil {
		return contacts, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactsRepositoryPostgreSQL) Delete(id uint) error {
	query := "DELETE FROM users WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
