// Problem 5: Named Types & Underlying Types
//
// Every type in Go has an UNDERLYING TYPE. Conversion rules between named
// types boil down to: "if the underlying types are identical, you may
// convert explicitly."
//
//   type UserID int64        // underlying = int64
//   type AccountID int64     // underlying = int64
//   type SecondsSinceEpoch UserID  // underlying = int64
//
// Tasks:
//   1. Define UserID, AccountID (both with underlying int64).
//   2. Show that direct assignment between them is rejected (uncomment
//      a line and observe the compile error).
//   3. Show that explicit conversion `AccountID(uid)` is accepted.
//   4. Show conversion to/from int64 is accepted.
//   5. Define a function `func ImpersonationKey(u UserID) string` and
//      explain (in a comment) why this prevents accidentally passing an
//      AccountID where a UserID is expected — type safety win.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: define UserID and AccountID

type UserID int64
type AccountID int64

func ImpersonationKey(u UserID) string {
	return fmt.Sprintf("user:%d", u)
}

func main() {
	var u UserID
	var a AccountID = 434
	u = UserID(a) // explicit type conversion is valid since underlying dtype is int64
	// ImpersonationKey(a)  This doesn't compile because even though the underlying dtype of both UserID and
	// AccountID is int64, but ImpersonationKey only accepts arguments of type UserID.
	fmt.Println(ImpersonationKey(u))
}
