package main

import (
	"errors"
	"fmt"
	"os"
)

func MustEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		panic(fmt.Sprintf("required env %q is missing — fix your config", key))
	}
	return v
}

func Env(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return "", errors.New("missing env: " + key)
	}
	return v, nil
}

func main() {
	// `Env` returns an error and is composable in retry/fallback logic:
	if v, err := Env("PORT"); err != nil {
		fmt.Println("PORT error:", err) // expected, gracefully handled
	} else {
		fmt.Println("PORT =", v)
	}

	// `MustEnv` is for STARTUP: misuse is a config bug. We catch it here
	// only to keep the demo running.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("MustEnv panicked: %v\n", r)
		}
	}()
	_ = MustEnv("DEFINITELY_NOT_SET")

	// Library-author rule of thumb:
	//   - Provide the error-returning version as the default.
	//   - Add a Must* wrapper for cases where the input is a constant
	//     known at compile time (regex, template, sql.Open driver name)
	//     and where panicking at startup is OK.
}
