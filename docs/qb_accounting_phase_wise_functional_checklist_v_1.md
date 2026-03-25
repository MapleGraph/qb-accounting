# QB Accounting Phase-wise Functional Checklist v1

## 1. Purpose

This document defines the functional scope of the **Accounting** module in the new QB ERP modular monolith.

The objective is to ensure that Accounting owns only the capabilities that must belong to the financial books layer, while avoiding duplication of master data and operational document ownership already assigned to other modules.

This checklist is intended to serve as the foundation for:
- feature planning
- phased development tracking
- schema design
- data dictionary preparation
- integration contract definition between Accounting and other QB ERP modules

---

## 2. Module Boundary

### 2.1 Modules already owning business capabilities

#### Setup owns
- organization
- company
- branch
- currency
- payment provider master data

#### Catalogue owns
- items
- item groups
- classifications
- income heads
- cost centres
- taxes
- tax groups
- tax codes

#### Employee owns
- employee registration
- roles
- departments
- payroll in future phase

#### CRM owns
- customer registration
- customer groups
- party master
- promotion master

#### Billing owns
- sales quotation
- sales order
- proforma invoice
- sales invoice
- refund invoice
- customer receipt

#### Purchase will own
- vendor registration
- purchase order
- vendor bill
- debit note
- related vendor-side operational documents

#### Expense will own
- organization-level expense capture
- expense workflows
- expense approvals
- expense source documents

---

## 3. Accounting Module Scope Statement

The Accounting module should **not** duplicate the schemas of Setup, Catalogue, CRM, Billing, Purchase, Expense, or Employee.

The Accounting module should be the **system of record for financial recognition and financial controls**.

Accounting should own the following domains:
- chart of accounts
- accounting books
- fiscal years and accounting periods
- journal engine
- posting engine
- source document to ledger traceability
- AR/AP accounting state
- bank accounting and reconciliation
- tax ledgering and tax reconciliation
- period close and year-end close
- financial control workflows
- inter-branch and inter-company accounting
- financial reporting and ledgers

---

## 4. Ownership Matrix

| Capability | Owning Module |
|---|---|
| Organization / Company / Branch / Currency | Setup |
| Tax masters / Tax codes / Tax groups | Catalogue |
| Cost centres / Income heads | Catalogue |
| Customer / Party master | CRM |
| Vendor master | Purchase or Party layer |
| Employee master | Employee |
| Sales commercial documents | Billing |
| Purchase commercial documents | Purchase |
| Expense source documents | Expense |
| Payroll source documents | Employee / Payroll |
| Chart of accounts | Accounting |
| Accounting periods | Accounting |
| Journals / vouchers | Accounting |
| Posting engine | Accounting |
| AR/AP open item accounting | Accounting |
| Bank reconciliation | Accounting |
| Tax ledgering | Accounting |
| Financial close | Accounting |
| Financial statements | Accounting |
| Inter-company accounting | Accounting |

---

## 5. Accounting Design Principles

### 5.1 Core principles
- Accounting owns the **financial effect**, not the business source document.
- Source modules remain the source of truth for document lifecycle.
- Accounting receives posting inputs and creates immutable financial postings.
- Posted journals must be append-only.
- Reversal is a first-class action; deletion of posted financial impact is not allowed.
- Accounting must store immutable snapshots required for audit.
- Every financial posting must be traceable back to its originating module and document.

### 5.2 Mandatory technical invariants
- every posting must include source module
- every posting must include source document type
- every posting must include source document id
- every posting must include source event id or idempotency key
- every journal must balance
- posting into a closed period must be blocked
- control account behavior must be enforced by rules, not convention alone
- automated posting rules must be versioned
- posted amounts must retain currency snapshot and exchange rate snapshot where applicable
- manual overrides must require audit trail and reason codes

---

## 6. Integration Contract Principles

Accounting should expose a posting interface that operational modules use when a financially relevant event occurs.

### 6.1 Source modules that must integrate with Accounting
- Billing
- Purchase
- Expense
- Employee / Payroll in later phase
- Setup and Catalogue for reference resolution only

### 6.2 Posting payload should conceptually include
- source module
- source document type
- source document id
- source event id
- company id
- branch id
- posting date
- document date
- currency and exchange metadata
- party references
- tax breakup snapshot
- dimension snapshot
- narration / reference metadata
- accounting impact payload or enough source detail for rule-based derivation

