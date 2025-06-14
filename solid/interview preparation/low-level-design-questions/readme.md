🧠 Low-Level Design (LLD) Questions – Beginner Friendly
This section includes a curated set of low-level design problems. These are ideal for interviews or practicing design fundamentals like OOP, design patterns, class modeling, and modularization.


1. 📚 Design a Library Management System
   Goal: Build a system that manages book borrowing, returns, inventory tracking, and user management for a library.

Entities: Book, User, Librarian, Library, Transaction

Features:

Search books by title/author

Borrow/return book

Maintain user borrow history

Penalty for late returns


2. 🚗 Design a Parking Lot System
   Goal: Simulate a multi-level parking system that can handle different types of vehicles.

Entities: Vehicle, Slot, Floor, ParkingTicket, Gate

Features:

Vehicle entry and exit

Parking charges calculation

Slot assignment based on vehicle type

Real-time availability tracking

3. 🛗 Design an Elevator (Lift) System
   Goal: Design a system to manage elevator movement in a multi-floor building.

Entities: Elevator, Button, Floor, Request

Features:

Handle multiple requests

Optimize direction of movement

Door open/close logic

Handle idle state

4. 🎬 Design a Movie Ticket Booking System
   Goal: Book movie tickets for various shows in different cinemas.

Entities: Movie, Theatre, Show, Seat, Booking

Features:

Show search and seat availability

Seat locking mechanism

Payment interface

Booking confirmation and cancellation

5. ☕ Design a Coffee Vending Machine
   Goal: Simulate a vending machine that serves beverages and manages ingredients.

Entities: Drink, Ingredient, Inventory, Machine

Features:

Select drink type

Refill ingredients

Handle low stock warnings

State transitions (Idle → Brewing → Ready)

6. 💸 Design a Splitwise App
   Goal: Create an expense-sharing system among a group of friends.

Entities: User, Expense, Group, BalanceSheet

Features:

Equal and custom splits

Settle up balances

Group-wise expenses

Transaction history

7. 🎮 Design a Tic-Tac-Toe Game
   Goal: Build a two-player game with a square grid and winning logic.

Entities: Game, Player, Board, Cell

Features:

Move validation

Win/draw detection

Support different grid sizes

8. 🐍 Design a Snake and Ladder Game
   Goal: Simulate the popular board game between multiple players.

Entities: Player, Board, Snake, Ladder, Dice

Features:

Position tracking

Dice roll simulation

Snake/ladder movement

Win detection

9. 🏨 Design a Hotel Room Booking System
   Goal: Book hotel rooms based on availability and preferences.

Entities: Room, Customer, Booking, Payment, Invoice

Features:

Search and filter rooms

Availability calendar

Reservation management

Cancellation and refund

10. 🛒 Design an Online Shopping Cart System
    Goal: Simulate the core components of an e-commerce site’s cart and order system.

Entities: Product, Cart, User, Order, Payment

Features:

Add/remove items from cart

Apply coupons

Checkout flow

Order confirmation and tracking

11. 🚕 Design a Ride Sharing (Uber) System
    Goal: Match drivers and riders in real-time for a ride.

Entities: Driver, Rider, Ride, Location, Vehicle

Features:

Real-time location matching

Ride request and allocation

Fare calculation

Trip summary and ratings

12. 📲 Design a Messaging App (like WhatsApp)
    Goal: Build the backend logic for sending and receiving messages.

Entities: User, Message, Chat, Group

Features:

Send/receive 1-on-1 and group messages

Message history

Group creation

Read receipts

13. 🗳️ Design a Voting (Polling) System
    Goal: Create a system to conduct polls and store user votes.

Entities: User, Poll, Option, Vote

Features:

Create new poll

Prevent duplicate voting

Show real-time results

Close poll automatically

14. 🔄 Implement Undo/Redo Functionality (like in a Text Editor)
    Goal: Support basic undo/redo features for a text editor.

Entities: Action, Editor, Stack

Features:

Command history tracking

Undo last operation

Redo previously undone operation

15. 🧠 Implement a Pub-Sub System
    Goal: Design a publisher-subscriber pattern system for sending events/messages.

Entities: Topic, Message, Publisher, Subscriber

Features:

Subscribe/unsubscribe to topics

Publish messages to all subscribers

Message queues

Thread-safe operation

16. 🧾 Implement a Logger using Singleton Pattern
    Goal: Create a logging system with exactly one instance across the app.

Concepts: Singleton pattern, thread-safe initialization

Features:

Global logging instance

Append log messages

(Optional) Write to file

17. 🛠️ Implement a Notification System
    Goal: Send different types of notifications (Email, SMS, Push) based on user preference.

Patterns: Strategy / Factory

Features:

Notification channels

Plug-in based extensibility

Retry/failure handling

18. 🧠 Design a Cache System (LRU Cache)
    Goal: Implement a cache that discards the least recently used items first.

Concepts: HashMap + Doubly Linked List

Features:

Get/set in O(1)

Fixed size

Eviction policy

19. 📁 Design a File System
    Goal: Simulate a basic file system with folders and files.

Patterns: Composite Pattern

Features:

Create/delete files/folders

List contents

Recursive structure traversal

Path navigation

20. 💳 Design a Payment Gateway System
    Goal: Simulate a system that supports multiple payment methods (credit card, UPI, wallet).

Patterns: Strategy / Factory

Features:

Multiple payment types

Payment failure handling

Refund support

Transaction tracking

