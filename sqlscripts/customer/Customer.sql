-- Create a new table called '[Customer]' in schema '[dbo]'
-- Drop the table if it already exists
IF OBJECT_ID('[dbo].[Customer]', 'U') IS NOT NULL
DROP TABLE [dbo].[Customer]
GO
-- Create the table in the specified schema
CREATE TABLE [dbo].[Customer]
(
    [Id] BIGINT IDENTITY(1,1) NOT NULL PRIMARY KEY, -- Primary Key column
    [CustomerCode] NVARCHAR(50) NOT NULL,
    [CustomerName] NVARCHAR(50) NULL,
    [CustomerAddress] NVARCHAR(50) NULL,
    [CustomerEmail] NVARCHAR(50) NULL,
    [CustomerPhone] NVARCHAR(50) NULL
);
GO