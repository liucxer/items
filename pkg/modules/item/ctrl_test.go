package item_test

import (
	"encoding/json"
	"testing"

	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
)

func TestBody(t *testing.T) {
	v, _ := json.MarshalIndent(models.Item{
		PrimaryID: models.PrimaryID{ID: 1},
		ItemRef:   models.ItemRef{ItemID: depends.GenUUID()},
		ItemBase: models.ItemBase{
			Code:          "CODE_TEST",
			ParentCode:    "PARENT_CODE_TEST",
			Name:          "NAME_TEST",
			AlphabetZH:    "ZH",
			AlphabetEN:    "EN",
			ImageURL:      "TEST_URL",
			RichText:      "REICH_TEST",
			Link:          "LINK",
			AttachmentURL: "ATTACHMENT",
			HasSub:        false,
		},
	}, "", "  ")
	t.Log("\n", string(v))

	v, _ = json.Marshal(models.UserBase{
		Username: "test",
		Password: "test",
	})
	t.Log("\n", string(v))
}
