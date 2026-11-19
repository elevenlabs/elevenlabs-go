package types

import (
	"fmt"
)

type ReviewResponseModelRejectReasonsItem string

const (
	ReviewResponseModelRejectReasonsItemLacksStructure       ReviewResponseModelRejectReasonsItem = "lacks_structure"
	ReviewResponseModelRejectReasonsItemDoesntOpen           ReviewResponseModelRejectReasonsItem = "doesnt_open"
	ReviewResponseModelRejectReasonsItemNotLiteraryWork      ReviewResponseModelRejectReasonsItem = "not_literary_work"
	ReviewResponseModelRejectReasonsItemLanguageNotSupported ReviewResponseModelRejectReasonsItem = "language_not_supported"
	ReviewResponseModelRejectReasonsItemTooShort             ReviewResponseModelRejectReasonsItem = "too_short"
	ReviewResponseModelRejectReasonsItemDuplicate            ReviewResponseModelRejectReasonsItem = "duplicate"
	ReviewResponseModelRejectReasonsItemPromotional          ReviewResponseModelRejectReasonsItem = "promotional"
	ReviewResponseModelRejectReasonsItemFormattingIssues     ReviewResponseModelRejectReasonsItem = "formatting_issues"
	ReviewResponseModelRejectReasonsItemLowQuality           ReviewResponseModelRejectReasonsItem = "low_quality"
	ReviewResponseModelRejectReasonsItemMetadataIncomplete   ReviewResponseModelRejectReasonsItem = "metadata_incomplete"
	ReviewResponseModelRejectReasonsItemMetadataInaccurate   ReviewResponseModelRejectReasonsItem = "metadata_inaccurate"
	ReviewResponseModelRejectReasonsItemTypos                ReviewResponseModelRejectReasonsItem = "typos"
	ReviewResponseModelRejectReasonsItemReviewError          ReviewResponseModelRejectReasonsItem = "review_error"
	ReviewResponseModelRejectReasonsItemSpam                 ReviewResponseModelRejectReasonsItem = "spam"
	ReviewResponseModelRejectReasonsItemLegalViolation       ReviewResponseModelRejectReasonsItem = "legal_violation"
	ReviewResponseModelRejectReasonsItemContentPolicy        ReviewResponseModelRejectReasonsItem = "content_policy"
	ReviewResponseModelRejectReasonsItemPublicDomain         ReviewResponseModelRejectReasonsItem = "public_domain"
	ReviewResponseModelRejectReasonsItemOther                ReviewResponseModelRejectReasonsItem = "other"
)

func NewReviewResponseModelRejectReasonsItemFromString(s string) (ReviewResponseModelRejectReasonsItem, error) {
	switch s {
	case "lacks_structure":
		return ReviewResponseModelRejectReasonsItemLacksStructure, nil
	case "doesnt_open":
		return ReviewResponseModelRejectReasonsItemDoesntOpen, nil
	case "not_literary_work":
		return ReviewResponseModelRejectReasonsItemNotLiteraryWork, nil
	case "language_not_supported":
		return ReviewResponseModelRejectReasonsItemLanguageNotSupported, nil
	case "too_short":
		return ReviewResponseModelRejectReasonsItemTooShort, nil
	case "duplicate":
		return ReviewResponseModelRejectReasonsItemDuplicate, nil
	case "promotional":
		return ReviewResponseModelRejectReasonsItemPromotional, nil
	case "formatting_issues":
		return ReviewResponseModelRejectReasonsItemFormattingIssues, nil
	case "low_quality":
		return ReviewResponseModelRejectReasonsItemLowQuality, nil
	case "metadata_incomplete":
		return ReviewResponseModelRejectReasonsItemMetadataIncomplete, nil
	case "metadata_inaccurate":
		return ReviewResponseModelRejectReasonsItemMetadataInaccurate, nil
	case "typos":
		return ReviewResponseModelRejectReasonsItemTypos, nil
	case "review_error":
		return ReviewResponseModelRejectReasonsItemReviewError, nil
	case "spam":
		return ReviewResponseModelRejectReasonsItemSpam, nil
	case "legal_violation":
		return ReviewResponseModelRejectReasonsItemLegalViolation, nil
	case "content_policy":
		return ReviewResponseModelRejectReasonsItemContentPolicy, nil
	case "public_domain":
		return ReviewResponseModelRejectReasonsItemPublicDomain, nil
	case "other":
		return ReviewResponseModelRejectReasonsItemOther, nil
	}
	var t ReviewResponseModelRejectReasonsItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReviewResponseModelRejectReasonsItem) Ptr() *ReviewResponseModelRejectReasonsItem {
	return &e
}
