package auth0

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwt"
	"go.opentelemetry.io/otel"
)

const (
	// CtxAuthenticated indicates if the requests is anonymous or not
	CtxAuthenticated = "jwt_authenticated"

	// CtxIsBCCMember indicates if the user has an active membership in BCC Norway
	// It is based on the "https://login.bcc.no/claims/hasMembership" claim
	// The reason for it's existence is that it can always exist and be checked more easily than
	// the presence of a PersonID for example
	CtxIsBCCMember = "jwt_is_bcc_member"

	// CtxPersonID is set if the user is not anonymous and has a BCC personId
	CtxPersonID = "jwt_person_id"

	// CtxJWTAudience is set if the person is not anonymous. It indicates what audience the
	// token was found valid for
	CtxJWTAudience = "jwt_audience"

	// CtxEmail is set to the email extracted from the token if it was present
	CtxEmail = "jwt_email"
)

// JWTConfig configures the JWT middleware
type JWTConfig struct {
	Domain    string
	Issuer    string
	Audiences []string
}

// JWT checks if there is a JWT in the Authorization header.
// If it is it will validate it, and set data in the context, or return a 403 forbidden
// If no JWT is found, TODO data in the context will be set to indicate an
// anonymous user
func JWT(ctx context.Context, config JWTConfig) gin.HandlerFunc {
	jwks := GetKeySet(ctx, config.Domain)

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := otel.Tracer("auth0").Start(ctx, "JWT Check")
		defer span.End()

		if c.Request.Header.Get("Authorization") == "" {
			c.Set(CtxAuthenticated, false)
			c.Set(CtxIsBCCMember, false)
			return
		}

		// middleware
		token, err := jwt.ParseRequest(
			c.Request,
			jwt.WithKeySet(jwks),
			jwt.WithHeaderKey("Authorization"),
		)

		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			log.L.Debug().Err(err).Msg("Error")
			return
		}

		valid := false
		errors := []error{}

		// Loop all allowed audiences, and collect errors
		// If none pass then print all errors for easier debugging,
		// else we can ignore the errors since we found a ok combo
		for _, audience := range config.Audiences {
			err := jwt.Validate(
				token,
				jwt.WithIssuer(config.Issuer),
				jwt.WithAudience(audience),
			)

			if err != nil {
				errors = append(errors, err)
			} else {
				valid = true
				c.Set(CtxJWTAudience, audience)
				break
			}
		}

		if !valid {
			log.L.Debug().
				Errs("Validation errors", errors).
				Msg("Validation")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// User is authenticated. Set the correct values and extract claims
		c.Set(CtxAuthenticated, true)

		// If you want to see all possible claims use the following line:
		// spew.Dump(token.PrivateClaims())

		// For now we manually add claims that are actually useful to avoid polluting the ctx
		// with data that we don't actually need and may be considered private under GDPR
		// If possible convert the claim to a string in order to make it easier to extract it later

		if privateClaims, ok := token.PrivateClaims()["https://members.bcc.no/app_metadata"]; ok {
			if val, ok := privateClaims.(map[string]interface{})["hasMembership"]; ok {
				c.Set(CtxIsBCCMember, val)
			}
		} else {
			// Make sure the key is set even if it does not exist in the token
			log.L.Warn().
				Str("claim", "https://members.bcc.no/app_metadata/").
				Msg("Unable to get claim")
			c.Set(CtxIsBCCMember, false)
		}

		if val, ok := token.Get("https://login.bcc.no/claims/personId"); ok {
			c.Set(CtxPersonID, fmt.Sprintf("%.0f", val))
		}

		if val, ok := token.Get("email"); ok {
			c.Set(CtxEmail, val)
		}
	}
}
