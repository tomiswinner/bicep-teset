CREATE TABLE [dbo].[kabu_transaction] (
    [kabu_transaction_id] VARCHAR (20) NOT NULL,
    [giver_id]            VARCHAR (20) NOT NULL,
    [receiver_id]         VARCHAR (20) NOT NULL,
    [num_sent]            INT          NOT NULL,
    [given_at]            ROWVERSION   NOT NULL,
    PRIMARY KEY CLUSTERED ([kabu_transaction_id] ASC),
    FOREIGN KEY ([giver_id]) REFERENCES [dbo].[employee] ([employee_id]),
    FOREIGN KEY ([receiver_id]) REFERENCES [dbo].[employee] ([employee_id])
);


GO

CREATE TABLE [dbo].[employee] (
    [employee_id]   VARCHAR (20) NOT NULL,
    [employee_name] VARCHAR (20) NOT NULL,
    [slack_id]      VARCHAR (20) NOT NULL,
    PRIMARY KEY CLUSTERED ([employee_id] ASC)
);


GO

