CREATE TABLE Account (
    AccountID       INT64 NOT NULL,
    Name            STRING(256) NOT NULL
) PRIMARY KEY (AccountID);

CREATE TABLE Singer (
    SingerID        STRING(36) NOT NULL,
    Name            STRING(256) NOT NULL
) PRIMARY KEY (SingerID);
