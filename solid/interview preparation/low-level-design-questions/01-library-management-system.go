//1. 📚 Design a Library Management System
//Goal: Build a system that manages book borrowing, returns, inventory tracking, and user management for a library.
//
//Entities: Book, User, Library, Transaction
//Features:
//Search books by title/author
//Borrow/return book
//Maintain user borrow history

package main

import (
	"fmt"
	"time"
)

type Book struct {
	ID       string
	Title    string
	Author   string
	IsIssued bool
	IssuedBy User
}

type User struct {
	ID            string
	Name          string
	BorrowedBooks []Book
	BorrowHistory []Transaction
}

type Transaction struct {
	BookID     string
	UserID     string
	BorrowDate time.Time
	ReturnDate *time.Time
	IsReturned bool
}

type Library struct {
	Name         string
	Books        map[string]Book
	Users        map[string]User
	Transactions []Transaction
}

type SearchBook interface {
	ISearch(input string) []Book
}

type SearchBookByTitle struct {
	books map[string]Book
}

type SearchBookByAuthor struct {
	books map[string]Book
}

func (s SearchBookByTitle) ISearch(title string) []Book {
	var result []Book
	for _, book := range s.books {
		if book.Title == title {
			result = append(result, book)
		}
	}
	return result
}

func (s SearchBookByAuthor) ISearch(author string) []Book {
	var result []Book
	for _, book := range s.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

func (l *Library) IssueBook(bookID string, userID string) error {
	book, bookExists := l.Books[bookID]
	user, userExists := l.Users[userID]

	if !bookExists || !userExists {
		return fmt.Errorf("Book or user not found")
	}
	if book.IsIssued {
		return fmt.Errorf("Book is already issued")
	}

	book.IsIssued = true
	book.IssuedBy = user
	l.Books[bookID] = book

	user.BorrowedBooks = append(user.BorrowedBooks, book)

	transaction := Transaction{
		BookID:     bookID,
		UserID:     userID,
		BorrowDate: time.Now(),
		IsReturned: false,
	}

	user.BorrowHistory = append(user.BorrowHistory, transaction)
	l.Users[userID] = user
	l.Transactions = append(l.Transactions, transaction)

	fmt.Printf("Book '%s' issued to user '%s'\n", book.Title, user.Name)
	return nil
}

func (l *Library) ReturnBook(bookID string, userID string) error {
	book, bookExists := l.Books[bookID]
	user, userExists := l.Users[userID]

	if !bookExists || !userExists {
		return fmt.Errorf("Book or user not found")
	}
	if !book.IsIssued || book.IssuedBy.ID != userID {
		return fmt.Errorf("Book not issued to this user")
	}

	book.IsIssued = false
	book.IssuedBy = User{}
	l.Books[bookID] = book

	newBorrowed := []Book{}
	for _, book := range user.BorrowedBooks {
		if book.ID != bookID {
			newBorrowed = append(newBorrowed, book)
		}
	}
	user.BorrowedBooks = newBorrowed

	now := time.Now()
	for i := range user.BorrowHistory {
		t := &user.BorrowHistory[i]
		if t.BookID == bookID && !t.IsReturned {
			t.IsReturned = true
			t.ReturnDate = &now
			break
		}
	}

	l.Users[userID] = user

	fmt.Printf("Book '%s' returned by user '%s'\n", book.Title, user.Name)
	return nil
}

func (l *Library) GetUserBorrowHistory(userID string) []Transaction {
	user, exists := l.Users[userID]
	if !exists {
		return []Transaction{}
	}
	return user.BorrowHistory
}

func main() {
	lib := Library{
		Name:  "City Library",
		Books: make(map[string]Book),
		Users: make(map[string]User),
	}

	lib.Books["B1"] = Book{ID: "B1", Title: "The Go Programming Language", Author: "Alan Donovan"}
	lib.Books["B2"] = Book{ID: "B2", Title: "Clean Code", Author: "Robert C. Martin"}

	lib.Users["U1"] = User{ID: "U1", Name: "Alice"}

	searchByTitle := SearchBookByTitle{books: lib.Books}
	fmt.Println("Search Result:", searchByTitle.ISearch("Clean Code"))

	err := lib.IssueBook("B2", "U1")
	if err != nil {
		fmt.Println("Issue Error:", err)
	}
	time.Sleep(1 * time.Second)
	err = lib.ReturnBook("B2", "U1")
	if err != nil {
		fmt.Println("Return Error:", err)
	}

	history := lib.GetUserBorrowHistory("U1")
	fmt.Println("Borrow History:", history)
}
