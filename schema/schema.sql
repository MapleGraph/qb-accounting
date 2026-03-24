-- QB Accounting Schema v1
-- Scope: accounting foundation, posting engine, source traceability, and AR/AP open items

create extension if not exists pgcrypto;
set search_path to public;

do $$ begin
    create type book_type_enum as enum ('PRIMARY', 'STATUTORY', 'MANAGEMENT');
exception when duplicate_object then null; end $$;

do $$ begin
    create type book_status_enum as enum ('ACTIVE', 'INACTIVE');
exception when duplicate_object then null; end $$;

do $$ begin
    create type fiscal_year_status_enum as enum ('DRAFT', 'OPEN', 'CLOSED', 'ARCHIVED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type period_status_enum as enum ('DRAFT', 'OPEN', 'SOFT_LOCKED', 'HARD_LOCKED', 'CLOSED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type account_nature_enum as enum ('ASSET', 'LIABILITY', 'EQUITY', 'INCOME', 'EXPENSE');
exception when duplicate_object then null; end $$;

do $$ begin
    create type account_usage_enum as enum ('HEADER', 'POSTABLE');
exception when duplicate_object then null; end $$;

do $$ begin
    create type journal_batch_status_enum as enum ('OPEN', 'POSTED', 'FAILED', 'CANCELLED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type journal_kind_enum as enum ('MANUAL', 'OPENING', 'ADJUSTMENT', 'REVERSAL', 'SYSTEM', 'CLOSING');
exception when duplicate_object then null; end $$;

do $$ begin
    create type journal_status_enum as enum ('DRAFT', 'POSTED', 'REVERSED', 'CANCELLED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type posting_rule_status_enum as enum ('DRAFT', 'ACTIVE', 'RETIRED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type posting_request_status_enum as enum ('RECEIVED', 'VALIDATED', 'POSTED', 'FAILED', 'REVERSED', 'IGNORED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type journal_link_role_enum as enum ('PRIMARY', 'REVERSAL', 'ADJUSTMENT');
exception when duplicate_object then null; end $$;

do $$ begin
    create type party_type_enum as enum ('CUSTOMER', 'VENDOR', 'PARTY', 'EMPLOYEE', 'OTHER');
exception when duplicate_object then null; end $$;

do $$ begin
    create type open_item_side_enum as enum ('RECEIVABLE', 'PAYABLE');
exception when duplicate_object then null; end $$;

do $$ begin
    create type open_item_status_enum as enum ('OPEN', 'PARTIALLY_ALLOCATED', 'SETTLED', 'WRITTEN_OFF', 'CANCELLED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type allocation_status_enum as enum ('APPLIED', 'REVERSED');
exception when duplicate_object then null; end $$;

do $$ begin
    create type open_item_adjustment_type_enum as enum ('WRITE_OFF', 'ROUND_OFF', 'MANUAL_ADJUSTMENT', 'FX_REVALUATION');
exception when duplicate_object then null; end $$;

create or replace function touch_updated_at()
returns trigger
language plpgsql
as $$
begin
    new.updated_at = now();
    return new;
end;
$$;

create table if not exists books (
    id uuid default gen_random_uuid() not null,
    company_id uuid not null,
    code varchar(30) not null,
    name varchar(120) not null,
    book_type book_type_enum default 'PRIMARY' not null,
    base_currency_code char(3) not null,
    reporting_currency_code char(3),
    status book_status_enum default 'ACTIVE' not null,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_books primary key (id),
    constraint uq_books_1 unique (company_id, code)
);

create index if not exists ix_books_1 on books (company_id);
create index if not exists ix_books_2 on books (status);

comment on table books is 'Accounting book master owned by Accounting. Represents a primary, statutory, or management book for one company without duplicating the company master from Setup.';
comment on column books.id is 'Primary key.';
comment on column books.company_id is 'Soft reference to Setup.company.';
comment on column books.code is 'Business code for the book, such as PRIMARY.';
comment on column books.name is 'Display name of the accounting book.';
comment on column books.book_type is 'Book purpose classification.';
comment on column books.base_currency_code is 'ISO currency code used as the base reporting currency.';
comment on column books.reporting_currency_code is 'Optional reporting currency when management reporting differs from base currency.';
comment on column books.status is 'Operational status of the book.';
comment on column books.created_by is 'User who created the record.';
comment on column books.created_at is 'Record creation timestamp.';
comment on column books.updated_by is 'User who last updated the record.';
comment on column books.updated_at is 'Last update timestamp.';

create table if not exists fiscal_years (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    code varchar(20) not null,
    name varchar(80) not null,
    start_date date not null,
    end_date date not null,
    status fiscal_year_status_enum default 'DRAFT' not null,
    close_sequence integer default 0 not null,
    closed_at timestamptz,
    closed_by uuid,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_fiscal_years primary key (id),
    constraint uq_fiscal_years_1 unique (book_id, code),
    constraint uq_fiscal_years_2 unique (book_id, start_date, end_date),
    constraint ck_fiscal_years_1 check (end_date >= start_date),
    constraint fk_fiscal_years_1 foreign key (book_id) references books (id)
);

create index if not exists ix_fiscal_years_1 on fiscal_years (company_id);
create index if not exists ix_fiscal_years_2 on fiscal_years (status);

comment on table fiscal_years is 'Book-specific fiscal years owned by Accounting. Supports year open/close without relying on an external financial calendar.';
comment on column fiscal_years.id is 'Primary key.';
comment on column fiscal_years.book_id is 'Reference to books.';
comment on column fiscal_years.company_id is 'Denormalized company reference for query efficiency.';
comment on column fiscal_years.code is 'Short fiscal year code such as FY2026.';
comment on column fiscal_years.name is 'Display name such as FY 2026-27.';
comment on column fiscal_years.start_date is 'Fiscal year start date.';
comment on column fiscal_years.end_date is 'Fiscal year end date.';
comment on column fiscal_years.status is 'Lifecycle status.';
comment on column fiscal_years.close_sequence is 'Version-like close counter; increments when year-end close is rerun.';
comment on column fiscal_years.closed_at is 'Timestamp when the year was closed.';
comment on column fiscal_years.closed_by is 'User who closed the year.';
comment on column fiscal_years.created_by is 'User who created the record.';
comment on column fiscal_years.created_at is 'Record creation timestamp.';
comment on column fiscal_years.updated_by is 'User who last updated the record.';
comment on column fiscal_years.updated_at is 'Last update timestamp.';

create table if not exists accounting_periods (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    fiscal_year_id uuid not null,
    company_id uuid not null,
    period_no smallint not null,
    period_name varchar(40) not null,
    start_date date not null,
    end_date date not null,
    status period_status_enum default 'DRAFT' not null,
    is_adjustment_period boolean default false not null,
    soft_locked_at timestamptz,
    soft_locked_by uuid,
    hard_locked_at timestamptz,
    hard_locked_by uuid,
    closed_at timestamptz,
    closed_by uuid,
    lock_reason text,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_accounting_periods primary key (id),
    constraint uq_accounting_periods_1 unique (book_id, fiscal_year_id, period_no, is_adjustment_period),
    constraint ck_accounting_periods_1 check (end_date >= start_date),
    constraint fk_accounting_periods_1 foreign key (book_id) references books (id),
    constraint fk_accounting_periods_2 foreign key (fiscal_year_id) references fiscal_years (id)
);

create index if not exists ix_accounting_periods_1 on accounting_periods (company_id);
create index if not exists ix_accounting_periods_2 on accounting_periods (status);
create index if not exists ix_accounting_periods_3 on accounting_periods (start_date, end_date);

comment on table accounting_periods is 'Period master for month-end control, soft lock, hard lock, and close status. This is a core Accounting-owned feature.';
comment on column accounting_periods.id is 'Primary key.';
comment on column accounting_periods.book_id is 'Reference to books.';
comment on column accounting_periods.fiscal_year_id is 'Reference to fiscal_years.';
comment on column accounting_periods.company_id is 'Denormalized company reference.';
comment on column accounting_periods.period_no is 'Sequential period number inside the fiscal year.';
comment on column accounting_periods.period_name is 'Label such as Apr-2026.';
comment on column accounting_periods.start_date is 'Period start date.';
comment on column accounting_periods.end_date is 'Period end date.';
comment on column accounting_periods.status is 'Lifecycle and lock status.';
comment on column accounting_periods.is_adjustment_period is 'Flags 13th-period or special adjustment periods.';
comment on column accounting_periods.soft_locked_at is 'Soft-lock timestamp.';
comment on column accounting_periods.soft_locked_by is 'User who soft-locked the period.';
comment on column accounting_periods.hard_locked_at is 'Hard-lock timestamp.';
comment on column accounting_periods.hard_locked_by is 'User who hard-locked the period.';
comment on column accounting_periods.closed_at is 'Closed timestamp.';
comment on column accounting_periods.closed_by is 'User who closed the period.';
comment on column accounting_periods.lock_reason is 'Optional business reason for lock or close.';
comment on column accounting_periods.created_by is 'User who created the record.';
comment on column accounting_periods.created_at is 'Record creation timestamp.';
comment on column accounting_periods.updated_by is 'User who last updated the record.';
comment on column accounting_periods.updated_at is 'Last update timestamp.';

create table if not exists account_groups (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    code varchar(30) not null,
    name varchar(120) not null,
    parent_group_id uuid,
    account_nature account_nature_enum not null,
    sort_order integer default 0 not null,
    is_system boolean default false not null,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_account_groups primary key (id),
    constraint uq_account_groups_1 unique (book_id, code),
    constraint fk_account_groups_1 foreign key (book_id) references books (id),
    constraint fk_account_groups_2 foreign key (parent_group_id) references account_groups (id)
);

create index if not exists ix_account_groups_1 on account_groups (parent_group_id);
create index if not exists ix_account_groups_2 on account_groups (account_nature);

comment on table account_groups is 'Hierarchical grouping structure for the chart of accounts. Keeps reporting rollups explicit and owned by Accounting.';
comment on column account_groups.id is 'Primary key.';
comment on column account_groups.book_id is 'Reference to books.';
comment on column account_groups.code is 'Business code for the account group.';
comment on column account_groups.name is 'Display name.';
comment on column account_groups.parent_group_id is 'Self-reference to parent group.';
comment on column account_groups.account_nature is 'Natural classification inherited by child accounts unless overridden.';
comment on column account_groups.sort_order is 'Display ordering hint.';
comment on column account_groups.is_system is 'Marks seeded groups managed by the platform.';
comment on column account_groups.created_by is 'User who created the record.';
comment on column account_groups.created_at is 'Record creation timestamp.';
comment on column account_groups.updated_by is 'User who last updated the record.';
comment on column account_groups.updated_at is 'Last update timestamp.';

create table if not exists accounts (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    code varchar(40) not null,
    name varchar(150) not null,
    display_name varchar(200),
    group_id uuid not null,
    parent_account_id uuid,
    account_nature account_nature_enum not null,
    usage account_usage_enum default 'POSTABLE' not null,
    normal_balance char(1) not null,
    control_type varchar(40),
    allow_manual_posting boolean default true not null,
    require_party boolean default false not null,
    require_branch boolean default false not null,
    require_cost_center boolean default false not null,
    require_employee boolean default false not null,
    require_tax_breakup boolean default false not null,
    is_system boolean default false not null,
    is_active boolean default true not null,
    external_dimension_policy jsonb default '{}'::jsonb not null,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_accounts primary key (id),
    constraint uq_accounts_1 unique (book_id, code),
    constraint ck_accounts_1 check (normal_balance in ('D','C')),
    constraint fk_accounts_1 foreign key (book_id) references books (id),
    constraint fk_accounts_2 foreign key (group_id) references account_groups (id),
    constraint fk_accounts_3 foreign key (parent_account_id) references accounts (id)
);

create index if not exists ix_accounts_1 on accounts (company_id);
create index if not exists ix_accounts_2 on accounts (group_id);
create index if not exists ix_accounts_3 on accounts (parent_account_id);
create index if not exists ix_accounts_4 on accounts (account_nature);
create index if not exists ix_accounts_5 on accounts (control_type);
create index if not exists ix_accounts_6 on accounts (is_active);

comment on table accounts is 'Core chart of accounts table. Stores postable and header accounts, control-account flags, and dimension requirements used by the posting engine.';
comment on column accounts.id is 'Primary key.';
comment on column accounts.book_id is 'Reference to books.';
comment on column accounts.company_id is 'Denormalized company reference.';
comment on column accounts.code is 'Account code visible to finance users.';
comment on column accounts.name is 'Account name.';
comment on column accounts.display_name is 'Optional longer display label.';
comment on column accounts.group_id is 'Reference to account_groups.';
comment on column accounts.parent_account_id is 'Self-reference for header/postable trees.';
comment on column accounts.account_nature is 'Asset, liability, equity, income, or expense.';
comment on column accounts.usage is 'Header or postable account.';
comment on column accounts.normal_balance is 'D for debit-normal, C for credit-normal.';
comment on column accounts.control_type is 'Optional semantic flag such as AR_CONTROL or AP_CONTROL.';
comment on column accounts.allow_manual_posting is 'Whether manual journals may post directly to this account.';
comment on column accounts.require_party is 'Whether a party reference is mandatory on journal lines.';
comment on column accounts.require_branch is 'Whether a branch reference is mandatory on journal lines.';
comment on column accounts.require_cost_center is 'Whether a cost centre reference is mandatory on journal lines.';
comment on column accounts.require_employee is 'Whether an employee reference is mandatory on journal lines.';
comment on column accounts.require_tax_breakup is 'Whether tax breakup metadata is mandatory for posting.';
comment on column accounts.is_system is 'Marks system-seeded accounts.';
comment on column accounts.is_active is 'Inactive accounts remain historical but cannot be selected for new posting.';
comment on column accounts.external_dimension_policy is 'Additional validation hints for external dimensions.';
comment on column accounts.created_by is 'User who created the record.';
comment on column accounts.created_at is 'Record creation timestamp.';
comment on column accounts.updated_by is 'User who last updated the record.';
comment on column accounts.updated_at is 'Last update timestamp.';

create table if not exists voucher_sequences (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    branch_id uuid,
    voucher_type varchar(40) not null,
    prefix varchar(20),
    suffix varchar(10),
    padding smallint default 5 not null,
    next_number bigint default 1 not null,
    reset_policy varchar(20) default 'FY' not null,
    is_active boolean default true not null,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_voucher_sequences primary key (id),
    constraint ck_voucher_sequences_1 check (padding between 1 and 10),
    constraint ck_voucher_sequences_2 check (next_number > 0),
    constraint ck_voucher_sequences_3 check (reset_policy in ('FY','NEVER')),
    constraint fk_voucher_sequences_1 foreign key (book_id) references books (id)
);

create index if not exists ix_voucher_sequences_1 on voucher_sequences (book_id, voucher_type);
create index if not exists ix_voucher_sequences_2 on voucher_sequences (company_id);
create index if not exists ix_voucher_sequences_3 on voucher_sequences (branch_id);
create unique index if not exists ux_voucher_sequences_branch on voucher_sequences (book_id, company_id, branch_id, voucher_type) where branch_id is not null;
create unique index if not exists ux_voucher_sequences_nobranch on voucher_sequences (book_id, company_id, voucher_type) where branch_id is null;

comment on table voucher_sequences is 'Accounting-owned voucher numbering configuration. Supports manual and system voucher number generation without relying on another module.';
comment on column voucher_sequences.id is 'Primary key.';
comment on column voucher_sequences.book_id is 'Reference to books.';
comment on column voucher_sequences.company_id is 'Denormalized company reference.';
comment on column voucher_sequences.branch_id is 'Optional soft reference to Setup.branch for branch-specific numbering.';
comment on column voucher_sequences.voucher_type is 'Voucher bucket such as JV, RV, PV, SV.';
comment on column voucher_sequences.prefix is 'Prefix applied before the numeric portion.';
comment on column voucher_sequences.suffix is 'Suffix applied after the numeric portion.';
comment on column voucher_sequences.padding is 'Zero-padding width for the numeric sequence.';
comment on column voucher_sequences.next_number is 'Next sequence value to allocate.';
comment on column voucher_sequences.reset_policy is 'Reset policy such as FY or NEVER.';
comment on column voucher_sequences.is_active is 'Whether the sequence is available for use.';
comment on column voucher_sequences.created_by is 'User who created the record.';
comment on column voucher_sequences.created_at is 'Record creation timestamp.';
comment on column voucher_sequences.updated_by is 'User who last updated the record.';
comment on column voucher_sequences.updated_at is 'Last update timestamp.';

create table if not exists journal_batches (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    branch_id uuid,
    batch_no varchar(40) not null,
    source_module varchar(50),
    batch_type varchar(30) default 'STANDARD' not null,
    posting_date date not null,
    status journal_batch_status_enum default 'OPEN' not null,
    narration text,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_journal_batches primary key (id),
    constraint uq_journal_batches_1 unique (book_id, batch_no),
    constraint fk_journal_batches_1 foreign key (book_id) references books (id)
);

create index if not exists ix_journal_batches_1 on journal_batches (company_id);
create index if not exists ix_journal_batches_2 on journal_batches (status);
create index if not exists ix_journal_batches_3 on journal_batches (posting_date);

comment on table journal_batches is 'Optional grouping layer for journals produced together, especially by system posting events or bulk close operations.';
comment on column journal_batches.id is 'Primary key.';
comment on column journal_batches.book_id is 'Reference to books.';
comment on column journal_batches.company_id is 'Denormalized company reference.';
comment on column journal_batches.branch_id is 'Optional soft reference to branch.';
comment on column journal_batches.batch_no is 'Human-readable batch identifier.';
comment on column journal_batches.source_module is 'Originating module such as BILLING or PURCHASE.';
comment on column journal_batches.batch_type is 'Business grouping classification.';
comment on column journal_batches.posting_date is 'Posting date shared by the batch.';
comment on column journal_batches.status is 'Batch lifecycle status.';
comment on column journal_batches.narration is 'Optional high-level description.';
comment on column journal_batches.created_by is 'User or service that created the batch.';
comment on column journal_batches.created_at is 'Record creation timestamp.';
comment on column journal_batches.updated_by is 'User who last updated the record.';
comment on column journal_batches.updated_at is 'Last update timestamp.';

create table if not exists journals (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    batch_id uuid,
    fiscal_year_id uuid not null,
    period_id uuid not null,
    company_id uuid not null,
    branch_id uuid,
    journal_no varchar(40) not null,
    journal_kind journal_kind_enum not null,
    status journal_status_enum default 'DRAFT' not null,
    source_module varchar(50),
    source_document_type varchar(50),
    source_document_id uuid,
    source_event_id varchar(100),
    idempotency_key varchar(150),
    journal_date date not null,
    posting_date date not null,
    currency_code char(3) not null,
    exchange_rate numeric(20,8) default 1 not null,
    reference_no varchar(80),
    external_reference_no varchar(80),
    narration text,
    reversal_of_journal_id uuid,
    reversed_by_journal_id uuid,
    metadata jsonb default '{}'::jsonb not null,
    posted_at timestamptz,
    posted_by uuid,
    reversed_at timestamptz,
    reversed_by uuid,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_journals primary key (id),
    constraint uq_journals_1 unique (book_id, journal_no),
    constraint ck_journals_1 check (exchange_rate > 0),
    constraint fk_journals_1 foreign key (book_id) references books (id),
    constraint fk_journals_2 foreign key (batch_id) references journal_batches (id),
    constraint fk_journals_3 foreign key (fiscal_year_id) references fiscal_years (id),
    constraint fk_journals_4 foreign key (period_id) references accounting_periods (id),
    constraint fk_journals_5 foreign key (reversal_of_journal_id) references journals (id),
    constraint fk_journals_6 foreign key (reversed_by_journal_id) references journals (id)
);

create index if not exists ix_journals_1 on journals (company_id);
create index if not exists ix_journals_2 on journals (period_id);
create index if not exists ix_journals_3 on journals (status);
create index if not exists ix_journals_4 on journals (source_module, source_document_type, source_document_id);
create index if not exists ix_journals_5 on journals (journal_date);
create index if not exists ix_journals_6 on journals (posting_date);
create unique index if not exists ux_journals_idempotency_key on journals (book_id, idempotency_key) where idempotency_key is not null;

comment on table journals is 'Journal header table for manual and system-generated vouchers. This is the core immutable accounting event once posted.';
comment on column journals.id is 'Primary key.';
comment on column journals.book_id is 'Reference to books.';
comment on column journals.batch_id is 'Optional reference to journal_batches.';
comment on column journals.fiscal_year_id is 'Reference to fiscal_years.';
comment on column journals.period_id is 'Reference to accounting_periods.';
comment on column journals.company_id is 'Denormalized company reference.';
comment on column journals.branch_id is 'Optional soft reference to branch.';
comment on column journals.journal_no is 'Voucher number.';
comment on column journals.journal_kind is 'Manual, system, opening, reversal, and so on.';
comment on column journals.status is 'Journal lifecycle status.';
comment on column journals.source_module is 'Originating module when system-generated.';
comment on column journals.source_document_type is 'Originating document type such as SALES_INVOICE.';
comment on column journals.source_document_id is 'Soft reference to originating source document.';
comment on column journals.source_event_id is 'Event or message identifier for idempotent posting.';
comment on column journals.idempotency_key is 'Optional dedupe key at journal level.';
comment on column journals.journal_date is 'Document date visible to finance.';
comment on column journals.posting_date is 'Effective posting date for books.';
comment on column journals.currency_code is 'Transaction currency at header level.';
comment on column journals.exchange_rate is 'Header-level exchange rate snapshot.';
comment on column journals.reference_no is 'Internal or external finance reference.';
comment on column journals.external_reference_no is 'Original source module reference number.';
comment on column journals.narration is 'Narration or description.';
comment on column journals.reversal_of_journal_id is 'Reference to original journal when this is a reversal.';
comment on column journals.reversed_by_journal_id is 'Reference to reversal journal when this row has been reversed.';
comment on column journals.metadata is 'Additional immutable snapshot metadata.';
comment on column journals.posted_at is 'Timestamp when the journal was posted.';
comment on column journals.posted_by is 'User or service that posted the journal.';
comment on column journals.reversed_at is 'Timestamp when the journal was reversed.';
comment on column journals.reversed_by is 'User or service that initiated reversal.';
comment on column journals.created_by is 'User or service that created the draft row.';
comment on column journals.created_at is 'Record creation timestamp.';
comment on column journals.updated_by is 'User who last updated the record.';
comment on column journals.updated_at is 'Last update timestamp.';

create table if not exists journal_lines (
    id uuid default gen_random_uuid() not null,
    journal_id uuid not null,
    line_no integer not null,
    account_id uuid not null,
    company_id uuid not null,
    branch_id uuid,
    party_id uuid,
    employee_id uuid,
    cost_center_id uuid,
    income_head_id uuid,
    tax_code_id uuid,
    project_id uuid,
    item_id uuid,
    description text,
    debit_amount_txn numeric(20,6) default 0 not null,
    credit_amount_txn numeric(20,6) default 0 not null,
    debit_amount_base numeric(20,6) default 0 not null,
    credit_amount_base numeric(20,6) default 0 not null,
    exchange_rate numeric(20,8) default 1 not null,
    line_metadata jsonb default '{}'::jsonb not null,
    created_by uuid,
    created_at timestamptz default now() not null,
    constraint pk_journal_lines primary key (id),
    constraint uq_journal_lines_1 unique (journal_id, line_no),
    constraint ck_journal_lines_1 check (debit_amount_txn >= 0 and credit_amount_txn >= 0 and debit_amount_base >= 0 and credit_amount_base >= 0),
    constraint ck_journal_lines_2 check (exchange_rate > 0),
    constraint ck_journal_lines_3 check (((debit_amount_txn > 0 and credit_amount_txn = 0) or (credit_amount_txn > 0 and debit_amount_txn = 0))),
    constraint fk_journal_lines_1 foreign key (journal_id) references journals (id),
    constraint fk_journal_lines_2 foreign key (account_id) references accounts (id)
);

create index if not exists ix_journal_lines_1 on journal_lines (account_id);
create index if not exists ix_journal_lines_2 on journal_lines (party_id);
create index if not exists ix_journal_lines_3 on journal_lines (branch_id);
create index if not exists ix_journal_lines_4 on journal_lines (cost_center_id);
create index if not exists ix_journal_lines_5 on journal_lines (tax_code_id);

comment on table journal_lines is 'Journal line table containing the debits, credits, external dimensions, and snapshots needed for the ledger and downstream subledger processing.';
comment on column journal_lines.id is 'Primary key.';
comment on column journal_lines.journal_id is 'Reference to journals.';
comment on column journal_lines.line_no is 'Sequential line number within the journal.';
comment on column journal_lines.account_id is 'Reference to accounts.';
comment on column journal_lines.company_id is 'Denormalized company reference.';
comment on column journal_lines.branch_id is 'Optional soft reference to branch.';
comment on column journal_lines.party_id is 'Soft reference to customer, vendor, or party master.';
comment on column journal_lines.employee_id is 'Soft reference to employee.';
comment on column journal_lines.cost_center_id is 'Soft reference to Catalogue cost centre.';
comment on column journal_lines.income_head_id is 'Soft reference to Catalogue income head.';
comment on column journal_lines.tax_code_id is 'Soft reference to Catalogue tax code.';
comment on column journal_lines.project_id is 'Soft reference to project if introduced later.';
comment on column journal_lines.item_id is 'Soft reference to item when source traceability needs it.';
comment on column journal_lines.description is 'Line-level description.';
comment on column journal_lines.debit_amount_txn is 'Debit amount in transaction currency.';
comment on column journal_lines.credit_amount_txn is 'Credit amount in transaction currency.';
comment on column journal_lines.debit_amount_base is 'Debit amount in base currency.';
comment on column journal_lines.credit_amount_base is 'Credit amount in base currency.';
comment on column journal_lines.exchange_rate is 'Line-level exchange rate snapshot.';
comment on column journal_lines.line_metadata is 'Snapshot payload for extra dimensions or tax breakup.';
comment on column journal_lines.created_by is 'User or service that created the line.';
comment on column journal_lines.created_at is 'Record creation timestamp.';

create table if not exists posting_rule_versions (
    id uuid default gen_random_uuid() not null,
    source_module varchar(50) not null,
    source_document_type varchar(50) not null,
    version_no integer not null,
    name varchar(120) not null,
    status posting_rule_status_enum default 'DRAFT' not null,
    effective_from timestamptz default now() not null,
    effective_to timestamptz,
    rule_payload jsonb not null,
    notes text,
    created_by uuid,
    created_at timestamptz default now() not null,
    updated_by uuid,
    updated_at timestamptz default now() not null,
    constraint pk_posting_rule_versions primary key (id),
    constraint uq_posting_rule_versions_1 unique (source_module, source_document_type, version_no),
    constraint ck_posting_rule_versions_1 check (effective_to is null or effective_to >= effective_from)
);

create index if not exists ix_posting_rule_versions_1 on posting_rule_versions (status);
create index if not exists ix_posting_rule_versions_2 on posting_rule_versions (effective_from, effective_to);

comment on table posting_rule_versions is 'Versioned posting-rule registry that lets Accounting own financial recognition logic without owning source documents.';
comment on column posting_rule_versions.id is 'Primary key.';
comment on column posting_rule_versions.source_module is 'Originating module such as BILLING or PURCHASE.';
comment on column posting_rule_versions.source_document_type is 'Originating document type.';
comment on column posting_rule_versions.version_no is 'Monotonic rule version number.';
comment on column posting_rule_versions.name is 'Display name for the rule set.';
comment on column posting_rule_versions.status is 'Draft, active, or retired status.';
comment on column posting_rule_versions.effective_from is 'Rule start timestamp.';
comment on column posting_rule_versions.effective_to is 'Optional end timestamp.';
comment on column posting_rule_versions.rule_payload is 'Rule definition payload consumed by the posting engine.';
comment on column posting_rule_versions.notes is 'Human notes for finance and engineering.';
comment on column posting_rule_versions.created_by is 'User who created the version.';
comment on column posting_rule_versions.created_at is 'Record creation timestamp.';
comment on column posting_rule_versions.updated_by is 'User who last updated the version.';
comment on column posting_rule_versions.updated_at is 'Last update timestamp.';

create table if not exists posting_requests (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    branch_id uuid,
    source_module varchar(50) not null,
    source_document_type varchar(50) not null,
    source_document_id uuid not null,
    source_event_id varchar(100) not null,
    idempotency_key varchar(150) not null,
    request_hash char(64) not null,
    request_status posting_request_status_enum default 'RECEIVED' not null,
    requested_posting_date date not null,
    requested_by uuid,
    rule_version_id uuid,
    current_journal_id uuid,
    error_code varchar(60),
    error_message text,
    retry_count integer default 0 not null,
    first_received_at timestamptz default now() not null,
    last_processed_at timestamptz,
    last_failed_at timestamptz,
    request_payload jsonb not null,
    created_at timestamptz default now() not null,
    updated_at timestamptz default now() not null,
    constraint pk_posting_requests primary key (id),
    constraint uq_posting_requests_1 unique (book_id, idempotency_key),
    constraint ck_posting_requests_1 check (retry_count >= 0),
    constraint fk_posting_requests_1 foreign key (book_id) references books (id),
    constraint fk_posting_requests_2 foreign key (rule_version_id) references posting_rule_versions (id),
    constraint fk_posting_requests_3 foreign key (current_journal_id) references journals (id)
);

create index if not exists ix_posting_requests_1 on posting_requests (company_id);
create index if not exists ix_posting_requests_2 on posting_requests (request_status);
create index if not exists ix_posting_requests_3 on posting_requests (source_module, source_document_type, source_document_id);
create index if not exists ix_posting_requests_4 on posting_requests (source_event_id);
create index if not exists ix_posting_requests_5 on posting_requests (last_processed_at);

comment on table posting_requests is 'Idempotent inbound posting envelope for documents coming from Billing, Purchase, Expense, and later Payroll.';
comment on column posting_requests.id is 'Primary key.';
comment on column posting_requests.book_id is 'Reference to books.';
comment on column posting_requests.company_id is 'Denormalized company reference.';
comment on column posting_requests.branch_id is 'Optional soft reference to branch.';
comment on column posting_requests.source_module is 'Originating module.';
comment on column posting_requests.source_document_type is 'Originating document type.';
comment on column posting_requests.source_document_id is 'Soft reference to the source document.';
comment on column posting_requests.source_event_id is 'Source event identifier.';
comment on column posting_requests.idempotency_key is 'Dedupe key used by the posting engine.';
comment on column posting_requests.request_hash is 'Hash of the canonical posting payload for replay diagnostics.';
comment on column posting_requests.request_status is 'Current processing status.';
comment on column posting_requests.requested_posting_date is 'Posting date requested by the source event.';
comment on column posting_requests.requested_by is 'User responsible for the request, when available.';
comment on column posting_requests.rule_version_id is 'Reference to the posting rule version chosen for processing.';
comment on column posting_requests.current_journal_id is 'Latest primary journal created for the request.';
comment on column posting_requests.error_code is 'Normalized posting failure code.';
comment on column posting_requests.error_message is 'Verbose posting failure message.';
comment on column posting_requests.retry_count is 'How many times the request has been retried.';
comment on column posting_requests.first_received_at is 'Timestamp when the request first entered accounting.';
comment on column posting_requests.last_processed_at is 'Last processing attempt timestamp.';
comment on column posting_requests.last_failed_at is 'Last failure timestamp.';
comment on column posting_requests.request_payload is 'Canonical request payload.';
comment on column posting_requests.created_at is 'Record creation timestamp.';
comment on column posting_requests.updated_at is 'Last update timestamp.';

create table if not exists posting_request_snapshots (
    id uuid default gen_random_uuid() not null,
    posting_request_id uuid not null,
    snapshot_type varchar(40) not null,
    snapshot_version integer default 1 not null,
    document_number varchar(60),
    document_date date,
    counterparty_name varchar(200),
    currency_code char(3),
    gross_amount_txn numeric(20,6),
    net_amount_txn numeric(20,6),
    tax_amount_txn numeric(20,6),
    snapshot_payload jsonb not null,
    captured_at timestamptz default now() not null,
    constraint pk_posting_request_snapshots primary key (id),
    constraint uq_posting_request_snapshots_1 unique (posting_request_id, snapshot_type, snapshot_version),
    constraint fk_posting_request_snapshots_1 foreign key (posting_request_id) references posting_requests (id)
);

create index if not exists ix_posting_request_snapshots_1 on posting_request_snapshots (snapshot_type);

comment on table posting_request_snapshots is 'Immutable source snapshots captured at posting time. Prevents later source-document edits from mutating the accounting audit trail.';
comment on column posting_request_snapshots.id is 'Primary key.';
comment on column posting_request_snapshots.posting_request_id is 'Reference to posting_requests.';
comment on column posting_request_snapshots.snapshot_type is 'Logical snapshot bucket such as HEADER or TAX.';
comment on column posting_request_snapshots.snapshot_version is 'Version number for the stored snapshot type.';
comment on column posting_request_snapshots.document_number is 'Source document number captured at posting time.';
comment on column posting_request_snapshots.document_date is 'Source document date captured at posting time.';
comment on column posting_request_snapshots.counterparty_name is 'Immutable customer or vendor display name snapshot.';
comment on column posting_request_snapshots.currency_code is 'Snapshot currency code.';
comment on column posting_request_snapshots.gross_amount_txn is 'Gross amount in transaction currency.';
comment on column posting_request_snapshots.net_amount_txn is 'Net amount in transaction currency.';
comment on column posting_request_snapshots.tax_amount_txn is 'Tax amount in transaction currency.';
comment on column posting_request_snapshots.snapshot_payload is 'Full immutable source snapshot.';
comment on column posting_request_snapshots.captured_at is 'Capture timestamp.';

create table if not exists journal_source_links (
    id uuid default gen_random_uuid() not null,
    posting_request_id uuid not null,
    journal_id uuid not null,
    link_role journal_link_role_enum default 'PRIMARY' not null,
    is_reversal boolean default false not null,
    created_at timestamptz default now() not null,
    constraint pk_journal_source_links primary key (id),
    constraint uq_journal_source_links_1 unique (posting_request_id, journal_id, link_role),
    constraint fk_journal_source_links_1 foreign key (posting_request_id) references posting_requests (id),
    constraint fk_journal_source_links_2 foreign key (journal_id) references journals (id)
);

create index if not exists ix_journal_source_links_1 on journal_source_links (journal_id);
create index if not exists ix_journal_source_links_2 on journal_source_links (link_role);

comment on table journal_source_links is 'Explicit many-to-many traceability between posting requests and generated journals. Supports primary, reversal, and adjustment linkage.';
comment on column journal_source_links.id is 'Primary key.';
comment on column journal_source_links.posting_request_id is 'Reference to posting_requests.';
comment on column journal_source_links.journal_id is 'Reference to journals.';
comment on column journal_source_links.link_role is 'How the journal relates to the posting request.';
comment on column journal_source_links.is_reversal is 'Convenience flag for reversal links.';
comment on column journal_source_links.created_at is 'Record creation timestamp.';

create table if not exists open_items (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    branch_id uuid,
    item_side open_item_side_enum not null,
    item_status open_item_status_enum default 'OPEN' not null,
    party_id uuid not null,
    party_type party_type_enum not null,
    control_account_id uuid not null,
    source_module varchar(50) not null,
    source_document_type varchar(50) not null,
    source_document_id uuid not null,
    source_line_ref varchar(100),
    journal_id uuid not null,
    journal_line_id uuid,
    document_no varchar(60) not null,
    document_date date not null,
    due_date date,
    currency_code char(3) not null,
    exchange_rate numeric(20,8) default 1 not null,
    original_amount_txn numeric(20,6) not null,
    original_amount_base numeric(20,6) not null,
    open_amount_txn numeric(20,6) not null,
    open_amount_base numeric(20,6) not null,
    settled_amount_txn numeric(20,6) default 0 not null,
    settled_amount_base numeric(20,6) default 0 not null,
    writeoff_amount_txn numeric(20,6) default 0 not null,
    writeoff_amount_base numeric(20,6) default 0 not null,
    remarks text,
    created_at timestamptz default now() not null,
    updated_at timestamptz default now() not null,
    constraint pk_open_items primary key (id),
    constraint ck_open_items_1 check (exchange_rate > 0),
    constraint ck_open_items_2 check (original_amount_txn >= 0 and original_amount_base >= 0 and open_amount_txn >= 0 and open_amount_base >= 0 and settled_amount_txn >= 0 and settled_amount_base >= 0 and writeoff_amount_txn >= 0 and writeoff_amount_base >= 0),
    constraint fk_open_items_1 foreign key (book_id) references books (id),
    constraint fk_open_items_2 foreign key (control_account_id) references accounts (id),
    constraint fk_open_items_3 foreign key (journal_id) references journals (id),
    constraint fk_open_items_4 foreign key (journal_line_id) references journal_lines (id)
);

create index if not exists ix_open_items_1 on open_items (party_id);
create index if not exists ix_open_items_2 on open_items (control_account_id);
create index if not exists ix_open_items_3 on open_items (item_side, item_status);
create index if not exists ix_open_items_4 on open_items (due_date);
create index if not exists ix_open_items_5 on open_items (source_module, source_document_type, source_document_id);
create unique index if not exists ux_open_items_journal_line on open_items (journal_line_id) where journal_line_id is not null;

comment on table open_items is 'Unified AR/AP open-item ledger owned by Accounting. Tracks receivables, payables, advances, and outstanding balances without duplicating Billing or Purchase documents.';
comment on column open_items.id is 'Primary key.';
comment on column open_items.book_id is 'Reference to books.';
comment on column open_items.company_id is 'Denormalized company reference.';
comment on column open_items.branch_id is 'Optional soft reference to branch.';
comment on column open_items.item_side is 'Receivable or payable.';
comment on column open_items.item_status is 'Open-item lifecycle status.';
comment on column open_items.party_id is 'Soft reference to customer, vendor, or party.';
comment on column open_items.party_type is 'Classification of the referenced party.';
comment on column open_items.control_account_id is 'Reference to the controlling AR or AP account.';
comment on column open_items.source_module is 'Originating module.';
comment on column open_items.source_document_type is 'Originating document type.';
comment on column open_items.source_document_id is 'Soft reference to the originating source document.';
comment on column open_items.source_line_ref is 'Optional source line identifier for line-level open items.';
comment on column open_items.journal_id is 'Reference to journals.';
comment on column open_items.journal_line_id is 'Reference to journal_lines when open item comes from one line.';
comment on column open_items.document_no is 'Source-facing document number snapshot.';
comment on column open_items.document_date is 'Source document date.';
comment on column open_items.due_date is 'Due date for aging and credit control.';
comment on column open_items.currency_code is 'Transaction currency.';
comment on column open_items.exchange_rate is 'Exchange rate snapshot.';
comment on column open_items.original_amount_txn is 'Original open-item amount in transaction currency.';
comment on column open_items.original_amount_base is 'Original open-item amount in base currency.';
comment on column open_items.open_amount_txn is 'Current outstanding amount in transaction currency.';
comment on column open_items.open_amount_base is 'Current outstanding amount in base currency.';
comment on column open_items.settled_amount_txn is 'Cumulative applied settlement amount in transaction currency.';
comment on column open_items.settled_amount_base is 'Cumulative applied settlement amount in base currency.';
comment on column open_items.writeoff_amount_txn is 'Cumulative write-off amount in transaction currency.';
comment on column open_items.writeoff_amount_base is 'Cumulative write-off amount in base currency.';
comment on column open_items.remarks is 'Optional remarks for finance users.';
comment on column open_items.created_at is 'Record creation timestamp.';
comment on column open_items.updated_at is 'Last update timestamp.';

create table if not exists open_item_allocations (
    id uuid default gen_random_uuid() not null,
    book_id uuid not null,
    company_id uuid not null,
    allocation_status allocation_status_enum default 'APPLIED' not null,
    allocation_date date not null,
    from_open_item_id uuid not null,
    to_open_item_id uuid not null,
    allocation_currency_code char(3) not null,
    allocation_amount_txn numeric(20,6) not null,
    allocation_amount_base numeric(20,6) not null,
    allocation_journal_id uuid,
    reference_no varchar(80),
    reversal_of_allocation_id uuid,
    notes text,
    created_by uuid,
    created_at timestamptz default now() not null,
    constraint pk_open_item_allocations primary key (id),
    constraint ck_open_item_allocations_1 check (from_open_item_id <> to_open_item_id),
    constraint ck_open_item_allocations_2 check (allocation_amount_txn > 0 and allocation_amount_base >= 0),
    constraint fk_open_item_allocations_1 foreign key (book_id) references books (id),
    constraint fk_open_item_allocations_2 foreign key (from_open_item_id) references open_items (id),
    constraint fk_open_item_allocations_3 foreign key (to_open_item_id) references open_items (id),
    constraint fk_open_item_allocations_4 foreign key (allocation_journal_id) references journals (id),
    constraint fk_open_item_allocations_5 foreign key (reversal_of_allocation_id) references open_item_allocations (id)
);

create index if not exists ix_open_item_allocations_1 on open_item_allocations (from_open_item_id);
create index if not exists ix_open_item_allocations_2 on open_item_allocations (to_open_item_id);
create index if not exists ix_open_item_allocations_3 on open_item_allocations (allocation_status);
create index if not exists ix_open_item_allocations_4 on open_item_allocations (allocation_date);

comment on table open_item_allocations is 'Settlement table that links open items together for invoice-to-receipt, bill-to-payment, advance application, and partial settlement flows.';
comment on column open_item_allocations.id is 'Primary key.';
comment on column open_item_allocations.book_id is 'Reference to books.';
comment on column open_item_allocations.company_id is 'Denormalized company reference.';
comment on column open_item_allocations.allocation_status is 'Allocation lifecycle status.';
comment on column open_item_allocations.allocation_date is 'Date on which the settlement is considered effective.';
comment on column open_item_allocations.from_open_item_id is 'Open item providing the value being applied.';
comment on column open_item_allocations.to_open_item_id is 'Open item receiving the application.';
comment on column open_item_allocations.allocation_currency_code is 'Currency of the applied value.';
comment on column open_item_allocations.allocation_amount_txn is 'Applied amount in transaction currency.';
comment on column open_item_allocations.allocation_amount_base is 'Applied amount in base currency.';
comment on column open_item_allocations.allocation_journal_id is 'Optional journal backing the allocation or reclassification.';
comment on column open_item_allocations.reference_no is 'Reference number or settlement note.';
comment on column open_item_allocations.reversal_of_allocation_id is 'Self-reference when this row reverses an earlier allocation.';
comment on column open_item_allocations.notes is 'Operational note for finance users.';
comment on column open_item_allocations.created_by is 'User or service that created the allocation.';
comment on column open_item_allocations.created_at is 'Record creation timestamp.';

create table if not exists open_item_adjustments (
    id uuid default gen_random_uuid() not null,
    open_item_id uuid not null,
    adjustment_type open_item_adjustment_type_enum not null,
    adjustment_date date not null,
    adjustment_journal_id uuid,
    amount_txn numeric(20,6) not null,
    amount_base numeric(20,6) not null,
    reason_code varchar(50),
    notes text,
    created_by uuid,
    created_at timestamptz default now() not null,
    constraint pk_open_item_adjustments primary key (id),
    constraint ck_open_item_adjustments_1 check (amount_txn > 0 and amount_base >= 0),
    constraint fk_open_item_adjustments_1 foreign key (open_item_id) references open_items (id),
    constraint fk_open_item_adjustments_2 foreign key (adjustment_journal_id) references journals (id)
);

create index if not exists ix_open_item_adjustments_1 on open_item_adjustments (open_item_id);
create index if not exists ix_open_item_adjustments_2 on open_item_adjustments (adjustment_type);
create index if not exists ix_open_item_adjustments_3 on open_item_adjustments (adjustment_date);

comment on table open_item_adjustments is 'Explicit write-off and adjustment ledger for open items. Keeps non-settlement balance changes auditable and separate from normal allocations.';
comment on column open_item_adjustments.id is 'Primary key.';
comment on column open_item_adjustments.open_item_id is 'Reference to open_items.';
comment on column open_item_adjustments.adjustment_type is 'Type of balance adjustment.';
comment on column open_item_adjustments.adjustment_date is 'Effective date of the adjustment.';
comment on column open_item_adjustments.adjustment_journal_id is 'Journal created for the adjustment, where applicable.';
comment on column open_item_adjustments.amount_txn is 'Adjustment amount in transaction currency.';
comment on column open_item_adjustments.amount_base is 'Adjustment amount in base currency.';
comment on column open_item_adjustments.reason_code is 'Normalized finance reason code.';
comment on column open_item_adjustments.notes is 'Human explanation.';
comment on column open_item_adjustments.created_by is 'User or service that created the adjustment.';
comment on column open_item_adjustments.created_at is 'Record creation timestamp.';

drop trigger if exists trg_books_updated_at on books;
create trigger trg_books_updated_at before update on books for each row execute function touch_updated_at();

drop trigger if exists trg_fiscal_years_updated_at on fiscal_years;
create trigger trg_fiscal_years_updated_at before update on fiscal_years for each row execute function touch_updated_at();

drop trigger if exists trg_accounting_periods_updated_at on accounting_periods;
create trigger trg_accounting_periods_updated_at before update on accounting_periods for each row execute function touch_updated_at();

drop trigger if exists trg_account_groups_updated_at on account_groups;
create trigger trg_account_groups_updated_at before update on account_groups for each row execute function touch_updated_at();

drop trigger if exists trg_accounts_updated_at on accounts;
create trigger trg_accounts_updated_at before update on accounts for each row execute function touch_updated_at();

drop trigger if exists trg_voucher_sequences_updated_at on voucher_sequences;
create trigger trg_voucher_sequences_updated_at before update on voucher_sequences for each row execute function touch_updated_at();

drop trigger if exists trg_journal_batches_updated_at on journal_batches;
create trigger trg_journal_batches_updated_at before update on journal_batches for each row execute function touch_updated_at();

drop trigger if exists trg_journals_updated_at on journals;
create trigger trg_journals_updated_at before update on journals for each row execute function touch_updated_at();

drop trigger if exists trg_posting_rule_versions_updated_at on posting_rule_versions;
create trigger trg_posting_rule_versions_updated_at before update on posting_rule_versions for each row execute function touch_updated_at();

drop trigger if exists trg_posting_requests_updated_at on posting_requests;
create trigger trg_posting_requests_updated_at before update on posting_requests for each row execute function touch_updated_at();

drop trigger if exists trg_open_items_updated_at on open_items;
create trigger trg_open_items_updated_at before update on open_items for each row execute function touch_updated_at();

create or replace view v_journal_balance_check as
select
    j.id as journal_id,
    j.journal_no,
    j.status,
    sum(jl.debit_amount_base) as total_debit_base,
    sum(jl.credit_amount_base) as total_credit_base,
    sum(jl.debit_amount_base) - sum(jl.credit_amount_base) as imbalance_base
from journals j
join journal_lines jl
  on jl.journal_id = j.id
group by j.id, j.journal_no, j.status;

create or replace view v_open_item_outstanding as
select
    oi.id,
    oi.book_id,
    oi.company_id,
    oi.branch_id,
    oi.item_side,
    oi.item_status,
    oi.party_id,
    oi.party_type,
    oi.document_no,
    oi.document_date,
    oi.due_date,
    oi.currency_code,
    oi.original_amount_txn,
    oi.open_amount_txn,
    oi.original_amount_base,
    oi.open_amount_base
from open_items oi
where oi.item_status in ('OPEN', 'PARTIALLY_ALLOCATED');