### 6.3 Accounting response / state should include
- posting status
- journal id(s)
- failure reason if rejected
- reversal linkage if reversed
- posting rule version used

---

## 7. Phase Plan Overview

The Accounting module should be built in the following phases:

1. Phase 0 - Scope and posting architecture
2. Phase 1 - Accounting foundation
3. Phase 2 - Automated posting engine
4. Phase 3 - AR/AP accounting layer
5. Phase 4 - Banking and reconciliation
6. Phase 5 - Tax accounting
7. Phase 6 - Period close and financial controls
8. Phase 7 - Inter-branch and inter-company accounting
9. Phase 8 - Financial reporting and intelligence

---

# 8. Detailed Phase-wise Checklist

## Phase 0 - Scope Lock and Architecture Decisions

### Objective
Freeze the module boundary and posting architecture before schema design.

### Checklist
- [ ] Finalize accounting scope statement
- [ ] Finalize ownership matrix across Setup, Catalogue, CRM, Billing, Purchase, Expense, Employee, and Accounting
- [ ] Define which documents are source-owned vs accounting-owned
- [ ] Define posting lifecycle states: draft, preview, posted, failed, reversed, cancelled
- [ ] Define whether posting is synchronous, asynchronous, or hybrid
- [ ] Define idempotency strategy for posting requests
- [ ] Define source event model for document posting
- [ ] Define chart of accounts design standard
- [ ] Define control account policy
- [ ] Define dimension strategy: branch, cost centre, department, income head, project, etc.
- [ ] Define branch-level balancing rules
- [ ] Define inter-company first-release scope
- [ ] Define foreign currency handling rules
- [ ] Define reversal rules
- [ ] Define audit log requirements
- [ ] Define manual journal permissions and approval requirements
- [ ] Define accounting error handling / retry / replay workflow
- [ ] Define opening balance migration approach

### Deliverables
- Accounting scope note
- integration contract v1
- posting lifecycle note
- finance invariants checklist

### Exit Criteria
- accounting ownership is unambiguous
- integration contract is frozen for Phase 1 and Phase 2 work

---

## Phase 1 - Accounting Foundation

### Objective
Create the bookkeeping base that can operate independently of downstream automations.

### Functional Areas
#### 1. Fiscal structure
- [ ] Define fiscal year setup
- [ ] Define accounting periods
- [ ] Support monthly period creation
- [ ] Support period statuses: open, soft-locked, hard-locked, closed
- [ ] Support book-level or company-level period policy

#### 2. Books and chart of accounts
- [ ] Create accounting book / ledger definition
- [ ] Create account groups
- [ ] Create chart of accounts
- [ ] Support parent-child account hierarchy
- [ ] Support account types: asset, liability, equity, income, expense
- [ ] Support posting permissions on accounts
- [ ] Mark control accounts
- [ ] Mark system-reserved accounts
- [ ] Support account activation / deactivation policy
- [ ] Support account mapping to branch / dimension restrictions where required

#### 3. Manual journal engine
- [ ] Create manual journal entry
- [ ] Support journal date and posting date
- [ ] Support narration and reference fields
- [ ] Support multi-line debit / credit entry
- [ ] Enforce balanced journal validation
- [ ] Support draft journal save
- [ ] Support journal edit before posting
- [ ] Post journal to ledger
- [ ] Reverse posted journal
- [ ] Support recurring journal template marker for later phase
- [ ] Support adjustment journal type
- [ ] Support opening balance journal type

#### 4. Voucher control
- [ ] Define voucher types
- [ ] Define numbering strategy by company / branch / voucher type
- [ ] Support reference number uniqueness rules

#### 5. Audit basics
- [ ] Record journal creation metadata
- [ ] Record posting metadata
- [ ] Record reversal metadata
- [ ] Record manual edits before posting
- [ ] Record reason codes for manual accounting override actions

### Reports in Phase 1
- [ ] Trial balance
- [ ] General ledger
- [ ] Journal register
- [ ] Account ledger

### Deliverables
- chart of accounts capability
- manual journal capability
- base ledger reports

