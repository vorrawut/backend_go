package favorite_token

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGeTFavoriteToken(t *testing.T) {
	t.Run("Sucessful Request", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		getFavoriteToken(c)
		// if err != nil {
		// 	t.Fatal("Everything should be ok")
		// }
	})
}

func TestAddFavoriteToken(t *testing.T) {
	t.Run("Sucessful Request", func(t *testing.T) {
		// _, err := addFavoriteToken(events.APIGatewayProxyRequest{})
		// if err != nil {
		// 	t.Fatal("Everything should be ok")
		// }
	})
}

func TestRemoveFavoriteToken(t *testing.T) {
	t.Run("Sucessful Request", func(t *testing.T) {
		// _, err := removeFavoriteToken(events.APIGatewayProxyRequest{})
		// if err != nil {
		// 	t.Fatal("Everything should be ok")
		// }
	})
}

func TestLambdaHandler(t *testing.T) {
	t.Run("Sucessful Request", func(t *testing.T) {
		// _, err := Handler(events.APIGatewayProxyRequest{})
		// if err != nil {
		// 	t.Fatal("Everything should be ok")
		// }
	})
}
