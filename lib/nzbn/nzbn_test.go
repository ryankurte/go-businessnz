package nzbn

import (
	"os"
	"testing"

	"github.com/ryankurte/go-businessnz/lib/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNzbnApi(t *testing.T) {

	// Fetch API key from environment
	apiKey := os.Getenv("BUSINESSNZ_API_KEY")
	assert.NotNil(t, apiKey)
	apiSecret := os.Getenv("BUSINESSNZ_API_SECRET")
	assert.NotNil(t, apiSecret)

	// Setup base and API
	// TODO: it'd be nice to use the sandbox here but, the test items do not seem to exist
	base := base.NewBase(apiKey, apiSecret, false, true)
	nzbn := NzbnApi{Base: &base}

	t.Run("Lookup by NZBN", func(t *testing.T) {

		entity, err := nzbn.Lookup("9429045862298")
		require.Nil(t, err)

		assert.Equal(t, "9429045862298", entity.Nzbn)
	})

	t.Run("Lookup by name", func(t *testing.T) {

		entities, err := nzbn.Search(SearchQuery{SearchTerm: "ElectronPowered"})
		require.Nil(t, err)
		require.NotNil(t, entities)

		assert.Equal(t, "9429045862298", entities.Items[0].Nzbn)
	})
}
