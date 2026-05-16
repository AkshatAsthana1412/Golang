package main

import "fmt"

type UserID int64
type AccountID int64

func ImpersonationKey(u UserID) string {
	return fmt.Sprintf("user:%d", u)
}

func main() {
	uid := UserID(42)
	aid := AccountID(99)

	// aid = uid              // ❌ cannot use UserID as AccountID
	aid = AccountID(uid) //    ✅ explicit conversion: same underlying type
	_ = aid

	// To/from underlying int64 needs conversion too:
	var raw int64 = int64(uid)
	uid = UserID(raw)

	fmt.Println(ImpersonationKey(uid))

	// If ImpersonationKey took int64, ANY int64-shaped value (including
	// AccountID after conversion, or even mistyped variables) would be
	// accepted. Using a distinct UserID name makes the contract explicit
	// and lets the compiler catch the wrong-id-type bug for free.
}
