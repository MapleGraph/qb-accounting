package services

import (
	"encoding/json"
	"fmt"
	"time"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
)

func formatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func formatTimePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := formatTime(*t)
	return &s
}

func bookModelToDTO(b *models.Book) *dto.BookResponse {
	if b == nil {
		return nil
	}
	return &dto.BookResponse{
		ID:                    b.ID,
		CompanyID:             b.CompanyID,
		Code:                  b.Code,
		Name:                  b.Name,
		BookType:              string(b.BookType),
		BaseCurrencyCode:      b.BaseCurrencyCode,
		ReportingCurrencyCode: b.ReportingCurrencyCode,
		Status:                string(b.Status),
	}
}

func booksModelToDTO(list []*models.Book) []*dto.BookResponse {
	out := make([]*dto.BookResponse, 0, len(list))
	for _, b := range list {
		out = append(out, bookModelToDTO(b))
	}
	return out
}

func fiscalYearModelToDTO(fy *models.FiscalYear) *dto.FiscalYearResponse {
	if fy == nil {
		return nil
	}
	return &dto.FiscalYearResponse{
		ID:            fy.ID,
		BookID:        fy.BookID,
		CompanyID:     fy.CompanyID,
		Code:          fy.Code,
		Name:          fy.Name,
		StartDate:     formatTime(fy.StartDate),
		EndDate:       formatTime(fy.EndDate),
		Status:        string(fy.Status),
		CloseSequence: fy.CloseSequence,
		ClosedAt:      formatTimePtr(fy.ClosedAt),
		ClosedBy:      fy.ClosedBy,
	}
}

func fiscalYearsModelToDTO(list []*models.FiscalYear) []*dto.FiscalYearResponse {
	out := make([]*dto.FiscalYearResponse, 0, len(list))
	for _, fy := range list {
		out = append(out, fiscalYearModelToDTO(fy))
	}
	return out
}

func accountingPeriodModelToDTO(p *models.AccountingPeriod) *dto.AccountingPeriodResponse {
	if p == nil {
		return nil
	}
	return &dto.AccountingPeriodResponse{
		ID:                 p.ID,
		BookID:             p.BookID,
		FiscalYearID:       p.FiscalYearID,
		CompanyID:          p.CompanyID,
		PeriodNo:           p.PeriodNo,
		PeriodName:         p.PeriodName,
		StartDate:          formatTime(p.StartDate),
		EndDate:            formatTime(p.EndDate),
		Status:             string(p.Status),
		IsAdjustmentPeriod: p.IsAdjustmentPeriod,
		SoftLockedAt:       formatTimePtr(p.SoftLockedAt),
		SoftLockedBy:       p.SoftLockedBy,
		HardLockedAt:       formatTimePtr(p.HardLockedAt),
		HardLockedBy:       p.HardLockedBy,
		ClosedAt:           formatTimePtr(p.ClosedAt),
		ClosedBy:           p.ClosedBy,
		LockReason:         p.LockReason,
	}
}

func accountingPeriodsModelToDTO(list []*models.AccountingPeriod) []*dto.AccountingPeriodResponse {
	out := make([]*dto.AccountingPeriodResponse, 0, len(list))
	for _, p := range list {
		out = append(out, accountingPeriodModelToDTO(p))
	}
	return out
}

func accountGroupModelToDTO(g *models.AccountGroup) *dto.AccountGroupResponse {
	if g == nil {
		return nil
	}
	return &dto.AccountGroupResponse{
		ID:            g.ID,
		BookID:        g.BookID,
		Code:          g.Code,
		Name:          g.Name,
		ParentGroupID: g.ParentGroupID,
		AccountNature: string(g.AccountNature),
		SortOrder:     g.SortOrder,
		IsSystem:      g.IsSystem,
	}
}

func accountGroupsModelToDTO(list []*models.AccountGroup) []*dto.AccountGroupResponse {
	out := make([]*dto.AccountGroupResponse, 0, len(list))
	for _, g := range list {
		out = append(out, accountGroupModelToDTO(g))
	}
	return out
}

func accountModelToDTO(a *models.Account) *dto.AccountResponse {
	if a == nil {
		return nil
	}
	return &dto.AccountResponse{
		ID:                 a.ID,
		BookID:             a.BookID,
		CompanyID:          a.CompanyID,
		Code:               a.Code,
		Name:               a.Name,
		DisplayName:        a.DisplayName,
		GroupID:            a.GroupID,
		ParentAccountID:    a.ParentAccountID,
		AccountNature:      string(a.AccountNature),
		Usage:              string(a.Usage),
		NormalBalance:      a.NormalBalance,
		ControlType:        a.ControlType,
		AllowManualPosting: a.AllowManualPosting,
		RequireParty:       a.RequireParty,
		RequireBranch:      a.RequireBranch,
		RequireCostCenter:  a.RequireCostCenter,
		RequireEmployee:    a.RequireEmployee,
		RequireTaxBreakup:  a.RequireTaxBreakup,
		IsSystem:           a.IsSystem,
		IsActive:           a.IsActive,
	}
}

func accountsModelToDTO(list []*models.Account) []*dto.AccountResponse {
	out := make([]*dto.AccountResponse, 0, len(list))
	for _, a := range list {
		out = append(out, accountModelToDTO(a))
	}
	return out
}

