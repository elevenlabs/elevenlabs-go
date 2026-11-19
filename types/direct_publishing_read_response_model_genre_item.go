package types

import (
	"fmt"
)

type DirectPublishingReadResponseModelGenreItem string

const (
	DirectPublishingReadResponseModelGenreItemFantasy                   DirectPublishingReadResponseModelGenreItem = "Fantasy"
	DirectPublishingReadResponseModelGenreItemRomance                   DirectPublishingReadResponseModelGenreItem = "Romance"
	DirectPublishingReadResponseModelGenreItemScienceFiction            DirectPublishingReadResponseModelGenreItem = "Science Fiction"
	DirectPublishingReadResponseModelGenreItemMysteryAndThriller        DirectPublishingReadResponseModelGenreItem = "Mystery and Thriller"
	DirectPublishingReadResponseModelGenreItemActionAndAdventure        DirectPublishingReadResponseModelGenreItem = "Action and Adventure"
	DirectPublishingReadResponseModelGenreItemDystopia                  DirectPublishingReadResponseModelGenreItem = "Dystopia"
	DirectPublishingReadResponseModelGenreItemBusinessAndEconomics      DirectPublishingReadResponseModelGenreItem = "Business and Economics"
	DirectPublishingReadResponseModelGenreItemTechnology                DirectPublishingReadResponseModelGenreItem = "Technology"
	DirectPublishingReadResponseModelGenreItemChristianAndInspirational DirectPublishingReadResponseModelGenreItem = "Christian & Inspirational"
	DirectPublishingReadResponseModelGenreItemHorror                    DirectPublishingReadResponseModelGenreItem = "Horror"
	DirectPublishingReadResponseModelGenreItemBiographyAndMemoir        DirectPublishingReadResponseModelGenreItem = "Biography and Memoir"
	DirectPublishingReadResponseModelGenreItemEducationAndLearning      DirectPublishingReadResponseModelGenreItem = "Education and Learning"
	DirectPublishingReadResponseModelGenreItemHistory                   DirectPublishingReadResponseModelGenreItem = "History"
	DirectPublishingReadResponseModelGenreItemChildrenSLiterature       DirectPublishingReadResponseModelGenreItem = "Children's Literature"
	DirectPublishingReadResponseModelGenreItemYoungAdult                DirectPublishingReadResponseModelGenreItem = "Young Adult"
	DirectPublishingReadResponseModelGenreItemFairyTalesAndFolklore     DirectPublishingReadResponseModelGenreItem = "Fairy Tales and Folklore"
	DirectPublishingReadResponseModelGenreItemFanFiction                DirectPublishingReadResponseModelGenreItem = "Fan Fiction"
	DirectPublishingReadResponseModelGenreItemGeneralFiction            DirectPublishingReadResponseModelGenreItem = "General Fiction"
	DirectPublishingReadResponseModelGenreItemHealthAndWellness         DirectPublishingReadResponseModelGenreItem = "Health and Wellness"
	DirectPublishingReadResponseModelGenreItemHistoricalFiction         DirectPublishingReadResponseModelGenreItem = "Historical Fiction"
	DirectPublishingReadResponseModelGenreItemHumor                     DirectPublishingReadResponseModelGenreItem = "Humor"
	DirectPublishingReadResponseModelGenreItemLiteraryClassics          DirectPublishingReadResponseModelGenreItem = "Literary Classics"
	DirectPublishingReadResponseModelGenreItemPhilosophy                DirectPublishingReadResponseModelGenreItem = "Philosophy"
	DirectPublishingReadResponseModelGenreItemPoetry                    DirectPublishingReadResponseModelGenreItem = "Poetry"
	DirectPublishingReadResponseModelGenreItemPoliticsAndGovernment     DirectPublishingReadResponseModelGenreItem = "Politics and Government"
	DirectPublishingReadResponseModelGenreItemPsychology                DirectPublishingReadResponseModelGenreItem = "Psychology"
	DirectPublishingReadResponseModelGenreItemScienceAndNature          DirectPublishingReadResponseModelGenreItem = "Science and Nature"
	DirectPublishingReadResponseModelGenreItemSelfHelp                  DirectPublishingReadResponseModelGenreItem = "Self-Help"
	DirectPublishingReadResponseModelGenreItemSpiritualityAndReligion   DirectPublishingReadResponseModelGenreItem = "Spirituality and Religion"
	DirectPublishingReadResponseModelGenreItemTravel                    DirectPublishingReadResponseModelGenreItem = "Travel"
	DirectPublishingReadResponseModelGenreItemTrueCrime                 DirectPublishingReadResponseModelGenreItem = "True Crime"
	DirectPublishingReadResponseModelGenreItemOther                     DirectPublishingReadResponseModelGenreItem = "Other"
)

