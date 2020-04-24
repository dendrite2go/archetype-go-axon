package authentication

import (
    bytes "bytes"
    log "log"
    strings "strings"

    base64 "encoding/base64"
    rand "crypto/rand"
    sha256 "crypto/sha256"

    jwt "github.com/pascaldekloe/jwt"
    grpc "google.golang.org/grpc"
    ssh "golang.org/x/crypto/ssh"

    axon_server "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/axon_server"
    axon_utils "github.com/jeroenvm/archetype-go-axon/src/pkg/axon_utils"
    grpc_example "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/example"
    trusted "github.com/jeroenvm/archetype-go-axon/src/pkg/trusted"
)

var acceptedCredentials map[string]string

func Init() {
    acceptedCredentials = make(map[string]string)
}

func UnsafeSetCredentials(credentials *grpc_example.Credentials) {
    acceptedCredentials[credentials.Identifier] = credentials.Secret
}

func SetCredentials(credentials *grpc_example.Credentials, conn *grpc.ClientConn, clientInfo *axon_server.ClientIdentification) error {
    payload := credentials.Identifier + "=" + credentials.Secret
    signatureKey, e := trusted.GetKeyManagerKey(credentials.Signature.SignatureName)
    if e != nil {
        return e
    }
    signature := ssh.Signature{
        Format: credentials.Signature.Format,
        Blob: credentials.Signature.Blob,
        Rest: credentials.Signature.Rest,
    }
    signatureKey.Verify([]byte(payload), &signature)
    log.Printf("Set credentials: %v: %v", credentials.Identifier, credentials.Secret)

    var encryptedSecret string
    if credentials.Secret == "" {
        encryptedSecret = ""
    } else {
        encryptedSecret, e = trusted.EncryptString(credentials.Secret)
        if e != nil {
            return e
        }
    }
    commandCredentials := grpc_example.Credentials{
        Identifier: credentials.Identifier,
        Secret: encryptedSecret,
    }
    command := grpc_example.RegisterCredentialsCommand{
        Credentials: &commandCredentials,
    }
    e = axon_utils.SendCommand("RegisterCredentialsCommand", &command, conn, clientInfo)
    return e
}

func Authenticate(username string, password string) bool {
    encryptedHashedPassword := acceptedCredentials[username]
    log.Printf("Authenticate: %v: encryptedHashedPassword: '%v'", username, encryptedHashedPassword)
    hashedPassword, e := trusted.DecryptString(encryptedHashedPassword)
    if e != nil {
        log.Printf("Authenticate: %v: could not decrypt hashed password: %v", username, e)
        return false
    }

    log.Printf("Authenticate: %v: hashedPassword: '%v'", username, hashedPassword)
    parts := strings.Split(hashedPassword, ":")
    if len(parts) != 3 {
        log.Printf("Authenticate: %v: number of parts: %v", username, len(parts))
        return false
    }

    if parts[0] != "sha256" {
        log.Printf("Authenticate: %v: hash type: %v", username, parts[0])
        return false
    }

    salt, e := base64.RawStdEncoding.DecodeString(parts[1])
    if e != nil {
        log.Printf("Authenticate: %v: could not base64 decode salt: %v", username, parts[1])
        return false
    }

    storedHash, e := base64.RawStdEncoding.DecodeString(parts[2])
    if e != nil {
        log.Printf("Authenticate: %v: could not base64 decode store hash: %v", username, parts[2])
        return false
    }

    blob := append(salt, ([]byte(password))...)
    givenHash := sha256.Sum256(blob)
    log.Printf("Authenticate: %v: given hash: %v", username, givenHash)

    return bytes.Compare(givenHash[:], storedHash) == 0
}

func CheckKnown(credentials *grpc_example.Credentials, projection *Projection) bool {
    newHashedPassword, e := trusted.DecryptString(credentials.Secret)
    if e != nil {
        log.Printf("Authenticate: %v: cannot decrypt new hash: assuming unknown", credentials.Identifier)
        return false
    }
    oldHashedPassword, e := trusted.DecryptString((*projection).Credentials[credentials.Identifier])
    if e != nil {
        log.Printf("Authenticate: %v: cannot decrypt old hash: assuming unknown", credentials.Identifier)
        return false
    }
    return newHashedPassword == oldHashedPassword
}

func Encode(password string) string {
    salt := make([]byte, 32)
    rand.Reader.Read(salt)
    passwordBytes := []byte(password)
    blob := append(salt, passwordBytes...)
    hash := sha256.Sum256(blob)
    var builder strings.Builder
    builder.WriteString("sha256:")
    builder.WriteString(base64.RawStdEncoding.EncodeToString(salt))
    builder.WriteString(":")
    builder.WriteString(base64.RawStdEncoding.EncodeToString(hash[:]))
    return builder.String()
}

func Verify(accessToken *grpc_example.AccessToken) bool {
    publicKey, e := trusted.GetRsaPublicKey()
    if e != nil {
        return false
    }
    _, e = jwt.RSACheck([]byte(accessToken.Jwt), publicKey)
    return e == nil
}
