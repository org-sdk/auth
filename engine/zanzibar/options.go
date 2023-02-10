package zanzibar

import (
    "github.com/org-sdk/auth/engine/zanzibar/keto"
    "github.com/org-sdk/auth/engine/zanzibar/openfga"
)

type OptFunc func(*State)

func WithKeto(readUrl, writeUrl string, useGRPC bool) OptFunc {
    return func(s *State) {
        s.ketoClient = keto.NewClient(readUrl, writeUrl, useGRPC)
    }
}

func WithOpenFga(scheme, host, storeId, token string) OptFunc {
    return func(s *State) {
        s.openfgaClient = openfga.NewClient(scheme, host, storeId, token)
    }
}
