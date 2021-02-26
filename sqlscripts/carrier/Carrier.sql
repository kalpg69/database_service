-- Create a new table called '[Carrier]' in schema '[dbo]'
-- Drop the table if it already exists
IF OBJECT_ID('[dbo].[Carrier]', 'U') IS NOT NULL
DROP TABLE [dbo].[Carrier]
GO
-- Create the table in the specified schema
CREATE TABLE [dbo].[Carrier]
(
    [Id] BIGINT IDENTITY(1,1) NOT NULL PRIMARY KEY, -- Primary Key column
    [CarrierCode] NVARCHAR(50) NOT NULL,
    [CarrierName] NVARCHAR(50) NULL,
    [CarrierAddress] NVARCHAR(50) NULL,
    [CarrierEmail] NVARCHAR(50) NULL,
    [CarrierPhone] NVARCHAR(50) NULL
);
GO