### Exit Criteria
- finance can create chart of accounts
- finance can open periods and post manual journals
- finance can view TB and GL

---

## Phase 2 - Automated Posting Engine

### Objective
Allow Billing, Purchase, and Expense to create business documents while Accounting owns the books impact.

### Functional Areas
#### 1. Posting framework
- [ ] Define posting source registry
- [ ] Define source document type registry
- [ ] Define posting template / rule framework
- [ ] Define posting inbox / staging table or queue abstraction
- [ ] Validate required references before posting
- [ ] Support idempotent posting request handling
- [ ] Prevent duplicate financial posting for same source event
- [ ] Link source document to posted journal(s)
- [ ] Support posting preview mode
- [ ] Support posting retry / replay
- [ ] Support posting dead-letter / failure queue
- [ ] Support posting rule versioning
- [ ] Support posting exception diagnostics

#### 2. Billing integrations
- [ ] Post sales invoice to accounting
- [ ] Post refund invoice / credit impact
- [ ] Post customer receipt
- [ ] Support receipt against invoice settlement creation trigger
- [ ] Support customer advance receipt recognition

#### 3. Purchase integrations
- [ ] Post vendor bill
- [ ] Post debit note
- [ ] Post vendor payment
- [ ] Support vendor advance payment recognition

#### 4. Expense integrations
- [ ] Post expense booking
- [ ] Post expense payment / reimbursement effect
- [ ] Support expense accrual if expense module triggers accrual events later

#### 5. Posting traceability
- [ ] Store source document number snapshot
- [ ] Store party snapshot
- [ ] Store tax snapshot
- [ ] Store dimension snapshot
- [ ] Store exchange rate snapshot if applicable
- [ ] Store original source status at time of posting

### Deliverables
- accounting posting API / internal service contract
- posting rule engine v1
- billing / purchase / expense posting integration v1

### Exit Criteria
- operational modules can trigger financial postings without accounting duplicating their schemas
- accounting can reject invalid postings with clear diagnostics

---

## Phase 3 - AR/AP Accounting Layer

### Objective
Move from raw ledger postings to proper receivable and payable accounting state.

### Functional Areas
#### 1. Receivables accounting
- [ ] Define AR control accounts
- [ ] Create customer open item ledger
- [ ] Track invoice-wise outstanding
- [ ] Track receipt-wise unapplied balance
- [ ] Support partial allocations
- [ ] Support overpayment / advance balance
- [ ] Support credit adjustment against outstanding
- [ ] Support receivable write-off
- [ ] Support allocation reversal / unapply

#### 2. Payables accounting
- [ ] Define AP control accounts
- [ ] Create vendor open item ledger
- [ ] Track bill-wise outstanding
- [ ] Track unapplied vendor payment balance
- [ ] Support partial bill settlement
- [ ] Support vendor advance tracking
- [ ] Support debit adjustment against payables
- [ ] Support payable write-off where policy allows
- [ ] Support settlement reversal / unapply

#### 3. Reconciliation and statements
- [ ] Customer statement from accounting view
- [ ] Vendor statement from accounting view
- [ ] AR aging buckets
- [ ] AP aging buckets
- [ ] Control account vs open item reconciliation report
- [ ] Source document vs accounting outstanding comparison report
- [ ] Party ledger

### Deliverables
- AR accounting layer
- AP accounting layer
- aging and statement reports

### Exit Criteria
- receivable and payable balances are operationally traceable and reconcilable

---

## Phase 4 - Banking and Reconciliation

### Objective
Enable finance teams to close bank and cash positions accurately from the accounting module.

### Functional Areas
#### 1. Bank accounting setup
- [ ] Map bank accounts to GL accounts
- [ ] Map company / branch to bank account usage
- [ ] Define clearing account support
- [ ] Define settlement account support

#### 2. Statement ingestion
- [ ] Bank statement import staging
- [ ] Statement line normalization
- [ ] Duplicate statement detection
- [ ] Statement source tracking

#### 3. Reconciliation workbench
- [ ] Auto-match accounting entries to statement lines
- [ ] Manual match support
- [ ] Split match support
- [ ] Merge match support
- [ ] Unmatched line handling
- [ ] Bank charge adjustment handling
- [ ] Interest posting support
- [ ] Miscellaneous bank adjustment support
- [ ] Reconciliation status tracking
- [ ] Reconciliation cutoff by period

