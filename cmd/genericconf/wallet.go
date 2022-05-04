// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package genericconf

import (
	"github.com/spf13/pflag"
	"path"
	"path/filepath"
)

const PASSWORD_NOT_SET = "PASSWORD_NOT_SET"

type WalletConfig struct {
	Pathname     string `koanf:"pathname"`
	PasswordImpl string `koanf:"password"`
	PrivateKey   string `koanf:"private-key"`
	Account      string `koanf:"account"`
}

func (w *WalletConfig) Password() *string {
	if w.PasswordImpl == PASSWORD_NOT_SET {
		return nil
	}
	return &w.PasswordImpl
}

var WalletConfigDefault = WalletConfig{
	Pathname:     "",
	PasswordImpl: "",
	PrivateKey:   "",
	Account:      "",
}

func WalletConfigAddOptions(prefix string, f *pflag.FlagSet, defaultPathname string) {
	f.String(prefix+".pathname", defaultPathname, "pathname for wallet")
	f.String(prefix+".password", WalletConfigDefault.PasswordImpl, "wallet passphrase")
	f.String(prefix+".private-key", WalletConfigDefault.PasswordImpl, "private key for wallet")
	f.String(prefix+".account", WalletConfigDefault.Account, "account to use (default is first account in keystore)")
}

func (w *WalletConfig) ResolveDirectoryNames(chain string) {
	// Make wallet directories relative to chain directory if specified and not already absolute
	if len(w.Pathname) != 0 && !filepath.IsAbs(w.Pathname) {
		w.Pathname = path.Join(chain, w.Pathname)
	}
}