func NewDirectPublishingReadResponseModelGenreItemFromString(s string) (DirectPublishingReadResponseModelGenreItem, error) {
	switch s {
	case "Fantasy":
		return DirectPublishingReadResponseModelGenreItemFantasy, nil
	case "Romance":
		return DirectPublishingReadResponseModelGenreItemRomance, nil
	case "Science Fiction":
		return DirectPublishingReadResponseModelGenreItemScienceFiction, nil
	case "Mystery and Thriller":
		return DirectPublishingReadResponseModelGenreItemMysteryAndThriller, nil
	case "Action and Adventure":
		return DirectPublishingReadResponseModelGenreItemActionAndAdventure, nil
	case "Dystopia":
		return DirectPublishingReadResponseModelGenreItemDystopia, nil
	case "Business and Economics":
		return DirectPublishingReadResponseModelGenreItemBusinessAndEconomics, nil
	case "Technology":
		return DirectPublishingReadResponseModelGenreItemTechnology, nil
	case "Christian & Inspirational":
		return DirectPublishingReadResponseModelGenreItemChristianAndInspirational, nil
	case "Horror":
		return DirectPublishingReadResponseModelGenreItemHorror, nil
	case "Biography and Memoir":
		return DirectPublishingReadResponseModelGenreItemBiographyAndMemoir, nil
	case "Education and Learning":
		return DirectPublishingReadResponseModelGenreItemEducationAndLearning, nil
	case "History":
		return DirectPublishingReadResponseModelGenreItemHistory, nil
	case "Children's Literature":
		return DirectPublishingReadResponseModelGenreItemChildrenSLiterature, nil
	case "Young Adult":
		return DirectPublishingReadResponseModelGenreItemYoungAdult, nil
	case "Fairy Tales and Folklore":
		return DirectPublishingReadResponseModelGenreItemFairyTalesAndFolklore, nil
	case "Fan Fiction":
		return DirectPublishingReadResponseModelGenreItemFanFiction, nil
	case "General Fiction":
		return DirectPublishingReadResponseModelGenreItemGeneralFiction, nil
	case "Health and Wellness":
		return DirectPublishingReadResponseModelGenreItemHealthAndWellness, nil
	case "Historical Fiction":
		return DirectPublishingReadResponseModelGenreItemHistoricalFiction, nil
	case "Humor":
		return DirectPublishingReadResponseModelGenreItemHumor, nil
	case "Literary Classics":
		return DirectPublishingReadResponseModelGenreItemLiteraryClassics, nil
	case "Philosophy":
		return DirectPublishingReadResponseModelGenreItemPhilosophy, nil
	case "Poetry":
		return DirectPublishingReadResponseModelGenreItemPoetry, nil
	case "Politics and Government":
		return DirectPublishingReadResponseModelGenreItemPoliticsAndGovernment, nil
	case "Psychology":
		return DirectPublishingReadResponseModelGenreItemPsychology, nil
	case "Science and Nature":
		return DirectPublishingReadResponseModelGenreItemScienceAndNature, nil
	case "Self-Help":
		return DirectPublishingReadResponseModelGenreItemSelfHelp, nil
	case "Spirituality and Religion":
		return DirectPublishingReadResponseModelGenreItemSpiritualityAndReligion, nil
	case "Travel":
		return DirectPublishingReadResponseModelGenreItemTravel, nil
	case "True Crime":
		return DirectPublishingReadResponseModelGenreItemTrueCrime, nil
	case "Other":
		return DirectPublishingReadResponseModelGenreItemOther, nil
	}
	var t DirectPublishingReadResponseModelGenreItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DirectPublishingReadResponseModelGenreItem) Ptr() *DirectPublishingReadResponseModelGenreItem {
	return &e
}