#### 4. Optional cash instruments
- [ ] Cheque handling if required
- [ ] PDC handling if required
- [ ] cheque / PDC realization accounting state

### Reports in Phase 4
- [ ] Bank reconciliation statement
- [ ] Unreconciled bank transactions report
- [ ] Cleared vs uncleared entries report
- [ ] period-wise reconciliation summary

### Deliverables
- bank reconciliation workbench
- bank accounting reports

### Exit Criteria
- monthly bank reconciliation can be completed inside the product

---

## Phase 5 - Tax Accounting

### Objective
Translate source-module tax logic into ledger-owned tax accounting and reconciliation.

### Functional Areas
#### 1. Tax ledger mapping
- [ ] Map tax codes and tax groups from Catalogue to tax ledger accounts
- [ ] Support input tax accounting
- [ ] Support output tax accounting
- [ ] Support tax payable / tax receivable accounts
- [ ] Support TDS accounting
- [ ] Support TCS accounting if needed
- [ ] Support reverse charge accounting hooks if applicable

#### 2. Tax accounting records
- [ ] Store tax breakup snapshot on posted entries
- [ ] Maintain tax register by company / branch / registration
- [ ] Maintain period-wise tax summaries
- [ ] Support tax adjustment journals
- [ ] Support tax settlement tracking

#### 3. Tax reconciliation
- [ ] Source tax vs ledger tax comparison
- [ ] registration-wise reconciliation
- [ ] period-wise mismatch detection
- [ ] tax suspense / unresolved mismatch tracking

### Reports in Phase 5
- [ ] Tax ledger register
- [ ] input vs output summary
- [ ] TDS/TCS ledger summary
- [ ] tax mismatch report
- [ ] tax payable / receivable position report

### Deliverables
- tax accounting register layer
- tax reconciliation workbench v1

### Exit Criteria
- accounting can explain tax balances from ledger-owned data without owning duplicate tax masters

---

## Phase 6 - Period Close and Financial Controls

### Objective
Make Accounting suitable for controlled enterprise finance operations.

### Functional Areas
#### 1. Close workflow
- [ ] Define close checklist template
- [ ] Define month-end close workflow
- [ ] Define quarter-end close workflow
- [ ] Define year-end close workflow
- [ ] Support close status tracking
- [ ] Detect blockers to close
- [ ] Support soft close
- [ ] Support hard close
- [ ] Support reopen with authorization

#### 2. Period-end accounting
- [ ] Accrual journal support
- [ ] Prepaid accounting support
- [ ] Deferred income support
- [ ] Deferred expense support
- [ ] Provision journal support
- [ ] Recurring journals
- [ ] reversal-on-next-period options
- [ ] year-end closing entries
- [ ] retained earnings roll-forward
- [ ] adjustment journal management

#### 3. Financial controls
- [ ] Manual journal approval workflow
- [ ] Sensitive account posting approval workflow
- [ ] Reason codes for overrides
- [ ] Suspense account monitoring
- [ ] Unposted source document report
- [ ] backdated posting detection
- [ ] closed-period exception report
- [ ] maker-checker controls for manual accounting actions

### Reports in Phase 6
- [ ] Close checklist status report
- [ ] Unposted source docs report
- [ ] Suspense account report
- [ ] backdated entry report
- [ ] manual override report

### Deliverables
- period close engine
- financial control workflow layer

### Exit Criteria
- finance can run month-end close in the system with traceable controls

---

## Phase 7 - Inter-Branch and Inter-Company Accounting

### Objective
Support multi-entity organizations such as Wonderla without manual side-book practices.

### Functional Areas
#### 1. Inter-branch accounting
- [ ] Define branch balancing rules
- [ ] Support branch-level balancing entries
- [ ] Support branch transfer accounting hooks
- [ ] Support branch-level trial balance

#### 2. Inter-company accounting
- [ ] Define due-to / due-from framework
- [ ] Define inter-company control accounts
- [ ] Support cross-company posting templates
- [ ] Support mirrored inter-company entries where required
- [ ] Support inter-company mismatch detection
- [ ] Support elimination journal type
- [ ] Support consolidation adjustment journal type