func voucherSequenceModelToDTO(v *models.VoucherSequence) *dto.VoucherSequenceResponse {
	if v == nil {
		return nil
	}
	return &dto.VoucherSequenceResponse{
		ID:          v.ID,
		BookID:      v.BookID,
		CompanyID:   v.CompanyID,
		BranchID:    v.BranchID,
		VoucherType: v.VoucherType,
		Prefix:      v.Prefix,
		Suffix:      v.Suffix,
		Padding:     v.Padding,
		NextNumber:  v.NextNumber,
		ResetPolicy: v.ResetPolicy,
		IsActive:    v.IsActive,
	}
}

func journalBatchModelToDTO(b *models.JournalBatch) *dto.JournalBatchResponse {
	if b == nil {
		return nil
	}
	return &dto.JournalBatchResponse{
		ID:           b.ID,
		BookID:       b.BookID,
		CompanyID:    b.CompanyID,
		BranchID:     b.BranchID,
		BatchNo:      b.BatchNo,
		SourceModule: b.SourceModule,
		BatchType:    b.BatchType,
		PostingDate:  formatTime(b.PostingDate),
		Status:       string(b.Status),
		Narration:    b.Narration,
	}
}

func journalBatchesModelToDTO(list []*models.JournalBatch) []*dto.JournalBatchResponse {
	out := make([]*dto.JournalBatchResponse, 0, len(list))
	for _, b := range list {
		out = append(out, journalBatchModelToDTO(b))
	}
	return out
}

func journalLineModelToDTO(ln *models.JournalLine) *dto.JournalLineResponse {
	if ln == nil {
		return nil
	}
	return &dto.JournalLineResponse{
		ID:               ln.ID,
		JournalID:        ln.JournalID,
		LineNo:           ln.LineNo,
		AccountID:        ln.AccountID,
		CompanyID:        ln.CompanyID,
		BranchID:         ln.BranchID,
		PartyID:          ln.PartyID,
		EmployeeID:       ln.EmployeeID,
		CostCenterID:     ln.CostCenterID,
		IncomeHeadID:     ln.IncomeHeadID,
		TaxCodeID:        ln.TaxCodeID,
		Description:      ln.Description,
		DebitAmountTxn:   ln.DebitAmountTxn,
		CreditAmountTxn:  ln.CreditAmountTxn,
		DebitAmountBase:  ln.DebitAmountBase,
		CreditAmountBase: ln.CreditAmountBase,
		ExchangeRate:     ln.ExchangeRate,
	}
}

func journalLinesModelToDTO(lines []*models.JournalLine) []dto.JournalLineResponse {
	out := make([]dto.JournalLineResponse, 0, len(lines))
	for _, ln := range lines {
		out = append(out, *journalLineModelToDTO(ln))
	}
	return out
}

func journalModelToDTO(j *models.Journal, lines []*models.JournalLine) *dto.JournalResponse {
	if j == nil {
		return nil
	}
	var lineDTOs []dto.JournalLineResponse
	if len(lines) > 0 {
		lineDTOs = journalLinesModelToDTO(lines)
	} else {
		lineDTOs = []dto.JournalLineResponse{}
	}
	return &dto.JournalResponse{
		ID:                  j.ID,
		BookID:              j.BookID,
		BatchID:             j.BatchID,
		FiscalYearID:        j.FiscalYearID,
		PeriodID:            j.PeriodID,
		CompanyID:           j.CompanyID,
		BranchID:            j.BranchID,
		JournalNo:           j.JournalNo,
		JournalKind:         string(j.JournalKind),
		Status:              string(j.Status),
		SourceModule:        j.SourceModule,
		SourceDocumentType:  j.SourceDocumentType,
		SourceDocumentID:    j.SourceDocumentID,
		SourceEventID:       j.SourceEventID,
		IdempotencyKey:      j.IdempotencyKey,
		JournalDate:         formatTime(j.JournalDate),
		PostingDate:         formatTime(j.PostingDate),
		CurrencyCode:        j.CurrencyCode,
		ExchangeRate:        j.ExchangeRate,
		ReferenceNo:         j.ReferenceNo,
		ExternalReferenceNo: j.ExternalReferenceNo,
		Narration:           j.Narration,
		ReversalOfJournalID: j.ReversalOfJournalID,
		ReversedByJournalID: j.ReversedByJournalID,
		PostedAt:            formatTimePtr(j.PostedAt),
		PostedBy:            j.PostedBy,
		ReversedAt:          formatTimePtr(j.ReversedAt),
		ReversedBy:          j.ReversedBy,
		Lines:               lineDTOs,
	}
}

func journalsModelToDTO(list []*models.Journal) []*dto.JournalResponse {
	out := make([]*dto.JournalResponse, 0, len(list))
	for _, j := range list {
		out = append(out, journalModelToDTO(j, nil))
	}
	return out
}

func postingRuleVersionModelToDTO(m *models.PostingRuleVersion) (*dto.PostingRuleVersionResponse, error) {
	if m == nil {
		return nil, nil
	}
	var raw json.RawMessage
	if len(m.RulePayload) > 0 {
		b, err := json.Marshal(m.RulePayload)
		if err != nil {
			return nil, fmt.Errorf("rule_payload: %w", err)
		}
		raw = json.RawMessage(b)
	} else {
		raw = json.RawMessage("{}")
	}
	return &dto.PostingRuleVersionResponse{
		ID:                 m.ID,
		SourceModule:       m.SourceModule,
		SourceDocumentType: m.SourceDocumentType,
		VersionNo:          m.VersionNo,
		Name:               m.Name,
		Status:             string(m.Status),
		EffectiveFrom:      formatTime(m.EffectiveFrom),
		EffectiveTo:        formatTimePtr(m.EffectiveTo),
		RulePayload:        raw,
		Notes:              m.Notes,
	}, nil
}
