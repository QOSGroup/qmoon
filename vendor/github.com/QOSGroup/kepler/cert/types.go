package cert

import (
	"github.com/tendermint/tendermint/crypto"
	"time"
)

const (
	CommonSubjAminoRoute = "certificate/CommonSubject"
	QSCSubjAminoRoute    = "certificate/QSCSubject"
	QCPSubjAminoRoute    = "certificate/QCPSubject"
)

type Subject interface{}

type CommonSubject struct {
	CN string `json:"cn"`
}

type QSCSubject struct {
	ChainId string        `json:"chain_id"`
	Name    string        `json:"name"`
	Banker  crypto.PubKey `json:"banker"`
}

type QCPSubject struct {
	ChainId  string `json:"chain_id"`
	QCPChain string `json:"qcp_chain"`
}

type CertificateSigningRequest struct {
	Subj      Subject       `json:"subj"`
	IsCa      bool          `json:"is_ca"`
	NotBefore time.Time     `json:"not_before"`
	NotAfter  time.Time     `json:"not_after"`
	PublicKey crypto.PubKey `json:"public_key"`
}

type Issuer struct {
	Subj      Subject       `json:"subj"`
	PublicKey crypto.PubKey `json:"public_key"`
}

type Certificate struct {
	CSR       CertificateSigningRequest `json:"csr"`
	CA        Issuer                    `json:"ca"`
	Signature []byte                    `json:"signature"`
}

func (crt Certificate) PublicKey() crypto.PubKey {
	return crt.CSR.PublicKey
}

type TrustCrts struct {
	PublicKeys []crypto.PubKey `json:"public_keys"`
}

func VerityCrt(caPublicKeys []crypto.PubKey, crt Certificate) bool {
	ok := false

	// Check issuer
	for _, value := range caPublicKeys {
		if value.Equals(crt.CA.PublicKey) {
			ok = crt.CA.PublicKey.VerifyBytes(MustMarshalJson(crt.CSR), crt.Signature)
			break
		}
	}

	// Check timestamp
	now := time.Now().Unix()
	if now <= crt.CSR.NotBefore.Unix() || now >= crt.CSR.NotAfter.Unix() {
		ok = false
	}

	return ok
}