#### 3. Group finance visibility
- [ ] Combined multi-entity trial balance
- [ ] entity-level and branch-level view switching
- [ ] inter-company outstanding report
- [ ] elimination tracking report

### Deliverables
- inter-company accounting framework
- consolidation-ready journal capability

### Exit Criteria
- multi-company books can be kept aligned without uncontrolled manual journals

---

## Phase 8 - Financial Reporting and Intelligence

### Objective
Establish Accounting as the source of truth for finance reporting.

### Functional Areas
#### 1. Core financial reports
- [ ] Trial balance
- [ ] General ledger
- [ ] Journal register
- [ ] Day book
- [ ] Account ledger
- [ ] Party ledger
- [ ] Balance sheet
- [ ] Profit and loss
- [ ] Cash flow statement

#### 2. Subledger and control reports
- [ ] AR aging
- [ ] AP aging
- [ ] customer statement
- [ ] vendor statement
- [ ] open item reports
- [ ] control account reconciliation reports

#### 3. Tax and bank reports
- [ ] tax ledger reports
- [ ] bank reconciliation reports
- [ ] unreconciled bank report

#### 4. Close and exceptions
- [ ] posting failure report
- [ ] unposted source docs report
- [ ] suspense report
- [ ] close status dashboard

#### 5. Dimensional finance reporting
- [ ] branch-wise P&L
- [ ] cost centre-wise view
- [ ] income head-wise drilldown
- [ ] department-wise finance view where applicable
- [ ] project-wise extension hooks if introduced later

#### 6. Historical stability
- [ ] snapshot-based reporting behavior
- [ ] historical period consistency rules
- [ ] export-friendly reporting views

### Deliverables
- finance report suite v1
- dimensional reporting v1

### Exit Criteria
- finance reporting can be served from accounting-owned structures with consistent books logic

---

# 9. Recommended Release Grouping

## Release 1
- Phase 0
- Phase 1
- Phase 2 foundation subset
- Reports: TB, GL, journal register

## Release 2
- remaining Phase 2
- Phase 3
- basic AR/AP reports

## Release 3
- Phase 4
- Phase 5
- Phase 6

## Release 4
- Phase 7
- Phase 8

---

# 10. Priority Features That Should Not Be Deferred

The following should be built early because they are hard to retrofit later:
- source document to journal traceability
- idempotent posting
- immutable posting snapshot strategy
- chart of accounts hierarchy and control account design
- accounting period lock model
- reversal framework
- AR/AP open item model
- audit trail model
- reason code model for manual overrides
- posting rule versioning
- branch / dimension propagation in postings

---

# 11. Features Explicitly Out of Scope for Accounting Ownership

The Accounting module should not own the primary schemas for:
- company master
- branch master
- currency master
- customer master
- vendor master
- employee master
- item master
- tax code / tax group master
- sales invoice source tables
- refund invoice source tables
- customer receipt source tables
- purchase order source tables
- vendor bill source tables
- expense source tables
- payroll source tables

Accounting may store only immutable references or snapshots required for posting and audit.

---

# 12. Definition of Done Before Schema Design Starts

The team should not start schema finalization until the following are agreed:
- [ ] module boundary finalized
- [ ] ownership matrix finalized
- [ ] posting contract finalized
- [ ] phase plan approved
- [ ] control account policy approved
- [ ] period lock policy approved
- [ ] reversal policy approved
- [ ] dimensional reporting policy approved
- [ ] AR/AP open item behavior approved
- [ ] inter-company first release scope approved

---

# 13. Next Document to Prepare

Once this checklist is approved, the next document should be:

## QB Accounting Schema and Data Dictionary v1

Recommended schema design order:
1. accounting foundation
2. chart of accounts
3. journals and ledger postings
4. posting source registry and traceability
5. AR/AP open items and allocations
6. bank reconciliation
7. tax ledgers
8. close controls and audit trail
9. inter-company support
10. reporting views / materialized reporting structures

---

# 14. Suggested Immediate Next Step

Use this checklist to finalize the Accounting module scope.

After approval, start schema design for only these first domains:
- accounting periods
- chart of accounts
- journals
- posting engine
- source reference traceability
- AR/AP open items

These form the backbone for every later accounting feature.

