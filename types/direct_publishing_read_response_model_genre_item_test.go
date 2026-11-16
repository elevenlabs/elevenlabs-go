package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDirectPublishingReadResponseModelGenreItemVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant DirectPublishingReadResponseModelGenreItem
		wire    string
	}
	cases := []enumCase{
		{variant: DirectPublishingReadResponseModelGenreItemFantasy, wire: "\"Fantasy\""},
		{variant: DirectPublishingReadResponseModelGenreItemRomance, wire: "\"Romance\""},
		{variant: DirectPublishingReadResponseModelGenreItemScienceFiction, wire: "\"Science Fiction\""},
		{variant: DirectPublishingReadResponseModelGenreItemMysteryAndThriller, wire: "\"Mystery and Thriller\""},
		{variant: DirectPublishingReadResponseModelGenreItemActionAndAdventure, wire: "\"Action and Adventure\""},
		{variant: DirectPublishingReadResponseModelGenreItemDystopia, wire: "\"Dystopia\""},
		{variant: DirectPublishingReadResponseModelGenreItemBusinessAndEconomics, wire: "\"Business and Economics\""},
		{variant: DirectPublishingReadResponseModelGenreItemTechnology, wire: "\"Technology\""},
		{variant: DirectPublishingReadResponseModelGenreItemChristianAndInspirational, wire: "\"Christian & Inspirational\""},
		{variant: DirectPublishingReadResponseModelGenreItemHorror, wire: "\"Horror\""},
		{variant: DirectPublishingReadResponseModelGenreItemBiographyAndMemoir, wire: "\"Biography and Memoir\""},
		{variant: DirectPublishingReadResponseModelGenreItemEducationAndLearning, wire: "\"Education and Learning\""},
		{variant: DirectPublishingReadResponseModelGenreItemHistory, wire: "\"History\""},
		{variant: DirectPublishingReadResponseModelGenreItemChildrenSLiterature, wire: "\"Children's Literature\""},
		{variant: DirectPublishingReadResponseModelGenreItemYoungAdult, wire: "\"Young Adult\""},
		{variant: DirectPublishingReadResponseModelGenreItemFairyTalesAndFolklore, wire: "\"Fairy Tales and Folklore\""},
		{variant: DirectPublishingReadResponseModelGenreItemFanFiction, wire: "\"Fan Fiction\""},
		{variant: DirectPublishingReadResponseModelGenreItemGeneralFiction, wire: "\"General Fiction\""},
		{variant: DirectPublishingReadResponseModelGenreItemHealthAndWellness, wire: "\"Health and Wellness\""},
		{variant: DirectPublishingReadResponseModelGenreItemHistoricalFiction, wire: "\"Historical Fiction\""},
		{variant: DirectPublishingReadResponseModelGenreItemHumor, wire: "\"Humor\""},
		{variant: DirectPublishingReadResponseModelGenreItemLiteraryClassics, wire: "\"Literary Classics\""},
		{variant: DirectPublishingReadResponseModelGenreItemPhilosophy, wire: "\"Philosophy\""},
		{variant: DirectPublishingReadResponseModelGenreItemPoetry, wire: "\"Poetry\""},
		{variant: DirectPublishingReadResponseModelGenreItemPoliticsAndGovernment, wire: "\"Politics and Government\""},
		{variant: DirectPublishingReadResponseModelGenreItemPsychology, wire: "\"Psychology\""},
		{variant: DirectPublishingReadResponseModelGenreItemScienceAndNature, wire: "\"Science and Nature\""},
		{variant: DirectPublishingReadResponseModelGenreItemSelfHelp, wire: "\"Self-Help\""},
		{variant: DirectPublishingReadResponseModelGenreItemSpiritualityAndReligion, wire: "\"Spirituality and Religion\""},
		{variant: DirectPublishingReadResponseModelGenreItemTravel, wire: "\"Travel\""},
		{variant: DirectPublishingReadResponseModelGenreItemTrueCrime, wire: "\"True Crime\""},
		{variant: DirectPublishingReadResponseModelGenreItemOther, wire: "\"Other\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded DirectPublishingReadResponseModelGenreItem
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
